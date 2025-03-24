package model

import "SYSUCODER/boot/entity"

// 提交记录（提交信息+评测结果）
type Record struct {
	Submission entity.Submission  `json:"submission,omitempty"`
	Judgements []entity.Judgement `json:"judgements,omitempty"`
}
