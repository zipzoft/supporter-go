package supporter_test

import (
	"reflect"
	"testing"

	"github.com/zipzoft/supporter-go"
)

func Test_InArray(t *testing.T) {
	type args struct {
		val interface{}
		arr []interface{}
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "InArray_nil",
			args: args{
				val: nil,
				arr: []interface{}{},
			},
			want: false,
		},
		{
			name: "InArray_zero",
			args: args{
				val: 0,
				arr: []interface{}{},
			},
			want: false,
		},
		{
			name: "InArray_empty_string",
			args: args{
				val: "",
				arr: []interface{}{},
			},
			want: false,
		},
		{
			name: "InArray_empty_slice",
			args: args{
				val: []interface{}{},
				arr: []interface{}{},
			},
			want: false,
		},
		{
			name: "InArray_empty_map",
			args: args{
				val: map[string]interface{}{},
				arr: []interface{}{},
			},
			want: false,
		},
		{
			name: "InArray_not_empty_string",
			args: args{
				val: "test",
				arr: []interface{}{},
			},
			want: false,
		},
		{
			name: "InArray_not_empty_slice",
			args: args{
				val: []interface{}{1, 2, 3},
				arr: []interface{}{},
			},
			want: false,
		},
		{
			name: "InArray_found_string",
			args: args{
				val: "test",
				arr: []interface{}{"test", "test2"},
			},
			want: true,
		},
		{
			name: "InArray_not_empty_map",
			args: args{
				val: map[string]interface{}{"test": 1},
				arr: []interface{}{},
			},
			want: false,
		},
		{
			name: "InArray_not_empty_map_key",
			args: args{
				val: map[string]interface{}{"test": 1},
				arr: []interface{}{map[string]interface{}{"test": 1}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := supporter.InArray(tt.args.val, tt.args.arr); got != tt.want {
				t.Errorf("InArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_InArray(b *testing.B) {
	type args struct {
		val interface{}
		arr []interface{}
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "InArray_nil",
			args: args{
				val: nil,
				arr: []interface{}{},
			},
			want: false,
		},
		{
			name: "InArray_zero",
			args: args{
				val: 0,
				arr: []interface{}{},
			},
			want: false,
		},
		{
			name: "InArray_empty_string",
			args: args{
				val: "",
				arr: []interface{}{},
			},
			want: false,
		},
		{
			name: "InArray_empty_slice",
			args: args{
				val: []interface{}{},
				arr: []interface{}{},
			},
			want: false,
		},
		{
			name: "InArray_empty_map",
			args: args{
				val: map[string]interface{}{},
				arr: []interface{}{},
			},
			want: false,
		},
		{
			name: "InArray_not_empty_string",
			args: args{
				val: "test",
				arr: []interface{}{},
			},
			want: false,
		},
		{
			name: "InArray_not_empty_slice",
			args: args{
				val: []interface{}{1, 2, 3},
				arr: []interface{}{},
			},
			want: false,
		},
		{
			name: "InArray_found_string",
			args: args{
				val: "test",
				arr: []interface{}{"test", "test2"},
			},
			want: true,
		},
		{
			name: "InArray_not_empty_map",
			args: args{
				val: map[string]interface{}{"test": 1},
				arr: []interface{}{},
			},
			want: false,
		},
		{
			name: "InArray_not_empty_map_key",
			args: args{
				val: map[string]interface{}{"test": 1},
				arr: []interface{}{map[string]interface{}{"test": 1}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				supporter.InArray(tt.args.val, tt.args.arr)
			}
		})
	}
}

func Test_First(t *testing.T) {
	type args struct {
		arr []interface{}
	}

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "First_nil",
			args: args{
				arr: nil,
			},
			want: nil,
		},
		{
			name: "First_empty",
			args: args{
				arr: []interface{}{},
			},
			want: nil,
		},
		{
			name: "First_one",
			args: args{
				arr: []interface{}{1},
			},
			want: 1,
		},
		{
			name: "First_two",
			args: args{
				arr: []interface{}{1, 2},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := supporter.First(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("First() = %v, want %v", got, tt.want)
			}
		})
	}
}
