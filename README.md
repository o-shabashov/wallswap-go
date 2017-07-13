# Wallswap-Go

[![BCH compliance](https://bettercodehub.com/edge/badge/o-shabashov/wallswap-go?branch=master)](https://bettercodehub.com/)

The project is just for fun and test programming skills. It consists of two parts - the server and crawler.

### Server:
1. Runs on 8080 port;
2. Gets the list of wallpapers from database;
3. Render the html page with the list of wallpaper;
4. Allows you to authenticate the user through OAuth2 Dropbox, gets access_token and stores the user in a database;

### Crawler:
1. Parses website [https://wallhaven.cc](https://wallhaven.cc) and collects direct URL's at the wallpaper;
2. Saves list of wallpapers in the database;
3. Upload wallpapers for each user in Dropbox directory.

## Installation
* Create MySQL database `wallswap` and import `wallswap.sql`
* Create [Dropbox App](https://www.dropbox.com/developers/apps/create) and fill `env.env`
```env
export DROPBOX_CLIENT_ID='APP_KEY_HERE'
export DROPBOX_CLIENT_SECRET='APP_SECRET_HERE'
```
* Redirect URL for Dropbox callback:
```
http://localhost:8080/oauth2callback
```
* Install [glide](https://glide.sh/) dependencies:
```bash
cd wallswap-go
glide update
```
* Run server:
```bash
go run http.go
```
* Run crawl once a week:
```bash
go run crawl.go
```

## Made with
1. [Macaron](https://go-macaron.com)
2. MySQL
3. [golang.org/x/oauth2](https://godoc.org/golang.org/x/oauth2)
