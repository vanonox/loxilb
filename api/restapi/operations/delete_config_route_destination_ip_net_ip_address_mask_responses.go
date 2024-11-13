// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// DeleteConfigRouteDestinationIPNetIPAddressMaskNoContentCode is the HTTP code returned for type DeleteConfigRouteDestinationIPNetIPAddressMaskNoContent
const DeleteConfigRouteDestinationIPNetIPAddressMaskNoContentCode int = 204

/*
DeleteConfigRouteDestinationIPNetIPAddressMaskNoContent OK

swagger:response deleteConfigRouteDestinationIpNetIpAddressMaskNoContent
*/
type DeleteConfigRouteDestinationIPNetIPAddressMaskNoContent struct {
}

// NewDeleteConfigRouteDestinationIPNetIPAddressMaskNoContent creates DeleteConfigRouteDestinationIPNetIPAddressMaskNoContent with default headers values
func NewDeleteConfigRouteDestinationIPNetIPAddressMaskNoContent() *DeleteConfigRouteDestinationIPNetIPAddressMaskNoContent {

	return &DeleteConfigRouteDestinationIPNetIPAddressMaskNoContent{}
}

// WriteResponse to the client
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteConfigRouteDestinationIPNetIPAddressMaskBadRequestCode is the HTTP code returned for type DeleteConfigRouteDestinationIPNetIPAddressMaskBadRequest
const DeleteConfigRouteDestinationIPNetIPAddressMaskBadRequestCode int = 400

/*
DeleteConfigRouteDestinationIPNetIPAddressMaskBadRequest Malformed arguments for API call

swagger:response deleteConfigRouteDestinationIpNetIpAddressMaskBadRequest
*/
type DeleteConfigRouteDestinationIPNetIPAddressMaskBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigRouteDestinationIPNetIPAddressMaskBadRequest creates DeleteConfigRouteDestinationIPNetIPAddressMaskBadRequest with default headers values
func NewDeleteConfigRouteDestinationIPNetIPAddressMaskBadRequest() *DeleteConfigRouteDestinationIPNetIPAddressMaskBadRequest {

	return &DeleteConfigRouteDestinationIPNetIPAddressMaskBadRequest{}
}

// WithPayload adds the payload to the delete config route destination Ip net Ip address mask bad request response
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskBadRequest) WithPayload(payload *models.Error) *DeleteConfigRouteDestinationIPNetIPAddressMaskBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config route destination Ip net Ip address mask bad request response
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigRouteDestinationIPNetIPAddressMaskUnauthorizedCode is the HTTP code returned for type DeleteConfigRouteDestinationIPNetIPAddressMaskUnauthorized
const DeleteConfigRouteDestinationIPNetIPAddressMaskUnauthorizedCode int = 401

/*
DeleteConfigRouteDestinationIPNetIPAddressMaskUnauthorized Invalid authentication credentials

swagger:response deleteConfigRouteDestinationIpNetIpAddressMaskUnauthorized
*/
type DeleteConfigRouteDestinationIPNetIPAddressMaskUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigRouteDestinationIPNetIPAddressMaskUnauthorized creates DeleteConfigRouteDestinationIPNetIPAddressMaskUnauthorized with default headers values
func NewDeleteConfigRouteDestinationIPNetIPAddressMaskUnauthorized() *DeleteConfigRouteDestinationIPNetIPAddressMaskUnauthorized {

	return &DeleteConfigRouteDestinationIPNetIPAddressMaskUnauthorized{}
}

// WithPayload adds the payload to the delete config route destination Ip net Ip address mask unauthorized response
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskUnauthorized) WithPayload(payload *models.Error) *DeleteConfigRouteDestinationIPNetIPAddressMaskUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config route destination Ip net Ip address mask unauthorized response
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigRouteDestinationIPNetIPAddressMaskForbiddenCode is the HTTP code returned for type DeleteConfigRouteDestinationIPNetIPAddressMaskForbidden
const DeleteConfigRouteDestinationIPNetIPAddressMaskForbiddenCode int = 403

/*
DeleteConfigRouteDestinationIPNetIPAddressMaskForbidden Capacity insufficient

swagger:response deleteConfigRouteDestinationIpNetIpAddressMaskForbidden
*/
type DeleteConfigRouteDestinationIPNetIPAddressMaskForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigRouteDestinationIPNetIPAddressMaskForbidden creates DeleteConfigRouteDestinationIPNetIPAddressMaskForbidden with default headers values
func NewDeleteConfigRouteDestinationIPNetIPAddressMaskForbidden() *DeleteConfigRouteDestinationIPNetIPAddressMaskForbidden {

	return &DeleteConfigRouteDestinationIPNetIPAddressMaskForbidden{}
}

// WithPayload adds the payload to the delete config route destination Ip net Ip address mask forbidden response
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskForbidden) WithPayload(payload *models.Error) *DeleteConfigRouteDestinationIPNetIPAddressMaskForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config route destination Ip net Ip address mask forbidden response
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigRouteDestinationIPNetIPAddressMaskNotFoundCode is the HTTP code returned for type DeleteConfigRouteDestinationIPNetIPAddressMaskNotFound
const DeleteConfigRouteDestinationIPNetIPAddressMaskNotFoundCode int = 404

/*
DeleteConfigRouteDestinationIPNetIPAddressMaskNotFound Resource not found

swagger:response deleteConfigRouteDestinationIpNetIpAddressMaskNotFound
*/
type DeleteConfigRouteDestinationIPNetIPAddressMaskNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigRouteDestinationIPNetIPAddressMaskNotFound creates DeleteConfigRouteDestinationIPNetIPAddressMaskNotFound with default headers values
func NewDeleteConfigRouteDestinationIPNetIPAddressMaskNotFound() *DeleteConfigRouteDestinationIPNetIPAddressMaskNotFound {

	return &DeleteConfigRouteDestinationIPNetIPAddressMaskNotFound{}
}

// WithPayload adds the payload to the delete config route destination Ip net Ip address mask not found response
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskNotFound) WithPayload(payload *models.Error) *DeleteConfigRouteDestinationIPNetIPAddressMaskNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config route destination Ip net Ip address mask not found response
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigRouteDestinationIPNetIPAddressMaskConflictCode is the HTTP code returned for type DeleteConfigRouteDestinationIPNetIPAddressMaskConflict
const DeleteConfigRouteDestinationIPNetIPAddressMaskConflictCode int = 409

/*
DeleteConfigRouteDestinationIPNetIPAddressMaskConflict Resource Conflict. VLAN already exists OR dependency VRF/VNET not found

swagger:response deleteConfigRouteDestinationIpNetIpAddressMaskConflict
*/
type DeleteConfigRouteDestinationIPNetIPAddressMaskConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigRouteDestinationIPNetIPAddressMaskConflict creates DeleteConfigRouteDestinationIPNetIPAddressMaskConflict with default headers values
func NewDeleteConfigRouteDestinationIPNetIPAddressMaskConflict() *DeleteConfigRouteDestinationIPNetIPAddressMaskConflict {

	return &DeleteConfigRouteDestinationIPNetIPAddressMaskConflict{}
}

// WithPayload adds the payload to the delete config route destination Ip net Ip address mask conflict response
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskConflict) WithPayload(payload *models.Error) *DeleteConfigRouteDestinationIPNetIPAddressMaskConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config route destination Ip net Ip address mask conflict response
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigRouteDestinationIPNetIPAddressMaskInternalServerErrorCode is the HTTP code returned for type DeleteConfigRouteDestinationIPNetIPAddressMaskInternalServerError
const DeleteConfigRouteDestinationIPNetIPAddressMaskInternalServerErrorCode int = 500

/*
DeleteConfigRouteDestinationIPNetIPAddressMaskInternalServerError Internal service error

swagger:response deleteConfigRouteDestinationIpNetIpAddressMaskInternalServerError
*/
type DeleteConfigRouteDestinationIPNetIPAddressMaskInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigRouteDestinationIPNetIPAddressMaskInternalServerError creates DeleteConfigRouteDestinationIPNetIPAddressMaskInternalServerError with default headers values
func NewDeleteConfigRouteDestinationIPNetIPAddressMaskInternalServerError() *DeleteConfigRouteDestinationIPNetIPAddressMaskInternalServerError {

	return &DeleteConfigRouteDestinationIPNetIPAddressMaskInternalServerError{}
}

// WithPayload adds the payload to the delete config route destination Ip net Ip address mask internal server error response
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskInternalServerError) WithPayload(payload *models.Error) *DeleteConfigRouteDestinationIPNetIPAddressMaskInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config route destination Ip net Ip address mask internal server error response
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigRouteDestinationIPNetIPAddressMaskServiceUnavailableCode is the HTTP code returned for type DeleteConfigRouteDestinationIPNetIPAddressMaskServiceUnavailable
const DeleteConfigRouteDestinationIPNetIPAddressMaskServiceUnavailableCode int = 503

/*
DeleteConfigRouteDestinationIPNetIPAddressMaskServiceUnavailable Maintenance mode

swagger:response deleteConfigRouteDestinationIpNetIpAddressMaskServiceUnavailable
*/
type DeleteConfigRouteDestinationIPNetIPAddressMaskServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigRouteDestinationIPNetIPAddressMaskServiceUnavailable creates DeleteConfigRouteDestinationIPNetIPAddressMaskServiceUnavailable with default headers values
func NewDeleteConfigRouteDestinationIPNetIPAddressMaskServiceUnavailable() *DeleteConfigRouteDestinationIPNetIPAddressMaskServiceUnavailable {

	return &DeleteConfigRouteDestinationIPNetIPAddressMaskServiceUnavailable{}
}

// WithPayload adds the payload to the delete config route destination Ip net Ip address mask service unavailable response
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskServiceUnavailable) WithPayload(payload *models.Error) *DeleteConfigRouteDestinationIPNetIPAddressMaskServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config route destination Ip net Ip address mask service unavailable response
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigRouteDestinationIPNetIPAddressMaskServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
