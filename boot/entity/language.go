package entity

type LanguageStatus uint8

const (
	LanguageDeprecated LanguageStatus = 1
	LanguageDisabled   LanguageStatus = 2
	LanguageEnabled    LanguageStatus = 3
)

// 类型 LanguageStatus 返回字符串的方法
func (s LanguageStatus) String() string {
	switch s {
	case LanguageDeprecated:
		return "弃用"
	case LanguageDisabled:
		return "禁用"
	case LanguageEnabled:
		return "启用"
	default:
		return "未知"
	}
}

// Language 编程语言
type Language struct {
	Id     uint64         `gorm:"primaryKey;autoIncrement;comment:语言ID" json:"id"`
	Name   string         `gorm:"type:varchar(255);not null;comment:语言名" json:"name"`
	Serial uint16         `gorm:"not null;default:0;comment:排序序号" json:"serial"`
	MapId  uint32         `gorm:"not null;default:0;comment:映射ID" json:"map_id,omitempty"`
	Status LanguageStatus `gorm:"not null;default:3;comment:状态" json:"status"`
}

// TableName 为 Language 实现表名方法
func (Language) TableName() string {
	return "tbl_language"
}
