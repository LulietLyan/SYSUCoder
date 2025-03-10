package DAO

import (
	"SYSUCODER/boot/entity"
	"SYSUCODER/boot/model"
)

type auxiliaryBlog struct {
	entity.Blog
	model.BriefUser
	model.BriefProblem
}
