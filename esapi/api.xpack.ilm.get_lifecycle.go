// Code generated from specification version 8-0-0-SNAPSHOT: DO NOT EDIT

package esapi

import (
	"context"
	"strings"
)

func newILMGetLifecycleFunc(t Transport) ILMGetLifecycle {
	return func(o ...func(*ILMGetLifecycleRequest)) (*Response, error) {
		var r = ILMGetLifecycleRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-get-lifecycle.html.
//
type ILMGetLifecycle func(o ...func(*ILMGetLifecycleRequest)) (*Response, error)

// ILMGetLifecycleRequest configures the Ilm  Get Lifecycle API request.
//
type ILMGetLifecycleRequest struct {
	Policy string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r ILMGetLifecycleRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(1 + len("_ilm") + 1 + len("policy") + 1 + len(r.Policy))
	path.WriteString("/")
	path.WriteString("_ilm")
	path.WriteString("/")
	path.WriteString("policy")
	if r.Policy != "" {
		path.WriteString("/")
		path.WriteString(r.Policy)
	}

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, _ := newRequest(method, path.String(), nil)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f ILMGetLifecycle) WithContext(v context.Context) func(*ILMGetLifecycleRequest) {
	return func(r *ILMGetLifecycleRequest) {
		r.ctx = v
	}
}

// WithPolicy - the name of the index lifecycle policy.
//
func (f ILMGetLifecycle) WithPolicy(v string) func(*ILMGetLifecycleRequest) {
	return func(r *ILMGetLifecycleRequest) {
		r.Policy = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f ILMGetLifecycle) WithPretty() func(*ILMGetLifecycleRequest) {
	return func(r *ILMGetLifecycleRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f ILMGetLifecycle) WithHuman() func(*ILMGetLifecycleRequest) {
	return func(r *ILMGetLifecycleRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f ILMGetLifecycle) WithErrorTrace() func(*ILMGetLifecycleRequest) {
	return func(r *ILMGetLifecycleRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f ILMGetLifecycle) WithFilterPath(v ...string) func(*ILMGetLifecycleRequest) {
	return func(r *ILMGetLifecycleRequest) {
		r.FilterPath = v
	}
}
