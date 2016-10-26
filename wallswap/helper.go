package wallswap

import (
    "math/rand"
    "database/sql"
)

/* @see http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang */
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}

func RandString(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
    }
    return string(b)
}

func GetDBConnection() (*sql.DB) {
    db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/wallswap")
    CheckErr(err)
    return db
}
