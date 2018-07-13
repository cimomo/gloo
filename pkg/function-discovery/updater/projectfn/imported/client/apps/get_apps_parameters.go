// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAppsParams creates a new GetAppsParams object
// with the default values initialized.
func NewGetAppsParams() *GetAppsParams {
	var ()
	return &GetAppsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppsParamsWithTimeout creates a new GetAppsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAppsParamsWithTimeout(timeout time.Duration) *GetAppsParams {
	var ()
	return &GetAppsParams{

		timeout: timeout,
	}
}

// NewGetAppsParamsWithContext creates a new GetAppsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAppsParamsWithContext(ctx context.Context) *GetAppsParams {
	var ()
	return &GetAppsParams{

		Context: ctx,
	}
}

// NewGetAppsParamsWithHTTPClient creates a new GetAppsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAppsParamsWithHTTPClient(client *http.Client) *GetAppsParams {
	var ()
	return &GetAppsParams{
		HTTPClient: client,
	}
}

/*GetAppsParams contains all the parameters to send to the API endpoint
for the get apps operation typically these are written to a http.Request
*/
type GetAppsParams struct {

	/*Cursor
	  Cursor from previous response.next_cursor to begin results after, if any.

	*/
	Cursor *string
	/*PerPage
	  Number of results to return, defaults to 30. Max of 100.

	*/
	PerPage *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get apps params
func (o *GetAppsParams) WithTimeout(timeout time.Duration) *GetAppsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get apps params
func (o *GetAppsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get apps params
func (o *GetAppsParams) WithContext(ctx context.Context) *GetAppsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get apps params
func (o *GetAppsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get apps params
func (o *GetAppsParams) WithHTTPClient(client *http.Client) *GetAppsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get apps params
func (o *GetAppsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the get apps params
func (o *GetAppsParams) WithCursor(cursor *string) *GetAppsParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the get apps params
func (o *GetAppsParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithPerPage adds the perPage to the get apps params
func (o *GetAppsParams) WithPerPage(perPage *int64) *GetAppsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get apps params
func (o *GetAppsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string
		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {
			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}

	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int64
		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {
			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}