package middleware

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mehrdadnekopour/go-tools/templates"
)

// Interface ...
type Interface struct {
	// another stuff , may be needed by middleware
}

// Init ...
func Init() *Interface {
	return &Interface{}
}

// NewProxyTargets ...
func (i *Interface) NewProxyTargets(urls ...*url.URL) (proxyTargets []*middleware.ProxyTarget) {
	c := len(urls)

	for i := 0; i < c; i++ {
		url := urls[i]
		target := &middleware.ProxyTarget{
			URL: url,
		}

		proxyTargets = append(proxyTargets, target)
	}

	return
}

// CORS ...
func (i *Interface) CORS() echo.MiddlewareFunc {
	return middleware.CORS()
}

// Secure ...
func (i *Interface) Secure() echo.MiddlewareFunc {
	return middleware.Secure()
}

// Proxy ...
func (i *Interface) Proxy(targets []*middleware.ProxyTarget) echo.MiddlewareFunc {
	return middleware.Proxy(middleware.NewRoundRobinBalancer(targets))
}

// RemoveTrailingSlash ...
func (i *Interface) RemoveTrailingSlash() echo.MiddlewareFunc {
	return middleware.RemoveTrailingSlash()
}

// AddTrailingSlash ...
func (i *Interface) AddTrailingSlash() echo.MiddlewareFunc {
	return middleware.AddTrailingSlash()
}

// LoggerWithConfig ...
func (i *Interface) LoggerWithConfig() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	})
}

//JWT ...
func (i *Interface) JWT(key string) echo.MiddlewareFunc {

	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(key),
		Skipper:    i.RegisterSkipper,
	})
}

// RegisterSkipper returns true if path =  /api/v1/register
func (i *Interface) RegisterSkipper(ctx echo.Context) bool {
	if ctx.Request().URL.String() == "/api/v1/auth/login" && ctx.Request().Method == "POST" {
		return true
	}

	// if ctx.Path() == "/api/v1/auth/authorize" && ctx.Request().Method == "GET" {
	// 	return true
	// }
	return false
}

// HTTPErrorHandler ...
func HTTPErrorHandler(err error, ctx echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	// ctx.Logger().Error(err)

	template := templates.GetWithCode(code, err)

	ctx.JSON(code, template)
}
