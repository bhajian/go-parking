package lot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/user/todo/models"
)

// FindLotsOKCode is the HTTP code returned for type FindLotsOK
const FindLotsOKCode int = 200

/*FindLotsOK returns a list the lots

swagger:response findLotsOK
*/
type FindLotsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Lot `json:"body,omitempty"`
}

// NewFindLotsOK creates FindLotsOK with default headers values
func NewFindLotsOK() *FindLotsOK {
	return &FindLotsOK{}
}

// WithPayload adds the payload to the find lots o k response
func (o *FindLotsOK) WithPayload(payload []*models.Lot) *FindLotsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find lots o k response
func (o *FindLotsOK) SetPayload(payload []*models.Lot) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindLotsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Lot, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*FindLotsDefault generic error response

swagger:response findLotsDefault
*/
type FindLotsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindLotsDefault creates FindLotsDefault with default headers values
func NewFindLotsDefault(code int) *FindLotsDefault {
	if code <= 0 {
		code = 500
	}

	return &FindLotsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find lots default response
func (o *FindLotsDefault) WithStatusCode(code int) *FindLotsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find lots default response
func (o *FindLotsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find lots default response
func (o *FindLotsDefault) WithPayload(payload *models.Error) *FindLotsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find lots default response
func (o *FindLotsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindLotsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
