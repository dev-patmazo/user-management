package models

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/dev-patmazo/user-management/utils"
	"gorm.io/gorm"
)

var DbClient *gorm.DB

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Role     string `json:"role"`
}

// CreateUser accepts map abnd unmarshals it to User struct
// It then validates the username and email to ensure they are unique
// it then hashes the password and inserts the user into the database
func CreateUser(newUser map[string]interface{}) (User, error) {

	var user User

	jsonData, err := json.Marshal(newUser)
	if err != nil {
		return User{}, err
	}

	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		return User{}, err
	}

	var existingUser User
	result := DbClient.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser)
	if result.Error == nil {
		return User{}, errors.New("username or email already exists")
	}

	user.Password = utils.BasicAuthGenerator(user.Username, "hopetogetthisjob")

	// Insert the user into the database
	result = DbClient.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

// GetUser simply retrieves a user from the database using id
func GetUser(id string) (User, error) {
	var user User
	result := DbClient.First(&user, id)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

// UpdateUser accepts a map and unmarshals it to User struct
// It then updates the user in the database
func UpdateUser(id string, updatedUser map[string]interface{}) (User, error) {

	var user User
	result := DbClient.First(&user, id)
	if result.Error != nil {
		return User{}, result.Error
	}

	jsonData, err := json.Marshal(updatedUser)
	if err != nil {
		return User{}, err
	}

	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		return User{}, err
	}

	user.Password = utils.BasicAuthGenerator(user.Username, updatedUser["password"].(string))

	result = DbClient.Updates(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

// DeleteUser deletes a user from the database
// It accepts an id and converts it to uint
// It then deletes the user from the database
// Note: Gorm has 2 delete methods, one is soft delete
// and the other is hard delete
func DeleteUser(id string) error {

	// Convert id to uint
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}

	var user User
	user.ID = uint(uintID)

	// result := DbClient.Delete(&user) // This will soft delete data from the database.
	result := DbClient.Unscoped().Delete(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func CheckUserRole(username, password string) (User, error) {
	var user User
	result := DbClient.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}
