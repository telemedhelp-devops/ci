package sms

type SMSGW interface {
	Send(sender, destination, message string) error
}

var sms SMSGW

func Send(sender, destination, message string) error {
	return sms.Send(sender, destination, message)
}
