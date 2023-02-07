package permission

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miniyus/gofiber/auth"
	"github.com/miniyus/gofiber/database"
	"github.com/miniyus/gofiber/entity"
	"github.com/miniyus/gofiber/utils"
	"gorm.io/gorm"
	"strings"
)

type HasPermissionParameter struct {
	DB           *gorm.DB
	DefaultPerms Collection
	FilterFunc   func(ctx *fiber.Ctx, groupId uint, p Permission) bool
}

// HasPermission
// check has permissions middleware
func HasPermission(parameter HasPermissionParameter, permissions ...Permission) fiber.Handler {
	return func(c *fiber.Ctx) error {
		pass := false

		var permCollection Collection

		db := parameter.DB
		if db == nil {
			db = database.GetDB()
		}

		authUser, err := auth.GetAuthUser(c)

		repo := NewRepository(db)

		get, err := repo.Get(*authUser.GroupId)
		if err == nil {
			permCollection = NewPermissionCollection()
			utils.NewCollection(get).For(func(v entity.Permission, i int) {
				permCollection.Add(EntityToPermission(v))
			})
		}

		if permCollection == nil {
			permCollection = parameter.DefaultPerms
		}

		permCollection.For(func(perm Permission, i int) {
			_, err = repo.Save(ToPermissionEntity(perm))
		})

		if err != nil {
			return err
		}

		if len(permissions) != 0 {
			permCollection.Concat(permissions)
		}

		userHasPerm := permCollection.Filter(func(p Permission, i int) bool {
			if parameter.FilterFunc != nil {
				return parameter.FilterFunc(c, *authUser.GroupId, p)
			}

			if *authUser.GroupId != 0 {
				return *authUser.GroupId == p.GroupId
			}

			return false
		})

		pass = checkPermissionFromCtx(userHasPerm.Items(), c)

		if pass {
			return c.Next()
		}

		return fiber.ErrForbidden
	}
}

func checkPermissionFromCtx(hasPerm []Permission, c *fiber.Ctx) bool {
	if len(hasPerm) == 0 {
		return false
	}

	pass := false
	utils.NewCollection(hasPerm).For(func(perm Permission, i int) {
		utils.NewCollection(perm.Actions).For(func(action Action, j int) {
			routePath := c.Path()
			if strings.Contains(routePath, action.Resource) {
				method := c.Method()
				if method == "OPTION" {
					method = "GET"
				}

				filtered := utils.NewCollection(action.Methods).Filter(func(v Method, i int) bool {
					return string(v) == method
				})

				if filtered.Count() != 0 {
					pass = true
				}
			}
		})
	})

	return pass
}
