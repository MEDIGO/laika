package notifier

import "github.com/lytics/slackhook"

// SlackNotifier is a notifier that send messages to Slack.
type SlackNotifier struct {
	client *slackhook.Client
}

// NewSlackNotifier creates a new SlackNotifier.
func NewSlackNotifier(url string) Notifier {
	return &SlackNotifier{slackhook.New(url)}
}

// NotifyStatusChange notifies a change in the status of a flag.
func (n *SlackNotifier) NotifyStatusChange(feature string, status bool, environment string) error {
	text := "disabled"
	color := "#e74c3c"

	if status {
		text = "enabled"
		color = "#27ae60"
	}

	return n.client.Send(&slackhook.Message{
		Text: "WOOF! WOFF! ARH-WOOOOOOOO!",
		Attachments: []*slackhook.Attachment{
			&slackhook.Attachment{
				Title: "Laika Flag Update!",
				Color: color,
				Fields: []slackhook.Field{
					slackhook.Field{
						Title: "Flag",
						Value: feature,
						Short: false,
					},
					slackhook.Field{
						Title: "Environment",
						Value: environment,
						Short: true,
					},
					slackhook.Field{
						Title: "Status",
						Value: text,
						Short: true,
					},
				},
			},
		},
	})
}
