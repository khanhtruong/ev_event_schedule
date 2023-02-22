package handler

import (
	"event_schedule/data/entity"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Events(c *gin.Context) {
	data, err := h.Repo.GetEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"message": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

func (h *Handler) EventDetails(c *gin.Context) {
	eventID := c.Param("event_id")
	if eventID == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"message": "event_id not found",
			},
		})
		return
	}

	data, err := h.Repo.GetEvent(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"message": err.Error(),
			},
		})
		return
	}

	sessionsFullFilled := []entity.EventSessionInDetail{}
	for i := range data.EventSessions {
		sessionID := fmt.Sprint(data.EventSessions[i].ID)
		count, err := h.Repo.SubscriptionAmount(sessionID)
		if err != nil {
			fmt.Printf("Unable to get subscription amount for session: %s, error: %s\n", sessionID, err.Error())
			continue
		}
		sessionsFullFilled = append(sessionsFullFilled, entity.EventSessionInDetail{EventSession: data.EventSessions[i], Available: count})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": entity.EventInDetail{
			ID:            data.ID,
			Title:         data.Title,
			Description:   data.Description,
			Location:      data.Location,
			Type:          data.Type,
			EventSessions: sessionsFullFilled,
		},
	})
}
