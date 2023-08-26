package reflect

import "reflect"

func HasField(data interface{}, field string) bool {
	metaValue := reflect.ValueOf(data).Elem()

	fieldValue := metaValue.FieldByName(field)
	if fieldValue == (reflect.Value{}) {
		return false
	}

	return true
}

func GetType(v interface{}) (res string) {
	t := reflect.TypeOf(v)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
		res += "*"
	}
	return res + t.Name()
}
