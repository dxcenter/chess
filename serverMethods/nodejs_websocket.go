package serverMethods

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func ProxyToNodejs(c *gin.Context) {
	director := func(req *http.Request) {
		r := c.Request
		req = r
		req.URL.Scheme = "http"
		req.URL.Host = "localhost:3000"
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(c.Writer, c.Request)
}
