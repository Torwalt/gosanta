package awards

type UserId int64
type CompanyId int64

// A user that can receive an award.
type User struct {
	Id        UserId
	CompanyId CompanyId
	Awards    []PhishingAward
}

// Return award received for the interaction with a phishing mail.
func (u *User) FindRelatedAward(emailRef string) *PhishingAward {
	for _, a := range u.Awards {
		if a.EmailRef == emailRef {
			return &a
		}
	}
	return nil
}

// A company to which users belong.
type Company struct {
	id    CompanyId
	users []User
}
