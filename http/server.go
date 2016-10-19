package main

import (
    "net/http"
    //"html/template"
    "gopkg.in/gin-gonic/gin.v1"
)

func main() {
    router := gin.Default()

    // Static files
    router.Static("/assets", "http/assets")

    router.LoadHTMLGlob("http/templates/*")

    var wallpapers map[string]string

    router.GET("/", func(c *gin.Context) {
        wallpapers = make(map[string]string)
        wallpapers["http://1"] = "http://2"

        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "Wallswap Golang",
            "wallpapers": wallpapers,
        })
    })

    router.Run(":8081")
}
