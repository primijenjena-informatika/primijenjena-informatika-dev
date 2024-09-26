document.addEventListener("DOMContentLoaded", () => {
    document.querySelector("#login-button").addEventListener("click", () => {
        window.location.href = "http://localhost:8080/oauth";
    });
});