package base

import (
	"gorm.io/gorm"
)

type FileName struct {
	gorm.Model
	File_name string	`json:"file_name"`
	Id uint	`gorm:"-;primary_key;AUTO_INCREMENT"`
	Column1 string	`json:"column_1"`
	Column2 string	`json:"column_2"`
	Column3 string	`json:"column_3"`
	Column4 string	`json:"column_4"`
	Column5 string	`json:"column_5"`
	Column6 string	`json:"column_6"`
	Column7 string	`json:"column_7"`
	Column8 string	`json:"column_8"`
	Column9 string	`json:"column_9"`
	Column10 string	`json:"column_10"`
	Column11 string	`json:"column_11"`
	Column12 string	`json:"column_12"`
	Column13 string	`json:"column_13"`
	Column14 string	`json:"column_14"`
	Column15 string	`json:"column_15"`
	Column16 string	`json:"column_16"`
	Column17 string	`json:"column_17"`
	Column18 string	`json:"column_18"`
	Column19 string	`json:"column_19"`
	Column20 string	`json:"column_20"`

}
type GetFileName struct {
	File_name string
	Id uint
}
type Body struct {
	 Data 	[][]string	`json:"data"`
	 Name	string	`json:"file_name"`
}