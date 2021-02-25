package main

import (
    "flag"

    _ "github.com/jinzhu/gorm/dialects/mysql"
    "{{ .ProjectPath }}/config"
    "{{ .ProjectPath }}/database"
    "{{ .ProjectPath }}/server"
)

func main() {

    env := flag.String("e", "development", "")
    flag.Parse()

    config.Init(*env)
    database.Init(false)
    defer database.Close()
    if err := server.Init(); err != nil {
        panic(err)
    }
}