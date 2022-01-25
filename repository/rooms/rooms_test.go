package rooms

import (
	"crypto/sha256"
	"fmt"
	"project-airbnb/configs"
	"project-airbnb/entities"
	"project-airbnb/repository/users"
	"project-airbnb/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(&entities.Room{})
	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Room{})
	roomRepo := NewRoomsRepo(db)
	userRepo := users.NewUsersRepo(db)
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
	t.Run("Insert Room into Database", func(t *testing.T) {

		var mockRoom entities.Room
		mockRoom.Name = "Room1"
		mockRoom.Location = "Jakarta"
		mockRoom.Price = 300000
		mockRoom.Duration = 7
		mockRoom.User_id = 1
		mockRoom.Status = "Already Booked"

		_, err := roomRepo.Create(mockRoom)
		assert.Nil(t, err)

	})
	t.Run("Update Room into Database", func(t *testing.T) {

		var mockRoom entities.Room
		mockRoom.Name = "Room2"
		mockRoom.Location = "Jakarta"
		mockRoom.Price = 300000
		mockRoom.Duration = 6
		mockRoom.User_id = 1
		mockRoom.Status = "Already Booked"

		_, err := roomRepo.Update(mockRoom, 1)
		assert.Nil(t, err)

	})
	t.Run("Get All Room into Database", func(t *testing.T) {

		_, err := roomRepo.Gets(1)
		assert.Nil(t, err)

	})
	t.Run("Delete Room into Database", func(t *testing.T) {

		_, err := roomRepo.Delete(1, 1)
		assert.Nil(t, err)

	})
}
