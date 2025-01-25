package postgresql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/internal/configs"
	"main/internal/models"
)

type Connection struct {
	Db *gorm.DB
}

func (c *Connection) Connect(config *configs.Config) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
			config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort),
	}), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to database")
	}
	c.Db = db
}

func (c *Connection) InitModels() {
	err := c.Db.AutoMigrate(models.Role{}, models.User{})
	if err != nil {
		panic("Couldn't initialize database models")
	}
}

func NewDatabaseConnection(config *configs.Config) *Connection {
	connection := &Connection{}
	connection.Connect(config)
	return connection
}
