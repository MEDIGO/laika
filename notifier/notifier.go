package notifier

// Notifier describes a notification sender.
type Notifier interface {
	// NotifyStatusChange notifies a change in the status of a flag.
	NotifyStatusChange(feature string, status bool, environment string) error
}
