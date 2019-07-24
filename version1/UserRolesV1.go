package version1

import (
	"time"
)

type UserRolesV1 struct {
	Id         string    `json:"id"`
	Roles      []string  `json:"roles"`
	UpdateTime time.Time `json:"update_time"`
}

func EmptyUserRolesV1() *UserRolesV1 {
	return &UserRolesV1{}
}

func NewUserRolesV1(id string, roles []string) *UserRolesV1 {
	return &UserRolesV1{
		Id:         id,
		Roles:      roles,
		UpdateTime: time.Now(),
	}
}
