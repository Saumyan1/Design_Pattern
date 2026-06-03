package main

import "fmt"

//Notification system



//strategy interface
type Notification interface{
	Send(message string)
}


//Concrete strategy


//email-notification
type EmailNotification struct{
}
func(e *EmailNotification) Send(message string){
	fmt.Println("Sending Email -> ", message)	
}

//sms notification
type SMSNotification struct{
}
func(s *SMSNotification) Send(message string){
	fmt.Println("Sending sms -> ", message)
}

//push notification
type PushNotification struct{
}
func(p *PushNotification) Send(message string){
	fmt.Println("Sending pushNotification -> ", message)
}

//Context - NotificationServices

type NotificationService struct{
	Notif Notification
}

//constructor

func NewNotificationService(n Notification) *NotificationService{
	return &NotificationService{
		Notif: n,	
	}

}

func(ns *NotificationService) SetNotification(n Notification){
	ns.Notif = n
}

func (ns *NotificationService) Notify(message string){
	 ns.Notif.Send(message)
}

//CLIENT 

func main(){
	email := &EmailNotification{}
	sm := &SMSNotification{}

	service := NewNotificationService(email)
	service.Notify("Order Placed")
	service.SetNotification(sm)
	service.Notify("Order Placed")
}