package queue

import (
	"bytes"
	"encoding/json"
)

// DumpString parse json
func DumpString(v interface{}) string {
	data, err := json.Marshal(v)
	buf := bytes.Buffer{}
	if err != nil {
		buf.WriteString("{err:\"json format error.")
		buf.WriteString(err.Error())
		buf.WriteString("\"}")
	} else {
		buf.Write(data)
	}
	return buf.String()
}
