package exceltool

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

const (
	blankLine = "\r\n"
)

//for server
func exportToGo(e *ExcelConfigExport) error {
	//first clear target dir
	filepath.Walk(e.TargetPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, "_config.go") {
			os.Remove(path)
		}
		return nil
	})

	failedMap := make(map[string]string)

	failedCount := e.doExport(func(xlsxFile *xlsx.File, filePath string) bool {
		for _, sheet := range xlsxFile.Sheets {
			name := strings.ToLower(sheet.Name)
			name = strings.ToUpper(name[0:1]) + name[1:]
			clsName := name
			subClsName := name + "_Item"
			fieldNameList := make([]string, 0, 32)
			fieldTypeList := make([]string, 0, 32)
			fieldIndexList := make([]int, 0, 32)
			cellMax := 0
			if len(sheet.Rows) >= leastRowCount {
				cellMax = len(sheet.Rows[flagRowIndex].Cells)
			}

			for colIdx := 0; colIdx < cellMax; colIdx++ {
				strFlag, _ := sheet.Rows[flagRowIndex].Cells[colIdx].String()
				strName := ""
				bNeedField := false

				if isServer(strFlag) {
					bNeedField = true
					strName, _ = sheet.Rows[nameRowIndex].Cells[colIdx].String()
				}

				if bNeedField {
					fieldIndexList = append(fieldIndexList, colIdx)
					fieldNameList = append(fieldNameList, strName)
					if isArrayMap(strFlag) {
						fieldTypeList = append(fieldTypeList, "string")

					} else if isArray2D(strFlag) {
						fieldTypeList = append(fieldTypeList, "array2d")
					} else if isArray(strFlag) {
						fieldTypeList = append(fieldTypeList, "array")
					} else if isMap(strFlag) {
						fieldTypeList = append(fieldTypeList, "map")
					} else if isString(strFlag) {
						fieldTypeList = append(fieldTypeList, "string")
					} else if isFloat(strFlag) {
						fieldTypeList = append(fieldTypeList, "float")
					} else if isAuto(strFlag) {
						v, _ := sheet.Rows[valueRowIndex].Cells[colIdx].String()
						//fmt.Println("col = " + strName + "  v  =  " + v + "  type =" + checkAutoType(v))
						fieldTypeList = append(fieldTypeList, checkAutoType(v))
					} else {
						fieldTypeList = append(fieldTypeList, "int")

					}
				}

			}

			fieldCount := len(fieldIndexList)
			if fieldCount == 0 {
				e.log.Info("parse sheet " + name + "  -无需导出")
				continue
			} else {
				e.log.Info("parse sheet " + name)
			}

			//分析完数据头部，开始分析数据本身
			buf := bytes.NewBuffer([]byte("package " + e.GoPackage + "\r\n"))
			buf.WriteString("type " + clsName + " struct{\r\n")
			buf.WriteString("\tItems []*" + subClsName + "\r\n")
			buf.WriteString("}\r\n")
			buf.WriteString(blankLine)
			buf.WriteString("type " + subClsName + " struct{\r\n")
			for k := 0; k < fieldCount; k++ {
				buf.WriteString("\t" + convertGoPublicField(fieldNameList[k]) + " " + convertGoType(fieldTypeList[k]) + "\r\n")
			}
			buf.WriteString("}\r\n")
			buf.WriteString(blankLine)
			buf.WriteString("var Inst" + clsName + " *" + clsName + "\r\n")
			buf.WriteString(blankLine)
			buf.WriteString("func init(){\r\n")

			maxRow := len(sheet.Rows)
			bFirstRow := true

			for index := valueRowIndex; index < maxRow; index++ {
				strLine := ""
				if len(sheet.Rows[index].Cells) <= fieldIndexList[0] {
					continue
				}
				firstValue, _ := sheet.Rows[index].Cells[fieldIndexList[0]].String()
				if len(firstValue) == 0 {
					continue
				}
				if bFirstRow {
					strLine = "\titem := &" + subClsName + "{}\r\n"
					bFirstRow = false
				} else {
					strLine = "\titem = &" + subClsName + "{}\r\n"
				}
				for idx := 0; idx < fieldCount; idx++ {
					strLine += "\titem." + convertGoPublicField(fieldNameList[idx]) + " = "
					colIndex := fieldIndexList[idx]
					strVal := ""
					if colIndex < len(sheet.Rows[index].Cells) {
						strVal, _ = sheet.Rows[index].Cells[colIndex].String()
					}

					strTmp, err := parseGoValue(fieldTypeList[idx], strVal)
					if err != nil {
						e.log.Error("parse sheet %s ,row %d col %d error, type %s", name, index, colIndex, fieldTypeList[idx])
						if v, ok := failedMap[filePath]; ok {
							failedMap[filePath] = v + "; parse sheet[" + name + "] error ,row = " + strconv.Itoa(index+1) + " ,col = " + strconv.Itoa(colIndex) + ",type =" + fieldTypeList[idx]
						} else {
							failedMap[filePath] = " parse sheet[" + name + "] error ,row = " + strconv.Itoa(index+1) + " ,col = " + strconv.Itoa(colIndex) + ",type =" + fieldTypeList[idx]
						}
						return false
					}
					strLine += strTmp + "\r\n"
				}
				buf.WriteString(strLine)
				buf.WriteString("\tInst" + clsName + ".Items = append(Inst" + clsName + ".Items,item)\r\n")
			}
			buf.WriteString("}\r\n") //init end

			//对分析所得数据进行存储
			outputFilePath := e.TargetPath
			if strings.HasSuffix(outputFilePath, string(filepath.Separator)) {
				outputFilePath += strings.ToLower(name) + "_config.go"
			} else {
				outputFilePath += "" + string(filepath.Separator) + strings.ToLower(name) + "_config.go"
			}
			file, err := os.OpenFile(outputFilePath, os.O_CREATE|os.O_RDWR, 0660)
			if err != nil {
				e.log.Error("write file error %s", outputFilePath)
				return false
			}

			file.WriteString(buf.String())
			file.Close()

		} //~sheet

		return true
	})

	if failedCount > 0 {
		for k, v := range failedMap {
			e.log.Error(k + "--" + v)
		}
		return errors.New(strconv.Itoa(failedCount) + " file export failed")
	}

	return nil
}

func convertGoPublicField(name string) string {
	if len(name) > 1 {
		name = strings.ToUpper(name[0:1]) + name[1:]
	}
	return name
}

func convertGoType(fieldType string) string {
	switch fieldType {
	case "int":
		return "int"
	case "string":
		return "string"
	case "float":
		return "float32"
	case "array":
		return "[]int"
	case "map":
		return "map[string]int"
	case "array2d":
		return "[][]int"
	case "arraymap":
		return "[]map[int]int"
	default:
		return "string"
	}
}

func parseGoValue(fieldType string, strValue string) (string, error) {
	switch fieldType {
	case "array2d":
		if len(strValue) == 0 {
			return "make([][]int,0)", nil
		}
		valueList := strings.Split(strValue, ",")
		strReturn := "[][]int{"
		for _, value := range valueList {
			strReturn += "{"
			arrayList := strings.Split(value, "|")
			for _, item := range arrayList {
				strReturn += item + ","
			}
			if strings.HasSuffix(strReturn, ",") {
				strReturn = strReturn[0 : len(strReturn)-1]
			}
			strReturn += "},"

		}
		if strings.HasSuffix(strReturn, ",") {
			strReturn = strReturn[0 : len(strReturn)-1]
		}
		strReturn += "}"
		return strReturn, nil
	case "array":
		if len(strValue) == 0 {
			return "make([]int,0)", nil
		}
		strReturn := "[]int{"
		arrayList := strings.Split(strValue, "|")
		for _, item := range arrayList {
			strReturn += item + ","
		}
		if strings.HasSuffix(strReturn, ",") {
			strReturn = strReturn[0 : len(strReturn)-1]
		}
		strReturn += "}"
		return strReturn, nil
	case "map":
		if len(strValue) == 0 {
			return "make(map[string]int)", nil
		}
		strReturn := `map[string]int{`
		valueList := strings.Split(strValue, ",")
		for _, value := range valueList {
			mapList := strings.Split(value, "|")
			if len(mapList) != 2 {
				return "make(map[string]int)", errors.New("_m value invalid")
			}
			if mapList[1] == "" {
				strReturn += `"` + mapList[0] + `":` + "0,"
			} else {
				strReturn += `"` + mapList[0] + `":` + mapList[1] + ","

			}
		}
		if strings.HasSuffix(strReturn, ",") {
			strReturn = strReturn[0 : len(strReturn)-1]
		}
		strReturn += `}`
		return strReturn, nil
	case "string":
		return `"` + strValue + `"`, nil
	case "float":
		if len(strValue) == 0 {
			return "0.0", nil
		}
		_, err := strconv.ParseFloat(strValue, 10)
		if err != nil {
			return "", err
		}
		return strValue, nil
	case "int":
		if len(strValue) == 0 {
			return "0", nil
		}
		_, err := strconv.Atoi(strValue)
		if err != nil {
			return "", err
		}
		return strValue, nil
	}
	return "", errors.New("unsupport type " + fieldType)
}
