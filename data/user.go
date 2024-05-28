package data

type User struct {
	Name    string // capitalizing means public
	Id      string
	Contact Contact
	Status  string
}

type Contact struct {
	Phone1  string
	Phone2  string
	Address string
}

type Student struct {
	User User
	CGPA float64 // json default uses float64
}

//TODO: Add Course, Teacher, and Fee structs
