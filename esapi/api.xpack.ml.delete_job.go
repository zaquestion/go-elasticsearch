// Code generated from specification version 8-0-0-SNAPSHOT: DO NOT EDIT

package esapi

import (
	"context"
	"strconv"
	"strings"
)

func newMLDeleteJobFunc(t Transport) MLDeleteJob {
	return func(job_id string, o ...func(*MLDeleteJobRequest)) (*Response, error) {
		var r = MLDeleteJobRequest{JobID: job_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-job.html.
//
type MLDeleteJob func(job_id string, o ...func(*MLDeleteJobRequest)) (*Response, error)

// MLDeleteJobRequest configures the Ml  Delete Job API request.
//
type MLDeleteJobRequest struct {
	JobID string

	Force             *bool
	WaitForCompletion *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MLDeleteJobRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "DELETE"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	path.WriteString("/")
	path.WriteString(r.JobID)

	params = make(map[string]string)

	if r.Force != nil {
		params["force"] = strconv.FormatBool(*r.Force)
	}

	if r.WaitForCompletion != nil {
		params["wait_for_completion"] = strconv.FormatBool(*r.WaitForCompletion)
	}

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
func (f MLDeleteJob) WithContext(v context.Context) func(*MLDeleteJobRequest) {
	return func(r *MLDeleteJobRequest) {
		r.ctx = v
	}
}

// WithForce - true if the job should be forcefully deleted.
//
func (f MLDeleteJob) WithForce(v bool) func(*MLDeleteJobRequest) {
	return func(r *MLDeleteJobRequest) {
		r.Force = &v
	}
}

// WithWaitForCompletion - should this request wait until the operation has completed before returning.
//
func (f MLDeleteJob) WithWaitForCompletion(v bool) func(*MLDeleteJobRequest) {
	return func(r *MLDeleteJobRequest) {
		r.WaitForCompletion = &v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MLDeleteJob) WithPretty() func(*MLDeleteJobRequest) {
	return func(r *MLDeleteJobRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MLDeleteJob) WithHuman() func(*MLDeleteJobRequest) {
	return func(r *MLDeleteJobRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MLDeleteJob) WithErrorTrace() func(*MLDeleteJobRequest) {
	return func(r *MLDeleteJobRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MLDeleteJob) WithFilterPath(v ...string) func(*MLDeleteJobRequest) {
	return func(r *MLDeleteJobRequest) {
		r.FilterPath = v
	}
}
