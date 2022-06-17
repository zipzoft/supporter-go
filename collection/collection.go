package collection

import "github.com/zipzoft/supporter-go/collection/pipe"

type Pipeline struct {
	data     interface{}
	pipeline []pipe.Pipe
}

func (c *Pipeline) Pipe(pipes ...pipe.Pipe) *Pipeline {
	c.pipeline = append(c.pipeline, pipes...)
	return c
}

func (c *Pipeline) Get() (newData interface{}, err error) {
	newData = c.data

	for _, pipe := range c.pipeline {
		newData, err = pipe.Handle(newData)
		if err != nil {
			return nil, err
		}
	}

	return newData, nil
}

func NewPipeline(data interface{}) *Pipeline {
	return &Pipeline{
		data:     data,
		pipeline: make([]pipe.Pipe, 0),
	}
}
