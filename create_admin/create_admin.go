package create_admin

import (
	"errors"
	"github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/database"
	"github.com/miniyus/gofiber/entity"
	"github.com/miniyus/gofiber/permission"
	"github.com/miniyus/gofiber/utils"
	"gorm.io/gorm"
	"log"
	"time"
)

func existsAdmin(db *gorm.DB) bool {
	admin := &entity.User{}
	rs := db.Where(entity.User{Role: entity.Admin}).Find(admin)
	rs, err := database.HandleResult(rs)
	if rs.RowsAffected == 0 {
		return false
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
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

	hashedPassword, err := utils.HashPassword(password)
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

	entPerms = append(entPerms, permission.ToPermissionEntity(*get))

	group := &entity.Group{
		Name:        "Admin",
		Permissions: entPerms,
		Users:       []entity.User{*user},
	}

	rs := db.Debug().Create(group)
	_, err = database.HandleResult(rs)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Success create admin")
	return
}
