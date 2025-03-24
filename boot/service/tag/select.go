package tag

import (
	"SYSUCODER/boot/DAO"
	"SYSUCODER/boot/entity"
	"SYSUCODER/boot/model"
	"errors"
	"log"
)

type TagPage struct {
	Tags []entity.Tag `json:"tags"`
	model.Page
}

// 根据ID查询标签
func SelectById(id uint64) (entity.Tag, error) {
	t, err := DAO.SelectTagById(id)
	if err != nil {
		return entity.Tag{}, err
	}

	return t, nil
}

// 根据题目ID查询标签
func SelectByProblemId(pid uint64) ([]entity.Tag, error) {
	tags, err := DAO.SelectTagsByProblemId(pid)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

// 查询题目标签关系是否存在
func CountProblemTag(pid uint64, tid uint64) (int64, error) {
	pt := entity.ProblemTag{
		ProblemId: pid,
		TagId:     tid,
	}

	count, err := DAO.CountProblemTag(pt)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// 查询所有用户
func Select(condition model.TagWhere) (TagPage, error) {
	if !condition.Page.Exist() {
		condition.Page.Set(1)
	}
	if !condition.Size.Exist() {
		condition.Size.Set(10)
	}
	users, err := DAO.SelectTags(condition)
	if err != nil {
		log.Println(err)
		return TagPage{}, errors.New("查询用户失败")
	}

	count, err := DAO.CountTags(condition)
	if err != nil {
		log.Println(err)
		return TagPage{}, errors.New("查询统计失败")
	}
	uPage := TagPage{
		Tags: users,
		Page: model.Page{
			Total: count,
			Page:  condition.Page.Value(),
			Size:  condition.Size.Value(),
		},
	}

	return uPage, nil
}
