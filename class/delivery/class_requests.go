package delivery

type ClassCreationRequest struct {
	Name          string `json:"name"`
	StudentNumber int    `json:"student_number"`
}

type ClassLoginRequest struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}
