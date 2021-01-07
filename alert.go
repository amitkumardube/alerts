package main

import (
	"fmt"
	db "dbconn"
)

func main(){
// Calling the function in dbconn package to see the product with expiry dates
    fmt.Println("Getting the list of products approaching expiry")
	db.Get_expiry_data()
	fmt.Println("Getting the list of products with limited stock")
	db.Get_stock_data()
}
