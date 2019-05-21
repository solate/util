package reply

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	Success(c, gin.H{"foo": "bar"})

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"err_code\":200000,\"err_msg\":\"success\",\"data\":{\"foo\":\"bar\"}}", w.Body.String())
}

func TestError(t *testing.T) {

	// 未加 data
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	Error(c, http.StatusForbidden, 400403, "params error", nil)

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Equal(t, "{\"err_code\":400403,\"err_msg\":\"params error\",\"data\":null}", w.Body.String())

	// 加 data
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	Error(c, http.StatusForbidden, 400403, "params error", gin.H{"foo": "bar"})

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Equal(t, "{\"err_code\":400403,\"err_msg\":\"params error\",\"data\":{\"foo\":\"bar\"}}", w.Body.String())

}

func TestErrorFuc(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	ErrorFunc(c, http.StatusForbidden, SetCodeAndMsg(400403, "params error"))

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Equal(t, "{\"err_code\":400403,\"err_msg\":\"params error\",\"data\":null}", w.Body.String())

	// 加 data
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	ErrorFunc(c, http.StatusForbidden, SetCodeAndMsg(400403, "params error"), SetData(gin.H{"foo": "bar"}))

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Equal(t, "{\"err_code\":400403,\"err_msg\":\"params error\",\"data\":{\"foo\":\"bar\"}}", w.Body.String())

}
