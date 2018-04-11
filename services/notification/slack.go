package notification

import (
	"fmt"
	"golang-firestore/config"

	slack "github.com/ashwanthkumar/slack-go-webhook"
)

/**
 * @brief      Sends a notification error.
 * @param      err   The error
 */
func SendNotificationError(err error) {
	if config.Env("APP_ENV") != "local" {
		data_error := fmt.Sprintf("==> Error: %s\n", err.Error())
		attachment1 := slack.Attachment{}
		attachment1.AddField(slack.Field{Title: "Environment", Value: config.Env("APP_ENV")}).AddField(slack.Field{Title: "Error", Value: data_error})
		payload := slack.Payload{
			Text:        "Hello Error from Qasir Firestore",
			Username:    "Qasir Firestore Error",
			Channel:     config.Env("SLACK_CHANNEL"),
			IconEmoji:   ":qasir:",
			Attachments: []slack.Attachment{attachment1},
		}
		data_err := slack.Send(config.Env("SLACK_WEBHOOK"), "", payload)
		if len(data_err) > 0 {
			fmt.Printf("error: %s\n", data_err)
		}
	}
}
