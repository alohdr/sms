// Code generated by go-swagger; DO NOT EDIT.

package template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"hanoman-id/xendit-payment/internal/models"
)

// DeleteTemplateTemplateIDOKCode is the HTTP code returned for type DeleteTemplateTemplateIDOK
const DeleteTemplateTemplateIDOKCode int = 200

/*DeleteTemplateTemplateIDOK Success

swagger:response deleteTemplateTemplateIdOK
*/
type DeleteTemplateTemplateIDOK struct {

	/*
	  In: Body
	*/
	Payload *DeleteTemplateTemplateIDOKBody `json:"body,omitempty"`
}

// NewDeleteTemplateTemplateIDOK creates DeleteTemplateTemplateIDOK with default headers values
func NewDeleteTemplateTemplateIDOK() *DeleteTemplateTemplateIDOK {

	return &DeleteTemplateTemplateIDOK{}
}

// WithPayload adds the payload to the delete template template Id o k response
func (o *DeleteTemplateTemplateIDOK) WithPayload(payload *DeleteTemplateTemplateIDOKBody) *DeleteTemplateTemplateIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete template template Id o k response
func (o *DeleteTemplateTemplateIDOK) SetPayload(payload *DeleteTemplateTemplateIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTemplateTemplateIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteTemplateTemplateIDBadRequestCode is the HTTP code returned for type DeleteTemplateTemplateIDBadRequest
const DeleteTemplateTemplateIDBadRequestCode int = 400

/*DeleteTemplateTemplateIDBadRequest Unauthorized

swagger:response deleteTemplateTemplateIdBadRequest
*/
type DeleteTemplateTemplateIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTemplateTemplateIDBadRequest creates DeleteTemplateTemplateIDBadRequest with default headers values
func NewDeleteTemplateTemplateIDBadRequest() *DeleteTemplateTemplateIDBadRequest {

	return &DeleteTemplateTemplateIDBadRequest{}
}

// WithPayload adds the payload to the delete template template Id bad request response
func (o *DeleteTemplateTemplateIDBadRequest) WithPayload(payload *models.Error) *DeleteTemplateTemplateIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete template template Id bad request response
func (o *DeleteTemplateTemplateIDBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTemplateTemplateIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteTemplateTemplateIDUnauthorizedCode is the HTTP code returned for type DeleteTemplateTemplateIDUnauthorized
const DeleteTemplateTemplateIDUnauthorizedCode int = 401

/*DeleteTemplateTemplateIDUnauthorized Unauthorized

swagger:response deleteTemplateTemplateIdUnauthorized
*/
type DeleteTemplateTemplateIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTemplateTemplateIDUnauthorized creates DeleteTemplateTemplateIDUnauthorized with default headers values
func NewDeleteTemplateTemplateIDUnauthorized() *DeleteTemplateTemplateIDUnauthorized {

	return &DeleteTemplateTemplateIDUnauthorized{}
}

// WithPayload adds the payload to the delete template template Id unauthorized response
func (o *DeleteTemplateTemplateIDUnauthorized) WithPayload(payload *models.Error) *DeleteTemplateTemplateIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete template template Id unauthorized response
func (o *DeleteTemplateTemplateIDUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTemplateTemplateIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteTemplateTemplateIDNotFoundCode is the HTTP code returned for type DeleteTemplateTemplateIDNotFound
const DeleteTemplateTemplateIDNotFoundCode int = 404

/*DeleteTemplateTemplateIDNotFound The specified resource was not found

swagger:response deleteTemplateTemplateIdNotFound
*/
type DeleteTemplateTemplateIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTemplateTemplateIDNotFound creates DeleteTemplateTemplateIDNotFound with default headers values
func NewDeleteTemplateTemplateIDNotFound() *DeleteTemplateTemplateIDNotFound {

	return &DeleteTemplateTemplateIDNotFound{}
}

// WithPayload adds the payload to the delete template template Id not found response
func (o *DeleteTemplateTemplateIDNotFound) WithPayload(payload *models.Error) *DeleteTemplateTemplateIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete template template Id not found response
func (o *DeleteTemplateTemplateIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTemplateTemplateIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
