package internal

import (
	"net/url"

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
