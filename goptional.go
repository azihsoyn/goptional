package goptional

import "reflect"

func GetOrElse(p, def interface{}) interface{} {
	rv := reflect.ValueOf(p)
	if rv.Type().Kind() != reflect.Ptr {
		return p
	}
	if rv.IsNil() {
		return def
	}
	return rv.Elem().Interface()
}
