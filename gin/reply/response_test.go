package reply

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {

	// success 成功
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	Response(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"code\":200000,\"message\":\"success\",\"data\":null}", w.Body.String())
}

// 测试设置http状态码
func TestSetHttpCode(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	Response(c, SetHttpCode(http.StatusForbidden))

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Equal(t, "{\"code\":200000,\"message\":\"success\",\"data\":null}", w.Body.String())
}

// 测试设置错误码
func TestSetCode(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	Response(c, SetCode(400403))

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"code\":400403,\"message\":\"success\",\"data\":null}", w.Body.String())

}

func TestSetMsg(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	Response(c, SetMessage("params error"))

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"code\":200000,\"message\":\"params error\",\"data\":null}", w.Body.String())
}

func TestSetCodeAndMsg(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	Response(c, SetCodeMessage(400403, "params error"))

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"code\":400403,\"message\":\"params error\",\"data\":null}", w.Body.String())
}

func TestSetData(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	Response(c, SetData(gin.H{"foo": "bar"}))

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"code\":200000,\"message\":\"success\",\"data\":{\"foo\":\"bar\"}}", w.Body.String())

}
