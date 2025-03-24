package comment

import (
	"SYSUCODER/boot/DAO"
	"SYSUCODER/boot/entity"
	"SYSUCODER/boot/model"
	"errors"
	"log"
)

type CommentPage struct {
	Comments []entity.Comment `json:"comments"`
	model.Page
}

func Select(condition model.CommentWhere, userId uint64, role entity.Role) (CommentPage, error) {
	if !condition.Page.Exist() {
		condition.Page.Set(1)
	}
	if !condition.Size.Exist() {
		condition.Size.Set(10)
	}
	if !condition.Status.Exist() {
		condition.Status.Set([]uint64{uint64(entity.CommentPublic)})
	} else {
		for _, v := range condition.Status.Value() {
			if entity.CommentStatus(v) < entity.CommentPublic {
				if role < entity.RoleAdmin {
					condition.UserId.Set(userId)
				}
			}
		}
	}
	comments, err := DAO.SelectComments(condition)
	if err != nil {
		log.Println(err)
		return CommentPage{}, errors.New("获取评论失败")
	}

	count, err := DAO.CountComments(condition)
	if err != nil {
		log.Println(err)
		return CommentPage{}, errors.New("获取统计失败")
	}
	cPage := CommentPage{
		Comments: comments,
		Page: model.Page{
			Page:  condition.Page.Value(),
			Size:  condition.Size.Value(),
			Total: count,
		},
	}

	return cPage, nil
}
