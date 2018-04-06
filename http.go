package util

// util
// Copyright (C) 2018 Maximilian Pachl

// MIT License
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// ---------------------------------------------------------------------------------------
//  imports
// ---------------------------------------------------------------------------------------

import (
	"encoding/json"
	"errors"
	"net"
	"net/http"

	"github.com/gorilla/schema"
)

// ---------------------------------------------------------------------------------------
//  errors
// ---------------------------------------------------------------------------------------

var (
	ErrInvalidContentType = errors.New("invalid content type")
)

// ---------------------------------------------------------------------------------------
//  public functions
// ---------------------------------------------------------------------------------------

// GetRemoteAddr returns the remote address of an http request.
// If an X-Forwarded-For Header is present the headers content is returned.
// Otherwise the src host of the ip packet is returned.
func GetRemoteAddr(r *http.Request) string {
	remote := r.RemoteAddr
	if fwd := r.Header.Get("X-Forwarded-For"); len(fwd) > 0 {
		return fwd
	}

	// remove the port from the remote address
	if host, _, err := net.SplitHostPort(remote); err == nil {
		remote = host
	}

	return remote
}

// Jsonify writes the JSON representation of v to the supplied
// http.ResposeWriter. If an error occours while marshalling the
// http response will be an internal server error.
func Jsonify(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// ParseBody reads the body of the request and parses it into v.
func ParseBody(r *http.Request, v interface{}) error {
	switch r.Header.Get("Content-Type") {
	case "application/x-www-form-urlencoded":
		err := r.ParseForm()
		if err != nil {
			return err
		}

		return schema.NewDecoder().Decode(v, r.Form)

	case "application/json":
		return json.NewDecoder(r.Body).Decode(v)

	default:
		return ErrInvalidContentType
	}
}