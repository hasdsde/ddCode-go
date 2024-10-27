package resp

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Data interface{} `json:"data"`
	// 消息
	Msg       string `json:"message"`
	RequestID string `json:"requestID"`
	Code      int    `json:"code" example:"200"`
}

type Page struct {
	List      interface{} `json:"list"`
	Count     int         `json:"total"`
	PageIndex int         `json:"pageNo"`
	PageSize  int         `json:"pageSize"`
}

type PageData struct {
	List      interface{} `json:"list"`
	Data      interface{} `json:"data"`
	Count     int         `json:"total"`
	PageIndex int         `json:"pageNo"`
	PageSize  int         `json:"pageSize"`
}
type Graph struct {
	List interface{} `json:"list"`
}

type OkList struct {
	List interface{} `json:"list"`
}

// ReturnOK 正常返回
func (res *Response) ReturnOK() *Response {
	res.Code = 0
	return res
}

// ReturnError 错误返回
func (res *Response) ReturnError(code int) *Response {
	res.Code = code
	return res
}

// Error 通常错误数据处理
func Error(c *gin.Context, code int, err error, msg string) {
	var res Response

	if msg != "" {
		res.Msg = msg
	}
	if err != nil {
		res.Msg += ":" + err.Error()
	}
	Return := res.ReturnError(code)
	err = setResult(c, Return, res.RequestID, code)
	if err != nil {
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, Return)
}

// ErrorWithMessage 用户登录等位置，返回报错的message，而不是直接弹出422
func ErrorWithMessage(c *gin.Context, code int, err error, msg string) {
	var res Response
	if msg != "" {
		res.Msg = msg
	}
	if err != nil {
		res.Msg += ":" + err.Error()
	}
	Return := res.ReturnError(code)
	err = setResult(c, Return, res.RequestID, code)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, Return)
}

// ErrorWithData 常错误数据处理 带数据返回
func ErrorWithData(c *gin.Context, code int, data interface{}, err error, msg string) {
	var res Response
	res.Data = data
	if err != nil {
		res.Msg = err.Error()
	}
	if msg != "" {
		res.Msg = msg
	}
	Return := res.ReturnError(code)
	err = setResult(c, Return, res.RequestID, code)
	if err != nil {
		return
	}
	c.AbortWithStatusJSON(code, Return)
}

// OK 通常成功数据处理
func OK(c *gin.Context, data interface{}, msg string) {
	var res Response
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	Return := res.ReturnOK()
	err := setResult(c, Return, res.RequestID, 200)
	if err != nil {
		return
	}
	c.AbortWithStatusJSON(200, Return)
}

func OKList(c *gin.Context, data interface{}, msg string) {
	var res OkList
	res.List = data
	OK(c, res, msg)
}

// PageOK 分页数据处理
func PageOK(c *gin.Context, result interface{}, count int, pageIndex, pageSize int, msg string) {
	var res Page
	res.List = result
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	OK(c, res, msg)
}
func PageOKData(c *gin.Context, data, list interface{}, count, pageIndex, pageSize int, msg string) {
	var res PageData
	res.Data = data
	res.List = list
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	OK(c, res, msg)
}
func GraphOK(c *gin.Context, result interface{}, msg string) {
	var res Graph
	res.List = result
	OK(c, res, msg)
}

func setResult(c *gin.Context, returning interface{}, msgID string, status int) error {
	var jsonStr []byte
	jsonStr, err := json.Marshal(returning)
	if err != nil {
		return fmt.Errorf("setResult error", err)
	}
	c.Set("result", string(jsonStr))
	c.Set("status", status)
	return nil
}

// Custom 兼容函数
func Custom(c *gin.Context, data gin.H) error {
	Return := data
	var jsonStr []byte
	jsonStr, err := json.Marshal(Return)
	if err != nil {
		return fmt.Errorf("MsgID[%s] ShouldBind error: %#v", err.Error())
	}
	c.Set("result", string(jsonStr))
	c.AbortWithStatusJSON(0, Return)
	return nil
}
