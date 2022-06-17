package supporter

import (
	"reflect"
	"strconv"
	"strings"
)

// DataGet Get an item from an array or object using "dot" notation.
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
