package middleware

import (
	"github.com/go-macaron/gzip"
	"gopkg.in/macaron.v1"
	"net/http"
)

const staticDir = "public"

var statikFS http.FileSystem


func RegisterMiddleware(m *macaron.Macaron)  {
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	if macaron.Env != macaron.DEV {
		m.Use(gzip.Gziper())
	}
	m.Use(macaron.Static("public"))
}