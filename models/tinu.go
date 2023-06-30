package model

func GetAllTinus() ([]Tinu, error) {
	var tinus []Tinu
	// search for all tinus in db
	tx := db.Find(&tinus)
	if tx.Error != nil {
		return []Tinu{}, tx.Error
	}

	return tinus, nil
}

func GetOneTinu(id string) (Tinu, error) {
	var tinu Tinu
	// search tinu with id in db
	tx := db.Where("id = ?", id).First(&tinu)
	if tx.Error != nil {
		return Tinu{}, tx.Error
	}

	return tinu, nil
}

func CreateTinu(tinu Tinu) error {
	// add new tinu to db
	tx := db.Create(&tinu)
	return tx.Error
}

func UpdateTinu(tinu Tinu) error {
	// save updated tinu to db
	tx := db.Save(&tinu)
	return tx.Error
}

func DeleteTinu(id string) error {
	tx := db.Unscoped().Delete(&Tinu{}, "id=?", id)
	return tx.Error
}
