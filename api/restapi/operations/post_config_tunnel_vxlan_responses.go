// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// PostConfigTunnelVxlanNoContentCode is the HTTP code returned for type PostConfigTunnelVxlanNoContent
const PostConfigTunnelVxlanNoContentCode int = 204

/*
PostConfigTunnelVxlanNoContent OK

swagger:response postConfigTunnelVxlanNoContent
*/
type PostConfigTunnelVxlanNoContent struct {
}

// NewPostConfigTunnelVxlanNoContent creates PostConfigTunnelVxlanNoContent with default headers values
func NewPostConfigTunnelVxlanNoContent() *PostConfigTunnelVxlanNoContent {

	return &PostConfigTunnelVxlanNoContent{}
}

// WriteResponse to the client
func (o *PostConfigTunnelVxlanNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PostConfigTunnelVxlanUnauthorizedCode is the HTTP code returned for type PostConfigTunnelVxlanUnauthorized
const PostConfigTunnelVxlanUnauthorizedCode int = 401

/*
PostConfigTunnelVxlanUnauthorized Invalid authentication credentials

swagger:response postConfigTunnelVxlanUnauthorized
*/
type PostConfigTunnelVxlanUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigTunnelVxlanUnauthorized creates PostConfigTunnelVxlanUnauthorized with default headers values
func NewPostConfigTunnelVxlanUnauthorized() *PostConfigTunnelVxlanUnauthorized {

	return &PostConfigTunnelVxlanUnauthorized{}
}

// WithPayload adds the payload to the post config tunnel vxlan unauthorized response
func (o *PostConfigTunnelVxlanUnauthorized) WithPayload(payload *models.Error) *PostConfigTunnelVxlanUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config tunnel vxlan unauthorized response
func (o *PostConfigTunnelVxlanUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigTunnelVxlanUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigTunnelVxlanConflictCode is the HTTP code returned for type PostConfigTunnelVxlanConflict
const PostConfigTunnelVxlanConflictCode int = 409

/*
PostConfigTunnelVxlanConflict Resource Conflict. VxLAN already exists OR dependency VRF/VNET not found

swagger:response postConfigTunnelVxlanConflict
*/
type PostConfigTunnelVxlanConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigTunnelVxlanConflict creates PostConfigTunnelVxlanConflict with default headers values
func NewPostConfigTunnelVxlanConflict() *PostConfigTunnelVxlanConflict {

	return &PostConfigTunnelVxlanConflict{}
}

// WithPayload adds the payload to the post config tunnel vxlan conflict response
func (o *PostConfigTunnelVxlanConflict) WithPayload(payload *models.Error) *PostConfigTunnelVxlanConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config tunnel vxlan conflict response
func (o *PostConfigTunnelVxlanConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigTunnelVxlanConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigTunnelVxlanInternalServerErrorCode is the HTTP code returned for type PostConfigTunnelVxlanInternalServerError
const PostConfigTunnelVxlanInternalServerErrorCode int = 500

/*
PostConfigTunnelVxlanInternalServerError Internal service error

swagger:response postConfigTunnelVxlanInternalServerError
*/
type PostConfigTunnelVxlanInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigTunnelVxlanInternalServerError creates PostConfigTunnelVxlanInternalServerError with default headers values
func NewPostConfigTunnelVxlanInternalServerError() *PostConfigTunnelVxlanInternalServerError {

	return &PostConfigTunnelVxlanInternalServerError{}
}

// WithPayload adds the payload to the post config tunnel vxlan internal server error response
func (o *PostConfigTunnelVxlanInternalServerError) WithPayload(payload *models.Error) *PostConfigTunnelVxlanInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config tunnel vxlan internal server error response
func (o *PostConfigTunnelVxlanInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigTunnelVxlanInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigTunnelVxlanServiceUnavailableCode is the HTTP code returned for type PostConfigTunnelVxlanServiceUnavailable
const PostConfigTunnelVxlanServiceUnavailableCode int = 503

/*
PostConfigTunnelVxlanServiceUnavailable Maintenance mode

swagger:response postConfigTunnelVxlanServiceUnavailable
*/
type PostConfigTunnelVxlanServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigTunnelVxlanServiceUnavailable creates PostConfigTunnelVxlanServiceUnavailable with default headers values
func NewPostConfigTunnelVxlanServiceUnavailable() *PostConfigTunnelVxlanServiceUnavailable {

	return &PostConfigTunnelVxlanServiceUnavailable{}
}

// WithPayload adds the payload to the post config tunnel vxlan service unavailable response
func (o *PostConfigTunnelVxlanServiceUnavailable) WithPayload(payload *models.Error) *PostConfigTunnelVxlanServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config tunnel vxlan service unavailable response
func (o *PostConfigTunnelVxlanServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigTunnelVxlanServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
