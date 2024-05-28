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
	User `json:"user"`
	Cgpa float64 // json default uses float64
}

//TODO: Add Course, Teacher, and Fee structs
