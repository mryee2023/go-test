package http

import (
	"fmt"
	"testing"
)

func TestHttpGet(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args

		wantErr bool
	}{
		{
			name: "test-get",
			args: args{
				url: "http://cp.cloudflare.com/generate_204",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HttpGet(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)

		})
	}
}
