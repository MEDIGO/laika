package notifier

type Notifier interface {
	NotifyStatusChange(featureName string, status bool, environmentName string) error
}
