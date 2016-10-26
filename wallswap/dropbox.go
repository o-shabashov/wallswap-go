package wallswap

import (
    "gopkg.in/macaron.v1"
    goauth2 "golang.org/x/oauth2"
    "github.com/go-macaron/oauth2"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// DropBox API `get_current_account` json answer will be converted to this struct
type dropBoxUser struct {
    AccountID     string `json:"account_id"`
    Name          struct {
                      GivenName       string `json:"given_name"`
                      Surname         string `json:"surname"`
                      FamiliarName    string `json:"familiar_name"`
                      DisplayName     string `json:"display_name"`
                      AbbreviatedName string `json:"abbreviated_name"`
                  } `json:"name"`
    Email         string `json:"email"`
    EmailVerified bool `json:"email_verified"`
    Disabled      bool `json:"disabled"`
    Country       string `json:"country"`
    Locale        string `json:"locale"`
    ReferralLink  string `json:"referral_link"`
    IsPaired      bool `json:"is_paired"`
    AccountType   struct {
                      Tag string `json:".tag"`
                  } `json:"account_type"`
    AccessToken   string
}

var dropBoxEndpoints = goauth2.Endpoint{
    AuthURL:  "https://www.dropbox.com/1/oauth2/authorize",
    TokenURL: "https://api.dropbox.com/1/oauth2/token",
}

func OAuthProvider(conf *goauth2.Config) macaron.Handler {
    conf.Endpoint = dropBoxEndpoints
    return oauth2.NewOAuth2Provider(conf)
}

func AuthUser(tokens oauth2.Tokens) User {
    var dBUser dropBoxUser

    // Request to Dropbox API
    req, err := http.NewRequest("POST", "https://api.dropboxapi.com/2/users/get_current_account", nil)
    if err != nil {
        panic(err)
    }
    req.Header.Set("Authorization", "Bearer " + tokens.Access())
    resp, err := http.DefaultClient.Do(req)
    CheckErr(err)

    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    // Convert json answer to struct
    jsonRaw := []byte(body);
    if err := json.Unmarshal(jsonRaw, &dBUser); err != nil {
        panic(err)
    }

    dBUser.AccessToken = tokens.Access()

    return dbSearch(dBUser)
}

func dbSearch(dBUser dropBoxUser) User {
    var user User

    db := GetDBConnection()
    defer db.Close()

    // Search user in DB
    err := db.QueryRow(
        "SELECT dropbox_id,access_token from user WHERE dropbox_id=?",
        dBUser.AccountID).Scan(&user.DropboxId, &user.AccessToken)
    switch {
    case err == sql.ErrNoRows:
        stmt, _ := db.Prepare("INSERT user SET email=?,access_token=?,auth_key=?,dropbox_id=?")
        stmt.Exec(dBUser.Email, dBUser.AccessToken, RandString(32), dBUser.AccountID)
        user.DropboxId = dBUser.AccountID
        user.AccessToken = dBUser.AccessToken
    }
    return user
}
