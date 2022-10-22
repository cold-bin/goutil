// @author cold bin
// @date 2022/10/22

package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResJson http api 请求返回json数据封装
type ResJson struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

var (
	resOk = ResJson{
		Code: 10000,
		Msg:  "请求成功",
		Data: nil,
	}

	resData = ResJson{
		Code: 10001,
		Msg:  "获取数据成功",
		Data: nil, // 待填充
	}

	resInterErr = ResJson{
		Code: 10002,
		Msg:  "服务繁忙",
		Data: nil,
	}

	resOpInfo = ResJson{
		Code: 10003,
		Msg:  "", // 待填充
		Data: nil,
	}
)

func ResOk(c *gin.Context) {
	c.JSON(http.StatusOK, resOk)
}

func ResOkWithData(c *gin.Context, data any) {
	resData.Data = data
	c.JSON(http.StatusOK, resData)
}

// ResInternalErr 内部出现错误
func ResInternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, resInterErr)
}

// ResOpInfo 返回友好的操作提示
func ResOpInfo(c *gin.Context, msg string) {
	resOpInfo.Msg = msg
	c.JSON(http.StatusOK, resOpInfo)
}
