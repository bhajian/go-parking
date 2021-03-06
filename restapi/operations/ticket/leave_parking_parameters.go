package ticket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewLeaveParkingParams creates a new LeaveParkingParams object
// with the default values initialized.
func NewLeaveParkingParams() LeaveParkingParams {
	var ()
	return LeaveParkingParams{}
}

// LeaveParkingParams contains all the bound params for the leave parking operation
// typically these are obtained from a http.Request
//
// swagger:parameters leaveParking
type LeaveParkingParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  Required: true
	  In: path
	*/
	LotID int64
	/*
	  Required: true
	  In: path
	*/
	TicketID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *LeaveParkingParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rLotID, rhkLotID, _ := route.Params.GetOK("lotId")
	if err := o.bindLotID(rLotID, rhkLotID, route.Formats); err != nil {
		res = append(res, err)
	}

	rTicketID, rhkTicketID, _ := route.Params.GetOK("ticketId")
	if err := o.bindTicketID(rTicketID, rhkTicketID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *LeaveParkingParams) bindLotID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("lotId", "path", "int64", raw)
	}
	o.LotID = value

	return nil
}

func (o *LeaveParkingParams) bindTicketID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("ticketId", "path", "int64", raw)
	}
	o.TicketID = value

	return nil
}
