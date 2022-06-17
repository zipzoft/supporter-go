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

	if IsEmpty(key) {
		return target
	}

	if IsEmpty(target) {
		return value
	}

	keys := strings.Split(key, ".")

	if index, err := strconv.ParseInt(keys[0], 10, 64); err == nil {
		if err == nil {
			if reflect.TypeOf(target).Kind() == reflect.Slice {
				s := reflect.ValueOf(target)

				if s.Len() > int(index) {
					return DataGet(s.Index(int(index)).Interface(), strings.Join(keys[1:], "."), value)
				}

				return value
			}

			return value
		}
	}

	switch target.(type) {

	case *map[string]interface{}:
		return DataGet(*target.(*map[string]interface{}), key, value)

	case map[string]interface{}:
		for _, segment := range keys {
			value = target.(map[string]interface{})[segment]

			if len(keys) > 1 {
				if IsEmpty(defaultValues) {
					return DataGet(value, strings.Join(keys[1:], "."))
				}

				return DataGet(value, strings.Join(keys[1:], "."), defaultValues...)
			}

			return value
		}
	}

	return value
}
