// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Validations for urn:tch
package messages

import (
	"github.com/moov-io/base"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD Element validations

func (v Message) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	rtp.AddError(&errs, v.AppHdr.Validate())
	if v.CreditTransfer != nil {
		rtp.AddError(&errs, v.CreditTransfer.Validate())
	}
	if v.MessageStatusReport != nil {
		rtp.AddError(&errs, v.MessageStatusReport.Validate())
	}
	if v.FICreditTransfer != nil {
		rtp.AddError(&errs, v.FICreditTransfer.Validate())
	}
	if v.Acknowledgement != nil {
		rtp.AddError(&errs, v.Acknowledgement.Validate())
	}
	if v.ResponseRequestForInformation != nil {
		rtp.AddError(&errs, v.ResponseRequestForInformation.Validate())
	}
	if v.RequestForInformation != nil {
		rtp.AddError(&errs, v.RequestForInformation.Validate())
	}
	if v.ReturnOfFunds != nil {
		rtp.AddError(&errs, v.ReturnOfFunds.Validate())
	}
	if v.PaymentRequest != nil {
		rtp.AddError(&errs, v.PaymentRequest.Validate())
	}
	if v.ResponsePaymentRequest != nil {
		rtp.AddError(&errs, v.ResponsePaymentRequest.Validate())
	}
	if v.ResponseReturnOfFunds != nil {
		rtp.AddError(&errs, v.ResponseReturnOfFunds.Validate())
	}
	if v.EchoRequest != nil {
		rtp.AddError(&errs, v.EchoRequest.Validate())
	}
	if v.EchoResponse != nil {
		rtp.AddError(&errs, v.EchoResponse.Validate())
	}
	if v.SignOffRequest != nil {
		rtp.AddError(&errs, v.SignOffRequest.Validate())
	}
	if v.SignOffResponse != nil {
		rtp.AddError(&errs, v.SignOffResponse.Validate())
	}
	if v.SignOnRequest != nil {
		rtp.AddError(&errs, v.SignOnRequest.Validate())
	}
	if v.SignOnResponse != nil {
		rtp.AddError(&errs, v.SignOnResponse.Validate())
	}
	if v.StandaloneRemittance != nil {
		rtp.AddError(&errs, v.StandaloneRemittance.Validate())
	}
	if v.SystemNotificationEvent != nil {
		rtp.AddError(&errs, v.SystemNotificationEvent.Validate())
	}
	if v.MessageReject != nil {
		rtp.AddError(&errs, v.MessageReject.Validate())
	}
	if v.TokenIdentification != nil {
		rtp.AddError(&errs, v.TokenIdentification.Validate())
	}
	if v.ParticipantReport != nil {
		rtp.AddError(&errs, v.ParticipantReport.Validate())
	}
	if v.ParticipantReportResponse != nil {
		rtp.AddError(&errs, v.ParticipantReportResponse.Validate())
	}
	if v.PaymentStatusRequest != nil {
		rtp.AddError(&errs, v.PaymentStatusRequest.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD ComplexType validations
