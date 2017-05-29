package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"
	"github.com/go-openapi/swag"

	"github.com/user/go-parking/restapi/operations"
	"github.com/user/go-parking/restapi/operations/lot"
	"github.com/user/go-parking/restapi/operations/ticket"
	"github.com/user/go-parking/service"
	"github.com/user/go-parking/models"
)

// This file is safe to edit. Once it exists it will not be overwritten

var lotService service.LotService


func configureFlags(api *operations.ParkingAppAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ParkingAppAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	lotService := service.LotService{}
	lotService.Init()

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.LotAddOneHandler = lot.AddOneHandlerFunc(func(params lot.AddOneParams) middleware.Responder {
		if err := lotService.AddLot(params.Body); err != nil {
			return lot.NewAddOneDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return lot.NewAddOneCreated().WithPayload(params.Body)
	})
	api.LotDestroyOneHandler = lot.DestroyOneHandlerFunc(func(params lot.DestroyOneParams) middleware.Responder {
		if err := lotService.DeleteLot(params.ID); err != nil {
			return lot.NewDestroyOneDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return lot.NewDestroyOneNoContent()
	})
	api.LotFindLotsHandler = lot.FindLotsHandlerFunc(func(params lot.FindLotsParams) middleware.Responder {
		mergedParams := lot.NewFindLotsParams()
		mergedParams.Since = swag.Int64(0)
		if params.Since != nil {
			mergedParams.Since = params.Since
		}
		if params.Limit != nil {
			mergedParams.Limit = params.Limit
		}
		return lot.NewFindLotsOK().WithPayload(lotService.AllLots(*mergedParams.Since, *mergedParams.Limit))
	})
	api.LotGetStstusHandler = lot.GetStstusHandlerFunc(func(params lot.GetStstusParams) middleware.Responder {
		return middleware.NotImplemented("operation lot.GetStstus has not yet been implemented")
	})
	api.TicketGetTicketHandler = ticket.GetTicketHandlerFunc(func(params ticket.GetTicketParams) middleware.Responder {
		var res, err = lotService.GetTicket(params.LotID, params.CarSize)
		if err != nil {
			return lot.NewAddOneDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return ticket.NewGetTicketOK().WithPayload(res)
	})
	api.TicketLeaveParkingHandler = ticket.LeaveParkingHandlerFunc(func(params ticket.LeaveParkingParams) middleware.Responder {
		var res = lotService.CarLeaves(params.LotID, params.TicketID)
		return ticket.NewLeaveParkingOK().WithPayload(res)
	})
	api.LotUpdateOneHandler = lot.UpdateOneHandlerFunc(func(params lot.UpdateOneParams) middleware.Responder {
		return middleware.NotImplemented("operation lot.UpdateOne has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
