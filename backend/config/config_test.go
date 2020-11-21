package config

import (
	"reflect"
	"testing"
)

func TestConfig_New(t *testing.T) {
	type fields struct {
		DBPass string
	}
	type args struct {
		env string
	}
	tests := []struct {
		name string
		args args
		want struct {
			conf Config
			err  string
		}
	}{
		{
			name: "Config is found",
			args: args{env: "test"},
			want: struct {
				conf Config
				err  string
			}{Config{"this_is_a_password", "localhost:3306"}, ""},
		},
		{
			name: "Config is not found and fallback is triggered",
			args: args{env: "doesnotexist"},
			want: struct {
				conf Config
				err  string
			}{Config{"devpass", "localhost:3306"}, ""},
		},
		{
			name: "Config found, but is invalid",
			args: args{env: "testInvalid"},
			want: struct {
				conf Config
				err  string
			}{Config{}, "invalid character"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotConfig, gotErr := New(tt.args.env); !reflect.DeepEqual(gotConfig, tt.want.conf) || !ErrorContains(gotErr, tt.want.err) {
				t.Errorf("New() = %v, err = %v, want %v", gotConfig, gotErr, tt.want)
			}
		})
	}
}
