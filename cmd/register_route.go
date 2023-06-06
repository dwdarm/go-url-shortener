package cmd

import (
	"github.com/dwdarm/go-url-shortener/src/registry"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(ctn *registry.Container, r *gin.Engine) {
	linkHandler := ctn.GetLinkHandler()
	r.POST("api/links/", linkHandler.LinkCreate)
	r.GET("api/links/:slug", linkHandler.LinkGet)
}
