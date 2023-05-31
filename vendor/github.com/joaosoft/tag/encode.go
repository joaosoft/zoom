package tag

import (
	"fmt"
	"net/url"
	"strings"
)

func (v Values) Encode() string {
	if v == nil {
		return ""
	}

	var buf strings.Builder
	for key, value := range v {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(url.QueryEscape(key))
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(fmt.Sprintf("%+v", value)))
	}

	return buf.String()
}
