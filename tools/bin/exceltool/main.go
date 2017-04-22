package main

import (
	"flag"
	"fmt"

	"github.com/colefan/sailfish/tools/exceltool"
)

var paramVersion = flag.Bool("version", false, "Show version")
var paramServer = flag.Bool("s", false, "Export server config")
var paramClient = flag.Bool("c", false, "Export client config")
var paramSrc = flag.String("src", "", "src: excel path")
var paramOutput = flag.String("output", "", "output: export output dir")
var paramType = flag.String("t", "", "t: export file type support json|lua|go")
var paramPackage = flag.String("pkg", "config", "pkg: go config package")

//Version define export tools version
const Version = "1.0.0"

func main() {
	flag.Parse()
	if *paramVersion {
		fmt.Println(Version)
		return
	}
	tool := exceltool.NewExcelConfigExport()
	if *paramSrc == "" {
		fmt.Println("-src can't be empty")
		return
	}
	tool.SourcePath = *paramSrc
	if *paramOutput == "" {
		fmt.Println("-output can't be empty")
		return
	}
	tool.TargetPath = *paramOutput
	if *paramType == "" {
		fmt.Println("use default export file type json")
		tool.TargetType = "json"
	} else {
		tool.TargetType = *paramType
	}

	if *paramPackage == "" {
		tool.GoPackage = "gsconfig"
	} else {
		tool.GoPackage = *paramPackage
	}

	if *paramServer {
		err := tool.Export4Server()
		if err != nil {
			fmt.Println("Export server config error " + err.Error())
		} else {
			fmt.Println("Export server config success:)")
		}
	}

	if *paramClient {
		err := tool.Export4Client()
		if err != nil {
			fmt.Println("Export client config error " + err.Error())
		} else {
			fmt.Println("Export client config success:)")
		}
	}

}
