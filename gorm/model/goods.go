package model

import "time"

// Goods 商品表
type Goods struct {
	UID         string    `gorm:"primary_key;type:char(16);not null"`             // 主键ID
	Name        string    `gorm:"type:varchar(64);not null"`                      // 商品名
	Description string    `gorm:"type:varchar(255);not null"`                     // 描述
	Price       float64   `gorm:"type:decimal(10,2);not null"`                    // 商品价格
	Creator     string    `gorm:"type:varchar(64);not null"`                      // 创建人
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;type:datetime"` // 创建时间
}

func (g Goods) TableName() string {
	return "t_goods"
}
