package pbtool

import "github.com/tealeg/xlsx"
import "github.com/colefan/sailfish/logg"
import "errors"
import "os"

var defaultType = []string{"string", "uint32", "int32", "uint64", "bool"}

//PbTool protocol buffer tool
type PbTool struct {
	SourcePath   string //描述pb的excel所在的文件
	TargetPath   string //pb格式文件生成路径
	ClientHeader string //客户端协议头文件
	ServerHeader string //服务端协议头文件
	log          *logg.BaseLogger
}

//NewPbTool create PbTool
func NewPbTool() *PbTool {
	pb := &PbTool{}
	pb.log = logg.NewLogger(100)
	pb.log.SetAppender("console", "")
	return pb
}

func (pb *PbTool) loadSheet(sheet *xlsx.Sheet) {

}

func (pb *PbTool) parseProto() {

}

func (pb *PbTool) generateProto() {

}

func (pb *PbTool) generateClientHeader() {

}

func (pb *PbTool) generateServerHeader() {

}

//Execute exec generate proto message and headers
func (pb *PbTool) Execute() error {
	pb.log.Info("Begin to execute proto message parsing")
	if pb.SourcePath == "" {
		pb.log.Error("SourcePath is nil")
		return errors.New("SourcePath is nil")
	}
	if pb.TargetPath == "" {
		pb.log.Error("TargetPath is nil")
		return errors.New("TargetPath is nil")
	}
	if pb.ClientHeader == "" {
		pb.log.Error("ClientHeader is nil")
		return errors.New("ClientPath is nil")
	}
	if pb.ServerHeader == "" {
		pb.log.Error("ServerHeader is nil")
		return errors.New("ServerHeader is nil")
	}
	if _, err := os.Stat(pb.SourcePath); err != nil {
		pb.log.Error("SourcePath not exists " + pb.SourcePath)
		return errors.New("SourcePath not exists")
	}
	if _, err := os.Stat(pb.TargetPath); err != nil {
		pb.log.Error("TargetPath not exists " + pb.TargetPath)
		return errors.New("TargetPath not exists")
	}
    

	return nil

}
