package member

import "time"

type Member interface {
	GetMemberEmail() string
	GetMemberName() string
	GetUserCreationDate() time.Time
}

type member struct {
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	CreationTime time.Time `json:"creation_time"`
}

func (m *member) GetMemberEmail() string {
	return m.Email
}

func (m *member) GetMemberName() string {
	return m.Name
}

func (m *member) GetUserCreationDate() time.Time {
	return m.CreationTime
}
func NewMember(name, email string) *member {
	return &member{
		Name:         name,
		Email:        email,
		CreationTime: time.Now(),
	}
}
