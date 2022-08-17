package test

import (
	"github.com/google/uuid"
	user2 "github.com/lambsroad/go-shop/users"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUserModel(t *testing.T) {
	t.Run("유저 한명을 생성해야 합니다.", func(t *testing.T) {
		uuid := uuid.New().String()
		user := user2.User{UUID: uuid, UserName: "jeffrey", Password: "test"}
		expected := user.CreateUser(user)

		assert.Equal(t, user.UserName, expected.UserName, "같은 유저가 아닙니다.")
	})
}

//func TestDeleteUser(t *testing.T) {}
//func TestUpdateUser(t *testing.T) {}
//func TestSearchUser(t *testing.T) {}
