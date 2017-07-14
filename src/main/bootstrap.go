/*
程序入口,初始化
 */
package main

import (
	"os"
	"fmt"
)

//初始化函数
func init() {
	
}

//入口
func main() {
	file, err := os.OpenFile("./resources/test.txt", os.O_WRONLY | os.O_CREATE, 0)
	if err != nil {
		fmt.Println("err :",err)
	}
	file.WriteString("nihao")
	if err != nil {
		fmt.Println("err :",err)
	}
}

//加载配置文件
func loadConfig()  {

}