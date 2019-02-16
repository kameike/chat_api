// Code generated by go-swagger; DO NOT EDIT.

package deploy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetHealthOKCode is the HTTP code returned for type GetHealthOK
const GetHealthOKCode int = 200

/*GetHealthOK ok

swagger:response getHealthOK
*/
type GetHealthOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetHealthOK creates GetHealthOK with default headers values
func NewGetHealthOK() *GetHealthOK {

	return &GetHealthOK{}
}

// WithPayload adds the payload to the get health o k response
func (o *GetHealthOK) WithPayload(payload string) *GetHealthOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get health o k response
func (o *GetHealthOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetHealthServiceUnavailableCode is the HTTP code returned for type GetHealthServiceUnavailable
const GetHealthServiceUnavailableCode int = 503

/*GetHealthServiceUnavailable notready

swagger:response getHealthServiceUnavailable
*/
type GetHealthServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetHealthServiceUnavailable creates GetHealthServiceUnavailable with default headers values
func NewGetHealthServiceUnavailable() *GetHealthServiceUnavailable {

	return &GetHealthServiceUnavailable{}
}

// WithPayload adds the payload to the get health service unavailable response
func (o *GetHealthServiceUnavailable) WithPayload(payload string) *GetHealthServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get health service unavailable response
func (o *GetHealthServiceUnavailable) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
