package copy

import (
	"errors"
	"reflect"
)

// convert struct to map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// copy struct field to another
func SimpleCopy(src, dst interface{}) (err error) {
	if src == nil {
		return
	}
	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)

	// dst must a pointer
	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
		return errors.New("dst type should be a struct pointer")
	}

	// src must a pointer
	if srcType.Kind() == reflect.Ptr {
		srcType, srcValue = srcType.Elem(), srcValue.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		return errors.New("src type should be a struct or a struct pointer")
	}

	// Get content
	dstType, dstValue = dstType.Elem(), dstValue.Elem()

	// property number
	propertyNums := dstType.NumField()

	for i := 0; i < propertyNums; i++ {
		// property
		property := dstType.Field(i)
		// property filled
		propertyValue := srcValue.FieldByName(property.Name)

		// Invalid, src not have property || same name but different type
		if !propertyValue.IsValid() || property.Type != propertyValue.Type() {
			continue
		}

		if dstValue.Field(i).CanSet() {
			dstValue.Field(i).Set(propertyValue)
		}
	}
	return nil
}
