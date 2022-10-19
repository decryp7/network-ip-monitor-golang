package main

import (
	"fmt"

	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func sendEmailUsingOutlook(subject string, message string) {
	//code to recover from panic and log
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(fmt.Errorf("error occurred during sending email: %v", err))
		}
	}()

	//get outlook ole interface
	ole.CoInitialize(0)
	unknown, _ := oleutil.CreateObject("Outlook.Application")
	outlook, _ := unknown.QueryInterface(ole.IID_IDispatch)

	//get current user email address
	session := oleutil.MustGetProperty(outlook, "Session").ToIDispatch()
	currentUser := oleutil.MustGetProperty(session, "CurrentUser").ToIDispatch()
	addressEntry := oleutil.MustGetProperty(currentUser, "AddressEntry").ToIDispatch()
	exchangeUser := oleutil.MustCallMethod(addressEntry, "GetExchangeUser").ToIDispatch()
	currentUserEmail := oleutil.MustGetProperty(exchangeUser, "PrimarySmtpAddress").Value()

	//create messge
	email := oleutil.MustCallMethod(outlook, "createitem", 0).ToIDispatch()
	emailRecipents := oleutil.MustCallMethod(email, "Recipients").ToIDispatch()
	oleutil.MustCallMethod(emailRecipents, "Add", currentUserEmail).ToIDispatch()
	oleutil.MustPutProperty(email, "Subject", subject).ToIDispatch()
	oleutil.MustPutProperty(email, "Body", message).ToIDispatch()

	//send the message!
	oleutil.MustCallMethod(email, "Send").ToIDispatch()
}
