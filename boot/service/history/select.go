package history

import (
	"SYSUCODER/boot/DAO"
	"SYSUCODER/boot/entity"
	"errors"
	"log"
)

// 根据题目ID查询历史记录
func SelectHistoriesByProblemId(pid uint64) ([]entity.History, error) {
	histories, err := DAO.SelectHistoriesByProblemId(pid)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取题目历史记录失败")
	}

	return histories, nil
}
