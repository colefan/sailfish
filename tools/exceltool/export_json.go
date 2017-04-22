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

func exportToServerJSON(e *ExcelConfigExport) error {
	return exportJSON(e, true)
}

func exportToClientJSON(e *ExcelConfigExport) error {
	return exportJSON(e, false)
}

func exportJSON(e *ExcelConfigExport, bServer bool) error {
	//first clear target dir
	filepath.Walk(e.TargetPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".json") {
			os.Remove(path)
		}
		return nil
	})

	failedMap := make(map[string]string)

	failedCount := e.doExport(func(xlsxFile *xlsx.File, filePath string) bool {
		for _, sheet := range xlsxFile.Sheets {
			name := strings.ToLower(sheet.Name)
			name = strings.ToUpper(name[0:1]) + name[1:]
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
				if bServer {
					if isServer(strFlag) {
						bNeedField = true
						strName, _ = sheet.Rows[nameRowIndex].Cells[colIdx].String()
					}
				} else {
					if isClient(strFlag) {
						bNeedField = true
						strName, _ = sheet.Rows[nameRowIndex].Cells[colIdx].String()
					}
				}
				if bNeedField {
					fieldIndexList = append(fieldIndexList, colIdx)
					fieldNameList = append(fieldNameList, strName)
					if isArray2D(strFlag) {
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
			buf := bytes.NewBuffer([]byte(`{"` + name + `":` + "\r\n"))
			buf.WriteString("[\r\n")
			maxRow := len(sheet.Rows)

			for index := valueRowIndex; index < maxRow; index++ {
				strLine := "\t{"
				if len(sheet.Rows[index].Cells) <= fieldIndexList[0] {
					continue
				}
				firstValue, _ := sheet.Rows[index].Cells[fieldIndexList[0]].String()
				if len(firstValue) == 0 {
					continue
				}
				for idx := 0; idx < fieldCount; idx++ {
					strLine += `"` + fieldNameList[idx] + `":`
					colIndex := fieldIndexList[idx]
					strVal := ""
					if colIndex < len(sheet.Rows[index].Cells) {
						strVal, _ = sheet.Rows[index].Cells[colIndex].String()
					}

					strTmp, err := parseJSONValue(fieldTypeList[idx], strVal)
					if err != nil {
						e.log.Error("parse sheet %s ,row %d col %d error, type %s", name, index, colIndex, fieldTypeList[idx])
						if v, ok := failedMap[filePath]; ok {
							failedMap[filePath] = v + "; parse sheet[" + name + "] error ,row = " + strconv.Itoa(index+1) + " ,col = " + strconv.Itoa(colIndex) + ",type =" + fieldTypeList[idx]

						} else {
							failedMap[filePath] = " parse sheet[" + name + "] error ,row = " + strconv.Itoa(index+1) + " ,col = " + strconv.Itoa(colIndex) + ",type =" + fieldTypeList[idx]
						}
						return false
					}
					strLine += strTmp + `, `
				}
				if len(strLine) > 1 {
					strLine = strLine[0 : len(strLine)-2]
				}
				strLine += "},\r\n"
				buf.WriteString(strLine)
			}
			buf.Truncate(buf.Len() - 3)
			buf.WriteString("\r\n")
			buf.WriteString("]\r\n")
			buf.WriteString("}\r\n")

			//对分析所得数据进行存储
			outputFilePath := e.TargetPath
			if strings.HasSuffix(outputFilePath, string(filepath.Separator)) {
				outputFilePath += strings.ToLower(name) + ".json"
			} else {
				outputFilePath += "" + string(filepath.Separator) + strings.ToLower(name) + ".json"
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

func parseJSONValue(fieldType string, strValue string) (string, error) {
	switch fieldType {
	case "array2d":
		if len(strValue) == 0 {
			return "[]", nil
		}
		valueList := strings.Split(strValue, ",")
		strReturn := "["
		for _, value := range valueList {
			strReturn += "["
			arrayList := strings.Split(value, "|")
			for _, item := range arrayList {
				strReturn += item + ","
			}
			if strings.HasSuffix(strReturn, ",") {
				strReturn = strReturn[0 : len(strReturn)-1]
			}
			strReturn += "],"

		}
		if strings.HasSuffix(strReturn, ",") {
			strReturn = strReturn[0 : len(strReturn)-1]
		}
		strReturn += "]"
		return strReturn, nil
	case "array":
		strReturn := "["
		arrayList := strings.Split(strValue, "|")
		for _, item := range arrayList {
			strReturn += item + ","
		}
		if strings.HasSuffix(strReturn, ",") {
			strReturn = strReturn[0 : len(strReturn)-1]
		}
		strReturn += "]"
		return strReturn, nil
	case "map":
		if len(strValue) == 0 {
			return `{}`, nil
		}
		strReturn := `{`
		valueList := strings.Split(strValue, ",")
		for _, value := range valueList {
			mapList := strings.Split(value, "|")
			if len(mapList) != 2 {
				return "{}", errors.New("_m value invalid")
			}
			strReturn += `"` + mapList[0] + `":` + mapList[1] + ","
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
