package sacloud

import (
	"bytes"
	"text/template"
)

func buildURL(pathFormat string, param interface{}) (string, error) {
	buf := bytes.NewBufferString("")
	t := template.New("buildURL")
	template.Must(t.Parse(pathFormat))
	err := t.Execute(buf, param)
	return buf.String(), err
}
