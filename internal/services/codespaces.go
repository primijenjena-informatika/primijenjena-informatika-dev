package services

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"primijenjena-informatika-dev/ent"
	"primijenjena-informatika-dev/ent/codespace"
	"strconv"
	"strings"

	"github.com/cenkalti/backoff/v4"
	"github.com/google/go-github/v64/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func OpenCodespace(db *ent.Client, code string, clientIp string) (string, error) {
	fmt.Println(code)
	userClient, orgClient, err := GetClients(code)
	if err != nil {
		return "", err
	}

	user, _, err := userClient.Users.Get(context.Background(), "")
	if err != nil {
		log.Println(err)
		return "", err
	}

	codespace, err := db.Codespace.Query().Where(codespace.GithubUserID(strconv.Itoa(int(user.GetID())))).First(context.Background())
	if err == nil {
		return codespace.GithubCodespaceURL, nil
	}

	createdCodespace, err := NewCodespace(db, userClient, orgClient, code, clientIp)
	if err != nil {
		return "", err
	}

	return createdCodespace.GetWebURL(), nil
}

func GetClients(code string) (*github.Client, *github.Client, error) {
	fmt.Println(viper.GetString("github_client_id"))
	fmt.Println(viper.GetString("github_client_secret"))
	// Create an OAuth2 config
	conf := &oauth2.Config{
		ClientID:     viper.GetString("github_client_id"),
		ClientSecret: viper.GetString("github_client_secret"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}

	// Exchange the code for an access token
	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Error exchanging code for access token: ", err)
		return nil, nil, err
	}

	// Use the access token to create a new GitHub client
	userClient := github.NewClient(conf.Client(context.Background(), token))

	githubToken := viper.GetString("github_org_token")
	orgClient := github.NewClient(conf.Client(context.Background(), &oauth2.Token{AccessToken: githubToken}))

	return userClient, orgClient, nil
}

func NewCodespace(db *ent.Client, userClient *github.Client, orgClient *github.Client, code string, clientIp string) (*github.Codespace, error) {
	user, _, err := userClient.Users.Get(context.Background(), "")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	repoName := rand.IntN(int(math.Exp2(40)))

	templateOrg := strings.Split(viper.GetString("github_org_name"), "/")[0]
	templateRepo := strings.Split(viper.GetString("github_org_name"), "/")[1]

	_, _, err = orgClient.Repositories.CreateFromTemplate(context.Background(), templateOrg, templateRepo, &github.TemplateRepoRequest{
		Name:    github.String(strconv.Itoa(repoName)),
		Owner:   github.String(viper.GetString("github_org_name")),
		Private: github.Bool(true),
	})
	if err != nil {
		log.Println("Error creating repository:", err)
		return nil, err
	}

	invite, _, err := orgClient.Repositories.AddCollaborator(context.Background(), viper.GetString("github_org_name"), strconv.Itoa(repoName), user.GetLogin(), &github.RepositoryAddCollaboratorOptions{
		Permission: "write",
	})
	if err != nil {
		log.Println("Error adding collaborator: ", err)
		return nil, err
	}

	_, err = userClient.Users.AcceptInvitation(context.Background(), invite.GetID())
	fmt.Println(err.Error())
	if err != nil {
		log.Println("Error accepting invitation: ", err)
		return nil, err
	}

	var createdCodespace *github.Codespace

	operation := func() error {
		createdCodespace, _, err = userClient.Codespaces.CreateInRepo(context.Background(), viper.GetString("github_org_name"), strconv.Itoa(repoName), &github.CreateCodespaceOptions{
			Ref:                    github.String(viper.GetString("codespace_branch")),
			ClientIP:               github.String(clientIp),
			RetentionPeriodMinutes: github.Int(viper.GetInt("codespace_retention_period")),
			Machine:                github.String(viper.GetString("codespace_machine")),
		})
		if err != nil {
			return err
		}

		return nil
	}

	err = backoff.Retry(operation, backoff.NewExponentialBackOff())
	if err != nil {
		log.Println("Error creating Codespace after retries:", err)
		return nil, err
	}

	go func() {
		db.Codespace.Create().
			SetMachineType(createdCodespace.GetMachine().GetName()).
			SetClientIP(clientIp).
			SetGithubUserID(strconv.Itoa(int(user.GetID()))).
			SetGithubCodespaceID(strconv.Itoa(int(createdCodespace.GetID()))).
			SetGithubCodespaceURL(createdCodespace.GetWebURL()).
			Exec(context.Background())
	}()

	return createdCodespace, nil
}
