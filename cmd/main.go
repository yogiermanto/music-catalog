package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"music-catalog/internal/configs"
	"music-catalog/pkg/internalsql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := configs.Init()
	if err != nil {
		log.Panic().Err(err).Msg("Failed to initialize configs")
	}
	conf := configs.Get()

	db, err := internalsql.Connect(conf.Database.DSN)
	if err != nil {
		log.Panic().Err(err).Msg("Failed to connect to database")
	}

	fmt.Println(db)

	router := gin.Default()
	router.GET("/api/health", func(c *gin.Context) {
		s, errDb := db.DB()
		if errDb != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"db": false,
			})
			return
		}

		errPing := s.Ping()
		if errPing != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"db": false,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"db": true,
		})
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Handler(),
	}
	log.Info().Msg("Listen and serve port :8080")

	gracefullyShutdown(srv)
}

func gracefullyShutdown(srv *http.Server) {
	go func() {
		// service connection
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) && err != nil {
			log.Fatal().Err(err).Msg("Failed to listen and server")
		}
	}()

	// wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server Shutdown")
	}

	// catching ctx.Done(), timeout of 15 seconds
	select {
	case <-ctx.Done():
		log.Info().Msg("timeout of 15 seconds.")
	default:
		log.Info().Msg("waiting to gracefully shutdown")
	}
	log.Info().Msg("Server exiting")
}
