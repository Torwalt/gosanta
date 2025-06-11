package usernotifying

import "gosanta/internal"

type UserNotifyer struct{}

func New() UserNotifyer {
	return UserNotifyer{}
}

func (u *UserNotifyer) SendToUser(awards.UserAwardEvent) error {
	// STUB TODO
	return nil
}
