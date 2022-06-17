package supporter

import "reflect"

func InArray(val interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				return true
			}
		}
	}

	return false
}

func First(val interface{}) interface{} {
	if IsEmpty(val) {
		return nil
	}

	switch reflect.TypeOf(val).Kind() {
	case reflect.Array:
	case reflect.Slice:
		reflectVal := reflect.ValueOf(val)
		for i := 0; i < reflectVal.Len(); {
			return reflectVal.Index(i).Interface()
		}
		return nil

	case reflect.Map:
		reflectVal := reflect.ValueOf(val)
		for _, value := range reflectVal.MapKeys() {
			key := value.Interface().(string)
			return val.(map[string]interface{})[key]
		}
	}

	return nil
}

func ToSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}
