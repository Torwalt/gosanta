package awards

import "fmt"

type UserId int64
type CompanyId int64

// A user that can receive an award.
type User struct {
	Id        UserId
	FirstName string
	LastName  string
	CompanyId CompanyId
	Awards    []PhishingAward
}

func (u *User) FullName() string {
	return fmt.Sprintf("%v %v", u.FirstName, u.LastName)
}

// A company to which users belong.
type Company struct {
	id    CompanyId
	users []User
}
