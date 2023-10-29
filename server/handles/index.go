package handles

import (
	"api/dto/resp"
	"api/i18n"
	"api/server/common"
)

func GetToken(c *common.ApiContext) {
	token, err := c.GetToken()
	if err != nil {
		c.JSONI(500, resp.T{
			Code: 500,
			Msg:  i18n.RedisError,
		})
		return
	}
	c.JSONI(200, resp.T{
		Code: 200,
		Msg:  i18n.RequestSuccess,
		Data: resp.M{
			token: token,
		},
	})
}

func Index(c *common.ApiContext) {
	c.JSON(200, resp.CodeMsg(200, "ok"))
}
