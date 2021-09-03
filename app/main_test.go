package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testName = "Testman"
)

func Test_helloWorldHandlersome(t *testing.T) {

	method := http.MethodGet
	url := "http://localtest/"
	emptyRequest, _ := http.NewRequest("GET", url, nil)
	validRequest, _ := http.NewRequest(method, fmt.Sprintf("%s?name=%s", url, testName), nil)
	validEmptyResponse := "Hello, World!"
	validResponse := fmt.Sprintf("Hello, %s!", testName)

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		wantResp string
	}{
		{
			name: "valid hello world",
			args: args{
				w: httptest.NewRecorder(),
				r: emptyRequest,
			},
			wantErr:  false,
			wantResp: validEmptyResponse,
		},
		{
			name: "valid hello with arg",
			args: args{
				w: httptest.NewRecorder(),
				r: validRequest,
			},
			wantErr:  false,
			wantResp: validResponse,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(helloWorldHandler)
			handler.ServeHTTP(tt.args.w, tt.args.r)

			if tt.args.w.Body.String() != tt.wantResp {
				t.Errorf("helloWorldHandler: got resp= %v, want resp %v", tt.args.w.Body.String(), tt.wantResp)
				return
			}
		})
	}
}
