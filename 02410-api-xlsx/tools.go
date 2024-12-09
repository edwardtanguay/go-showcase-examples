package main

import (
	"github.com/tealeg/xlsx"
)

func generateXslxFile() (*xlsx.File, error) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Info")
	if err != nil {
		return nil, err
	}

	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "Product Number"
	cell = row.AddCell()
	cell.Value = "837482"

	return file, nil
}
