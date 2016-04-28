package notifier

import (
	"github.com/nlopes/slack"
)

type SlackNotifier struct {
	token   string
	channel string
}

func NewSlackNotifier(token, channel string) *SlackNotifier {
	return &SlackNotifier{token, channel}
}

func (n *SlackNotifier) NotifyStatusChange(featureName string, status bool, environmentName string) error {
	statusChange := "disabled"
	color := "#e74c3c"
	if status {
		statusChange = "enabled"
		color = "#27ae60"
	}

	api := slack.New(n.token)
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Title: "Laika Flag Update!",
		Color: color,
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Flag",
				Value: featureName,
				Short: false,
			},
			slack.AttachmentField{
				Title: "Environment",
				Value: environmentName,
				Short: true,
			},
			slack.AttachmentField{
				Title: "Status Change",
				Value: statusChange,
				Short: true,
			},
		},
	}

	params.Attachments = []slack.Attachment{attachment}

	_, _, err := api.PostMessage(n.channel, "WOOF! WOFF! ARH-WOOOOOOOO!", params)
	return err
}
