// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: b8d375e642
// Version Date: Fri Aug 4 23:55:38 UTC 2017

// Package http provides an HTTP client for the Rello service.
package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/pkg/errors"

	// This Service
	pb "github.com/myambition/ambition/services/rello/rello-service"
	"github.com/myambition/ambition/services/rello/rello-service/svc"
)

var (
	_ = endpoint.Chain
	_ = httptransport.NewClient
	_ = fmt.Sprint
	_ = bytes.Compare
	_ = ioutil.NopCloser
)

// New returns a service backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(instance string, options ...ClientOption) (pb.RelloServer, error) {
	var cc clientConfig

	for _, f := range options {
		err := f(&cc)
		if err != nil {
			return nil, errors.Wrap(err, "cannot apply option")
		}
	}

	clientOptions := []httptransport.ClientOption{
		httptransport.ClientBefore(
			contextValuesToHttpHeaders(cc.headers)),
	}

	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	_ = u

	var CheckListWebhookZeroEndpoint endpoint.Endpoint
	{
		CheckListWebhookZeroEndpoint = httptransport.NewClient(
			"POST",
			copyURL(u, "/"),
			EncodeHTTPCheckListWebhookZeroRequest,
			DecodeHTTPCheckListWebhookResponse,
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		CheckListWebhookEndpoint: CheckListWebhookZeroEndpoint,
	}, nil
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}

type clientConfig struct {
	headers []string
}

// ClientOption is a function that modifies the client config
type ClientOption func(*clientConfig) error

// CtxValuesToSend configures the http client to pull the specified keys out of
// the context and add them to the http request as headers.  Note that keys
// will have net/http.CanonicalHeaderKey called on them before being send over
// the wire and that is the form they will be available in the server context.
func CtxValuesToSend(keys ...string) ClientOption {
	return func(o *clientConfig) error {
		o.headers = keys
		return nil
	}
}

func contextValuesToHttpHeaders(keys []string) httptransport.RequestFunc {
	return func(ctx context.Context, r *http.Request) context.Context {
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				r.Header.Set(k, v)
			}
		}

		return ctx
	}
}

// HTTP Client Decode

// DecodeHTTPCheckListWebhookResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded Empty response from the HTTP response body.
// If the response has a non-200 status code, we will interpret that as an
// error and attempt to decode the specific error message from the response
// body. Primarily useful in a client.
func DecodeHTTPCheckListWebhookResponse(_ context.Context, r *http.Response) (interface{}, error) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read http body")
	}

	if len(buf) == 0 {
		return nil, errors.New("response http body empty")
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(errorDecoder(buf), "status code: '%d'", r.StatusCode)
	}

	var resp pb.Empty
	if err = json.Unmarshal(buf, &resp); err != nil {
		return nil, errorDecoder(buf)
	}

	return &resp, nil
}

// HTTP Client Encode

// EncodeHTTPCheckListWebhookZeroRequest is a transport/http.EncodeRequestFunc
// that encodes a checklistwebhook request into the various portions of
// the http request (path, query, and body).
func EncodeHTTPCheckListWebhookZeroRequest(_ context.Context, r *http.Request, request interface{}) error {
	strval := ""
	_ = strval
	req := request.(*pb.ChecklistUpdate)
	_ = req

	r.Header.Set("transport", "HTTPJSON")
	r.Header.Set("request-url", r.URL.Path)

	// Set the path parameters
	path := strings.Join([]string{
		"",
		"",
	}, "/")
	u, err := url.Parse(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't unmarshal path %q", path)
	}
	r.URL.RawPath = u.RawPath
	r.URL.Path = u.Path

	// Set the query parameters
	values := r.URL.Query()
	var tmp []byte
	_ = tmp

	r.URL.RawQuery = values.Encode()

	// Set the body parameters
	var buf bytes.Buffer
	toRet := request.(*pb.ChecklistUpdate)

	toRet.Model = req.Model

	toRet.Action = req.Action

	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(toRet); err != nil {
		return errors.Wrapf(err, "couldn't encode body as json %v", toRet)
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func errorDecoder(buf []byte) error {
	var w errorWrapper
	if err := json.Unmarshal(buf, &w); err != nil {
		const size = 8196
		if len(buf) > size {
			buf = buf[:size]
		}
		return fmt.Errorf("response body '%s': cannot parse non-json request body", buf)
	}

	return errors.New(w.Error)
}

type errorWrapper struct {
	Error string `json:"error"`
}
