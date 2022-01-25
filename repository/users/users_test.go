package users

import (
	"crypto/sha256"
	"fmt"
	"project-airbnb/configs"
	"project-airbnb/entities"
	"project-airbnb/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(&entities.User{})

	db.AutoMigrate(&entities.User{})
	userRepo := NewUsersRepo(db)
	db.Migrator().DropTable(&entities.User{})
	t.Run("Insert User into Database", func(t *testing.T) {
		hash := sha256.Sum256([]byte("ilham123"))
		password := fmt.Sprintf("%x", hash[:])
		var mockUser entities.User
		mockUser.Email = "ilham@yahoo.com"
		mockUser.Password = password
		mockUser.Name = "ilham"

		_, err := userRepo.Register(mockUser)
		assert.Error(t, err)

	})
	t.Run("Select All User from Database", func(t *testing.T) {
		_, err := userRepo.Gets()
		assert.Error(t, err)

	})
	db.AutoMigrate(&entities.User{})
	t.Run("Insert User into Database", func(t *testing.T) {
		hash := sha256.Sum256([]byte("ilham123"))
		password := fmt.Sprintf("%x", hash[:])
		var mockUser entities.User
		mockUser.Email = "ilham@yahoo.com"
		mockUser.Password = password
		mockUser.Name = "ilham"

		_, err := userRepo.Register(mockUser)
		assert.Nil(t, err)

	})
	t.Run("Error Login User into Database", func(t *testing.T) {
		hash := sha256.Sum256([]byte("ilhasdasdam123"))
		password := fmt.Sprintf("%x", hash[:])
		var mockUser entities.User
		mockUser.Email = "ilham@yashoo.com"
		mockUser.Password = password
		mockUser.Name = "ilham"

		_, err := userRepo.LoginUser(mockUser.Email, mockUser.Password)
		assert.Error(t, err)

	})
	t.Run("Login User into Database", func(t *testing.T) {
		hash := sha256.Sum256([]byte("ilham123"))
		password := fmt.Sprintf("%x", hash[:])
		var mockUser entities.User
		mockUser.Email = "ilham@yahoo.com"
		mockUser.Password = password

		_, err := userRepo.LoginUser(mockUser.Email, mockUser.Password)
		assert.Nil(t, err)

	})
	t.Run("Select All User from Database", func(t *testing.T) {
		_, err := userRepo.Gets()
		assert.Nil(t, err)

	})
	t.Run("Select User By Id from Database", func(t *testing.T) {
		_, err := userRepo.Get(1)
		assert.Nil(t, err)

	})
	t.Run("Select User By Id from Database", func(t *testing.T) {
		_, err := userRepo.Get(100)
		assert.Error(t, err)

	})
	t.Run("Update User into Database", func(t *testing.T) {
		hash := sha256.Sum256([]byte("junius123"))
		password := fmt.Sprintf("%x", hash[:])
		var mockUser entities.User
		mockUser.Email = "junius@yahoo.com"
		mockUser.Password = password
		mockUser.Name = "junius"

		_, err := userRepo.Update(mockUser, 1)
		assert.Nil(t, err)

	})
	t.Run("Error Update User into Database", func(t *testing.T) {
		hash := sha256.Sum256([]byte("junius123"))
		password := fmt.Sprintf("%x", hash[:])
		var mockUser entities.User
		mockUser.Email = "junius@yahoo.com"
		mockUser.Password = password
		mockUser.Name = "junius"

		_, err := userRepo.Update(mockUser, 100)
		assert.Error(t, err)

	})
	t.Run("Error Delete User into Database", func(t *testing.T) {

		_, err := userRepo.Delete(100)
		assert.Error(t, err)

	})
	t.Run("Delete User into Database", func(t *testing.T) {
		_, err := userRepo.Delete(1)
		assert.Nil(t, err)

	})

}
