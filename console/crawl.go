package main

import (
    "fmt"
    "golang.org/x/net/html"
    "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// Helper function to pull the id attribute from a Token
func getId(t html.Token) (ok bool, id string) {
    // Iterate over all of the Token's attributes until we find an "id"
    for _, attr := range t.Attr {
        if attr.Key == "data-wallpaper-id" {
            id = attr.Val
            ok = true
        }
    }
    return
}

// Extract all http** links from a given webpage
func crawl(url string, ch chan string, chFinished chan bool) {
    resp, err := http.Get(url)

    defer func() {
        // Notify that we're done after this function
        chFinished <- true
    }()

    if err != nil {
        fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
        return
    }

    b := resp.Body
    defer b.Close() // close Body when the function returns

    z := html.NewTokenizer(b)

    for {
        tt := z.Next()

        switch {
        case tt == html.ErrorToken:
            // End of the document, we're done
            return
        case tt == html.StartTagToken:
            t := z.Token()

            // Check if the token is an <img> tag
            isImg := t.Data == "figure"
            if !isImg {
                continue
            }

            // Extract the href value, if there is one
            ok, id := getId(t)
            if !ok {
                continue
            }
            ch <- id
        }
    }
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    url := "https://alpha.wallhaven.cc/search?categories=101&purity=110&sorting=random&order=desc"

    // Database init
    db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/wallswap")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // Truncate and prepare table
    db.Query("truncate wallpaper")
    stmt, err := db.Prepare("INSERT wallpaper SET thumb_url=?,url=?")
    checkErr(err)

    // Channels
    chIds := make(chan string)
    chFinished := make(chan bool)

    // Kick off the crawl process (concurrently)
    go crawl(url, chIds, chFinished)

    // Subscribe to both channels
    isFinished := false
    for isFinished == false {
        select {
        case id := <-chIds:
            fmt.Println(" - " + id)
            stmt.Exec(
                "https://alpha.wallhaven.cc/wallpapers/thumb/small/th-" + id + ".jpg",
                "https://wallpapers.wallhaven.cc/wallpapers/full/wallhaven-" + id + ".jpg")
            checkErr(err)

        case <-chFinished:
            isFinished = true
        }
    }
    close(chIds)
}
