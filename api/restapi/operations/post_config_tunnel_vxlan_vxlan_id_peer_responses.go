// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// PostConfigTunnelVxlanVxlanIDPeerOKCode is the HTTP code returned for type PostConfigTunnelVxlanVxlanIDPeerOK
const PostConfigTunnelVxlanVxlanIDPeerOKCode int = 200

/*
PostConfigTunnelVxlanVxlanIDPeerOK OK

swagger:response postConfigTunnelVxlanVxlanIdPeerOK
*/
type PostConfigTunnelVxlanVxlanIDPeerOK struct {

	/*
	  In: Body
	*/
	Payload *models.VxlanPeerEntry `json:"body,omitempty"`
}

// NewPostConfigTunnelVxlanVxlanIDPeerOK creates PostConfigTunnelVxlanVxlanIDPeerOK with default headers values
func NewPostConfigTunnelVxlanVxlanIDPeerOK() *PostConfigTunnelVxlanVxlanIDPeerOK {

	return &PostConfigTunnelVxlanVxlanIDPeerOK{}
}

// WithPayload adds the payload to the post config tunnel vxlan vxlan Id peer o k response
func (o *PostConfigTunnelVxlanVxlanIDPeerOK) WithPayload(payload *models.VxlanPeerEntry) *PostConfigTunnelVxlanVxlanIDPeerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config tunnel vxlan vxlan Id peer o k response
func (o *PostConfigTunnelVxlanVxlanIDPeerOK) SetPayload(payload *models.VxlanPeerEntry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigTunnelVxlanVxlanIDPeerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigTunnelVxlanVxlanIDPeerUnauthorizedCode is the HTTP code returned for type PostConfigTunnelVxlanVxlanIDPeerUnauthorized
const PostConfigTunnelVxlanVxlanIDPeerUnauthorizedCode int = 401

/*
PostConfigTunnelVxlanVxlanIDPeerUnauthorized Invalid authentication credentials

swagger:response postConfigTunnelVxlanVxlanIdPeerUnauthorized
*/
type PostConfigTunnelVxlanVxlanIDPeerUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigTunnelVxlanVxlanIDPeerUnauthorized creates PostConfigTunnelVxlanVxlanIDPeerUnauthorized with default headers values
func NewPostConfigTunnelVxlanVxlanIDPeerUnauthorized() *PostConfigTunnelVxlanVxlanIDPeerUnauthorized {

	return &PostConfigTunnelVxlanVxlanIDPeerUnauthorized{}
}

// WithPayload adds the payload to the post config tunnel vxlan vxlan Id peer unauthorized response
func (o *PostConfigTunnelVxlanVxlanIDPeerUnauthorized) WithPayload(payload *models.Error) *PostConfigTunnelVxlanVxlanIDPeerUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config tunnel vxlan vxlan Id peer unauthorized response
func (o *PostConfigTunnelVxlanVxlanIDPeerUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigTunnelVxlanVxlanIDPeerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigTunnelVxlanVxlanIDPeerInternalServerErrorCode is the HTTP code returned for type PostConfigTunnelVxlanVxlanIDPeerInternalServerError
const PostConfigTunnelVxlanVxlanIDPeerInternalServerErrorCode int = 500

/*
PostConfigTunnelVxlanVxlanIDPeerInternalServerError Internal service error

swagger:response postConfigTunnelVxlanVxlanIdPeerInternalServerError
*/
type PostConfigTunnelVxlanVxlanIDPeerInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigTunnelVxlanVxlanIDPeerInternalServerError creates PostConfigTunnelVxlanVxlanIDPeerInternalServerError with default headers values
func NewPostConfigTunnelVxlanVxlanIDPeerInternalServerError() *PostConfigTunnelVxlanVxlanIDPeerInternalServerError {

	return &PostConfigTunnelVxlanVxlanIDPeerInternalServerError{}
}

// WithPayload adds the payload to the post config tunnel vxlan vxlan Id peer internal server error response
func (o *PostConfigTunnelVxlanVxlanIDPeerInternalServerError) WithPayload(payload *models.Error) *PostConfigTunnelVxlanVxlanIDPeerInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config tunnel vxlan vxlan Id peer internal server error response
func (o *PostConfigTunnelVxlanVxlanIDPeerInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigTunnelVxlanVxlanIDPeerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigTunnelVxlanVxlanIDPeerServiceUnavailableCode is the HTTP code returned for type PostConfigTunnelVxlanVxlanIDPeerServiceUnavailable
const PostConfigTunnelVxlanVxlanIDPeerServiceUnavailableCode int = 503

/*
PostConfigTunnelVxlanVxlanIDPeerServiceUnavailable Maintenance mode

swagger:response postConfigTunnelVxlanVxlanIdPeerServiceUnavailable
*/
type PostConfigTunnelVxlanVxlanIDPeerServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigTunnelVxlanVxlanIDPeerServiceUnavailable creates PostConfigTunnelVxlanVxlanIDPeerServiceUnavailable with default headers values
func NewPostConfigTunnelVxlanVxlanIDPeerServiceUnavailable() *PostConfigTunnelVxlanVxlanIDPeerServiceUnavailable {

	return &PostConfigTunnelVxlanVxlanIDPeerServiceUnavailable{}
}

// WithPayload adds the payload to the post config tunnel vxlan vxlan Id peer service unavailable response
func (o *PostConfigTunnelVxlanVxlanIDPeerServiceUnavailable) WithPayload(payload *models.Error) *PostConfigTunnelVxlanVxlanIDPeerServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config tunnel vxlan vxlan Id peer service unavailable response
func (o *PostConfigTunnelVxlanVxlanIDPeerServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigTunnelVxlanVxlanIDPeerServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
