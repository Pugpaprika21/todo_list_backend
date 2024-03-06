package migrtion

import (
	"github.com/Pugpaprika21/go-fiber/db"
	"github.com/Pugpaprika21/go-fiber/models"
)

func Run() {
	db.Conn.AutoMigrate(
		&models.User{},
		&models.MasterRole{},
		&models.MasterSettingRole{},
		&models.FileStorageSystem{},
		&models.Todo{},
	)
}
