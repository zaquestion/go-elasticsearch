// Code generated from specification version 8-0-0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newILMPutLifecycleFunc(t Transport) ILMPutLifecycle {
	return func(o ...func(*ILMPutLifecycleRequest)) (*Response, error) {
		var r = ILMPutLifecycleRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-put-lifecycle.html.
//
type ILMPutLifecycle func(o ...func(*ILMPutLifecycleRequest)) (*Response, error)

// ILMPutLifecycleRequest configures the Ilm  Put Lifecycle API request.
//
type ILMPutLifecycleRequest struct {
	Body io.Reader

	Policy string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r ILMPutLifecycleRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

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

	req, _ := newRequest(method, path.String(), r.Body)

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if r.Body != nil {
		req.Header[headerContentType] = headerContentTypeJSON
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
func (f ILMPutLifecycle) WithContext(v context.Context) func(*ILMPutLifecycleRequest) {
	return func(r *ILMPutLifecycleRequest) {
		r.ctx = v
	}
}

// WithBody - The lifecycle policy definition to register.
//
func (f ILMPutLifecycle) WithBody(v io.Reader) func(*ILMPutLifecycleRequest) {
	return func(r *ILMPutLifecycleRequest) {
		r.Body = v
	}
}

// WithPolicy - the name of the index lifecycle policy.
//
func (f ILMPutLifecycle) WithPolicy(v string) func(*ILMPutLifecycleRequest) {
	return func(r *ILMPutLifecycleRequest) {
		r.Policy = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f ILMPutLifecycle) WithPretty() func(*ILMPutLifecycleRequest) {
	return func(r *ILMPutLifecycleRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f ILMPutLifecycle) WithHuman() func(*ILMPutLifecycleRequest) {
	return func(r *ILMPutLifecycleRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f ILMPutLifecycle) WithErrorTrace() func(*ILMPutLifecycleRequest) {
	return func(r *ILMPutLifecycleRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f ILMPutLifecycle) WithFilterPath(v ...string) func(*ILMPutLifecycleRequest) {
	return func(r *ILMPutLifecycleRequest) {
		r.FilterPath = v
	}
}
