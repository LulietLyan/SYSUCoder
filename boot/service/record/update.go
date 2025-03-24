package record

import (
	"SYSUCODER/boot/DAO"
	"SYSUCODER/boot/entity"
	"errors"
	"log"
	"time"
)

// 更新提交记录
func UpdateSubmissionById(s entity.Submission) error {
	updateTime := time.Now()
	s.UpdateTime = updateTime

	err := DAO.UpdateSubmissionById(s)
	if err != nil {
		return err
	}

	return nil
}

// 更新评测结果
func UpdateJudgementById(j entity.Judgement) error {
	err := DAO.UpdateJudgementById(j)
	if err != nil {
		return err
	}

	// 更新提交记录状态更新时间
	err = DAO.UpdateSubmissionUpdateTimeById(j.SubmissionId)
	if err != nil {
		log.Println(err)
		return errors.New("更新提交记录状态更新时间失败")
	}

	return nil
}
