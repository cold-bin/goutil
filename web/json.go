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

	resInterErr = ResJson{
		Code: 10002,
		Msg:  "服务繁忙",
		Data: nil,
	}

	resParamErr = ResJson{
		Code: 10003,
		Msg:  "请求参数错误",
		Data: nil,
	}

	resPermissionDenied = ResJson{
		Code: 10005,
		Msg:  "权限不足",
		Data: nil,
	}

	resVerifiedErr = ResJson{
		Code: 10006,
		Msg:  "身份校验失败",
		Data: nil,
	}
)

func ResOk(c *gin.Context) {
	c.JSON(http.StatusOK, resOk)
}

func ResOkWithData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, ResJson{
		Code: 10001,
		Msg:  "获取数据成功",
		Data: data,
	})
}

// ResInternalErr 内部出现错误，很少调用
func ResInternalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, resInterErr)
}

func ResParamErr(c *gin.Context) {
	c.JSON(http.StatusOK, resParamErr)
}

// ResOpInfo 返回友好的操作提示，能直接返回给前端
func ResOpInfo(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, ResJson{
		Code: 10004,
		Msg:  msg,
		Data: nil,
	})
}

func ResPermissionDenied(c *gin.Context) {
	c.JSON(http.StatusOK, resPermissionDenied)
}

func ResVerifiedErr(c *gin.Context) {
	c.JSON(http.StatusOK, resVerifiedErr)
}

func ResStd(c *gin.Context, res ResJson) {
	c.JSON(http.StatusOK, res)
}
