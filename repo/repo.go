package repo

import (
	"MemberClub/member"
	"fmt"
	"log"
)

type UserRepo interface {
	InsertNewMember(string, interface{}) error
	Size() int
}

type userRepo struct {
	userMap map[string]interface{}
}

func (ur *userRepo) InsertNewMember(email string, mem interface{}) error {
	newMember, ok := mem.(member.Member)
	if !ok {
		err := fmt.Errorf("got data of type %T but wanted memmer.Member", newMember)
		log.Println(err)
		return err
	}
	if _, present := ur.userMap[email]; present {
		return fmt.Errorf("this email is already taken")
	}
	ur.userMap[email] = newMember
	return nil
}
func (ur *userRepo) Size() int {
	return len(ur.userMap)
}
func NewRepo() *userRepo {
	m := make(map[string]interface{})
	return &userRepo{m}
}

var test UserRepo = NewRepo()

//func (ur *UserRepo) String() string {
//	return fmt.Sprintf("", ur.userMap)
//}
