package db

import "api/drivers/db/model"

type UserMethod interface {
	// FindById where(id=@id)
	FindById(id int) (model.User, error)
	// FindByPassword where(username=@username AND password=@password)
	FindByPassword(username string, password string) (model.User, error)
}

type UserAuthMethod interface {
}
