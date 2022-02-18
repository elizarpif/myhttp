package myhttp

import (
	"net/url"
	"reflect"
	"testing"
)

func Test_getUrl(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		{
			name: "usual url",
			args: args{
				addr: "http://www.adjust.com",
			},
			want: &url.URL{
				Scheme: "http",
				Host:   "www.adjust.com",
			},
			wantErr: false,
		},
		{
			name: "without http url",
			args: args{
				addr: "www.adjust.com",
			},
			want: &url.URL{
				Scheme: "http",
				Path:   "www.adjust.com",
			},
			wantErr: false,
		},
		{
			name: "only host",
			args: args{
				addr: "adjust.com",
			},
			want: &url.URL{
				Scheme: "http",
				Path:   "adjust.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getUrl(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUrl() got = %v, want %v", got.String(), tt.want.String())
			}
		})
	}
}
