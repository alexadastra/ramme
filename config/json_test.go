package config

import (
	"testing"
	"time"

	"github.com/maxatome/go-testdeep/td"
)

func Test_newTargetFromJSON(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *Target
		wantErr bool
	}{
		{
			args: args{
				filePath: "./assets/config.json",
			},
			wantErr: false,
			want: &Target{
				Basic: map[Name]*Entry{
					"host":                     {Val: "0.0.0.0", T: "string"},
					"grpc_port":                {Val: 6560, T: "int"},
					"http_port":                {Val: 8080, T: "int"},
					"http_write_timeout":       {Val: 15 * time.Second, T: "duration"},
					"http_admin_port":          {Val: 8081, T: "int"},
					"http_admin_read_timeout":  {Val: 15 * time.Second, T: "duration"},
					"http_admin_write_timeout": {Val: 15 * time.Second, T: "duration"},
					"log_level":                {Val: 1, T: "int"},
					"is_local_environment":     {Val: true, T: "bool"},
					"http_read_timeout":        {Val: 15 * time.Second, T: "duration"},
				},
				Advanced: map[Name]*Entry{
					"ping_message":  {Val: "Hello Blip Blop", T: "string"},
					"mongo_db_dsn":  {Val: "mongodb://admin:coolmongobongo1583@localhost:1491", T: "string"},
					"some_duration": {Val: 3 * time.Hour, T: "duration"},
					"some_bool":     {Val: true, T: "bool"},
					"some_int":      {Val: 42069, T: "int"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newTargetFromJSON(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("newTargetFromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !td.Cmp(t, got, tt.want) {
				t.Errorf("NewTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
