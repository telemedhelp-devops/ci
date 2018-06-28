package sms

import (
	"github.com/goodsign/gosmsc"
)

type goSMSC struct {
	impl *gosmsc.SenderFetcherImpl
}

func InitGoSMSC(login, password string) (err error) {
	var goSMSC goSMSC
	goSMSC.impl, err = gosmsc.NewSenderFetcherImpl(&gosmsc.SmscClientOptions{
		User: login,
		Password: password,
	})
	if err != nil {
		return
	}
	sms = &goSMSC
	return
}

func (sms goSMSC) Send(sender, destination, message string) error {
	_, err := sms.impl.Send(destination, message)
	return err
}
