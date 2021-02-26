package util

import (
	"github.com/gin-gonic/gin"
)

func GetURLFromReqContext(context *gin.Context) string {
	urlParam := context.Param("domain")
	return urlParam
}
func GetWordFromReqContext(context *gin.Context) string {
	wordParam := context.Param("word")
	return wordParam
}
