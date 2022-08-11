package pipe

import (
	"errors"
	"reflect"
)

var _ Pipe = (*firstPipe)(nil)

type firstPipe struct{}

// Handle implements Pipe
func (pipe *firstPipe) Handle(data interface{}) (interface{}, error) {
	if data == nil {
		return nil, errors.New("data is nil")
	}

	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		// Foreach item in slice
		for i := 0; i < reflect.ValueOf(data).Len(); i++ {
			// Get item
			item := reflect.ValueOf(data).Index(i).Interface()

			return item, nil
		}

		return nil, nil

	case reflect.Map:
		// Foreach item in map
		for _, key := range reflect.ValueOf(data).MapKeys() {
			// Get item
			item := reflect.ValueOf(data).MapIndex(key).Interface()

			return item, nil
		}

		return nil, nil

	default:
		return nil, errors.New("data is not a slice or map")

	}
}

func First() Pipe {
	return &firstPipe{}
}
