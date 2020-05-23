package gofile

import "testing"

//写内容测试
func TestWriteFile(t *testing.T) {

	err := WriteFile("./test.txt", "hello world!")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success")
}
