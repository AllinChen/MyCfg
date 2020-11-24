package mycfg

import (
	"fmt"
	"io/ioutil"

	"regexp"

	"github.com/romberli/log"
)

//Read 主模块，输入配置文件名和配置表达符号，读取配置返回配置表
func Read(fileName, oc string) map[string]string {
	//CfgMap 配置表
	var CfgMap map[string]string
	CfgMap = make(map[string]string)
	results := ReGet("test.cfg", "=")
	for _, result := range results {
		CfgMap[result[1]] = result[2]

	}

	// for _, result := range results {
	// 	fmt.Println(CfgMap[result[1]])
	// }
	return CfgMap

}

//StringRead 配置文件读取模块
func StringRead(fileName string) string {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Warnf("读取文件失败:%#v", err)
		return ""
	}
	log.Info("读取配置文件成功")
	return string(f)
}

//ReGet 配置匹配模块，输入文件名和配置表达式符号，可以匹配“=”与“：”
func ReGet(fileName, oc string) (res [][]string) {
	cfgfile := StringRead(fileName)
	re := regexp.MustCompile(`([^` + oc + `]+)` + oc + `([^\n]+)\n`)
	//解释规则，解析正则表达式，如果成功则返回
	if re == nil {
		fmt.Println("error regexp")
		log.Warnf("ReGet函数输入错误，正则表示式异常")
	}
	//根据规则提取信息
	result := re.FindAllStringSubmatch(cfgfile, -1)
	if result == nil {

		log.Warnf("读取配置失败")
		return nil
	}
	log.Infof("读取配置成功")
	return result
}
