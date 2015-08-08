package utils

import (
	"net/url"
	"reflect"
	"strconv"
)

//ToURLValues converts struct to url.Values - useful for making POST form requests
func ToURLValues(inter interface{}) url.Values {
	values := make(url.Values)
	iVal := reflect.ValueOf(inter).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		case bool:
			if f.Bool() {
				v = "1"
			} else {
				v = "0"
			}
		}
		tag := typ.Field(i).Tag.Get("uval")
		if tag == "-" {
			continue
		}
		if tag == "" {
			tag = typ.Field(i).Name
		}
		values.Set(tag, v)
	}
	return values
}
