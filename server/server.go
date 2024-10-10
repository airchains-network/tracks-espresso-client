package server

import (
	"context"
	"github.com/airchains-network/tracks-espresso-client/client"
	"github.com/airchains-network/tracks-espresso-client/config"
	"github.com/airchains-network/tracks-espresso-client/database"
	"github.com/airchains-network/tracks-espresso-client/server/espresso"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	*gin.Engine
}

func InitServer(ctx context.Context, mongoDB *database.DB, rc *client.Client) *Server {
	gin.SetMode(config.GinEvn)
	ginServer := gin.Default()
	ginServer.Use(cors.New(
		cors.Config{
			AllowAllOrigins: true,
		},
	))

	pprof.Register(ginServer)

	ginCall := &Server{Engine: ginServer}

	ginCall.Handler(ctx, mongoDB, rc)

	return ginCall
}

func (s *Server) Handler(ctx context.Context, mongoDB *database.DB, rc *client.Client) {
	s.PostHandle("/track/espresso", espresso.TracksEspressoDataLoad(ctx, mongoDB))
}

func (s *Server) PostHandle(routes string, handler gin.HandlerFunc) {
	s.Engine.POST(routes, handler)
}

func (s *Server) GetHandle(routes string, handler gin.HandlerFunc) {
	s.Engine.GET(routes, handler)
}

// Home is a handler function for the home page
func (s *Server) Home(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/")
}

// RedirectToHome is a handler function to redirect all requests to the home page ("/")
func (s *Server) RedirectToHome(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/")
}
