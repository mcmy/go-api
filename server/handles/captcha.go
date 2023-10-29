package handles

import (
	"api/drivers/redis"
	"api/dto/resp"
	"api/i18n"
	"api/server/common"
	"api/utils"
	"bytes"
	"context"
	"errors"
	"github.com/dchest/captcha"
	"net/http"
	"time"
)

type VerifyType string

const (
	IMAGE VerifyType = ".png"
	AUDIO VerifyType = ".wav"
)

func Captcha(c *common.ApiContext) {
	token, err := c.GetToken()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	code := utils.RandomStr(5, "0123456789")
	var verifyType VerifyType
	switch c.Param("type") {
	case "image":
		verifyType = IMAGE
		break
	case "audio":
		verifyType = AUDIO
		break
	default:
		verifyType = IMAGE
	}
	err = serve(c, verifyType, token, code, false)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	err = redis.Redis.Set(context.Background(), redis.GenKey(token, "verify_code"), code, 3*time.Second).Err()
	if err != nil {
		c.JSONI(500, resp.T{
			Code: 500,
			Msg:  i18n.RedisError,
		})
		c.Abort()
		return
	}
}

func (t *VerifyType) ToString() string {
	return string(*t)
}

func serve(c *common.ApiContext, verifyType VerifyType, id, code string, download bool) error {
	w := c.Writer
	r := c.Request
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	array, err := stringToIntArray(code)
	if err != nil {
		return err
	}

	var content bytes.Buffer
	switch verifyType {
	case IMAGE:
		w.Header().Set("Content-Type", "image/png")
		_, err = captcha.NewImage(id, array, 128, 64).WriteTo(&content)
		break
	case AUDIO:
		w.Header().Set("Content-Type", "audio/x-wav")
		_, err = captcha.NewAudio(captcha.New(), array, getAudioLang(c.GetAcceptLang())).WriteTo(&content)
		break
	default:
		return errors.New("not found verify type")
	}
	if err != nil {
		return err
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}

	http.ServeContent(w, r, id+verifyType.ToString(), time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

func stringToIntArray(s string) ([]byte, error) {
	result := make([]byte, len(s))
	for i, char := range s {
		digit := char - '0'
		if digit < 0 || digit > 9 {
			return nil, errors.New("convert error")
		}
		result[i] = byte(digit)
	}
	return result, nil
}

func getAudioLang(lang string) string {
	s := lang[:2]
	hasLang := []string{"en", "ja", "ru", "zh"}
	found := false
	for _, v := range hasLang {
		if v == s {
			found = true
			break
		}
	}
	if !found {
		s = i18n.DefaultLang[:2]
	}
	return s
}
