package service

import (
	"buyevent/internal/customers"
	"buyevent/internal/logger"
	"buyevent/internal/notifications"
	"fmt"
	"log"
	"strconv"
)

// Service a main part of programm
type Service struct {
	Customers         []customers.Customer
	NotificationsType []notifications.NotificationType
	Notifications     []notifications.Notification
	InfoLogger        *log.Logger
	ErrorLogger       *log.Logger
}

// Service - instant of Service type
var service Service

func init() {
	service = newService()
}

func newService() Service {
	return Service{
		Customers:         customers.Customers,
		NotificationsType: notifications.NotificationsType,
		Notifications:     notifications.Notifications,
		InfoLogger:        logger.InfoLogger,
		ErrorLogger:       logger.ErrorLogger,
	}
}

// Run - runs the service
func Run() {
	var (
		operator             string
		customerIDString     string
		customerID           int
		notificationType     string
		notificationIDString string
		notificationID       int
		err                  error
	)
	fmt.Printf("%s", "Please enter you name: ")
	if _, err = fmt.Scan(&operator); err != nil {
		service.ErrorLogger.Printf("%s\n", "Error on reading operator name: "+err.Error())
		fmt.Println("Opss, error on reading your name :(")
	}
	fmt.Println(operator + ", welcome to Buy-Event :)\n")
	service.InfoLogger.Println(operator + " loged in")
	// Customers list
	for {
		fmt.Println("Customers:")
		for _, v := range service.Customers {
			fmt.Printf("%d - %s - %s\n", v.ID, v.PhoneNumber, v.Email)
		}
		fmt.Printf("\n%s", "Please, choose a customer by id ('q'-quit): ")
		if _, err = fmt.Scan(&customerIDString); err != nil {
			service.ErrorLogger.Printf("%s\n", "Error on reading customer id: "+err.Error())
			fmt.Println("Opss, error on reading customer id. Please try again.")
			continue
		}
		if customerIDString == "q" {
			fmt.Println(operator + ", see you again :)\n")
			service.InfoLogger.Println(operator + " loged out")
			break
		}
		if customerID, err = strconv.Atoi(customerIDString); err != nil {
			service.ErrorLogger.Printf("%s\n", "Wrong customer id or command for quit "+customerIDString)
			fmt.Println("Opss, Please choose customer id from the list or type 'q' for quit.")
			continue
		}
		if customerID > len(service.Customers) || customerID <= 0 {
			service.ErrorLogger.Printf("%s\n", "Wrong customer id! No customer with id "+strconv.Itoa(customerID))
			fmt.Println("Opss, Wrong customer id. Please choose customer id from the list.")
			continue
		}
		service.InfoLogger.Println(operator + " choosed customer with id: " + strconv.Itoa(customerID))

		// Notifications type
		for {
			fmt.Println("NotificationsType:")
			for _, v := range service.NotificationsType {
				fmt.Printf("%s - %s\n", v.Name, v.Description)
			}
			fmt.Printf("\n%s", "Please, choose a notification type by name ('b'-back to customers): ")
			if _, err = fmt.Scan(&notificationType); err != nil {
				service.ErrorLogger.Printf("%s\n", "Error on reading notification type name: "+err.Error())
				fmt.Println("Opss, error on reading notification type name. Please try again.")
				continue
			}
			switch notificationType {
			case "sms":
				for {
					fmt.Println("Notifications:")
					for _, v := range service.Notifications {
						fmt.Printf("%d - %s - %s\n", v.ID, v.Name, v.Description)
					}
					fmt.Printf("\n%s", "Please, choose a notification by id ('b'-back to notifications type): ")
					if _, err = fmt.Scan(&notificationIDString); err != nil {
						service.ErrorLogger.Printf("%s\n", "Error on reading notification id: "+err.Error())
						fmt.Println("Opss, error on reading notification id. Please try again.")
						continue
					}
					if notificationIDString == "b" {
						service.InfoLogger.Println(operator + " returned to customers list")
						break
					}
					if notificationID, err = strconv.Atoi(notificationIDString); err != nil {
						service.ErrorLogger.Printf("%s\n", "Wrong notification id or command to back "+customerIDString)
						fmt.Println("Opss, Please choose notification id from the list or type 'b' for back.")
						continue
					}
					if notificationID > len(service.Notifications) || notificationID <= 0 {
						service.ErrorLogger.Printf("%s\n", "Wrong notification id! No notification with id "+strconv.Itoa(customerID))
						fmt.Println("Opss, Wrong notification id. Please choose notification id from the list.")
						continue
					}
					service.InfoLogger.Println(operator + " choosed notification with id: " + notificationIDString)
					notifications.SendSMS(service.Customers[customerID-1], service.Notifications[notificationID-1].Name)
					service.InfoLogger.Println(operator + " sent sms to customer with phonenumber: " + service.Customers[customerID].PhoneNumber)
				}
				continue
			case "email":
				for {
					fmt.Println("Notifications:")
					for _, v := range service.Notifications {
						fmt.Printf("%d - %s - %s\n", v.ID, v.Name, v.Description)
					}
					fmt.Printf("\n%s", "Please, choose a notification by id ('b'-back to notifications type): ")
					if _, err = fmt.Scan(&notificationIDString); err != nil {
						service.ErrorLogger.Printf("%s\n", "Error on reading notification id: "+err.Error())
						fmt.Println("Opss, error on reading notification id. Please try again.")
						continue
					}
					if notificationIDString == "b" {
						service.InfoLogger.Println(operator + " returned to notifications type")
						break
					}
					if notificationID, err = strconv.Atoi(notificationIDString); err != nil {
						service.ErrorLogger.Printf("%s\n", "Wrong notification id or command to back "+customerIDString)
						fmt.Println("Opss, Please choose notification id from the list or type 'b' for back.")
						continue
					}
					if notificationID > len(service.Notifications) || notificationID <= 0 {
						service.ErrorLogger.Printf("%s\n", "Wrong notification id! No notification with id "+strconv.Itoa(customerID))
						fmt.Println("Opss, Wrong notification id. Please choose notification id from the list.")
						continue
					}
					service.InfoLogger.Println(operator + " choosed notification with id: " + notificationIDString)
					notifications.SendEmail(service.Customers[customerID-1], service.Notifications[notificationID-1].Name)
					service.InfoLogger.Println(operator + " sent email to customer with email: " + service.Customers[customerID].Email)
				}
				continue
			case "b":
				service.InfoLogger.Println(operator + " returned to cusromers list")
			default:
				service.ErrorLogger.Println(operator + " choose the wrong notification type: " + notificationType)
				fmt.Println("Opss, wrong notifacation type. Please try again.")
				continue
			}
			break
		}
	}
}
