// Code generated from specification version 8-0-0-SNAPSHOT: DO NOT EDIT

package xpack

import (
	"context"
	"io"
	"strings"
)

func newMLUpdateModelSnapshotFunc(t Transport) MLUpdateModelSnapshot {
	return func(body io.Reader, job_id string, snapshot_id string, o ...func(*MLUpdateModelSnapshotRequest)) (*Response, error) {
		var r = MLUpdateModelSnapshotRequest{Body: body, JobID: job_id, SnapshotID: snapshot_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

//
// See full documentation at http://www.elastic.co/guide/en/elasticsearch/reference/current/ml-update-snapshot.html.
//
type MLUpdateModelSnapshot func(body io.Reader, job_id string, snapshot_id string, o ...func(*MLUpdateModelSnapshotRequest)) (*Response, error)

// MLUpdateModelSnapshotRequest configures the Ml   Update Model Snapshot API request.
//
type MLUpdateModelSnapshotRequest struct {
	Body io.Reader

	JobID      string
	SnapshotID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r MLUpdateModelSnapshotRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(1 + len("_ml") + 1 + len("anomaly_detectors") + 1 + len(r.JobID) + 1 + len("model_snapshots") + 1 + len(r.SnapshotID) + 1 + len("_update"))
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("anomaly_detectors")
	path.WriteString("/")
	path.WriteString(r.JobID)
	path.WriteString("/")
	path.WriteString("model_snapshots")
	path.WriteString("/")
	path.WriteString(r.SnapshotID)
	path.WriteString("/")
	path.WriteString("_update")

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
func (f MLUpdateModelSnapshot) WithContext(v context.Context) func(*MLUpdateModelSnapshotRequest) {
	return func(r *MLUpdateModelSnapshotRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f MLUpdateModelSnapshot) WithPretty() func(*MLUpdateModelSnapshotRequest) {
	return func(r *MLUpdateModelSnapshotRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f MLUpdateModelSnapshot) WithHuman() func(*MLUpdateModelSnapshotRequest) {
	return func(r *MLUpdateModelSnapshotRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f MLUpdateModelSnapshot) WithErrorTrace() func(*MLUpdateModelSnapshotRequest) {
	return func(r *MLUpdateModelSnapshotRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f MLUpdateModelSnapshot) WithFilterPath(v ...string) func(*MLUpdateModelSnapshotRequest) {
	return func(r *MLUpdateModelSnapshotRequest) {
		r.FilterPath = v
	}
}
