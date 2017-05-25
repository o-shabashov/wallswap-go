package main

import (
    "fmt"
    "golang.org/x/net/html"
    "net/http"
    _ "github.com/go-sql-driver/mysql"
    "wallswap-go/wallswap"
)

// Should be run once a week
func main() {
    wallpapers := storeWallpaperURLs()

    var accessToken string

    db := wallswap.GetDBConnection()
    defer db.Close()

    rows, err := db.Query("select access_token from user")
    wallswap.CheckErr(err)

    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&accessToken)
        wallswap.CheckErr(err)
        wallswap.DeleteFiles(accessToken)
        wallswap.UploadFiles(accessToken, wallpapers)
    }
    err = rows.Err()
    wallswap.CheckErr(err)
}

// Extract all figure.id's from a given webpage
func crawlWallpapers(url string, ch chan string, chFinished chan bool) {
    resp, err := http.Get(url)

    defer func() {
        // Notify that we're done after this function
        chFinished <- true
    }()

    if err != nil {
        fmt.Println("ERROR: Failed to crawl " + url)
        return
    }

    body := resp.Body
    defer body.Close() // close Body when the function returns

    z := html.NewTokenizer(body)

    for {
        tt := z.Next()

        switch {
        case tt == html.ErrorToken:
            // End of the document, we're done
            return
        case tt == html.StartTagToken:
            t := z.Token()

            // Check if the token is an <figure> tag
            isImg := t.Data == "figure"
            if !isImg {
                continue
            }

            // Extract the id value, if there is one
            id, ok := wallswap.GetId(t)
            if !ok {
                continue
            }
            ch <- id
        }
    }
}

// Parse wallheaven for wallpaper urls and store to MySQL
func storeWallpaperURLs() map[string]string {
    url := "https://alpha.wallhaven.cc/search?categories=101&purity=100&sorting=random&order=desc"

    wallpapers := make(map[string]string)

    // Database init
    db := wallswap.GetDBConnection()
    defer db.Close()

    // Truncate and prepare table
    db.Query("truncate wallpaper")
    stmt, err := db.Prepare("INSERT wallpaper SET thumb_url=?,url=?")
    wallswap.CheckErr(err)

    // Channels
    chIds := make(chan string)
    chFinished := make(chan bool)

    // Kick off the crawl process (concurrently)
    go crawlWallpapers(url, chIds, chFinished)

    // Subscribe to both channels
    isFinished := false
    for isFinished == false {
        select {
        case id := <-chIds:
            thumbUrl := "https://alpha.wallhaven.cc/wallpapers/thumb/small/th-" + id + ".jpg"
            fullUrl := "https://wallpapers.wallhaven.cc/wallpapers/full/wallhaven-" + id + ".jpg"

            stmt.Exec(thumbUrl, fullUrl)

            wallpapers[thumbUrl] = fullUrl
            wallswap.CheckErr(err)

        case <-chFinished:
            isFinished = true
        }
    }
    close(chIds)

    return wallpapers
}
