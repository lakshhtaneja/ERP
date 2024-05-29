package internal

import "github.com/lakshhtaneja/ERP/data"

var Peter = data.Student{
	User: data.User{
		Name: "Peter",
		Id:   "123",
		Contact: data.Contact{
			Phone:   "1234567890",
			Email:   "abc@xyz.com",
			Address: "123, 4th Street, New York",
		},
		Status: "active",
	},
	Cgpa:    9.8,
	Major:   "Computer Science",
	ClassOf: 2025,
}
