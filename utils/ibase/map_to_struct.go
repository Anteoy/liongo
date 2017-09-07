package ibase

import (
	"reflect"
	"fmt"
)

func MapToStruct(obj interface{}, objMap map[string]interface{}) error {
	for name, value := range objMap {
		structValue := reflect.ValueOf(obj).Elem()
		structFieldValue := structValue.FieldByName(name)

		if !structFieldValue.CanSet() {
			return fmt.Errorf("struct 字段 %s 不能set,请检查!!!", name)
		}

		if !structFieldValue.IsValid() {
			return fmt.Errorf("struct中找不到 %s 的对应字段列,请检查!!!", name)
		}

		structFieldType := structFieldValue.Type()
		val := reflect.ValueOf(value)
		if structFieldType != val.Type() {
			return fmt.Errorf("map struct 同名匹配列的类型不匹配,请检查!!! struct: %s,map value: %s\n", structFieldType, val.Type())
		}

		structFieldValue.Set(val)
	}

	return nil
}