package notification

type Notification interface {
	Send(from string, to []string, subject string, message string) error
}

func NewNotification(strategy Notification) Notification {
	if strategy == nil {
		panic("notification strategy is required")
	}

	return strategy
}
