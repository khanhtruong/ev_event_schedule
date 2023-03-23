package main

import (
	"event_schedule/config"
	"event_schedule/handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	config := config.LoadConfig()

	r := gin.Default()
	MountHandler(r, &config)

	r.Run(fmt.Sprintf(":%s", config.Port))
}

func MountHandler(r *gin.Engine, config *config.Config) {
	handler := handler.NewHandler(config)

	// Event handlers
	r.GET("/event", handler.Events)
	r.GET("/event/:event_id", handler.EventDetails)

	// Subscription handlers
	r.POST("/subscription", handler.Subscribe)
	r.PUT("/subscription/:subscription_id", handler.UpdateSubscribe)
	r.DELETE("/subscription/:subscription_id", handler.DeleteSubscribe)
	r.PUT("/subscription/:subscription_id/localcalendar", handler.UpdateSubscribeCalendarID)
	r.GET("/subscription/event", handler.SubscribedEvents)
}
