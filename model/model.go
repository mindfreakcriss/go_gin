package model

type User struct {
	ID       int    `gorm:"column:ID;type:int;primaryKey;autoIncrement"`
	Name     string `gorm:"column:Name;type:varchar(100);not null;comment:用户名;uniqueIndex:idx_name"`
	Password string `gorm:"column:Password;type:varchar(100);not null;comment:密码"`
	NickName string `gorm:"column:Nickname;type:varchar(100);not null;comment:昵称"`
}

func (User) TableName() string {
	// 设置表名
	return "t_users"
}
