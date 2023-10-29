package common

import (
	"api/i18n"
	"api/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetCtxValueString(ctx *gin.Context, key string, def ...string) string {
	s, exists := ctx.Get(key)
	if !exists {
		return ""
	}
	return utils.CString(s, def...)
}

func GetAcceptLang(ctx *gin.Context) string {
	lang := GetCtxValueString(ctx, "lang")
	if lang != "" {
		return lang
	}
	HLang := ctx.GetHeader("Accept-Language")
	if HLang != "" {
		split := strings.Split(HLang, ",")
		for _, val := range split {
			val = strings.TrimSpace(val)
			tl := i18n.L[val]
			if tl != nil {
				lang = val
				break
			}
		}
	}
	if lang == "" {
		lang = i18n.DefaultLang
	}
	ctx.Set("lang", lang)
	return lang
}
