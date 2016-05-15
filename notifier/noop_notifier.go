package notifier

// NOOPNotifier is a notifier that doesn't perform any notifications.
type NOOPNotifier struct{}

// NewNOOPNotifier creates a new NOOPNotifier.
func NewNOOPNotifier() Notifier {
	return &NOOPNotifier{}
}

// NotifyStatusChange doesn't perform any operation and always returns nil.
func (n *NOOPNotifier) NotifyStatusChange(featureName string, status bool, environmentName string) error {
	return nil
}
