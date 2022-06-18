package supporter

import (
	"reflect"
	"strconv"
	"strings"
)

// DataGet Get an item from an array or object using "dot" notation.
// Example:
// 		// Output: "a"
// 		fmt.Println(supporter.DataGet(map[string]interface{}{"a": "b"}, "a"))
//
// 		// Or get a nested value
// 		person := map[string]interface{}{
// 			"name": map[string]interface{}{
// 				"first": "John",
// 				"last":  "Doe",
// 			},
// 			"age": "23",
//			"phones" : []map[string]interface{}{
//				{"type": "home", "number": "555-555-5555"},
//				{"type": "work", "number": "555-666-6666"},
//			},
// 		}
//
// 		// Output: "John"
// 		fmt.Println(supporter.DataGet(person, "name.first"))
//
// 		// Output: "Doe"
// 		fmt.Println(supporter.DataGet(person, "name.last"))
//
// 		// Output: "23"
// 		fmt.Println(supporter.DataGet(person, "age"))
//
// 		// Output: "555-555-5555"
// 		fmt.Println(supporter.DataGet(person, "phones.0.number"))
func DataGet(target interface{}, key string, defaultValues ...interface{}) (value interface{}) {

	value = nil
	if len(defaultValues) > 0 {
		value = defaultValues[0]
	}

	if IsEmpty(key) || IsEmpty(target) {
		return target
	}

	keys := strings.Split(key, ".")

	if index, err := strconv.ParseInt(keys[0], 10, 64); err == nil {
		if err == nil {
			if reflect.TypeOf(target).Kind() == reflect.Slice {
				s := reflect.ValueOf(target)

				if s.Len() > int(index) {
					return DataGet(s.Index(int(index)).Interface(), strings.Join(keys[1:], "."), value)
				}
			}

			return value
		}
	}

	if reflect.TypeOf(target).Kind() == reflect.Map {
		m := reflect.ValueOf(target)

		if m.MapIndex(reflect.ValueOf(keys[0])).IsValid() {
			return DataGet(m.MapIndex(reflect.ValueOf(keys[0])).Interface(), strings.Join(keys[1:], "."), value)
		}
	}

	return value
}
