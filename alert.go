package main

import (
	"fmt"
	db "dbconn"
	"email"
)

func main(){
// Calling the function in dbconn package to see the product with expiry dates
    fmt.Println("Getting the list of products approaching expiry")
	db.Get_expiry_data()
	fmt.Println("Getting the list of products with limited stock")
	db.Get_stock_data()

// Calling the function from email package to send email
    email.Send_email()
}
