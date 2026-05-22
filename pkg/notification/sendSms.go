package notification

import "e-commerce/config"

type NotificationClient interface {
	SendSms(phoneNumber string, message string) error
}
type notificationClient struct {
	appConfig config.AppConfig
}

func (c notificationClient) SendSms(phoneNumber string, message string) error {
	return nil
}

func NewNotificationClient(appConfig config.AppConfig) NotificationClient {
	return &notificationClient{appConfig: appConfig}
}
