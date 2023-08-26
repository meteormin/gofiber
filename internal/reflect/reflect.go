package reflect

import (
	"fmt"
	"reflect"
)

func HasField(data interface{}, field string) bool {
	metaValue := reflect.ValueOf(data).Elem()

	fieldValue := metaValue.FieldByName(field)
	if fieldValue == (reflect.Value{}) {
		return false
	}

	return true
}

func GetType(v interface{}) string {
	objType := reflect.TypeOf(v)
	return getType(objType)
}

func getType(p reflect.Type) (res string) {
	objType := p
	switch objType.Kind() {
	case reflect.Struct:
		if objType.PkgPath() != "" {
			res = fmt.Sprintf("%s.%s", objType.PkgPath(), objType.Name())
		} else {
			res = objType.Name()
		}

	case reflect.Ptr:
		if objType.PkgPath() != "" {
			res = fmt.Sprintf("%s.%s", objType.PkgPath(), objType.Elem().Name())
		} else {
			res = objType.Elem().Name()
		}
	case reflect.Func:
		numOut := objType.NumOut()
		if numOut > 0 {
			returnType := objType.Out(0)
			res = getType(returnType)
		}
	default:
		if objType.PkgPath() != "" {
			res = fmt.Sprintf("%s.%s", objType.PkgPath(), objType.Name())
		} else {
			res = objType.Name()
		}
	}

	return res
}
