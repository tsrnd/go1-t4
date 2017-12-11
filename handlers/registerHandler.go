package handlers

type RegisterInfo struct {
	UserName string `schema:"userName"`
	Email    string `schema:"email"`
	Phone    string `schema:"phoneNumber"`
	Address  string `schema:"address"`
	Password string `schema:"password"`
	Errors   map[string]string
}

// func Register(w http.ResponseWriter, r *http.Request) {
// 	HomeVars := NewHomePageVars(r)
// 	HomeVars.Message = utils.ShowMessage(w, r, "register")
// 	HomeVars.PageTitle = "Login"
// 	utils.GenerateTemplate(w, HomeVars, "login_register", "login", "register")
// }

// func RegisterHandler(w http.ResponseWriter, r *http.Request) {
// 	var info RegisterInfo
// 	var HomeVars HomePageVars
// 	err := utils.MapFormValues(&info, r)
// 	if err != nil {
// 		fmt.Println("cannot decode register info: ", err)
// 	} else if info.Validate() == false {
// 		HomeVars.RegisterInfo = info
// 		utils.GenerateTemplate(w, HomeVars, "login_register", "login", "register")
// 	} else if utils.CheckUserExist(info.UserName) {
// 		mess := "User account already exist!!"
// 		utils.SetMessage(w, mess, "register")
// 	} else {
// 		user := SetUser(info)
// 		err2 := models.CreateUser(user) // save user info
// 		if err2 != nil {
// 			fmt.Println("Error occur when creating user")
// 		} else {
// 			mess := "register user successful"
// 			utils.SetMessage(w, mess, "register")
// 		}
// 	}
// 	http.Redirect(w, r, "/register", http.StatusSeeOther)
// }

// func (msg *RegisterInfo) Validate() bool {
// 	msg.Errors = make(map[string]string)

// 	emailRegex := regexp.MustCompile(".+@.+\\..+")
// 	emailMatched := emailRegex.Match([]byte(msg.Email))
// 	if emailMatched == false {
// 		msg.Errors["Email"] = "Please enter a valid email address"
// 	}

// 	if strings.TrimSpace(msg.Address) == "" {
// 		msg.Errors["Address"] = "Please write your address"
// 	}

// 	if strings.TrimSpace(msg.UserName) == "" {
// 		msg.Errors["Name"] = "Please write your name"
// 	}

// 	if strings.TrimSpace(msg.Password) == "" {
// 		msg.Errors["Password"] = "Please write your password"
// 	}

// 	phoneRegex := regexp.MustCompile("^[0-9]{10}")
// 	phoneMatched := phoneRegex.Match([]byte(msg.Phone))
// 	if phoneMatched == false {
// 		msg.Errors["Phone"] = "Please enter a valid phone number(with number only and 10 digits)"
// 	}

// 	return len(msg.Errors) == 0
// }

// func SetUser(userInfo RegisterInfo) (user models.User) {
// 	user.UserName = userInfo.UserName
// 	hashPass, _ := utils.HashPassword(userInfo.Password)
// 	user.Password = hashPass
// 	user.Address = userInfo.Address
// 	user.Phone = userInfo.Phone
// 	user.Email = userInfo.Email
// 	return user
// }
