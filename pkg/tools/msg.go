package tools

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	MSG_ERR             = -1
	MSG_OK              = 0
	UNAUTHORIZED_ACCESS = 10000
)

// JSONResult json
type JSONResult struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	RequestId interface{} `json:"request_id"`
}

type JSONResultQ struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Total     int         `json:"total"`
	RequestId interface{} `json:"request_id"`
}

// JSON 渲染
func JSON(c *gin.Context, Code int, message string, data ...interface{}) {
	result := new(JSONResult)
	result.Code = Code
	result.Message = message

	if len(data) > 0 {
		result.Data = data[0]
	} else {
		result.Data = ""
	}
	c.JSON(http.StatusOK, result)
}

// Json 失败返回
func JSONFailed(c *gin.Context, Code int, message string) {
	result := new(JSONResult)
	result.Code = Code
	result.Message = message

	c.JSON(http.StatusOK, result)
}

// Json 成功返回
func JSONOk(c *gin.Context, data ...interface{}) {
	var result interface{}
	if len(data) > 0 {
		result = data[0]
	} else {
		result = ""
	}

	c.JSON(http.StatusOK, JSONResult{
		Code:      MSG_OK,
		Message:   "Ok",
		Data:      result,
		RequestId: c.Request.Header.Get("X-Request-Id"),
	})
}

func JSONokQ(c *gin.Context, total int, data ...interface{}) {
	var result interface{}
	if len(data) > 0 {
		result = data[0]
	} else {
		result = ""
	}

	c.JSON(http.StatusOK, JSONResultQ{
		Code:      MSG_OK,
		Message:   "Ok",
		Data:      result,
		Total:     total,
		RequestId: c.Request.Header.Get("X-Request-Id"),
	})
}

type Err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Err) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}

func New(code int, msg string) *Err {
	return &Err{
		Code:    code,
		Message: msg,
	}
}
