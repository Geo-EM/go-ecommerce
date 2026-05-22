package notification

import (
	"e-commerce/config"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type NotificationClient interface {
	SendSms(phoneNumber string, message string) error
}
type notificationClient struct {
	appConfig config.AppConfig
}

func (c notificationClient) SendSms(phoneNumber string, message string) error {
	accountSid := c.appConfig.TwilioAccountSID
	authToken := c.appConfig.TwilioAuthToken

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phoneNumber)
	params.SetFrom(c.appConfig.TwilioFromPhoneNumber) // from Twilio number
	params.SetBody(message)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		return err
	}
	// } else {
	// 	response, _ := json.Marshal(*resp)
	// 	fmt.Println("Response: " + string(response))
	// }

	return nil
}

func NewNotificationClient(appConfig config.AppConfig) NotificationClient {
	return &notificationClient{appConfig: appConfig}
}
