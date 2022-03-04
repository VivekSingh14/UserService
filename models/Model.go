package models

type User struct {
	Id         string `json:id`
	FirstName  string `json:fName`
	MiddleName string `json:mName`
	LastName   string `json:lName`
	DOB        string `json:dob`
	Address    string `json:address`
	UserName   string `json:uName`
	UserType   string `json:uType`
}
