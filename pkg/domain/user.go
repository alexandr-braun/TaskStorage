package domain

type User struct {
	Id         int
	FirstName  string
	MiddleName string
	LastName   string
}

func (u User) GetFullName() string {
	var result = u.FirstName + " " + u.MiddleName + " " + u.LastName

	return result
}
