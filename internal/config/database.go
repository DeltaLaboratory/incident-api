package config

import (
	"fmt"
	"net/url"
)

type Database struct {
	Host     string  `hcl:"host"`
	Port     int     `hcl:"port"`
	Database *string `hcl:"database"`
	User     *string `hcl:"user"`
	Password *string `hcl:"password"`
}

func (db Database) URI() string {
	connect := url.URL{
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%d", db.Host, db.Port),
	}

	if db.User != nil {
		connect.User = url.UserPassword(*db.User, *db.Password)
	}

	if db.Database != nil {
		connect.Path = *db.Database
	}

	return connect.String()
}
