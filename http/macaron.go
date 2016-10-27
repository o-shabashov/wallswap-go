package main

import (
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/oauth2"
    goauth2 "golang.org/x/oauth2"
    "github.com/go-macaron/session"
    "github.com/o-shabashov/wallswap-go/wallswap"
    "os"
)

func main() {
    m := macaron.Classic()
    m.Use(session.Sessioner())

    m.Use(wallswap.OAuthProvider(
        &goauth2.Config{
            ClientID:     os.Getenv("DROPBOX_CLIENT_ID"),
            ClientSecret: os.Getenv("DROPBOX_CLIENT_SECRET"),
            RedirectURL:  "http://localhost:8081/oauth2callback",
        },
    ))

    // Tokens are injected to the handlers
    m.Get("/", func(tokens oauth2.Tokens) string {

        if tokens.Expired() {
            return "not logged in, or the access token is expired"
        } else {
            user := wallswap.AuthUser(tokens)
            return user.DropboxId
        }

        return tokens.Access()
    })

    m.Run(8081)
}
