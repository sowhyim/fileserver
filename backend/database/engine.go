package database

import (
	"time"

	"fileserver/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var engine *gorm.DB

func Init(config *model.DatabaseConfig) {
	if config == nil || len(config.Connections) == 0 {
		panic("init database got an nil or has no connection in config!")
	}

	var resolverConfig = dbresolver.Config{}
	for i := range config.Connections {
		var sources, replicas []gorm.Dialector
		for j := range config.Connections[i].Sources {
			sources = append(sources, mysql.Open(config.Connections[i].Sources[j].BuildUrl()))
		}
		for j := range config.Connections[i].Replicas {
			replicas = append(replicas, mysql.Open(config.Connections[i].Replicas[j].BuildUrl()))
		}
		resolverConfig.Sources = sources
		resolverConfig.Replicas = replicas
		resolverConfig.Policy = dbresolver.RandomPolicy{}
	}

	var err error
	// choose first config for default url
	engine, err = gorm.Open(mysql.Open(config.Connections[0].Sources[0].BuildUrl()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var maxIdleTime = config.MaxIdleTime
	if maxIdleTime == 0 {
		maxIdleTime = 60 * 60
	}
	var maxLifetime = config.MaxLifetime
	if maxLifetime == 0 {
		maxIdleTime = 60 * 60 * 24
	}
	var maxIdleConnections = config.MaxIdleConnections
	if maxIdleConnections == 0 {
		maxIdleConnections = 100
	}
	var maxOpenConnections = config.MaxOpenConnections
	if maxOpenConnections == 0 {
		maxOpenConnections = 200
	}

	engine.Use(
		dbresolver.Register(resolverConfig).
			SetConnMaxIdleTime(time.Duration(maxIdleTime) * time.Second).
			SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second).
			SetMaxOpenConns(maxOpenConnections).
			SetMaxIdleConns(maxIdleConnections),
	)
}

func AutoMigrate() {
	engine.AutoMigrate(
		&model.Userinfo{},
		&model.UserPassword{},
		&model.UserGroup{},
		&model.FileInfo{},
		&model.FileGroup{},
		&model.FileGroupRelationship{},
		&model.FileSharing{},
	)
}
