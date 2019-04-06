// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	apimodel "github.com/kameike/chat_api/swggen/apimodel"
)

// PostAuthReader is a Reader for the PostAuth structure.
type PostAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostAuthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAuthOK creates a PostAuthOK with default headers values
func NewPostAuthOK() *PostAuthOK {
	return &PostAuthOK{}
}

/*PostAuthOK handles this case with default header values.

ok
*/
type PostAuthOK struct {
	Payload *apimodel.AuthInfo
}

func (o *PostAuthOK) Error() string {
	return fmt.Sprintf("[POST /auth][%d] postAuthOK  %+v", 200, o.Payload)
}

func (o *PostAuthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(apimodel.AuthInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}