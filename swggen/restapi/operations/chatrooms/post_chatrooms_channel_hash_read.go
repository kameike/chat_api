// Code generated by go-swagger; DO NOT EDIT.

package chatrooms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostChatroomsChannelHashReadHandlerFunc turns a function with the right signature into a post chatrooms channel hash read handler
type PostChatroomsChannelHashReadHandlerFunc func(PostChatroomsChannelHashReadParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostChatroomsChannelHashReadHandlerFunc) Handle(params PostChatroomsChannelHashReadParams) middleware.Responder {
	return fn(params)
}

// PostChatroomsChannelHashReadHandler interface for that can handle valid post chatrooms channel hash read params
type PostChatroomsChannelHashReadHandler interface {
	Handle(PostChatroomsChannelHashReadParams) middleware.Responder
}

// NewPostChatroomsChannelHashRead creates a new http.Handler for the post chatrooms channel hash read operation
func NewPostChatroomsChannelHashRead(ctx *middleware.Context, handler PostChatroomsChannelHashReadHandler) *PostChatroomsChannelHashRead {
	return &PostChatroomsChannelHashRead{Context: ctx, Handler: handler}
}

/*PostChatroomsChannelHashRead swagger:route POST /chatrooms/{channel_hash}/read chatrooms postChatroomsChannelHashRead

全部既読にするやつ

*/
type PostChatroomsChannelHashRead struct {
	Context *middleware.Context
	Handler PostChatroomsChannelHashReadHandler
}

func (o *PostChatroomsChannelHashRead) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostChatroomsChannelHashReadParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}