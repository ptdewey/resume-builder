package internal

import (
	"net/url"
	"reflect"

	lua "github.com/yuin/gopher-lua"
)

func luaTableToStringSlice(tbl *lua.LTable) []string {
	var result []string
	tbl.ForEach(func(_, value lua.LValue) {
		result = append(result, value.String())
	})
	return result
}

func stringToURL(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		// TODO: determine what should happen in error case
		return ""
	}

	if u.Scheme == "" {
		// TODO: maybe add separate handler for email?
		s = "https://" + s
	}

	return s
}

func contains[T any](s []T, v T, field string) bool {
	for _, t := range s {
		vVal := reflect.ValueOf(v)
		tVal := reflect.ValueOf(t)

		vField := vVal.FieldByName(field)
		tField := tVal.FieldByName(field)

		if vField.IsValid() && tField.IsValid() && vField.Interface() == tField.Interface() {
			return true
		}
	}
	return false
}
