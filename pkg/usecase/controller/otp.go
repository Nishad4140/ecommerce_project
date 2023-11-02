package controller

import (
	"errors"
	"fmt"

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

func VerifyOTP(otp, phno string) (*openapi.VerifyV2VerificationCheck, error) {
	var cfg config.Config

	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.TWILIOACCOUNTSID,
		Password: cfg.TWILIOAUTHTOKEN,
	})
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo("+91 " + phno)
	fmt.Println(otp)
	params.SetCode(otp)
	resp, err := client.VerifyV2.CreateVerificationCheck(cfg.TWILIOSERVICESID, params)
	return resp, err
}
