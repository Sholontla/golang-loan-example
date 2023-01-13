package main

import (
	"regexp"
	"strings"

	"github.com/google/uuid"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type Clients struct {
	ClientId          uuid.UUID
	ClientFirstName   string `json:"client_first_name"`
	ClientLastName    string `json:"client_last_name"`
	ClientUserName    string `json:"client_user_name"`
	ClientEmail       string `json:"client_email"`
	ClientPhoneNumber string `json:"client_phone_number"`
	ClientStatus      string `json:"client_status"`
	ClientDateCreated string `json:"client_date_created"`
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

func (customer *Clients) UserValidation() utilerrors.RestErr {
	customer.ClientEmail = strings.TrimSpace(strings.ToLower(customer.ClientEmail))

	if customer.ClientEmail == "" {
		return utilerrors.NewBadRequestError("Invalid User Email")

	} else if !isEmailValid(customer.ClientEmail) {
		return utilerrors.NewBadRequestError("Invalid User Email")

	} else if customer.ClientFirstName == "" {
		return utilerrors.NewBadRequestError("Invalid User Name")
	}

	return nil
}
