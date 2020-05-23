package orm

import (
	"math"
	"time"
)

//替换gorm Model ,
type Model struct {
	Id        int        `form:"id" json:"id" binding:"omitempty" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//获取开始游标
func GetPageStart(page, perPage int) (cursor int) {
	return int(math.Max(float64(page-1), 0)) * perPage
}

//like添加前后百分号
func AddPercent(name string) string {
	return "%" + name + "%"
}
