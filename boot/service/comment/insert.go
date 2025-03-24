package comment

import (
	"SYSUCODER/boot/DAO"
	"SYSUCODER/boot/entity"
	"errors"
	"log"
	"time"
)

// 插入评论
func Insert(c entity.Comment) (uint64, error) {
	var err error

	updateTime := time.Now()
	c.UpdateTime = updateTime
	c.CreateTime = updateTime
	c.Status = entity.CommentPublic

	// 插入评论
	c.Id, err = DAO.InsertComment(c)
	if err != nil {
		log.Println(err)
		return 0, errors.New("插入评论失败")
	}

	return c.Id, nil
}
