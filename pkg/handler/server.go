package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "leadgen/docs"
	"leadgen/pkg/service"
	"net/http"
	"sync"
)

type ApiHandler struct {
	engine             *gin.Engine
	log                *logrus.Logger
	serv               *http.Server
	httpServerExitDone *sync.WaitGroup

	s    *service.S
	port int
}

func NewApiHandler(
	engine *gin.Engine,
	log *logrus.Logger,
	s *service.S,
	port int,
) *ApiHandler {
	api := &ApiHandler{
		engine:             engine,
		log:                log,
		httpServerExitDone: &sync.WaitGroup{},
		s:                  s,
		port:               port,
	}

	api.HandleRoutes()

	return api
}

func (h *ApiHandler) HandleRoutes() {
	h.engine.Use(gin.Recovery())
	h.engine.Use(CORSMiddleware())

	h.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h.engine.POST("/api/buildings", h.CreateBuilding)
	h.engine.GET("/api/buildings", h.ListBuildings)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Обработка OPTIONS-запроса
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (h *ApiHandler) Run() error {
	h.serv = &http.Server{
		Addr:    fmt.Sprintf(":%d", h.port),
		Handler: h.engine,
	}
	h.httpServerExitDone.Add(1)

	go func() {
		defer h.httpServerExitDone.Done()

		// always returns error. ErrServerClosed on graceful close
		if err := h.serv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			h.log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	h.log.Info("HTTP server started")
	return nil
}

func (h *ApiHandler) Shutdown() error {
	err := h.serv.Shutdown(context.Background())
	if err != nil {
		return err
	}
	h.httpServerExitDone.Wait()
	h.log.Info("HTTP server stoped")
	return nil
}
