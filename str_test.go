package supporter_test

import (
	"reflect"
	"testing"

	"github.com/zipzoft/supporter-go"
)

func Test_IsEmpty(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "IsEmpty_nil",
			args: args{
				val: nil,
			},
			want: true,
		},
		{
			name: "IsEmpty_zero",
			args: args{
				val: reflect.Zero(reflect.TypeOf(0)).Interface(),
			},
			want: true,
		},
		{
			name: "IsEmpty_empty_string",
			args: args{
				val: "",
			},
			want: true,
		},
		{
			name: "IsEmpty_empty_string_with_space",
			args: args{
				val: " ",
			},
			want: true,
		},
		{
			name: "IsEmpty_empty_bool",
			args: args{
				val: false,
			},
			want: true,
		},
		{
			name: "IsEmpty_empty_slice",
			args: args{
				val: []interface{}{},
			},
			want: true,
		},
		{
			name: "IsEmpty_empty_map",
			args: args{
				val: map[string]interface{}{},
			},
			want: true,
		},
		{
			name: "IsEmpty_not_empty_string",
			args: args{
				val: "test",
			},
			want: false,
		},
		{
			name: "IsEmpty_not_empty_slice",
			args: args{
				val: []interface{}{1, 2, 3},
			},
			want: false,
		},
		{
			name: "IsEmpty_not_empty_map",
			args: args{
				val: map[string]interface{}{"test": 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := supporter.IsEmpty(tt.args.val); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MatchGroupsAllSub(t *testing.T) {
	type args struct {
		pattern string
		str     string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{
			name: "MatchGroupsAllSub_empty_pattern",
			args: args{
				pattern: "",
				str:     "",
			},
			want: map[string][]string{},
		},
		{
			name: "MatchGroupsAllSub_empty_str",
			args: args{
				pattern: "test",
				str:     "",
			},
			want: map[string][]string{},
		},
		{
			name: "MatchGroupsAllSub_empty_pattern_and_str",
			args: args{
				pattern: "",
				str:     "",
			},
			want: map[string][]string{},
		},
		{
			name: "MatchGroupsAllSub_pattern_and_str",
			args: args{
				pattern: `(?P<name>\w+)`,
				str:     "test",
			},
			want: map[string][]string{
				"name": {"test"},
			},
		},
		{
			name: "MatchGroupsAllSub_pattern_and_str_with_multiple_groups",
			args: args{
				pattern: `name (?P<name>\w+) age (?P<age>\d+)`,
				str:     "name John age 20",
			},
			want: map[string][]string{
				"name": {"John"},
				"age":  {"20"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := supporter.MatchGroupsAllSub(tt.args.pattern, tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchGroupsAllSub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_IsNotEmpty(t *testing.T) {
	type args struct {
		val interface{}
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "IsNotEmpty_nil",
			args: args{
				val: nil,
			},
			want: false,
		},
		{
			name: "IsNotEmpty_zero",
			args: args{
				val: reflect.Zero(reflect.TypeOf(0)).Interface(),
			},
			want: false,
		},
		{
			name: "IsNotEmpty_empty_string",
			args: args{
				val: "",
			},
			want: false,
		},
		{
			name: "IsNotEmpty_empty_string_with_space",
			args: args{
				val: " ",
			},
			want: false,
		},
		{
			name: "IsNotEmpty_empty_bool",
			args: args{
				val: false,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := supporter.IsNotEmpty(tt.args.val); got != tt.want {
				t.Errorf("IsNotEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_StrRandom(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name   string
		args   args
		length int
	}{
		{
			name: "StrRandom_length_0",
			args: args{
				length: 0,
			},
			length: 0,
		},
		{
			name: "StrRandom_length_1",
			args: args{
				length: 1,
			},
			length: 1,
		},
		{
			name: "StrRandom_length_10",
			args: args{
				length: 10,
			},
			length: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := supporter.StrRandom(tt.args.length); len(got) != tt.length {
				t.Errorf("StrRandom() = %v, want %v", got, tt.length)
			}
		})
	}
}

func Test_MatchGroups(t *testing.T) {
	type args struct {
		pattern string
		str     string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "MatchGroups_empty_str",
			args: args{
				pattern: "test",
				str:     "",
			},
			want: map[string]string{},
		},
		{
			name: "MatchGroups_empty_pattern_and_str",
			args: args{
				pattern: "",
				str:     "",
			},
			want: map[string]string{},
		},
		{
			name: "MatchGroups_pattern_and_str",
			args: args{
				pattern: `(?P<name>\w+)`,
				str:     "test",
			},
			want: map[string]string{
				"name": "test",
			},
		},
		{
			name: "MatchGroups_pattern_and_str_with_multiple_groups",
			args: args{
				pattern: `name (?P<name>\w+) age (?P<age>\d+)`,
				str:     "name John age 20",
			},
			want: map[string]string{
				"name": "John",
				"age":  "20",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := supporter.MatchGroups(tt.args.pattern, tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
