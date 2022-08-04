package supporter_test

import (
	"reflect"
	"testing"

	"github.com/zipzoft/supporter-go"
)

func Test_DataGet(t *testing.T) {
	type args struct {
		target interface{}
		key    string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "DataGet_nil",
			args: args{
				target: nil,
				key:    "",
			},
			want: nil,
		},
		{
			name: "DataGet_zero",
			args: args{
				target: 0,
				key:    "",
			},
			want: 0,
		},
		{
			name: "DataGet_empty_string",
			args: args{
				target: "",
				key:    "",
			},
			want: "",
		},
		{
			name: "DataGet_empty_slice",
			args: args{
				target: []interface{}{},
				key:    "",
			},
			want: []interface{}{},
		},
		{
			name: "DataGet_empty_map",
			args: args{
				target: map[string]interface{}{},
				key:    "",
			},
			want: map[string]interface{}{},
		},
		{
			name: "DataGet_not_empty_string",
			args: args{
				target: "test",
				key:    "",
			},
			want: "test",
		},
		{
			name: "DataGet_not_empty_slice",
			args: args{
				target: []interface{}{1, 2, 3},
				key:    "",
			},
			want: []interface{}{1, 2, 3},
		},
		{
			name: "DataGet_not_empty_map",
			args: args{
				target: map[string]interface{}{"test": 1},
				key:    "",
			},
			want: map[string]interface{}{"test": 1},
		},
		{
			name: "DataGet_not_empty_map_key",
			args: args{
				target: map[string]interface{}{"test": 1},
				key:    "test",
			},
			want: 1,
		},
		{
			name: "DataGet_not_empty_map_key_not_found",
			args: args{
				target: map[string]interface{}{"test": 1},
				key:    "test2",
			},
			want: nil,
		},
		{
			name: "DataGet_indent_map",
			args: args{
				target: map[string]interface{}{
					"test": map[string]interface{}{
						"test2": 1,
					},
				},
				key: "test.test2",
			},
			want: 1,
		},
		{
			name: "DataGet_indent_map_not_found",
			args: args{
				target: map[string]interface{}{
					"test": map[string]interface{}{
						"test2": 1,
					},
				},
				key: "test.test3",
			},
			want: nil,
		},
		{
			name: "DataGet_slice_index",
			args: args{
				target: []interface{}{1, 2, 3},
				key:    "0",
			},
			want: 1,
		},
		{
			name: "DataGet_slice_index_not_found",
			args: args{
				target: []interface{}{1, 2, 3},
				key:    "4",
			},
			want: nil,
		},
		{
			name: "DataGet_indent_of_slice",
			args: args{
				target: map[string]interface{}{
					"test": []interface{}{
						map[string]interface{}{
							"test2": 1,
						},
					},
				},
				key: "test.0.test2",
			},
			want: 1,
		},
		{
			name: "DataGet_indent_of_slice_not_found",
			args: args{
				target: map[string]interface{}{
					"test": []interface{}{
						map[string]interface{}{
							"test2": 1,
						},
					},
				},
				key: "test.0.test3",
			},
			want: nil,
		},
		{
			name: "DataGet_indent_of_slice_index",
			args: args{
				target: map[string]interface{}{
					"test": []interface{}{
						map[string]interface{}{
							"test2": 1,
						},
						map[string]interface{}{
							"test2": 2,
						},
					},
				},
				key: "test.1.test2",
			},
			want: 2,
		},
		{
			name: "DataGet_indent_of_slice_index_not_found",
			args: args{
				target: map[string]interface{}{
					"test": []interface{}{
						map[string]interface{}{
							"test2": 1,
						},
						map[string]interface{}{
							"test2": 2,
						},
					},
				},
				key: "test.2.test2",
			},
			want: nil,
		},
		{
			name: "DataGet_index_in_slice",
			args: args{
				target: map[string]interface{}{
					"test": []interface{}{
						map[string]interface{}{
							"test2": 1,
						},
						map[string]interface{}{
							"test2": 2,
						},
					},
				},
				key: "test.1",
			},
			want: map[string]interface{}{
				"test2": 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := supporter.DataGet(tt.args.target, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDataGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		supporter.DataGet(map[string]interface{}{
			"test": map[string]interface{}{
				"test2": 1,
			},
		}, "test.test2")
	}
}
