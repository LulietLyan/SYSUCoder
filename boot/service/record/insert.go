package record

import (
	"SYSUCODER/boot/DAO"
	"SYSUCODER/boot/entity"
	"errors"
	"time"
)

// 插入提交记录
func InsertSubmission(s entity.Submission) (uint64, error) {
	var err error

	updateTime := time.Now()
	s.UpdateTime = updateTime
	s.CreateTime = updateTime

	s.Id, err = DAO.InsertSubmission(s)
	if err != nil {
		return 0, err
	}

	return s.Id, nil
}

// 插入评测结果
func InsertJudgement(j entity.Judgement) (uint64, error) {
	var err error

	j.Id, err = DAO.InsertJudgement(j)
	if err != nil {
		return 0, err
	}

	// 更新提交记录状态更新时间
	err = DAO.UpdateSubmissionUpdateTimeById(j.SubmissionId)
	if err != nil {
		return 0, errors.New("更新提交记录状态更新时间失败")
	}

	return j.Id, nil
}
