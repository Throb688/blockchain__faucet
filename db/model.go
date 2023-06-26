package db

import "time"

// 定义数据库的访问模型
type TODO struct {
	// 结构体成员tag，包含3种，一种是数据库，一种是json，一种是form表单
	ID      int        `sql:"id" json:"id" form:"id"`                // 唯一标识的
	ToAddress   string     `sql:"ToAddress" json:"title" form:"url"`       // 地址
	Created *time.Time `sql:"created" json:"created" form:"created"` // 交易的时间
}
