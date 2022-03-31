package OrderedValues

import (
	"fmt"
	"net/url"
)

func (v Values) Add(key string, value string) {
	key = url.QueryEscape(key)
	value = url.QueryEscape(value)

	if len(v.String) == 0 {
		v.String += fmt.Sprintf("%s=%s", key, value)
		return
	}

	v.String += fmt.Sprintf("&%s=%s", key, value)
}
