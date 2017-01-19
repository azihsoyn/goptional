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

func String(s string) *string {
	return &s
}

func Int(i int) *int {
	return &i
}

func Int64(i int64) *int64 {
	return &i
}

func Bool(b bool) *bool {
	return &b
}
