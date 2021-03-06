// Code generated by go-swagger; DO NOT EDIT.

package messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	apimodel "github.com/kameike/chat_api/swggen/apimodel"
)

// GetChatroomsChatroomHashMessagesOKCode is the HTTP code returned for type GetChatroomsChatroomHashMessagesOK
const GetChatroomsChatroomHashMessagesOKCode int = 200

/*GetChatroomsChatroomHashMessagesOK ok

swagger:response getChatroomsChatroomHashMessagesOK
*/
type GetChatroomsChatroomHashMessagesOK struct {

	/*
	  In: Body
	*/
	Payload *apimodel.MessagesResponse `json:"body,omitempty"`
}

// NewGetChatroomsChatroomHashMessagesOK creates GetChatroomsChatroomHashMessagesOK with default headers values
func NewGetChatroomsChatroomHashMessagesOK() *GetChatroomsChatroomHashMessagesOK {

	return &GetChatroomsChatroomHashMessagesOK{}
}

// WithPayload adds the payload to the get chatrooms chatroom hash messages o k response
func (o *GetChatroomsChatroomHashMessagesOK) WithPayload(payload *apimodel.MessagesResponse) *GetChatroomsChatroomHashMessagesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get chatrooms chatroom hash messages o k response
func (o *GetChatroomsChatroomHashMessagesOK) SetPayload(payload *apimodel.MessagesResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChatroomsChatroomHashMessagesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetChatroomsChatroomHashMessagesForbiddenCode is the HTTP code returned for type GetChatroomsChatroomHashMessagesForbidden
const GetChatroomsChatroomHashMessagesForbiddenCode int = 403

/*GetChatroomsChatroomHashMessagesForbidden error

swagger:response getChatroomsChatroomHashMessagesForbidden
*/
type GetChatroomsChatroomHashMessagesForbidden struct {

	/*
	  In: Body
	*/
	Payload *apimodel.Error `json:"body,omitempty"`
}

// NewGetChatroomsChatroomHashMessagesForbidden creates GetChatroomsChatroomHashMessagesForbidden with default headers values
func NewGetChatroomsChatroomHashMessagesForbidden() *GetChatroomsChatroomHashMessagesForbidden {

	return &GetChatroomsChatroomHashMessagesForbidden{}
}

// WithPayload adds the payload to the get chatrooms chatroom hash messages forbidden response
func (o *GetChatroomsChatroomHashMessagesForbidden) WithPayload(payload *apimodel.Error) *GetChatroomsChatroomHashMessagesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get chatrooms chatroom hash messages forbidden response
func (o *GetChatroomsChatroomHashMessagesForbidden) SetPayload(payload *apimodel.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetChatroomsChatroomHashMessagesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
