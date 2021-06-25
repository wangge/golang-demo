package main

import (
	"fmt"
	"reflect"
)

type DemoA struct {
	Name string
	Gender int
	Age int
	Rows []string
}

type DemoB struct {
	Name string
	Gender string
	Address string
	Rows []int
}


// 将同时存在于dest,from的字段，从from赋值到dest.
func structAssign(dest interface{}, from interface{})  {
	dVal := reflect.ValueOf(dest).Elem()
	fVal := reflect.ValueOf(from).Elem()
	fTypeOft := fVal.Type()
	for i:= 0;i< fVal.NumField();i++{
		name := fTypeOft.Field(i).Name
		// 检查目标struct是否存在name
		if reflect.DeepEqual(dVal.FieldByName(name), reflect.Value{}){
			continue
		}
		dFieldType := dVal.FieldByName(name).Type().Name()
		fFieldType := fVal.FieldByName(name).Type().Name()
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		if ok:= dVal.FieldByName(name).IsValid(); ok && dFieldType == fFieldType{
			dVal.FieldByName(name).Set(reflect.ValueOf(fVal.Field(i).Interface()))
		}
	}
}

func CopyStruct(src, dst interface{}) {
	dstVal := reflect.ValueOf(dst).Elem() //获取reflect.Type类型
	srcVal := reflect.ValueOf(src).Elem()   //获取reflect.Type类型
	vTypeOfT := srcVal.Type()
	for i := 0; i < srcVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		// 检查目标struct是否存在name
		if reflect.DeepEqual(dstVal.FieldByName(name), reflect.Value{}){
			continue
		}
		dFieldType := dstVal.FieldByName(name).Type()
		//dFieldTypeName := dFieldType.Name()
		sFieldType := srcVal.FieldByName(name).Type()
		//sFieldTypeName := sFieldType.Name()
		// 名字及类型一致，才能赋值
		fmt.Println(name,sFieldType.AssignableTo(dFieldType))
		if ok := dstVal.FieldByName(name).IsValid(); ok && sFieldType.AssignableTo(dFieldType) {
			dstVal.FieldByName(name).Set(reflect.ValueOf(srcVal.Field(i).Interface()))
		}
	}
}
func main()  {
	as := DemoA{Name: "a",Rows: []string{"1","2"}}
	bs := DemoB{Name: "b",Gender: "男", Address: "sssss"}
	fmt.Println(bs)
	CopyStruct(&as, &bs)
	fmt.Println(bs)

	fmt.Println("xxx", as.toString())
}