package db

import (
	"api/config"
	"api/utils"
)

func EncryptPassword(password string) string {
	return utils.CalculateMD5(password + config.T.Conf.Salt)
}
