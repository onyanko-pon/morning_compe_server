package DataBase

import (
  "os"
)

type Config struct {
    Host string
    Username string
    Password string
    DBName string
    Port string
}

func newConfig() *Config {
    c := new(Config)

    c.Host = os.Getenv("DB_HOST")
    c.Username = os.Getenv("DB_USERNAME")
    c.Password = os.Getenv("DB_PASSWORD")
    c.DBName = os.Getenv("DB_DBNAME")
    c.Port =  os.Getenv("DB_PORT")

    return c
}