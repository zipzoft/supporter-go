package pipe

import (
	"errors"
	"reflect"
)

var _ Pipe = (*findPipe)(nil)

type findPipe struct {
	_handler func(item interface{}, key interface{}) bool
}

// Handle implements Pipe
func (pipe *findPipe) Handle(data interface{}) (interface{}, error) {
	if data == nil {
		return nil, errors.New("data is nil")
	}
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		// Foreach item in slice
		for i := 0; i < reflect.ValueOf(data).Len(); i++ {
			// Get item
			item := reflect.ValueOf(data).Index(i).Interface()

			// Check if item is valid
			if !pipe._handler(item, i) {
				continue
			}

			return item, nil
		}

		return nil, nil

	case reflect.Map:
		// Foreach item in map
		for _, key := range reflect.ValueOf(data).MapKeys() {
			// Get item
			item := reflect.ValueOf(data).MapIndex(key).Interface()

			// Check if item is valid
			if !pipe._handler(item, key.Interface()) {
				continue
			}

			return item, nil
		}

		return nil, nil

	default:
		return nil, errors.New("data is not a slice or map")
	}

}

func Find(handler func(item interface{}, key interface{}) bool) Pipe {
	return &findPipe{
		_handler: handler,
	}
}
