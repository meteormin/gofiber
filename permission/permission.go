package permission

import (
	"fmt"
	"github.com/miniyus/gofiber/entity"
	cLog "github.com/miniyus/gofiber/log"
	"github.com/miniyus/gofiber/utils"
	"gorm.io/gorm"
	"log"
)

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	PATCH  Method = "PATCH"
	DELETE Method = "DELETE"
)

type Action struct {
	Methods  []Method
	Resource string
}

func NewAction(resource string, method []Method) Action {
	return Action{method, resource}
}

type Permission struct {
	GroupId uint
	Name    string
	Actions []Action
}

func NewPermission(groupId uint, name string, actions []Action) Permission {
	return Permission{
		groupId, name, actions,
	}
}

type user struct {
	Id      *uint `json:"id"`
	GroupId *uint `json:"group_id"`
	Role    *uint `json:"role"`
}

type includeUser struct {
	UserId  *uint `json:"user_id"`
	GroupId *uint `json:"group_id"`
	User    *user `json:"user"`
}

type Collection interface {
	utils.Collection[Permission]
	All() []Permission
	RemoveByName(name string) bool
	Get(name string) (*Permission, error)
}

type CollectionStruct struct {
	*utils.BaseCollection[Permission]
	items []Permission
}

func NewPermissionCollection(perms ...Permission) Collection {
	defaultPerms := make([]Permission, 0)
	if len(perms) == 0 {
		perms = defaultPerms
	}

	base := utils.NewCollection(perms).(*utils.BaseCollection[Permission])

	return &CollectionStruct{BaseCollection: base}
}

func (p *CollectionStruct) All() []Permission {
	return p.items
}

func (p *CollectionStruct) RemoveByName(name string) bool {
	filtered := p.Filter(func(v Permission, i int) bool {
		return v.Name == name
	})

	if len(filtered.Items()) == 0 {
		return false
	}

	var rmIndex int
	for i, perm := range p.items {
		if perm.Name == filtered.Items()[0].Name {
			rmIndex = i
		}
	}

	p.Remove(rmIndex)

	return true
}

func (p *CollectionStruct) Get(name string) (*Permission, error) {
	filtered := p.Filter(func(v Permission, i int) bool {
		return v.Name == name
	})

	if filtered.Count() == 0 {
		return nil, fmt.Errorf("can't found %s Permission", name)
	}

	return &filtered.Items()[0], nil
}

func ToPermissionEntity(perm Permission) entity.Permission {
	var ent entity.Permission
	ent.Permission = perm.Name
	ent.GroupId = perm.GroupId
	for _, action := range perm.Actions {
		for _, method := range action.Methods {
			ent.Actions = append(ent.Actions, entity.Action{
				Resource: action.Resource,
				Method:   string(method),
			})
		}
	}

	return ent
}

func EntityToPermission(permission entity.Permission) Permission {
	actions := make([]Action, 0)
	utils.NewCollection(permission.Actions).For(func(v entity.Action, i int) {
		filtered := utils.NewCollection(permission.Actions).Filter(func(a entity.Action, j int) bool {
			return a.PermissionId == v.PermissionId && a.Resource == v.Resource
		})

		methods := make([]Method, 0)
		filtered.For(func(f entity.Action, k int) {
			methods = append(methods, Method(f.Method))
		})

		actions = append(actions, Action{
			Resource: v.Resource,
			Methods:  methods,
		})
	})

	return Permission{
		GroupId: permission.GroupId,
		Name:    permission.Permission,
		Actions: actions,
	}
}

func CreateDefaultPermissions(db *gorm.DB, cfgs []Config) {
	perms := NewPermissionsFromConfig(cfgs)
	permCollection := NewPermissionCollection(perms...)

	repo := NewRepository(db)
	var entities []entity.Permission

	permCollection.For(func(perm Permission, i int) {
		entities = append(entities, ToPermissionEntity(perm))
	})

	all, err := repo.All()
	if err != nil {
		cLog.GetLogger().Error(err)
		log.Print(err)
	}

	if len(all) != 0 {
		return
	}

	_, err = repo.BatchCreate(entities)
	if err != nil {
		if err != nil {
			cLog.GetLogger().Error(err)
			log.Print(err)
		}
	}

}
