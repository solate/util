package orm

import "testing"

func TestNewOrm(t *testing.T) {

	orm := NewOrm("127.0.0.1", "3306", "test", "root", "root")
	orm.SetOption(10, 20, 300, true)
	db, err := orm.Init()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(db.DB())

}
