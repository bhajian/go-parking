package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/user/todo/restapi/operations/lot"
	"github.com/user/todo/restapi/operations/ticket"
)

// NewParkingAppAPI creates a new ParkingApp instance
func NewParkingAppAPI(spec *loads.Document) *ParkingAppAPI {
	return &ParkingAppAPI{
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
		spec:            spec,
		ServeError:      errors.ServeError,
		JSONConsumer:    runtime.JSONConsumer(),
		JSONProducer:    runtime.JSONProducer(),
		LotAddOneHandler: lot.AddOneHandlerFunc(func(params lot.AddOneParams) middleware.Responder {
			return middleware.NotImplemented("operation LotAddOne has not yet been implemented")
		}),
		LotDestroyOneHandler: lot.DestroyOneHandlerFunc(func(params lot.DestroyOneParams) middleware.Responder {
			return middleware.NotImplemented("operation LotDestroyOne has not yet been implemented")
		}),
		LotFindLotsHandler: lot.FindLotsHandlerFunc(func(params lot.FindLotsParams) middleware.Responder {
			return middleware.NotImplemented("operation LotFindLots has not yet been implemented")
		}),
		LotGetStstusHandler: lot.GetStstusHandlerFunc(func(params lot.GetStstusParams) middleware.Responder {
			return middleware.NotImplemented("operation LotGetStstus has not yet been implemented")
		}),
		TicketGetTicketHandler: ticket.GetTicketHandlerFunc(func(params ticket.GetTicketParams) middleware.Responder {
			return middleware.NotImplemented("operation TicketGetTicket has not yet been implemented")
		}),
		TicketLeaveParkingHandler: ticket.LeaveParkingHandlerFunc(func(params ticket.LeaveParkingParams) middleware.Responder {
			return middleware.NotImplemented("operation TicketLeaveParking has not yet been implemented")
		}),
		LotUpdateOneHandler: lot.UpdateOneHandlerFunc(func(params lot.UpdateOneParams) middleware.Responder {
			return middleware.NotImplemented("operation LotUpdateOne has not yet been implemented")
		}),
	}
}

/*ParkingAppAPI The product of a parking.io Behnam Hajian */
type ParkingAppAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	// JSONConsumer registers a consumer for a "application/io.goswagger.parking.v1+json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/io.goswagger.parking.v1+json" mime type
	JSONProducer runtime.Producer

	// LotAddOneHandler sets the operation handler for the add one operation
	LotAddOneHandler lot.AddOneHandler
	// LotDestroyOneHandler sets the operation handler for the destroy one operation
	LotDestroyOneHandler lot.DestroyOneHandler
	// LotFindLotsHandler sets the operation handler for the find lots operation
	LotFindLotsHandler lot.FindLotsHandler
	// LotGetStstusHandler sets the operation handler for the get ststus operation
	LotGetStstusHandler lot.GetStstusHandler
	// TicketGetTicketHandler sets the operation handler for the get ticket operation
	TicketGetTicketHandler ticket.GetTicketHandler
	// TicketLeaveParkingHandler sets the operation handler for the leave parking operation
	TicketLeaveParkingHandler ticket.LeaveParkingHandler
	// LotUpdateOneHandler sets the operation handler for the update one operation
	LotUpdateOneHandler lot.UpdateOneHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *ParkingAppAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ParkingAppAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ParkingAppAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ParkingAppAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ParkingAppAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ParkingAppAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ParkingAppAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ParkingAppAPI
func (o *ParkingAppAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.LotAddOneHandler == nil {
		unregistered = append(unregistered, "lot.AddOneHandler")
	}

	if o.LotDestroyOneHandler == nil {
		unregistered = append(unregistered, "lot.DestroyOneHandler")
	}

	if o.LotFindLotsHandler == nil {
		unregistered = append(unregistered, "lot.FindLotsHandler")
	}

	if o.LotGetStstusHandler == nil {
		unregistered = append(unregistered, "lot.GetStstusHandler")
	}

	if o.TicketGetTicketHandler == nil {
		unregistered = append(unregistered, "ticket.GetTicketHandler")
	}

	if o.TicketLeaveParkingHandler == nil {
		unregistered = append(unregistered, "ticket.LeaveParkingHandler")
	}

	if o.LotUpdateOneHandler == nil {
		unregistered = append(unregistered, "lot.UpdateOneHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ParkingAppAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ParkingAppAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *ParkingAppAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/io.goswagger.parking.v1+json":
			result["application/io.goswagger.parking.v1+json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *ParkingAppAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/io.goswagger.parking.v1+json":
			result["application/io.goswagger.parking.v1+json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ParkingAppAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the parking app API
func (o *ParkingAppAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ParkingAppAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/lot"] = lot.NewAddOne(o.context, o.LotAddOneHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/lot/{id}"] = lot.NewDestroyOne(o.context, o.LotDestroyOneHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/lot"] = lot.NewFindLots(o.context, o.LotFindLotsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/lot/status/{id}"] = lot.NewGetStstus(o.context, o.LotGetStstusHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/ticket/getTicket"] = ticket.NewGetTicket(o.context, o.TicketGetTicketHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/lot/{lotId}/ticket/{ticketId}/carLeaves"] = ticket.NewLeaveParking(o.context, o.TicketLeaveParkingHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/lot/{id}"] = lot.NewUpdateOne(o.context, o.LotUpdateOneHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ParkingAppAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *ParkingAppAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
