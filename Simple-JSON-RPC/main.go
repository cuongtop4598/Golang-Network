package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

// SMS Args
type SmsArgs struct {
	Number, Content string
}

//Email Args
type EmailArgs struct {
	To, Subject, Content string
}

//Response structure for email and SMS service
type Response struct {
	Result string
}

//Simple example of SMS service without logic
type EmailService struct{}
type SmsService struct{}

func (t *EmailService) SendEmail(r *http.Request, args *EmailArgs, result *Response) error {
	*result = Response{Result: fmt.Sprintf("Email sent to %s", args.To)}
	return nil
}

func (t *SmsService) SendSMS(r *http.Request, args *SmsArgs, result *Response) error {
	*result = Response{Result: fmt.Sprintf("Sms sent to %s", args.Number)}
	return nil
}

//Example how to register services and start rpc server with
func main() {
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	sms := new(SmsService)
	email := new(EmailService)

	rpcServer.RegisterService(sms, "sms")
	rpcServer.RegisterService(email, "email")

	router := mux.NewRouter()
	router.Handle("/delivery", rpcServer)
	http.ListenAndServe(":1337", router)
}
