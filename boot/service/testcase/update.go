package testcase

import (
	"SYSUCODER/boot/DAO"
	"SYSUCODER/boot/entity"
	"errors"
	"log"
)

// 根据ID更新评测点数据
func UpdateById(t entity.Testcase) error {
	// 查询评测点
	t0, err := DAO.SelectTestcaseById(t.Id)
	if err != nil {
		log.Println(err)
		return errors.New("评测点不存在")
	}

	t0.Serial = t.Serial
	t0.ProblemId = t.ProblemId
	t0.TestInput = t.TestInput
	t0.TestOutput = t.TestOutput

	// 更新题目更新时间
	err = DAO.UpdateProblemUpdateTimeById(t.ProblemId)
	if err != nil {
		log.Println(err)
		return errors.New("更新题目更新时间失败")
	}

	// 更新评测点
	err = DAO.UpdateTestcaseById(t0)
	if err != nil {
		log.Println(err)
		return errors.New("更新评测点失败")
	}

	return nil
}
