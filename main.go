package main

import (
	"Tickermaster/pkg/config"
	"Tickermaster/pkg/constants"
	"Tickermaster/pkg/database"
	"Tickermaster/pkg/middleware"
	"Tickermaster/pkg/router"
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"time"
)

func main() {

	config := config.GetConfig()

	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	// Create context that listens for the interrupt signal from the OS.
	ctx, cancel = signal.NotifyContext(context.Background(), constants.SignalsToShutdown...)

	dbConfig := config.DB
	db := database.New(&database.DBOptions{
		Driver:        dbConfig.Driver,
		User:          dbConfig.User,
		Password:      dbConfig.Password,
		Host:          dbConfig.Host,
		Port:          dbConfig.Port,
		Name:          dbConfig.Name,
		SlowThreshold: dbConfig.SlowThreshold,
		Colorful:      dbConfig.Log.Colorful,
	})

	r := router.New(db)

	versionGroup := r.Group("/v1")

	couponGroup := versionGroup.Group("/coupons")
	couponGroup.Use(middleware.VerifyToken)
	couponGroup.GET("/list", router.CouponList)
	couponGroup.POST("/register", router.CouponRegister)
	couponGroup.POST("/grab", router.CouponGrab)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port),
		Handler: r,
	}

	// Wait for the server to stop.
	<-ctx.Done()
	cancel()
	// graceful shutdown
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
