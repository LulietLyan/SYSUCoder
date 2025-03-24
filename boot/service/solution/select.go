package solution

import (
	"SYSUCODER/boot/DAO"
	"SYSUCODER/boot/entity"
	"errors"
)

// 根据ID查询题解
func SelectById(id uint64) (entity.Solution, error) {
	s, err := DAO.SelectSolutionById(id)
	if err != nil {
		return entity.Solution{}, errors.New("获取题解失败")
	}

	return s, nil
}

// 查询所有题解（不返回源代码）
func SelectAll() ([]entity.Solution, error) {
	solutions, err := DAO.SelectAllSolutions()
	if err != nil {
		return nil, err
	}

	return solutions, nil
}

// 根据题目ID查询题解（不返回源代码）
func SelectByProblemId(pid uint64) ([]entity.Solution, error) {
	solutions, err := DAO.SelectSolutionsByProblemId(pid)
	if err != nil {
		return nil, err
	}

	return solutions, nil
}

// 隐藏源代码
func hideSourceCode(solutions []entity.Solution) {
	for i := range solutions {
		solutions[i].SourceCode = ""
	}
}
