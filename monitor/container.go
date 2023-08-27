package monitor

import (
	"github.com/miniyus/gofiber/app"
	myReflect "github.com/miniyus/gofiber/internal/reflect"
	"reflect"
)

type bindType string

const (
	bind      bindType = "bindingInterface"
	singleton bindType = "singleton"
)

type containerInfo struct {
	Key          string
	BindType     bindType
	InstanceType string
}

func newContainerInfo(a app.Application) []containerInfo {
	instances := a.Instances()
	containerInfos := make([]containerInfo, 0)

	for _, inst := range instances {
		var bt bindType
		reflectType := reflect.TypeOf(inst)
		if reflectType.Kind() == reflect.Func {
			bt = bind
		} else {
			bt = singleton
		}

		instType := myReflect.GetType(inst)
		ci := containerInfo{
			Key:          instType,
			BindType:     bt,
			InstanceType: instType,
		}
		containerInfos = append(containerInfos, ci)
	}

	return containerInfos
}
