// Code generated by go-swagger; DO NOT EDIT.

package chat_rooms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetChatroomsIDOKCode is the HTTP code returned for type GetChatroomsIDOK
const GetChatroomsIDOKCode int = 200

/*GetChatroomsIDOK チャットルームを完全に取得する際にでてくるやつ

swagger:response getChatroomsIdOK
*/
type GetChatroomsIDOK struct {
}

// NewGetChatroomsIDOK creates GetChatroomsIDOK with default headers values
func NewGetChatroomsIDOK() *GetChatroomsIDOK {

	return &GetChatroomsIDOK{}
}

// WriteResponse to the client
func (o *GetChatroomsIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
