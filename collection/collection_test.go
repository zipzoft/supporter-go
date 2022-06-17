package collection_test

import (
	"testing"

	"github.com/zipzoft/supporter-go/collection"
	"github.com/zipzoft/supporter-go/collection/pipe"
)

func TestCollection(t *testing.T) {

	t.Run("Collection with filter pipe", func(t *testing.T) {
		expected := []int{2, 4, 6, 8, 10}

		pipeline := collection.NewPipeline([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

		pipeline.Pipe(
			pipe.Filter(func(i interface{}, key interface{}) bool {
				return i.(int)%2 == 0
			}),
		)

		data, err := pipeline.Get()
		if err != nil {
			t.Error(err)
		}

		// Convert interface to int
		items := make([]int, 0)
		for _, item := range data.([]interface{}) {
			items = append(items, item.(int))
		}

		for i, v := range items {
			if v != expected[i] {
				t.Errorf("Expected %v, got %v", expected[i], v)
			}
		}
	})

	t.Run("Collection with map pipe", func(t *testing.T) {
		pipeline := collection.NewPipeline([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		expected := []int{16, 18, 20}

		pipeline.Pipe(
			pipe.Map(func(i interface{}, key interface{}) interface{} {
				return i.(int) * 2
			}),

			pipe.Filter(func(i interface{}, key interface{}) bool {
				return i.(int) > 15
			}),
		)

		data, err := pipeline.Get()
		if err != nil {
			t.Error(err)
		}

		// Convert interface to int
		items := make([]int, 0)
		for _, item := range data.([]interface{}) {
			items = append(items, item.(int))
		}

		for i, v := range items {
			if v != expected[i] {
				t.Errorf("Expected %v, got %v", expected[i], v)
			}
		}
	})
}
