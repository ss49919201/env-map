package envmap

import (
	"os"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name       string
		want       map[string]string
		setEnvFunc func()
	}{
		{
			"empty",
			map[string]string{},
			func() {},
		},
		{
			"one",
			map[string]string{
				"foo": "bar",
			},
			func() {
				os.Setenv("foo", "bar")
			},
		},
		{
			"multiple",
			map[string]string{
				"foo": "bar",
				"baz": "qux",
			},
			func() {
				os.Setenv("foo", "bar")
				os.Setenv("baz", "qux")
			},
		},
	}
	for _, tt := range tests {
		os.Clearenv()

		tt.setEnvFunc()

		t.Run(tt.name, func(t *testing.T) {
			if got := Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
