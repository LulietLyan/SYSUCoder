package model

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

// 用于动态生成 OJ 判题系统测试用例的模块
// 其核心功能是根据配置的结构化规则递归展开并生成随机数据

var DatamakeLimit = uint64(10000)

var currentLimit uint64

// 定义值的生成规则，包括Type、最大/最小值及其对应的引用 ID(Min Max MinId MaxId)
type CommonTestcaseValue struct {
	ValueSizeId uint64  `json:"value_size_id,omitempty"`
	Type        string  `json:"type,omitempty"`
	Max         float64 `json:"max,omitempty"`
	Min         float64 `json:"min,omitempty"`
	MaxId       uint64  `json:"max_id,omitempty"`
	MinId       uint64  `json:"min_id,omitempty"`
}

// 表示一行数据，包含行大小标识和多个 CommonTestcaseValue
type CommonTestcaseRow struct {
	RowSizeId uint64                `json:"row_size_id,omitempty"`
	Values    []CommonTestcaseValue `json:"values,omitempty"`
}

// 表示整个测试用例的输入结构，包含多个 CommonTestcaseRow
type CommonTestcaseInput struct {
	Rows []CommonTestcaseRow `json:"rows,omitempty"`
}

// Unfold 生成单个随机值
func (c *CommonTestcaseValue) Unfold(hsh *[]float64) (DataMakerValue, error) {
	// 首先检查生成次数是否超过限制
	if currentLimit > DatamakeLimit {
		return DataMakerValue{}, errors.New("生成次数超过" + fmt.Sprint(DatamakeLimit))
	}
	currentLimit++

	// 通过 MaxId/MinId 从历史值数组中动态获取当前值的范围
	if c.MaxId > 0 && c.MaxId < uint64(len(*hsh)) {
		c.Max = (*hsh)[c.MaxId]
	}
	if c.MinId > 0 && c.MinId < uint64(len(*hsh)) {
		c.Min = (*hsh)[c.MinId]
	}

	// 获取类型信息并返回生成的值以及值的类型
	t := GetValueType(c.Type)
	v := rand.Float64()*(c.Max-c.Min) + c.Min
	*hsh = append(*hsh, v)
	return DataMakerValue{
		Type:  t,
		Value: v,
	}, nil
}

// Unfold 展开单行数据并生成多个值
func (c *CommonTestcaseRow) Unfold(hsh *[]float64) (DataMakerRow, error) {
	var row DataMakerRow

	// 遍历每个 Value
	for _, v := range c.Values {
		// 若 ValueSizeId 有效，根据历史值生成多个相同值
		// 否则生成单个值
		if v.ValueSizeId > 0 && v.ValueSizeId < uint64(len(*hsh)) {
			for i := 0; i < int((*hsh)[v.ValueSizeId]); i++ {
				value, err := v.Unfold(hsh)
				if err != nil {
					return row, err
				}
				row.AppendValue(value)
			}
		} else {
			value, err := v.Unfold(hsh)
			if err != nil {
				return row, err
			}
			row.AppendValue(value)
		}
	}
	return row, nil
}

// Unfold 展开整个输入结构，生成最终的测试数据
func (c *CommonTestcaseInput) Unfold() (DataMakerInput, error) {
	// 初始化全局计数器以及随机种子
	currentLimit = 0
	rand.Seed(uint64(time.Now().UnixNano()))
	var input DataMakerInput
	// 维护历史值数组，用于动态引用已生成的值
	var hsh []float64
	hsh = append(hsh, 0)

	// 遍历每个 Row
	for _, row := range c.Rows {
		// 若 RowSizeId 有效（存在对应的历史值），则根据历史值生成多行数据
		// 否则直接生成一行数据
		if row.RowSizeId > 0 && row.RowSizeId < uint64(len(hsh)) {
			for i := 0; i < int(hsh[row.RowSizeId]); i++ {
				newRow, err := row.Unfold(&hsh)
				if err != nil {
					return input, err
				}
				input.AppendRow(newRow)
			}
		} else {
			newRow, err := row.Unfold(&hsh)
			if err != nil {
				return input, err
			}
			input.AppendRow(newRow)
		}
	}

	return input, nil
}
