package model

import "gorm.io/gorm"

type SolutionWhere struct {
}

// 题解
type Solution struct {
	Language    string `json:"language"`
	SourceCode  string `json:"source_code"`
	Explanation string `json:"explanation"`
}

// 题解说明
type SolutionInstruction struct {
	Title        string   `json:"title" binding:"omitempty"`
	Description  string   `json:"description" binding:"omitempty"`
	Input        string   `json:"input" binding:"omitempty"`
	Output       string   `json:"output" binding:"omitempty"`
	SampleInput  string   `json:"sample_input" binding:"omitempty"`
	SampleOutput string   `json:"sample_output" binding:"omitempty"`
	Hint         string   `json:"hint" binding:"omitempty"`
	Tags         []string `json:"tags" binding:"omitempty"`
	Solution     string   `json:"solution" binding:"omitempty"`
	Language     string   `json:"language" binding:"omitempty"`
}

func (con *SolutionWhere) GenerateWhere() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}
		where := db.Where(whereClause)
		query := []string{"tbl_solution.*"}
		return where.Select(query)
	}
}
