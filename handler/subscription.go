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
				"errorCode": dto.ERR_INVALID_REQUEST_BODY,
				"errorMsg":  err.Error(),
			},
		})
		return
	}

	userID := c.Request.Header.Get("access_token")
	if userID == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INVALID_ACCESS_TOKEN,
				"errorMsg":  "AccessToken is nil, please try again",
			},
		})
		return
	}
	uID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INVALID_ACCESS_TOKEN,
				"errorMsg":  "UserID is not an integer",
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
				"errorCode": dto.ERR_INTERNAL_SERVER_ERROR,
				"errorMsg":  err.Error(),
			},
		})
		return
	}

	// Preload data
	if sub == nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INTERNAL_SERVER_ERROR,
				"errorMsg":  "Insert subcription failed",
			},
		})
		return
	}

	result, err := h.Repo.GetSubscription(fmt.Sprint(sub.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INTERNAL_SERVER_ERROR,
				"errorMsg":  "Fetch subscription failed",
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
				"errorCode": dto.ERR_INVALID_REQUEST_BODY,
				"errorMsg":  err.Error(),
			},
		})
		return
	}

	subscriptionID := c.Param("subscription_id")
	if subscriptionID == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INVALID_PATH_PARAMS,
				"errorMsg":  "SubscriptionID is nil, please try again",
			},
		})
		return
	}

	userID := c.Request.Header.Get("access_token")
	if userID == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INVALID_ACCESS_TOKEN,
				"errorMsg":  "AccessToken is nil, please try again",
			},
		})
		return
	}
	uID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INVALID_ACCESS_TOKEN,
				"errorMsg":  "UserID is not an integer",
			},
		})
		return
	}

	// Check ownership of this subscription_id
	// If the subscription is not belongs to user request
	// Cancel the request
	sub, err := h.Repo.GetSubscription(subscriptionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INTERNAL_SERVER_ERROR,
				"errorMsg":  "Unable to retreived subscription data",
			},
		})
		return
	}
	if sub.UserID != uint(uID) {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INTERNAL_SERVER_ERROR,
				"errorMsg":  "Subscription not found",
			},
		})
		return
	}

	_, err = h.Repo.UpdateSubscription(subscriptionID, fmt.Sprint(subscribeReq.Data.EventSessionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INTERNAL_SERVER_ERROR,
				"errorMsg":  err.Error(),
			},
		})
		return
	}

	result, err := h.Repo.GetSubscription(subscriptionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INTERNAL_SERVER_ERROR,
				"errorMsg":  "Fetch subscription failed",
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
				"errorCode": dto.ERR_INVALID_PATH_PARAMS,
				"errorMsg":  "SubscriptionID is nil, please try again",
			},
		})
		return
	}

	userID := c.Request.Header.Get("access_token")
	if userID == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INVALID_ACCESS_TOKEN,
				"errorMsg":  "AccessToken is nil, please try again",
			},
		})
		return
	}
	uID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INVALID_ACCESS_TOKEN,
				"errorMsg":  "UserID is not an integer",
			},
		})
		return
	}

	// Check ownership of this subscription_id
	// If the subscription is not belongs to user request
	// Cancel the request
	sub, err := h.Repo.GetSubscription(subscriptionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INTERNAL_SERVER_ERROR,
				"errorMsg":  "Invalid subscription_id, please try again",
			},
		})
		return
	}
	if sub.UserID != uint(uID) {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INTERNAL_SERVER_ERROR,
				"errorMsg":  "Invalid subscription_id, please try again",
			},
		})
		return
	}

	if err := h.Repo.DeleteSubscription(subscriptionID); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INTERNAL_SERVER_ERROR,
				"errorMsg":  err.Error(),
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
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INVALID_REQUEST_BODY,
				"errorMsg":  err.Error(),
			},
		})
		return
	}

	subscriptionID := c.Param("subscription_id")
	if subscriptionID == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INVALID_PATH_PARAMS,
				"errorMsg":  "SubscriptionID is nil, please try again",
			},
		})
		return
	}

	userID := c.Request.Header.Get("access_token")
	if userID == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INVALID_ACCESS_TOKEN,
				"errorMsg":  "AccessToken is nil, please try again",
			},
		})
		return
	}
	uID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INVALID_ACCESS_TOKEN,
				"errorMsg":  "UserID is not an integer",
			},
		})
		return
	}

	// Check ownership of this subscription_id
	// If the subscription is not belongs to user request
	// Cancel the request
	sub, err := h.Repo.GetSubscription(subscriptionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INTERNAL_SERVER_ERROR,
				"errorMsg":  "Invalid subscription_id, please try again",
			},
		})
		return
	}
	if sub.UserID != uint(uID) {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INTERNAL_SERVER_ERROR,
				"errorMsg":  "Invalid subscription_id, please try again",
			},
		})
		return
	}

	_, err = h.Repo.UpdateSubscriptionCalendarID(subscriptionID, subscribeReq.Data.LocalEventCalendarID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INTERNAL_SERVER_ERROR,
				"errorMsg":  "SubscriptionID is nil, please try again",
			},
		})
		return
	}

	result, err := h.Repo.GetSubscription(subscriptionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INTERNAL_SERVER_ERROR,
				"errorMsg":  "Fetch subscription failed",
			},
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
}

func (h *Handler) SubscribedEvents(c *gin.Context) {
	userID := c.Request.Header.Get("access_token")
	if userID == "" {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INVALID_ACCESS_TOKEN,
				"errorMsg":  "AccessToken is nil, please try again",
			},
		})
		return
	}
	uID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INVALID_ACCESS_TOKEN,
				"errorMsg":  "UserID is not an integer",
			},
		})
		return
	}

	result, err := h.Repo.GetSubscriptions(uint(uID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": map[string]interface{}{
				"errorCode": dto.ERR_INTERNAL_SERVER_ERROR,
				"errorMsg":  "Fetch subscription failed",
			},
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
}
