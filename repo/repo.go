package repo

import (
	"MemberClub/member"
	"fmt"
)

type UserRepo struct {
	userMap map[string]interface{}
}

func (ur *UserRepo) InsertNewMember(member member.Member) error {
	email := member.GetMemberEmail()
	if _, ok := ur.userMap[email]; ok {
		return fmt.Errorf("this email is already taken")
	}
	ur.userMap[email] = member
	return nil
}
func (ur UserRepo) Size() int {
	return len(ur.userMap)
}
func NewRepo() *UserRepo {
	return &UserRepo{}
}
func (ur *UserRepo) String() string {
	return fmt.Sprintf("", ur.userMap)
}
