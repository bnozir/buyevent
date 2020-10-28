package notifications

import (
	"buyevent/internal/configs"
	"buyevent/internal/customers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// NotificationType - notification type for customers
type NotificationType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// NotificationsType - array of notifications type
var NotificationsType []NotificationType

// Notification - notification for customers
type Notification struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Notifications - array of notifications
var Notifications []Notification

func init() {
	// Reading notifications type from the file
	sliceByteNotificationsType, err := ioutil.ReadFile(configs.Configs.NotificationsTypeDataFilePath)
	if err != nil {
		panic(err)
	}
	// Unmarshal sliceByteNotificationsType to NotificationsType struct
	err = json.Unmarshal(sliceByteNotificationsType, &NotificationsType)
	if err != nil {
		panic(err)
	}

	// Reading notifications from the file
	sliceByteNotifications, err1 := ioutil.ReadFile(configs.Configs.NotificationsDataFilePath)
	if err1 != nil {
		panic(err1)
	}
	// Unmarshal sliceByteNotifications to Notifications struct
	err = json.Unmarshal(sliceByteNotifications, &Notifications)
	if err != nil {
		panic(err)
	}
}

// SendSMS - sends SMS to customer
func SendSMS(customer customers.Customer, notification string) {
	fmt.Println("Please wait while SMS sending...")
	time.Sleep(2 * time.Second)
	fmt.Println(notification + " successfuly sent to customer with phonenumber: " + customer.PhoneNumber)
}

// SendEmail - sends email to customer
func SendEmail(customer customers.Customer, notification string) {
	fmt.Println("Please wait while email sending...")
	time.Sleep(2 * time.Second)
	fmt.Println(notification + " successfuly sent to customer with email: " + customer.Email)
}
