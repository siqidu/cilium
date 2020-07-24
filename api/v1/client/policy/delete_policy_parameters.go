// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// NewDeletePolicyParams creates a new DeletePolicyParams object
// with the default values initialized.
func NewDeletePolicyParams() *DeletePolicyParams {
	var ()
	return &DeletePolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePolicyParamsWithTimeout creates a new DeletePolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePolicyParamsWithTimeout(timeout time.Duration) *DeletePolicyParams {
	var ()
	return &DeletePolicyParams{

		timeout: timeout,
	}
}

// NewDeletePolicyParamsWithContext creates a new DeletePolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePolicyParamsWithContext(ctx context.Context) *DeletePolicyParams {
	var ()
	return &DeletePolicyParams{

		Context: ctx,
	}
}

// NewDeletePolicyParamsWithHTTPClient creates a new DeletePolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePolicyParamsWithHTTPClient(client *http.Client) *DeletePolicyParams {
	var ()
	return &DeletePolicyParams{
		HTTPClient: client,
	}
}

/*DeletePolicyParams contains all the parameters to send to the API endpoint
for the delete policy operation typically these are written to a http.Request
*/
type DeletePolicyParams struct {

	/*Labels*/
	Labels models.Labels

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete policy params
func (o *DeletePolicyParams) WithTimeout(timeout time.Duration) *DeletePolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete policy params
func (o *DeletePolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete policy params
func (o *DeletePolicyParams) WithContext(ctx context.Context) *DeletePolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete policy params
func (o *DeletePolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete policy params
func (o *DeletePolicyParams) WithHTTPClient(client *http.Client) *DeletePolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete policy params
func (o *DeletePolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLabels adds the labels to the delete policy params
func (o *DeletePolicyParams) WithLabels(labels models.Labels) *DeletePolicyParams {
	o.SetLabels(labels)
	return o
}

// SetLabels adds the labels to the delete policy params
func (o *DeletePolicyParams) SetLabels(labels models.Labels) {
	o.Labels = labels
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Labels != nil {
		if err := r.SetBodyParam(o.Labels); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
