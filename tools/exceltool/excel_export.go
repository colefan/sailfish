package exceltool

import (
	"strings"

	"errors"

	"os"
	"path/filepath"

	"strconv"

	"github.com/colefan/sailfish/logg"
	"github.com/tealeg/xlsx"
)

const (
	flagRowIndex  = 0
	descRowIndex  = 1
	nameRowIndex  = 2
	valueRowIndex = 3
)
const (
	leastRowCount = 4 //  最少行数:flag,desc,name,value
)

func isClient(flag string) bool {
	if strings.HasPrefix(flag, "A") || strings.HasPrefix(flag, "C") {
		return true
	}
	return false

}

func isServer(flag string) bool {
	if strings.HasPrefix(flag, "A") || strings.HasPrefix(flag, "S") {
		return true
	}
	return false
}

func isArray(flag string) bool {
	if strings.HasPrefix(flag, "A_a") || strings.HasPrefix(flag, "S_a") || strings.HasPrefix(flag, "C_a") {
		return true
	}

	return false
}

func isArray2D(flag string) bool {
	if strings.HasPrefix(flag, "A_aa") || strings.HasPrefix(flag, "S_aa") || strings.HasPrefix(flag, "C_aa") {
		return true
	}
	return false
}

func isArrayMap(flag string) bool {
	if strings.HasPrefix(flag, "A_am") ||
		strings.HasPrefix(flag, "C_am") ||
		strings.HasPrefix(flag, "S_am") {
		return true
	}
	return false
}

func isMap(flag string) bool {
	if strings.HasPrefix(flag, "A_m") ||
		strings.HasPrefix(flag, "S_m") ||
		strings.HasPrefix(flag, "C_m") {
		return true
	}
	return false
}

func isString(flag string) bool {
	if strings.HasPrefix(flag, "A_s") ||
		strings.HasPrefix(flag, "S_s") ||
		strings.HasPrefix(flag, "C_s") {
		return true
	}
	return false
}

func isFloat(flag string) bool {
	if strings.HasPrefix(flag, "A_f") ||
		strings.HasPrefix(flag, "S_f") ||
		strings.HasPrefix(flag, "C_f") {
		return true
	}
	return false
}

func isAuto(flag string) bool {
	if flag == "A" || flag == "C" || flag == "S" {
		return true
	}
	return false
}

func checkAutoType(value string) string {
	if value == "" {
		return "string"
	}
	value = strings.TrimSpace(value)
	_, err := strconv.Atoi(value)
	if err == nil {
		return "int"
	}

	_, err = strconv.ParseFloat(value, 10)
	if err == nil {
		return "float"
	}
	return "string"
}

//ExcelConfigExport export excel config
type ExcelConfigExport struct {
	SourcePath string
	TargetPath string
	TargetType string
	GoPackage  string
	log        *logg.BaseLogger
}

//NewExcelConfigExport create an instance
func NewExcelConfigExport() *ExcelConfigExport {
	e := &ExcelConfigExport{}
	e.log = logg.NewLogger(100)
	e.log.SetAppender("console", ``)
	return e
}

//Export4Server export excel config for server
func (e *ExcelConfigExport) Export4Server() error {
	e.log.Info("Begin to export excel for server.")
	err := os.Chdir(e.SourcePath)
	if err != nil {
		e.log.Error("SourcePath error " + err.Error())
		return err
	}
	err = os.Chdir(e.TargetPath)
	if err != nil {
		e.log.Error("TargetPath error " + err.Error())
		return err
	}
	if e.TargetType == "" {
		e.log.Error("TargetPath is empty")
		return errors.New("TargetPath is empty")
	}

	if e.TargetType == "json" {
		return exportToServerJSON(e)
	} else if e.TargetType == "go" {
		return exportToGo(e)
	} else {
		e.log.Error("Unsupport TargetType : " + e.TargetType)
		return errors.New("Unsupport TargetType : " + e.TargetType)
	}

}

//Export4Client export excel config for client
func (e *ExcelConfigExport) Export4Client() error {
	e.log.Info("Begin to export excel for client.")
	err := os.Chdir(e.SourcePath)
	if err != nil {
		e.log.Error("SourcePath error " + err.Error())
		return err
	}
	err = os.Chdir(e.TargetPath)
	if err != nil {
		e.log.Error("TargetPath error " + err.Error())
		return err
	}

	if e.TargetType == "" {
		e.log.Error("TargetPath is empty")
		return errors.New("TargetPath is empty")
	}

	if e.TargetType == "json" {
		return exportToClientJSON(e)

	} else if e.TargetType == "lua" {
		return exportToLua(e)

	} else {
		e.log.Error("Unsupport TargetType : " + e.TargetType)
		return errors.New("Unsupport TargetType : " + e.TargetType)
	}
}

type exportFunc func(xlsxFile *xlsx.File, filePath string) bool

func (e *ExcelConfigExport) doExport(export exportFunc) (failedCount int) {
	fileList := make([]string, 0, 32)
	filepath.Walk(e.SourcePath, func(filePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), ".xlsx") {
			fileList = append(fileList, filePath)
		}
		return nil
	})

	for _, filePath := range fileList {
		xlsxFile, err := xlsx.OpenFile(filePath)
		if err != nil {
			e.log.Error("open excel file " + filePath + " error :" + err.Error())
			failedCount++
			continue
		}
		e.log.Info("Begin to convert " + filePath)
		if !export(xlsxFile, filePath) {
			failedCount++
		}
	}
	return
}
