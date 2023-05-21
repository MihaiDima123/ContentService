package ds

import (
	"gorm.io/gorm"
)

type Datasource interface {
	GetConnection() *gorm.DB
	Initialize(configuration Configuration) Datasource
}

type Configuration struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}
