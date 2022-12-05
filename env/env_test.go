package env

import (
	"os"
	"reflect"
	"testing"
)

func TestStorage_ParseSecrets(t *testing.T) {
	appName := "RAMME-TEMPLATE"
	type args struct {
		appName string
		env     map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "no env",
			args: args{
				appName: appName,
				env:     map[string]string{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "correct env, empty val",
			args: args{
				appName: appName,
				env: map[string]string{
					appName + "-SECRET": "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "correct env, empty json",
			args: args{
				appName: appName,
				env: map[string]string{
					appName + "-SECRET": "{}",
				},
			},
			want:    map[string]interface{}{},
			wantErr: false,
		},
		{
			name: "correct env, json with values",
			args: args{
				appName: appName,
				env: map[string]string{
					appName + "-SECRET": "{\"fancy_password\":\"pass123\", \"another_secret\": 4756, \"more_secret\":true}",
				},
			},
			want: map[string]interface{}{
				"fancy_password": "pass123",
				"another_secret": 4756.0,
				"more_secret":    true,
			},
			wantErr: false,
		},
		{
			name: "correct env, multiline json with values ",
			args: args{
				appName: appName,
				env: map[string]string{
					appName + "-SECRET": "{\"fancy_password\":\"pass123\", \n \"another_secret\": 4756,\n \"more_secret\":true \n}",
				},
			},
			want: map[string]interface{}{
				"fancy_password": "pass123",
				"another_secret": 4756.0,
				"more_secret":    true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for name, val := range tt.args.env {
				if err := os.Setenv(name, val); err != nil {
					panic(err)
				}
			}

			s := &Storage{}
			got, err := s.ParseSecrets(tt.args.appName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.ParseSecrets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.ParseSecrets() = %v, want %v", got, tt.want)
			}
		})
	}
}
