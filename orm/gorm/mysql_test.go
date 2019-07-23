package gorm

import (
	"testing"
)

func TestInitMysql(t *testing.T) {

	db, err := InitMysql(&Mysql{
		Host:     "localhost",
		Port:     "3306",
		Database: "test",
		Username: "root",
		Password: "root",
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(db.DB())
	t.Log("success")

}

func TestInitMysqlOptions(t *testing.T) {

	db, err := InitMysql(&Mysql{
		Host:     "localhost",
		Port:     "3306",
		Database: "test",
		Username: "root",
		Password: "root",
		//options
		Options: &Options{
			MaxIdleConns:    1,
			MaxOpenConns:    2,
			ConnMaxLifetime: 3,
			IsDebug:         true,
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(db.DB())
	t.Log("success")

}
