package goutil

import "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func JsonMarshal(data interface{}) string {
	b, err := json.Marshal(&data)
	ChkErr(err)
	return string(b)
}
