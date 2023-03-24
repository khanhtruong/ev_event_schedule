package handler

import (
	"event_schedule/data/dto"
	"event_schedule/data/entity"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Events(c *gin.Context) {
	filter := c.Query("filter")
	from := c.Query("from")
	to := c.Query("to")

	filters := strings.Split(filter, ",")
	if len(filters) == 0 {
		filters = entity.ALL_FILTER
	}

	data, err := h.Repo.GetEvents(filters, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INTERNAL_SERVER_ERROR,
				"errorMsg":  err.Error(),
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
				"errorCode": dto.ERR_INVALID_PATH_PARAMS,
				"errorMsg":  "event_id not found",
			},
		})
		return
	}

	data, err := h.Repo.GetEvent(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INTERNAL_SERVER_ERROR,
				"errorMsg":  err.Error(),
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
