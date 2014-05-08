package notifier

import ("code.google.com/p/goexmpp")

func Send(from string, to []string, text string, messageType string) {
	m := &xmpp.Message{
		From: from,
		To:   to,
		Body: text,
		Type: messageType,
	}
	err := m.Send(c)
}
