package controller

import (
	"errors"

	"github.com/Nishad4140/ecommerce_project/pkg/config"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

func SendOTP(phno string) error {
	cfg, _ := config.LoadConfig()
	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.TWILIOACCOUNTSID,
		Password: cfg.TWILIOAUTHTOKEN,
	})
	params := &openapi.CreateVerificationParams{}
	params.SetTo("+91 " + phno)
	params.SetChannel("sms")
	_, err := client.VerifyV2.CreateVerification(cfg.TWILIOSERVICESID, params)
	if err != nil {
		return errors.New("can't send the otp")
	}
	return nil
}

func VerifyOTP(otp, phno string) error {
	var cfg config.Config

	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.TWILIOACCOUNTSID,
		Password: cfg.TWILIOAUTHTOKEN,
	})
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo("+91 " + phno)
	params.SetCode(otp)
	resp, err := client.VerifyV2.CreateVerificationCheck(cfg.TWILIOSERVICESID, params)
	if err != nil {
		return errors.New("not vrified")
	} else if *resp.Status != "approved" {
		return errors.New("not vrified")
	}
	return nil
}
