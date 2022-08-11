package pipe

import (
	"errors"
	"reflect"
)

var _ Pipe = (*reducePipe)(nil)

type reducePipe struct {
	initial  interface{}
	_handler func(carry interface{}, item interface{}) interface{}
}

// Handle implements Pipe
func (pipe *reducePipe) Handle(data interface{}) (interface{}, error) {
	if data == nil {
		return nil, errors.New("data is nil")
	}

	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		// Foreach item in slice
		carry := pipe.initial
		for i := 0; i < reflect.ValueOf(data).Len(); i++ {
			// Get item
			item := reflect.ValueOf(data).Index(i).Interface()

			carry = pipe._handler(carry, item)
		}

		return carry, nil

	case reflect.Map:
		// Foreach item in map
		carry := pipe.initial
		for _, key := range reflect.ValueOf(data).MapKeys() {
			// Get item
			item := reflect.ValueOf(data).MapIndex(key).Interface()

			carry = pipe._handler(carry, item)
		}

		return carry, nil

	default:
		return nil, errors.New("data is not a slice or map")
	}
}

func Reduce(handler func(carry interface{}, item interface{}) interface{}, initial interface{}) Pipe {
	return &reducePipe{
		_handler: handler,
		initial:  initial,
	}
}
