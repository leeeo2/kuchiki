package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leeexeo/kuchiki/pkg/common/exception"
)

func RespHandleMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		//取出output和error
		resp := gin.H{}
		e, ok := c.Get("error")
		if ok && e != nil {
			err := e.(exception.Error)
			resp["RespMeta"] = map[string]interface{}{
				"RequestId": "xxx",
			}
			resp["Error"] = map[string]interface{}{
				"ErrCode": err.ErrName,
				"Msg":     err.Error(),
			}
			//TODO:自定义error
			c.JSON(err.HttpCode, resp)
		} else {
			res, ok := c.Get("result")
			if ok && res != nil {
				resp["RespMeta"] = map[string]interface{}{
					"RequestId": "xxx",
				}
				resp["Result"] = res
				c.JSON(http.StatusOK, resp)
			}
		}
	}
}
