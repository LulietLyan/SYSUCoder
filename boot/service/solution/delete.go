package solution

import (
	"SYSUCODER/boot/DAO"
	"errors"
	"log"
)

// 根据ID删除题解
func DeleteById(id uint64) error {
	// 查询题解
	_, err := DAO.SelectSolutionById(id)
	if err != nil {
		log.Println(err)
		return errors.New("题解不存在")
	}

	// 删除题解
	err = DAO.DeleteSolutionById(id)
	if err != nil {
		log.Println(err)
		return errors.New("删除题解失败")
	}

	return nil
}
