package svc

// This file provides server-side bindings for the HTTP transport.
// It utilizes the transport/http.Server.

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
)

/*
var customHTTPBefore httptransport.RequestFunc

// MakeHTTPHandler returns a handler that makes a set of endpoints available
// on predefined paths.
func MakeHTTPHandler(endpoints Endpoints) http.Handler {
	serverOptions := []httptransport.ServerOption{
	httptransport.ServerBefore(headersToContext, customHTTPBefore),
	httptransport.ServerErrorEncoder(errorEncoder),
	httptransport.ServerAfter(httptransport.SetContentType(contentType)),

*/
func init() {
	customHTTPBefore = dumpBodyToContext
}

type HTTPBody struct{}

func dumpBodyToContext(ctx context.Context, r *http.Request) context.Context {
	var save []byte
	var err error
	save, r.Body, err = drainBody(r.Body)
	if err != nil {
		return ctx
	}
	ctx = context.WithValue(ctx, HTTPBody{}, save)
	return ctx
}

func drainBody(b io.ReadCloser) ([]byte, io.ReadCloser, error) {
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err := b.Close(); err != nil {
		return nil, b, err
	}

	return buf.Bytes(), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}
