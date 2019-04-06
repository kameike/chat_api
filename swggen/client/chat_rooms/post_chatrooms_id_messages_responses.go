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

// PostChatroomsIDMessagesReader is a Reader for the PostChatroomsIDMessages structure.
type PostChatroomsIDMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostChatroomsIDMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostChatroomsIDMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostChatroomsIDMessagesOK creates a PostChatroomsIDMessagesOK with default headers values
func NewPostChatroomsIDMessagesOK() *PostChatroomsIDMessagesOK {
	return &PostChatroomsIDMessagesOK{}
}

/*PostChatroomsIDMessagesOK handles this case with default header values.

ok
*/
type PostChatroomsIDMessagesOK struct {
	Payload *apimodel.ChatroomFull
}

func (o *PostChatroomsIDMessagesOK) Error() string {
	return fmt.Sprintf("[POST /chatrooms/{id}/messages][%d] postChatroomsIdMessagesOK  %+v", 200, o.Payload)
}

func (o *PostChatroomsIDMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(apimodel.ChatroomFull)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
