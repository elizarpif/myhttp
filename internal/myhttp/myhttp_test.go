package myhttp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"
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

func Test_getHashResponse(t *testing.T) {
	type args struct {
		client *http.Client
		addr   string
		server *httptest.Server
	}

	client := &http.Client{Timeout: time.Second}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "mock",
			args: args{
				client: client,
				server: func() *httptest.Server {
					s := httptest.NewServer(
						http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
							w.Write([]byte("hello"))
						}),
					)
					return s
				}(),
			},
			want: "5d41402abc4b2a76b9719d911017c592",
		},
		{
			name: "mock no response for 3 sec",
			args: args{
				client: client,
				server: func() *httptest.Server {
					s := httptest.NewServer(
						http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
							w.Write([]byte("hello"))
							time.Sleep(time.Second * 3)
						}),
					)
					return s
				}(),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := tt.args.server
			defer server.Close()

			if got := getHashResponse(tt.args.client, server.URL); !reflect.DeepEqual(fmt.Sprintf("%x", got), tt.want) {
				t.Errorf("getHashResponse() = %x, want %v", string(got), string(tt.want))
			}
		})
	}
}

func Test_hashBytes(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hello",
			args: args{
				bytes: []byte("hello"),
			},
			want: "5d41402abc4b2a76b9719d911017c592",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashBytes(tt.args.bytes); !reflect.DeepEqual(fmt.Sprintf("%x", got), tt.want) {
				t.Errorf("hashBytes() = %x, want %s", got, tt.want)
			}
		})
	}
}
