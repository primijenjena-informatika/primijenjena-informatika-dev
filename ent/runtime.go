// Code generated by ent, DO NOT EDIT.

package ent

import (
	"primijenjena-informatika-dev/ent/codespace"
	"primijenjena-informatika-dev/ent/schema"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	codespaceFields := schema.Codespace{}.Fields()
	_ = codespaceFields
	// codespaceDescCreatedAt is the schema descriptor for created_at field.
	codespaceDescCreatedAt := codespaceFields[1].Descriptor()
	// codespace.DefaultCreatedAt holds the default value on creation for the created_at field.
	codespace.DefaultCreatedAt = codespaceDescCreatedAt.Default.(func() time.Time)
	// codespaceDescUpdatedAt is the schema descriptor for updated_at field.
	codespaceDescUpdatedAt := codespaceFields[2].Descriptor()
	// codespace.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	codespace.DefaultUpdatedAt = codespaceDescUpdatedAt.Default.(func() time.Time)
	// codespace.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	codespace.UpdateDefaultUpdatedAt = codespaceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// codespaceDescMachineType is the schema descriptor for machine_type field.
	codespaceDescMachineType := codespaceFields[3].Descriptor()
	// codespace.MachineTypeValidator is a validator for the "machine_type" field. It is called by the builders before save.
	codespace.MachineTypeValidator = codespaceDescMachineType.Validators[0].(func(string) error)
	// codespaceDescClientIP is the schema descriptor for client_ip field.
	codespaceDescClientIP := codespaceFields[4].Descriptor()
	// codespace.ClientIPValidator is a validator for the "client_ip" field. It is called by the builders before save.
	codespace.ClientIPValidator = codespaceDescClientIP.Validators[0].(func(string) error)
	// codespaceDescGithubUserID is the schema descriptor for github_user_id field.
	codespaceDescGithubUserID := codespaceFields[5].Descriptor()
	// codespace.GithubUserIDValidator is a validator for the "github_user_id" field. It is called by the builders before save.
	codespace.GithubUserIDValidator = codespaceDescGithubUserID.Validators[0].(func(string) error)
	// codespaceDescGithubCodespaceID is the schema descriptor for github_codespace_id field.
	codespaceDescGithubCodespaceID := codespaceFields[6].Descriptor()
	// codespace.GithubCodespaceIDValidator is a validator for the "github_codespace_id" field. It is called by the builders before save.
	codespace.GithubCodespaceIDValidator = codespaceDescGithubCodespaceID.Validators[0].(func(string) error)
	// codespaceDescGithubCodespaceURL is the schema descriptor for github_codespace_url field.
	codespaceDescGithubCodespaceURL := codespaceFields[7].Descriptor()
	// codespace.GithubCodespaceURLValidator is a validator for the "github_codespace_url" field. It is called by the builders before save.
	codespace.GithubCodespaceURLValidator = codespaceDescGithubCodespaceURL.Validators[0].(func(string) error)
	// codespaceDescID is the schema descriptor for id field.
	codespaceDescID := codespaceFields[0].Descriptor()
	// codespace.DefaultID holds the default value on creation for the id field.
	codespace.DefaultID = codespaceDescID.Default.(func() uuid.UUID)
}
