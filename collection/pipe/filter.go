package pipe

import (
	"errors"
	"reflect"
)

var _ Pipe = (*filterPipe)(nil)

type filterPipe struct {
	_handler func(item interface{}, key interface{}) bool
}

// Handle implements Pipe
func (pipe *filterPipe) Handle(data interface{}) (interface{}, error) {
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

			// Check if item is valid
			if !pipe._handler(item, i) {
				continue
			}

			// Append item to new slice
			if newItems == nil {
				newItems = make([]interface{}, 0)
			}

			newItems = reflect.Append(reflect.ValueOf(newItems), reflect.ValueOf(item)).Interface()
		}

		return newItems, nil

	case reflect.Map:
		var newItems interface{}

		// Foreach item in map
		for _, key := range reflect.ValueOf(data).MapKeys() {
			// Get item
			item := reflect.ValueOf(data).MapIndex(key).Interface()

			// Check if item is valid
			if !pipe._handler(item, key.Interface()) {
				continue
			}

			// Append item to new slice
			if newItems == nil {
				newItems = make(map[string]interface{})
			}

			newItems.(map[string]interface{})[key.Interface().(string)] = item
		}

		return newItems, nil

	default:
		return nil, errors.New("data is not a slice or map")
	}
}

func Filter(handler func(item interface{}, key interface{}) bool) Pipe {
	return &filterPipe{
		_handler: handler,
	}
}
