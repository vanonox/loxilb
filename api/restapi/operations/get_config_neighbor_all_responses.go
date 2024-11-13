// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetConfigNeighborAllOKCode is the HTTP code returned for type GetConfigNeighborAllOK
const GetConfigNeighborAllOKCode int = 200

/*
GetConfigNeighborAllOK OK

swagger:response getConfigNeighborAllOK
*/
type GetConfigNeighborAllOK struct {

	/*
	  In: Body
	*/
	Payload *GetConfigNeighborAllOKBody `json:"body,omitempty"`
}

// NewGetConfigNeighborAllOK creates GetConfigNeighborAllOK with default headers values
func NewGetConfigNeighborAllOK() *GetConfigNeighborAllOK {

	return &GetConfigNeighborAllOK{}
}

// WithPayload adds the payload to the get config neighbor all o k response
func (o *GetConfigNeighborAllOK) WithPayload(payload *GetConfigNeighborAllOKBody) *GetConfigNeighborAllOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config neighbor all o k response
func (o *GetConfigNeighborAllOK) SetPayload(payload *GetConfigNeighborAllOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigNeighborAllOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigNeighborAllUnauthorizedCode is the HTTP code returned for type GetConfigNeighborAllUnauthorized
const GetConfigNeighborAllUnauthorizedCode int = 401

/*
GetConfigNeighborAllUnauthorized Invalid authentication credentials

swagger:response getConfigNeighborAllUnauthorized
*/
type GetConfigNeighborAllUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigNeighborAllUnauthorized creates GetConfigNeighborAllUnauthorized with default headers values
func NewGetConfigNeighborAllUnauthorized() *GetConfigNeighborAllUnauthorized {

	return &GetConfigNeighborAllUnauthorized{}
}

// WithPayload adds the payload to the get config neighbor all unauthorized response
func (o *GetConfigNeighborAllUnauthorized) WithPayload(payload *models.Error) *GetConfigNeighborAllUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config neighbor all unauthorized response
func (o *GetConfigNeighborAllUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigNeighborAllUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigNeighborAllInternalServerErrorCode is the HTTP code returned for type GetConfigNeighborAllInternalServerError
const GetConfigNeighborAllInternalServerErrorCode int = 500

/*
GetConfigNeighborAllInternalServerError Internal service error

swagger:response getConfigNeighborAllInternalServerError
*/
type GetConfigNeighborAllInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigNeighborAllInternalServerError creates GetConfigNeighborAllInternalServerError with default headers values
func NewGetConfigNeighborAllInternalServerError() *GetConfigNeighborAllInternalServerError {

	return &GetConfigNeighborAllInternalServerError{}
}

// WithPayload adds the payload to the get config neighbor all internal server error response
func (o *GetConfigNeighborAllInternalServerError) WithPayload(payload *models.Error) *GetConfigNeighborAllInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config neighbor all internal server error response
func (o *GetConfigNeighborAllInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigNeighborAllInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigNeighborAllServiceUnavailableCode is the HTTP code returned for type GetConfigNeighborAllServiceUnavailable
const GetConfigNeighborAllServiceUnavailableCode int = 503

/*
GetConfigNeighborAllServiceUnavailable Maintenance mode

swagger:response getConfigNeighborAllServiceUnavailable
*/
type GetConfigNeighborAllServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigNeighborAllServiceUnavailable creates GetConfigNeighborAllServiceUnavailable with default headers values
func NewGetConfigNeighborAllServiceUnavailable() *GetConfigNeighborAllServiceUnavailable {

	return &GetConfigNeighborAllServiceUnavailable{}
}

// WithPayload adds the payload to the get config neighbor all service unavailable response
func (o *GetConfigNeighborAllServiceUnavailable) WithPayload(payload *models.Error) *GetConfigNeighborAllServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config neighbor all service unavailable response
func (o *GetConfigNeighborAllServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigNeighborAllServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
