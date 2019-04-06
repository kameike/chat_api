// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostProfileParams creates a new PostProfileParams object
// with the default values initialized.
func NewPostProfileParams() *PostProfileParams {
	var ()
	return &PostProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostProfileParamsWithTimeout creates a new PostProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostProfileParamsWithTimeout(timeout time.Duration) *PostProfileParams {
	var ()
	return &PostProfileParams{

		timeout: timeout,
	}
}

// NewPostProfileParamsWithContext creates a new PostProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostProfileParamsWithContext(ctx context.Context) *PostProfileParams {
	var ()
	return &PostProfileParams{

		Context: ctx,
	}
}

// NewPostProfileParamsWithHTTPClient creates a new PostProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostProfileParamsWithHTTPClient(client *http.Client) *PostProfileParams {
	var ()
	return &PostProfileParams{
		HTTPClient: client,
	}
}

/*PostProfileParams contains all the parameters to send to the API endpoint
for the post profile operation typically these are written to a http.Request
*/
type PostProfileParams struct {

	/*ImageURL
	  画像のURL

	*/
	ImageURL *string
	/*Name
	  名前

	*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post profile params
func (o *PostProfileParams) WithTimeout(timeout time.Duration) *PostProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post profile params
func (o *PostProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post profile params
func (o *PostProfileParams) WithContext(ctx context.Context) *PostProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post profile params
func (o *PostProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post profile params
func (o *PostProfileParams) WithHTTPClient(client *http.Client) *PostProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post profile params
func (o *PostProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImageURL adds the imageURL to the post profile params
func (o *PostProfileParams) WithImageURL(imageURL *string) *PostProfileParams {
	o.SetImageURL(imageURL)
	return o
}

// SetImageURL adds the imageUrl to the post profile params
func (o *PostProfileParams) SetImageURL(imageURL *string) {
	o.ImageURL = imageURL
}

// WithName adds the name to the post profile params
func (o *PostProfileParams) WithName(name *string) *PostProfileParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post profile params
func (o *PostProfileParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PostProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ImageURL != nil {

		// query param imageUrl
		var qrImageURL string
		if o.ImageURL != nil {
			qrImageURL = *o.ImageURL
		}
		qImageURL := qrImageURL
		if qImageURL != "" {
			if err := r.SetQueryParam("imageUrl", qImageURL); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}