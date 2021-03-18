package GmailClient

func Send(m Mail) error {
	config := newGmailConfig()
	return send(m, config)
}