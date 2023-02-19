package admin

import (
	"github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/entity"
	"github.com/miniyus/gofiber/internal/hash"
	"github.com/miniyus/gofiber/permission"
	"gorm.io/gorm"
	"log"
	"time"
)

func existsAdmin(db *gorm.DB) bool {
	admin := &entity.User{}
	if err := db.Where(entity.User{Role: entity.Admin}).Find(admin).Error; err != nil {
		return false
	}

	return true
}

func CreateAdmin(db *gorm.DB, configs *config.Configs) {
	if existsAdmin(db) {
		log.Println("Skip create admin: already exists admin account")
		return
	}

	permCollection := permission.NewPermissionCollection(permission.NewPermissionsFromConfig(configs.Permission)...)

	caCfg := configs.CreateAdmin

	username := caCfg.Username
	email := caCfg.Email
	password := caCfg.Password

	if !caCfg.IsActive || username == "" || password == "" || email == "" {
		log.Println("Skip create admin: account info is empty")
		return
	}

	hashedPassword, err := hash.Bcrypt.HashPassword(password)
	if err != nil {
		log.Println(err)
		return
	}

	now := time.Now()

	user := &entity.User{
		Username:        username,
		Password:        hashedPassword,
		Email:           email,
		Role:            entity.Admin,
		EmailVerifiedAt: &now,
	}

	permissions := permCollection

	entPerms := make([]entity.Permission, 0)

	get, err := permissions.Get(string(entity.Admin))
	if err != nil {
		return
	}

	entPerms = append(entPerms, get.ToEntity())

	group := &entity.Group{
		Name:        "Admin",
		Permissions: entPerms,
		Users:       []entity.User{*user},
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		return tx.Debug().Create(group).Error
	})

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Success create admin")
	return
}
