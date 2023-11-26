package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Personalizations struct {
	To          []string
	DynamicData map[string]interface{}
	TemplateID  string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file -> ", err)
	}
	// Make a new reader that consumes from `test-topic`, partition 0.
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topix",
		GroupID: "group-id",
	})
	defer r.Close()

	// Read a message
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error while receiving message: %s", err.Error())
			continue
		}
		composeEmailForForgotPassword("meesaqnabi@gmail.com", string(m.Value))
		fmt.Printf("Message: of irfaq is %s\n", string(m.Value))
		// Process the message here
		// Optionally commit the offset
	}
}
func composeEmailForForgotPassword(recipient, message string) {

	logInLink := message
	pData := map[string]interface{}{
		"logInLink": logInLink,
	}

	requestData := Personalizations{
		To:          []string{recipient},
		DynamicData: pData,
		TemplateID:  "d-51c3d100bcbb4088a5dc53f8f2ee6e44",
	}
	ComposeDynamicTemplateEmail(requestData)

	//composeAndSend(recipient, subject, "./htmlJobsdive/templates/email_forgot_password.tmpl.html", data)

}

func ComposeDynamicTemplateEmail(personalizations Personalizations) {

	m := mail.NewV3Mail()

	//from
	e := mail.NewEmail(os.Getenv("SENDER_NAME"), os.Getenv("SENDER_EMAIL"))
	m.SetFrom(e)

	//template id
	m.SetTemplateID(personalizations.TemplateID)

	p := mail.NewPersonalization()

	for _, to := range personalizations.To {
		//mail.NewEmail("", to)
		p.AddTos(mail.NewEmail("", to))
	}

	// tos := []*mail.Email{
	// 	mail.NewEmail("Example User", "test1@example.com"),
	// 	mail.NewEmail("Example User", "test2@example.com"),
	// }
	// p.AddTos(tos...)

	//p.SetDynamicTemplateData("receipt", "true")
	p.DynamicTemplateData = personalizations.DynamicData

	m.AddPersonalizations(p)

	sendDynamicTemplateEmail(mail.GetRequestBody(m))

}
func sendDynamicTemplateEmail(body []byte) {
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = body
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
