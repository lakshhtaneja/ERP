package data

type User struct { // capitalizing means public
	Name    string `json:"name"`
	Id      string `json:"id"`
	Contact `json:"contact"`
	Status  string `json:"status"`
}

type Contact struct {
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type Student struct {
	User    `json:"user"`
	Major   string `json:"major"`
	ClassOf string `json:"ClassOf"`
	Cgpa    string
	Notes   []string `json:"notes"`
}

//TODO: Add Course, Teacher, and Fee structs
