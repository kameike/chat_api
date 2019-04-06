// Code generated by go-swagger; DO NOT EDIT.

package chat_rooms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	apimodel "github.com/kameike/chat_api/swggen/apimodel"
)

// PostChatroomsReader is a Reader for the PostChatrooms structure.
type PostChatroomsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostChatroomsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostChatroomsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostChatroomsOK creates a PostChatroomsOK with default headers values
func NewPostChatroomsOK() *PostChatroomsOK {
	return &PostChatroomsOK{}
}

/*PostChatroomsOK handles this case with default header values.

ok
*/
type PostChatroomsOK struct {
	Payload []*apimodel.Chatroom
}

func (o *PostChatroomsOK) Error() string {
	return fmt.Sprintf("[POST /chatrooms][%d] postChatroomsOK  %+v", 200, o.Payload)
}

func (o *PostChatroomsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
