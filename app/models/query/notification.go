package query

import (
	"github.com/Windscribe/fider/app/models"
	"github.com/Windscribe/fider/app/models/enum"
)

type CountUnreadNotifications struct {
	Result int
}

type GetNotificationByID struct {
	ID     int
	Result *models.Notification
}

type GetActiveNotifications struct {
	Result []*models.Notification
}

type GetActiveSubscribers struct {
	Number  int
	Channel enum.NotificationChannel
	Event   enum.NotificationEvent

	Result []*models.User
}
