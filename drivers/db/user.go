package db

func DeleteUserById(id int) (bool, error) {
	tx := Gorm.Exec("DELETE FROM user WHERE id = ?", id)
	return tx.RowsAffected > 0, tx.Error
}
