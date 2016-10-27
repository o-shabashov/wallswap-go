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
    m.Use(macaron.Renderer())
    m.Use(macaron.Recovery())
    m.Use(macaron.Static("assets"))

    m.Use(wallswap.OAuthProvider(
        &goauth2.Config{
            ClientID:     os.Getenv("DROPBOX_CLIENT_ID"),
            ClientSecret: os.Getenv("DROPBOX_CLIENT_SECRET"),
            RedirectURL:  "http://localhost:8081/oauth2callback",
        },
    ))

    // Tokens are injected to the handlers
    m.Get("/", func(tokens oauth2.Tokens, ctx *macaron.Context) {
        if !tokens.Expired() {
            wallswap.AuthUser(tokens)
        }

        ctx.Data["wallpapers"] = wallswap.GetWallpapers()
        ctx.Data["isGuest"] = tokens.Expired()
        ctx.HTML(200, "index")

    })

    m.Run(8081)
}
