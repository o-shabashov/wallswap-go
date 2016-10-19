package main

import (
    "net/http"
    //"html/template"
    "gopkg.in/gin-gonic/gin.v1"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("http/templates/*")

    router.Static("/assets", "http/assets")
    router.GET("/", index)

    router.Run(":8081")
}

func getWallpapers() (wallpapers map[string]string) {
    wallpapers = make(map[string]string)

    var (
        thumbUrl string
        url string
    )

    // Database init
    db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/wallswap")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    rows, err := db.Query("select thumb_url, url from wallpaper")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&thumbUrl, &url)
        if err != nil {
            log.Fatal(err)
        }
        wallpapers[thumbUrl] = url
    }
    err = rows.Err()

    if err != nil {
        log.Fatal(err)
    }

    return
}

func index(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
        "title": "Wallswap Golang",
        "wallpapers": getWallpapers(),
    })
}
