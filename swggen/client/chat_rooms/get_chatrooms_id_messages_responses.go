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

// GetChatroomsIDMessagesReader is a Reader for the GetChatroomsIDMessages structure.
type GetChatroomsIDMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChatroomsIDMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetChatroomsIDMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetChatroomsIDMessagesOK creates a GetChatroomsIDMessagesOK with default headers values
func NewGetChatroomsIDMessagesOK() *GetChatroomsIDMessagesOK {
	return &GetChatroomsIDMessagesOK{}
}

/*GetChatroomsIDMessagesOK handles this case with default header values.

ok
*/
type GetChatroomsIDMessagesOK struct {
	Payload []*apimodel.Message
}

func (o *GetChatroomsIDMessagesOK) Error() string {
	return fmt.Sprintf("[GET /chatrooms/{id}/messages][%d] getChatroomsIdMessagesOK  %+v", 200, o.Payload)
}

func (o *GetChatroomsIDMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}