package retry

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	err := Do(Register, 5, 1*time.Second)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("success")
	}

}

var i int

// retry func
func Register() (err error) {

	fmt.Println(i)
	err = errors.New("test")
	if i == 2 { //3rd success 0,1,2
		return nil
	}
	i++
	return
}
