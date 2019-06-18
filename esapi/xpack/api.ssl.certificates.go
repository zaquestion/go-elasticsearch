// Code generated from specification version 8-0-0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"strings"
)

func newSSLCertificatesFunc(t Transport) SSLCertificates {
	return func(o ...func(*SSLCertificatesRequest)) (*Response, error) {
		var r = SSLCertificatesRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-ssl.html.
//
type SSLCertificates func(o ...func(*SSLCertificatesRequest)) (*Response, error)

// SSLCertificatesRequest configures the Ssl Certificates API request.
//
type SSLCertificatesRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r SSLCertificatesRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(len("/_ssl/certificates"))
	path.WriteString("/_ssl/certificates")

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
func (f SSLCertificates) WithContext(v context.Context) func(*SSLCertificatesRequest) {
	return func(r *SSLCertificatesRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f SSLCertificates) WithPretty() func(*SSLCertificatesRequest) {
	return func(r *SSLCertificatesRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f SSLCertificates) WithHuman() func(*SSLCertificatesRequest) {
	return func(r *SSLCertificatesRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f SSLCertificates) WithErrorTrace() func(*SSLCertificatesRequest) {
	return func(r *SSLCertificatesRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f SSLCertificates) WithFilterPath(v ...string) func(*SSLCertificatesRequest) {
	return func(r *SSLCertificatesRequest) {
		r.FilterPath = v
	}
}
