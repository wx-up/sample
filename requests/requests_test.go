package requests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct{}

func (t *TestStruct) Validate() map[string][]string {
	return map[string][]string{
		"name": {"bob", "wx"},
	}
}

func Test_reflectCallValidateFunc(t *testing.T) {
	type args struct {
		obj  any
		name string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{
			name: "method not exist",
			args: args{
				obj:  &TestStruct{},
				name: "Print",
			},
			want: nil,
		},
		{
			name: "pass",
			args: args{
				obj:  &TestStruct{},
				name: "Validate",
			},
			want: map[string][]string{
				"name": {"bob", "wx"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reflectCallValidateFunc(tt.args.obj, tt.args.name)
			assert.Equal(t, tt.want, got)
		})
	}
}
