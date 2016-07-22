package notifier

import (
	"github.com/nlopes/slack"
)

// SlackNotifier is a notifier that send messages to Slack.
type SlackNotifier struct {
	client  *slack.Client
	channel string
}

// NewSlackNotifier creates a new SlackNotifier
func NewSlackNotifier(token, channel string) Notifier {
	return &SlackNotifier{slack.New(token), channel}
}

// NotifyStatusChange notifies a change in the status of a flag.
func (n *SlackNotifier) NotifyStatusChange(feature string, status bool, environment string) error {
	text := "disabled"
	color := "#e74c3c"

	if status {
		text = "enabled"
		color = "#27ae60"
	}

	attachment := slack.Attachment{
		Title: "Laika Flag Update!",
		Color: color,
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Flag",
				Value: feature,
				Short: false,
			},
			slack.AttachmentField{
				Title: "Environment",
				Value: environment,
				Short: true,
			},
			slack.AttachmentField{
				Title: "Status Change",
				Value: text,
				Short: true,
			},
		},
	}

	_, _, err := n.client.PostMessage(n.channel, "WOOF! WOFF! ARH-WOOOOOOOO!", slack.PostMessageParameters{
		AsUser:      true,
		Attachments: []slack.Attachment{attachment},
	})
	return err
}
