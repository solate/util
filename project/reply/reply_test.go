package reply

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	res, _ := Response()
	assert.Equal(t, "{\"code\":200000,\"message\":\"success\",\"data\":null}", string(res))
}

func TestSetCode(t *testing.T) {
	res, _ := Response(SetCode(400403))
	assert.Equal(t, "{\"code\":400403,\"message\":\"success\",\"data\":null}", string(res))
}

func TestSetMessage(t *testing.T) {
	res, _ := Response(SetMessage("params error"))
	assert.Equal(t, "{\"code\":200000,\"message\":\"params error\",\"data\":null}", string(res))
}

func TestSetCodeAndMessage(t *testing.T) {
	res, _ := Response(SetCodeAndMessage(400403, "params error"))
	assert.Equal(t, "{\"code\":400403,\"message\":\"params error\",\"data\":null}", string(res))
}

func TestSetData(t *testing.T) {
	res, _ := Response(SetData(map[string]interface{}{"foo": "bar"}))
	assert.Equal(t, "{\"code\":200000,\"message\":\"success\",\"data\":{\"foo\":\"bar\"}}", string(res))
}
