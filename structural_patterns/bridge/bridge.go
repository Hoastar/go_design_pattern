package bridge

// 通知接口
type IMsgSender interface {
	Send(msg string) error
}

// EmailMsgSender 发送邮件
// 另外还有电话、短信等各种实现
type EmailMsgSender struct {
	emails []string
}

// NewEmailMsgSender
func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{emails: emails}
}

// Send
func (s *EmailMsgSender) Send(msg string) error {
	// 具体实现
	return nil
}

// INotification 通知接口
type INotification interface {
	Notify(msg string) error
}

// ErrorNotification 错误通知
// 后面可能还有 warning 各种级别
type ErrorNotification struct {
	sender IMsgSender
}

// NewErrorNotification
func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}

// Notify 发送通知
func (n *ErrorNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}
