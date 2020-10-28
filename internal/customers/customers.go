package customers

import (
	"buyevent/internal/configs"
	"encoding/json"
	"io/ioutil"
)

// Customer customer, user.
type Customer struct {
	ID          int    `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}

// Customers array of customers, users
var Customers []Customer

func init() {
	// Reading customers from the file
	sliceByteCustomers, err := ioutil.ReadFile(configs.Configs.CustomersDataFilePath)
	if err != nil {
		panic(err)
	}
	// Unmarshal sliceByteCustomers to Customers struct
	json.Unmarshal(sliceByteCustomers, &Customers)
	if err != nil {
		panic(err)
	}
}
