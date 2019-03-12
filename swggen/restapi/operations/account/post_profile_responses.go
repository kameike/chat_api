// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	apimodel "github.com/kameike/chat_api/swggen/apimodel"
)

// PostProfileOKCode is the HTTP code returned for type PostProfileOK
const PostProfileOKCode int = 200

/*PostProfileOK ok

swagger:response postProfileOK
*/
type PostProfileOK struct {

	/*
	  In: Body
	*/
	Payload *apimodel.User `json:"body,omitempty"`
}

// NewPostProfileOK creates PostProfileOK with default headers values
func NewPostProfileOK() *PostProfileOK {

	return &PostProfileOK{}
}

// WithPayload adds the payload to the post profile o k response
func (o *PostProfileOK) WithPayload(payload *apimodel.User) *PostProfileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post profile o k response
func (o *PostProfileOK) SetPayload(payload *apimodel.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProfileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
