package email

import (
    "net/smtp"
    conf "config"
    "log"
    "fmt"
)

// the send sending contains following steps
/*
1. To dial the smtp server along with port to get the smtp client
2. Call Mail function on client to setup the from (the address which is sending the email)
3. Call Rcpt function on client to setup the to (the address on which email is being sent)
*/

func Send_email(){
    // fetching the email related global variable from config package
    smtp_server := conf.Smtp
    port := conf.Port
    user := conf.User
    password := conf.Password
    email_to := conf.Email_to
    emailBody := "Hello Mamma \n \n Mamma you are beautiful"
    subject := "Subject: My Beautiful Mamma \n"

    mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
    msg := []byte(subject + mime + emailBody)

    auth := smtp.PlainAuth("", user, password, smtp_server)
    smtp_server_port := fmt.Sprintf("%s:%s",smtp_server,port)
    err := smtp.SendMail(smtp_server_port, auth, user, email_to, msg)
    if err != nil {
    		log.Fatal(err)
    }
    log.Println("Email successfully sent to ")
    log.Println(email_to)
}