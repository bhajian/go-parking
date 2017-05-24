package lot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/user/todo/models"
)

// GetStstusOKCode is the HTTP code returned for type GetStstusOK
const GetStstusOKCode int = 200

/*GetStstusOK OK

swagger:response getStstusOK
*/
type GetStstusOK struct {

	/*
	  In: Body
	*/
	Payload *models.Lot `json:"body,omitempty"`
}

// NewGetStstusOK creates GetStstusOK with default headers values
func NewGetStstusOK() *GetStstusOK {
	return &GetStstusOK{}
}

// WithPayload adds the payload to the get ststus o k response
func (o *GetStstusOK) WithPayload(payload *models.Lot) *GetStstusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ststus o k response
func (o *GetStstusOK) SetPayload(payload *models.Lot) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStstusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetStstusDefault error

swagger:response getStstusDefault
*/
type GetStstusDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetStstusDefault creates GetStstusDefault with default headers values
func NewGetStstusDefault(code int) *GetStstusDefault {
	if code <= 0 {
		code = 500
	}

	return &GetStstusDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get ststus default response
func (o *GetStstusDefault) WithStatusCode(code int) *GetStstusDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get ststus default response
func (o *GetStstusDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get ststus default response
func (o *GetStstusDefault) WithPayload(payload *models.Error) *GetStstusDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ststus default response
func (o *GetStstusDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStstusDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}