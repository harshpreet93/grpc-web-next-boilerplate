package config

import (
	"log"
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
			err  error
		}
	}{
		{
			name: "Config is found",
			args: args{env: "test"},
			want: struct {
				conf Config
				err  error
			}{Config{"this_is_a_password"}, nil},
		},
		{
			name: "Config is not found and fallback is triggered",
			args: args{env: "doesnotexist"},
			want: struct {
				conf Config
				err  error
			}{Config{"devpass"}, nil},
		},
		// {
		// 	name: "Config found, but is invalid",
		// 	args: args{env: "testInvalid"},
		// 	want: struct {
		// 		conf Config
		// 		err  error
		// 	}{Config{}, nil},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{}
			log.Printf("env is %s", tt.args.env)
			if gotConfig, gotErr := c.New(tt.args.env); !reflect.DeepEqual(gotConfig, tt.want.conf) || !reflect.DeepEqual(gotErr, tt.want.err) {
				t.Errorf("Config.New() = %v, err = %v, want %v", gotConfig, gotErr, tt.want)
			}
		})
	}
}
