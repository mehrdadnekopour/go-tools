package router

import (
	"github.com/labstack/echo"
)

// MainGroupV1 Represents the main route groups for api version 1
var MainGroupV1 *echo.Group

//AssetsGroup .... /assets
var AssetsGroup *echo.Group

// RoutesGroup represent a group for a single resource
type RoutesGroup struct {
	GroupPattern string
	Routes       Routes
}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc echo.HandlerFunc
}

// Routes is an array of route
type Routes []Route

// New ...
func (rg *RoutesGroup) New() {
	group := MainGroupV1.Group(rg.GroupPattern)

	for _, route := range rg.Routes {
		var handler echo.HandlerFunc
		handler = route.HandlerFunc
		// handler = logger.Logger(handler, route.Name)

		group.Add(route.Method, route.Pattern, handler)
		// r.Add(route.Method, route.Pattern, handler)
	}
}
