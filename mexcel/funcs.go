package mexcel

import (
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/mehrdadnekopour/go-tools/helpers"
	"github.com/mehrdadnekopour/go-tools/mypes"
)

// CreateXLSX ...
func CreateXLSX(path, fileName, defaultSheet string) (xlsx *ExcelFile, merr mypes.Merror) {
	f := excelize.NewFile()
	// Create a new sheet.
	f.NewSheet(defaultSheet)

	if err := f.SaveAs(fileName); err != nil {
		merr.Set(true, err, mypes.HTTPInternalServerError)
		return
	}

	fullPath := fmt.Sprintf("%s/%s", path, fileName)
	xlsx = &ExcelFile{
		File: f,
		Path: fullPath,
	}

	return
}

// OpenFile ...
func OpenFile(path string) (xlsx *ExcelFile, merr mypes.Merror) {

	file, err := excelize.OpenFile(path)
	if err != nil {
		merr.Set(true, err, mypes.OpenFileFailed)
		return
	}

	xlsx = &ExcelFile{
		File: file,
		Path: path,
	}

	return
}

// Row ...
func Row(row int) int {
	return row - 1
}

// CheckEmptyRow ...
func CheckEmptyRow(row []string) bool {

	return CheckEmptyColumnsInRow(row, 5)
}

// CheckEmptyColumnsInRow ....
func CheckEmptyColumnsInRow(row []string, columnsCount int) bool {
	empty := false

	firstCols := row[0:columnsCount]
	s := strings.Join(firstCols, "")
	s = strings.Trim(s, " ")

	if len(s) == 0 {
		empty = true
	}
	return empty
}

// CheckEmptyColumns ....
func CheckEmptyColumns(row []string, cFrom, cTo int) bool {
	empty := false

	cols := row[cFrom:cTo]
	s := strings.Join(cols, "")
	s = strings.Trim(s, " ")

	if len(s) == 0 {
		empty = true
	}
	return empty
}

// ReadTable ... Table must have some standards :
// 1- First row as Title
// 2- Second row as Headers
// 3- All cells in table area must have value
// 4- if rowEnd = 0 > means read table until the last filled row
func (x *ExcelFile) ReadTable(sheet, alias string, id, rowBegin, rowEnd int, colBegin, colEnd ExcelColumn) (table *mypes.QuestionRow, merr mypes.Merror) {

	var e error
	rows := x.File.GetRows(sheet)
	table = &mypes.QuestionRow{ID: id, Alias: alias}

	fmt.Println(rows, table)

	rowBegin--

	if rowEnd == 0 {
		rowEnd = rowBegin + len(rows)
	}
	rowEnd--

	cBegin := int(colBegin)
	cEnd := int(colEnd)

	var questions mypes.QuestionRows

	headers := make(map[int]string)

	qrID := 1
	for i := rowBegin; i <= rowEnd; i++ {
		row := rows[i]
		if row[cBegin] == "" {
			break
		}

		if i == rowBegin {
			table.Title = row[cBegin]
			continue
		}

		questionRow := &mypes.QuestionRow{}
		questionRow.ID = qrID
		var properties mypes.Properties

		propertyID := 1
		for c := cBegin; c <= cEnd; c++ {
			cellValue := row[c]
			if cellValue == "" {
				break
			}
			fmt.Println(cellValue)

			if i == rowBegin+1 {
				headers[c] = cellValue
				continue
			} else {
				if c == cBegin {
					questionRow.Title = cellValue
					qrAlias := cellValue
					if questionRow.Title == "کل" {
						qrAlias = "total"
					}

					questionRow.Alias = qrAlias
				} else {
					property := &mypes.Property{}
					property.ID = propertyID
					property.Title = headers[c]
					alias := headers[c]
					if property.Title == "کل" {
						alias = "total"
					}

					property.Alias = alias

					if cellValue == "-" {
						// property.Value = cellValue
						continue
					} else if cellValue == "X" {
						property.Value = cellValue
					} else {
						property.Value, e = helpers.CastToDouble(cellValue)
						if e != nil {
							log := fmt.Sprintf("%s%d -> value= %s", ExcelColumn(c).String(), i+1, cellValue)
							fmt.Println(log)
						}
					}
					property.Alias = fmt.Sprintf("%s", ExcelColumn(c).String())

					properties = append(properties, property)
					propertyID++
				}
			}
		}

		if i == rowBegin+1 {
			continue
		}
		questionRow.Properties = properties
		questions = append(questions, questionRow)
		qrID++
	}

	table.Rows = questions
	return
}

// ReadTableWithVerticalTitles ...
func (x *ExcelFile) ReadTableWithVerticalTitles(sheet string, rowBegin, rowEnd int, colBegin, colEnd ExcelColumn) (tables map[string]*mypes.QuestionRow, merr mypes.Merror) {

	tables = make(map[string]*mypes.QuestionRow)
	rows := x.File.GetRows(sheet)

	rowBegin--

	if rowEnd == 0 {
		rowEnd = rowBegin + len(rows)
	}
	rowEnd--

	cBegin := int(colBegin)
	cEnd := int(colEnd)

	var headers []string

	currentTableHeader := ""
	var currentTableQr *mypes.QuestionRow
	var currentTableQrs mypes.QuestionRows

	currentTableID := 1
	currentTableRowsID := 1

	for r := rowBegin + 1; r <= rowEnd; r++ {
		row := rows[r]
		qr := &mypes.QuestionRow{}

		propID := 1
		var props mypes.Properties
		for c := cBegin; c <= cEnd; c++ {
			valCell := strings.TrimSpace(row[c])

			if r == rowBegin+1 {
				if c > cBegin {
					headers = append(headers, valCell)

				}
			} else {
				if c == cBegin {
					if valCell != "" {

						if currentTableHeader != "" {
							currentTableQr.Rows = currentTableQrs
							tables[currentTableQr.Title] = currentTableQr
						}

						currentTableHeader = valCell

						currentTableQr = new(mypes.QuestionRow)
						currentTableQr.ID = currentTableID
						currentTableQr.Title = valCell

						qrAlias := valCell
						if currentTableQr.Title == "کل" {
							qrAlias = "total"
						}
						currentTableQr.Alias = qrAlias

						currentTableQrs = mypes.QuestionRows{}

						currentTableRowsID = 1
						currentTableID++
					}
				}
				if c > cBegin {
					if c == cBegin+1 {
						qr.Title = valCell
						qr.Alias = valCell
						if valCell == "کل" {
							qr.Alias = "total"
						}

					} else {
						prop := &mypes.Property{ID: propID}
						prop.Title = headers[c-1]
						alias := headers[c-1]
						if prop.Title == "کل" {
							alias = "total"
						}
						prop.Alias = alias

						if valCell == "-" {
							continue
						} else if valCell == "X" {
							prop.Value = valCell
						} else {
							prop.Value, _ = helpers.CastToDouble(valCell)
						}

						props = append(props, prop)
						propID++
					}
				}
			}
		}

		qr.Properties = props

		qr.ID = currentTableRowsID
		currentTableQrs = append(currentTableQrs, qr)
		currentTableRowsID++
	}
	currentTableQr.Rows = currentTableQrs
	tables[currentTableQr.Title] = currentTableQr

	return
}
