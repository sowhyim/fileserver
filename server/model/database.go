package model

import "fmt"

type DatabaseConfig struct {
	Connections        []DatabaseConnection
	MaxIdleTime        int
	MaxLifetime        int
	MaxIdleConnections int
	MaxOpenConnections int
}

type DatabaseConnection struct {
	Sources  []DatabaseAccount
	Replicas []DatabaseAccount
}

type DatabaseAccount struct {
	Host     string
	Account  string
	Password string
	Database string
}

func (a DatabaseAccount) BuildUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=true",
		a.Account, a.Password, a.Host, a.Database)
}
