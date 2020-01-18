package mexcel

import (
	excelizeV2 "github.com/360EntSecGroup-Skylar/excelize/v2"
)

// GetCollName ...
func GetCollName(c int) (colName string) {
	colName, _ = excelizeV2.ColumnNumberToName(c)
	return
}

// GetCollNumber ...
func GetCollNumber(colName string) (number int) {
	number, _ = excelizeV2.ColumnNameToNumber(colName)
	return
}

// GetCellName ...
func GetCellName(col, row int) (cellName string) {
	cellName, _ = excelizeV2.CoordinatesToCellName(col, row)
	return
}
