package ibase

import (
	"fmt"
	"reflect"
)

func StructToStruct(obj interface{}, obj2 interface{}) (interface{}, error) {
	vt := reflect.TypeOf(obj).Elem()
	vv := reflect.ValueOf(obj).Elem()
	vt2 := reflect.TypeOf(obj2).Elem()
	vv2 := reflect.ValueOf(obj2).Elem()
	for i := 0; i < vt.NumField(); i++ {
		f := vt.Field(i)
		//f.Tag.Get("bson")
		for i := 0; i < vt2.NumField(); i++ {
			f2 := vt2.Field(i)
			if (f.Name == f2.Name || (f.Name == "Order_no" && f2.Name == "OrderNo") || (f.Name == "OrderNo" && f2.Name == "Order_no")) && f.Type == f2.Type {
				structFieldValue := vv2.FieldByName(f2.Name)
				if !structFieldValue.CanSet() {
					fmt.Errorf("struct 字段 %s 不能set,请检查!!!", f2.Name)
				} else {
					structFieldValue.Set(vv.FieldByName(f.Name))
				}
			}
		}
	}

	return obj2, nil
}
