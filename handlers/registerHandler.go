package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/goweb4/utils"
)

type Register struct {
	UserName string `schema:"userName"`
	Email    string `schema:"email"`
	Phone    string `schema:"phoneNumber"`
	Address  string `schema:"address"`
	Password string `schema:"password"`
	Errors   map[string]string
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var info Register
	var HomeVars HomePageVars
	err := utils.MapFormValues(&info, r)
	if err != nil {
		fmt.Println("cannot decode register info: ", err)
	} else {
		if info.Validate() == false {
			HomeVars.RegisterInfo = info
			utils.GenerateTemplate(w, HomeVars, "login_register", "login", "register")
		} else {
			// save user info
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
}

func (msg *Register) Validate() bool {
	msg.Errors = make(map[string]string)

	emailRegex := regexp.MustCompile(".+@.+\\..+")
	emailMatched := emailRegex.Match([]byte(msg.Email))
	if emailMatched == false {
		msg.Errors["Email"] = "Please enter a valid email address"
	}

	if strings.TrimSpace(msg.Address) == "" {
		msg.Errors["Address"] = "Please write your address"
	}

	if strings.TrimSpace(msg.UserName) == "" {
		msg.Errors["Name"] = "Please write your name"
	}

	if strings.TrimSpace(msg.Password) == "" {
		msg.Errors["Password"] = "Please write your password"
	}

	phoneRegex := regexp.MustCompile("^[0-9]{10}")
	phoneMatched := phoneRegex.Match([]byte(msg.Phone))
	if phoneMatched == false {
		msg.Errors["Phone"] = "Please enter a valid phone number(with number only and 9 digits)"
	}

	return len(msg.Errors) == 0
}
