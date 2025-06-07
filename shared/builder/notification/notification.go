package notification

type Notification interface {
	To(to string) Notification
	From(from string) Notification
	Cc(cc ...string) Notification
	Bcc(bcc ...string) Notification
	Subject(subject string) Notification
	Body(body string) Notification
	BodyHTML(body string) Notification
	Attachments(attachments ...string) Notification
	Send() error
}

func NewNotification(strategy Notification) Notification {
	if strategy == nil {
		panic("notification strategy cannot be nil")
	}

	return strategy
}
