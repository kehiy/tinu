package model

func GetAllUsers() ([]User, error) {
	var users []User
	// search for all users in db
	tx := db.Find(&users)
	if tx.Error != nil {
		return []User{}, tx.Error
	}

	return users, nil
}

func GetOneUser(id string) (User, error) {
	var user User
	// search user with id in db
	tx := db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return User{}, tx.Error
	}

	return user, nil
}

func CreateUser(user User) error {
	// add new user to db
	tx := db.Create(&user)
	return tx.Error
}

func UpdateUser(user User) error {
	// save updated user to db
	tx := db.Save(&user)
	return tx.Error
}

func DeleteUser(id string) error {
	tx := db.Unscoped().Delete(&User{}, "id=?", id)
	return tx.Error
}
