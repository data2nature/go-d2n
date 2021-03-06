// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Aborts a catchpoint catchup.
	// (DELETE /v2/catchup/{catchpoint})
	AbortCatchup(ctx echo.Context, catchpoint string) error
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error
	// Return a list of participation keys
	// (GET /v2/participation)
	GetParticipationKeys(ctx echo.Context) error
	// Add a participation key to the node
	// (POST /v2/participation)
	AddParticipationKey(ctx echo.Context) error
	// Delete a given participation key by ID
	// (DELETE /v2/participation/{participation-id})
	DeleteParticipationKeyByID(ctx echo.Context, participationId string) error
	// Get participation key info given a participation ID
	// (GET /v2/participation/{participation-id})
	GetParticipationKeyByID(ctx echo.Context, participationId string) error
	// Append state proof keys to a participation key
	// (POST /v2/participation/{participation-id})
	AppendKeys(ctx echo.Context, participationId string) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AbortCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) AbortCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AbortCatchup(ctx, catchpoint)
	return err
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// GetParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeys(ctx)
	return err
}

// AddParticipationKey converts echo context to params.
func (w *ServerInterfaceWrapper) AddParticipationKey(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddParticipationKey(ctx)
	return err
}

// DeleteParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteParticipationKeyByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteParticipationKeyByID(ctx, participationId)
	return err
}

// GetParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeyByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeyByID(ctx, participationId)
	return err
}

// AppendKeys converts echo context to params.
func (w *ServerInterfaceWrapper) AppendKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AppendKeys(ctx, participationId)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty":  true,
		"timeout": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------
	if paramValue := ctx.QueryParam("timeout"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE("/v2/catchup/:catchpoint", wrapper.AbortCatchup, m...)
	router.POST("/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.GET("/v2/participation", wrapper.GetParticipationKeys, m...)
	router.POST("/v2/participation", wrapper.AddParticipationKey, m...)
	router.DELETE("/v2/participation/:participation-id", wrapper.DeleteParticipationKeyByID, m...)
	router.GET("/v2/participation/:participation-id", wrapper.GetParticipationKeyByID, m...)
	router.POST("/v2/participation/:participation-id", wrapper.AppendKeys, m...)
	router.POST("/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
