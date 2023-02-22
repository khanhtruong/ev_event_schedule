package handler

import (
	"event_schedule/data/dto"
	"event_schedule/data/entity"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) Subscribe(c *gin.Context) {
	subscribeReq := dto.SubscribeDto{}
	if err := c.ShouldBindJSON(&subscribeReq); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"message": err.Error(),
			},
		})
		return
	}

	userID := c.Request.Header.Get("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"message": "UserID header is nil, please try again",
			},
		})
		return
	}
	uID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"message": "UserID is not an integer",
			},
		})
		return
	}

	// Create new subscription
	sub, err := h.Repo.Subscribe(&entity.Subscriptions{
		ID:                   uuid.New(),
		UserID:               uint(uID),
		LocalEventCalendarID: "",
		EventSessionID:       subscribeReq.Data.EventSessionID,
		Status:               entity.Pending,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"message": err.Error(),
			},
		})
		return
	}

	// Preload data
	if sub == nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"message": "Insert subcription failed",
			},
		})
		return
	}

	result, err := h.Repo.GetSubscription(fmt.Sprint(sub.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"message": "Fetch subscription failed",
			},
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
}

func (h *Handler) UpdateSubscribe(c *gin.Context) {
	subscribeReq := dto.SubscribeDto{}
	if err := c.ShouldBindJSON(&subscribeReq); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"message": err.Error(),
			},
		})
		return
	}

	subscriptionID := c.Param("subscription_id")
	if subscriptionID == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"message": "SubscriptionID is nil, please try again",
			},
		})
		return
	}

	_, err := h.Repo.UpdateSubscription(subscriptionID, fmt.Sprint(subscribeReq.Data.EventSessionID))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"message": err.Error(),
			},
		})
		return
	}

	result, err := h.Repo.GetSubscription(subscriptionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"message": "Fetch subscription failed",
			},
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
}

func (h *Handler) DeleteSubscribe(c *gin.Context) {
	subscriptionID := c.Param("subscription_id")
	if subscriptionID == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"message": "SubscriptionID is nil, please try again",
			},
		})
		return
	}

	if err := h.Repo.DeleteSubscription(subscriptionID); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"message": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"status": "OK",
		},
	})
}

func (h *Handler) UpdateSubscribeCalendarID(c *gin.Context) {
	subscribeReq := dto.SubscribeCalendarIDDto{}
	if err := c.ShouldBindJSON(&subscribeReq); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"message": err.Error(),
			},
		})
		return
	}

	subscriptionID := c.Param("subscription_id")
	if subscriptionID == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"message": "SubscriptionID is nil, please try again",
			},
		})
		return
	}

	_, err := h.Repo.UpdateSubscriptionCalendarID(subscriptionID, subscribeReq.Data.LocalEventCalendarID)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"message": "SubscriptionID is nil, please try again",
			},
		})
		return
	}

	result, err := h.Repo.GetSubscription(subscriptionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"message": "Fetch subscription failed",
			},
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
}
