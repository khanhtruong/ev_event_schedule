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

	r.GET("/events", handler.Events)
	r.GET("/events/:event_id", handler.EventDetails)

	r.POST("/subscriptions", handler.Subscribe)
	r.PUT("/subscriptions/:subscription_id", handler.UpdateSubscribe)
	r.DELETE("/subscriptions/:subscription_id", handler.DeleteSubscribe)
	r.PUT("/subscriptions/:subscription_id/localcalendar", handler.UpdateSubscribeCalendarID)
}
