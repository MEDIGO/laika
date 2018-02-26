package notifier

import (
	"fmt"

	"github.com/vsco/slackhook"
)

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
	text := fmt.Sprintf("Feature *%s* is now %s in *%s*.", feature, label(status), environment)

	return n.client.Send(&slackhook.Message{
		Attachments: []*slackhook.Attachment{
			&slackhook.Attachment{
				Text:       text,
				Color:      color(status),
				MarkdownIn: []string{"text"},
			},
		},
	})
}

func color(status bool) string {
	if status {
		return "good"
	}
	return "danger"
}

func label(status bool) string {
	if status {
		return "enabled"
	}
	return "disabled"
}
