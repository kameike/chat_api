// Code generated by go-swagger; DO NOT EDIT.

package chat_rooms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	apimodel "github.com/kameike/chat_api/swggen/apimodel"
)

// PostChatroomsIDMessagesOKCode is the HTTP code returned for type PostChatroomsIDMessagesOK
const PostChatroomsIDMessagesOKCode int = 200

/*PostChatroomsIDMessagesOK ok

swagger:response postChatroomsIdMessagesOK
*/
type PostChatroomsIDMessagesOK struct {

	/*
	  In: Body
	*/
	Payload *apimodel.ChatroomFull `json:"body,omitempty"`
}

// NewPostChatroomsIDMessagesOK creates PostChatroomsIDMessagesOK with default headers values
func NewPostChatroomsIDMessagesOK() *PostChatroomsIDMessagesOK {

	return &PostChatroomsIDMessagesOK{}
}

// WithPayload adds the payload to the post chatrooms Id messages o k response
func (o *PostChatroomsIDMessagesOK) WithPayload(payload *apimodel.ChatroomFull) *PostChatroomsIDMessagesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post chatrooms Id messages o k response
func (o *PostChatroomsIDMessagesOK) SetPayload(payload *apimodel.ChatroomFull) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostChatroomsIDMessagesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
