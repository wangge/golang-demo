package main

import (
	"fmt"
	"reflect"
)

type testT struct {
	A string `db:"a"`
	B int `db:"b"`
}
func structToMap(obj interface{}, tag string) map[string]interface{}  {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		key := obj1.Field(i).Name
		if tag != ""{
			key = obj1.Field(i).Tag.Get(tag)
		}
		data[key] = obj2.Field(i).Interface()
	}
	return data
}
func main()  {
	t := testT{
		A:"ssssss",
	}
	fmt.Println(structToMap(t, ""))
}