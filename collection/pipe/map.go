package pipe

import (
	"errors"
	"reflect"
)

var _ Pipe = (*mapPipe)(nil)

type mapPipe struct {
	_handle func(item interface{}, key interface{}) interface{}
}

// Handle implements Pipe
func (pipe *mapPipe) Handle(data interface{}) (interface{}, error) {
	if data == nil {
		return nil, errors.New("data is nil")
	}

	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		var newItems interface{}

		// Foreach item in slice
		for i := 0; i < reflect.ValueOf(data).Len(); i++ {
			// Get item
			item := reflect.ValueOf(data).Index(i).Interface()

			newItem := pipe._handle(item, i)

			// Append item to new slice
			if newItems == nil {
				newItems = make([]interface{}, 0)
			}

			newItems = reflect.Append(reflect.ValueOf(newItems), reflect.ValueOf(newItem)).Interface()
		}

		return newItems, nil

	case reflect.Map:
		var newItems interface{}

		// Foreach item in map
		for _, key := range reflect.ValueOf(data).MapKeys() {
			// Get item
			item := reflect.ValueOf(data).MapIndex(key).Interface()

			newItem := pipe._handle(item, key.Interface())
			// Append item to new slice
			if newItems == nil {
				newItems = make(map[string]interface{})
			}

			newItems = reflect.Append(reflect.ValueOf(newItems), reflect.ValueOf(newItem)).Interface()
		}

		return newItems, nil

	default:
		return nil, errors.New("data is not a slice or map")
	}
}

func Map(handler func(item interface{}, key interface{}) interface{}) Pipe {
	return &mapPipe{
		_handle: handler,
	}
}
