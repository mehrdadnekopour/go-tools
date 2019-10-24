package helpers

import (
	"net/url"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewProxyTargets ...
func NewProxyTargets(urls ...*url.URL) (proxyTargets []*middleware.ProxyTarget) {
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

// Proxy ...
func Proxy(targets []*middleware.ProxyTarget) echo.MiddlewareFunc {
	return middleware.Proxy(middleware.NewRoundRobinBalancer(targets))
}
