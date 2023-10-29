package handles

import (
	"api/drivers/db"
	"api/drivers/db/query"
	"api/dto/resp"
	"api/i18n"
	"api/server/common"
	"encoding/json"
)

type loginStruct struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func Login(c *common.ApiContext) {
	b, err := c.VerifyCode()
	if err != nil || !b {
		c.JSONI(200, resp.CodeMsg(0, i18n.VerificationCodeError))
		return
	}
	data, err := c.GetRawData()
	if err != nil {
		c.JSONI(200, resp.CodeMsg(0, i18n.RedisError))
	}
	var login loginStruct
	err = json.Unmarshal(data, &login)
	if err != nil {
		c.JSONI(200, resp.CodeMsg(0, i18n.RedisError))
	}
	user, err := query.Q.User.FindByPassword(login.Username, db.EncryptPassword(login.Password))
	if err != nil {
		c.JSONI(200, resp.CodeMsg(0, i18n.RedisError))
	}
	c.JSONI(200, resp.CodeMsg(200, i18n.RequestSuccess).Data(user))
}

func Register(c *common.ApiContext) {
	b, err := c.VerifyCode()
	if err != nil || !b {
		c.JSONI(200, resp.CodeMsg(0, i18n.VerificationCodeError))
		return
	}

}

func UserInfo(c *common.ApiContext) {

}
