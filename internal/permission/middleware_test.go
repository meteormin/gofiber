package permission_test

import (
	"github.com/miniyus/gofiber/internal/permission"
	"github.com/miniyus/gofiber/utils"
	"strings"
	"testing"
)

var hasPerm = []permission.Permission{
	{
		Actions: []permission.Action{
			{
				Methods: []permission.Method{
					"GET",
				},
				Resource: "/test",
			},
		},
		GroupId: 1,
		Name:    "TEST",
	},
}

func TestCheckPermission(t *testing.T) {
	pass := false
	method := "GET"
	utils.NewCollection(hasPerm).For(func(perm permission.Permission, i int) {
		utils.NewCollection(perm.Actions).For(func(action permission.Action, j int) {
			routePath := "/test"
			if strings.Contains(routePath, action.Resource) {

				if method == "OPTION" {
					method = "GET"
				}

				filtered := utils.NewCollection(action.Methods).Filter(func(v permission.Method, i int) bool {
					return string(v) == method
				})

				if len(filtered) != 0 {
					pass = true
				}
			}
		})
	})

	if !pass {
		t.Error(pass)
	}
}
