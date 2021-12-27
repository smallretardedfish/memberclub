package repo

import (
	"MemberClub/member"
	"fmt"
	"log"
)

type UserRepo interface {
	InsertNewMember(interface{}) error
	GetAllMembers() interface{}
	Size() int
}

type userRepo struct {
	userMap map[string]interface{}
}

func (ur *userRepo) InsertNewMember(mem interface{}) error {
	newMember, ok := mem.(member.Member)
	if !ok {
		err := fmt.Errorf("got data of type %T but wanted member.Member", newMember)
		log.Println(err)
		return err
	}
	email := newMember.GetMemberEmail()
	if _, present := ur.userMap[email]; present {
		return fmt.Errorf("this email is already taken")
	}
	ur.userMap[email] = newMember
	return nil
}
func (ur *userRepo) Size() int {
	return len(ur.userMap)
}
func (ur *userRepo) GetAllMembers() interface{} {

	return ur.userMap
}
func NewRepo() *userRepo {
	m := make(map[string]interface{})
	return &userRepo{m}
}
