package controllers

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"log"
	"os/exec"
	"time"
)

// CMS API
type CMSController struct {
	beego.Controller
}

func (c *CMSController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("AllBlock", c.AllBlock)
}

// @router /staticblock/:key [get]
func (this *CMSController) StaticBlock() {

}

// @router /all/:key [get]
func (this *CMSController) AllBlock() {

}

// @router /info.html [get]
func (this *CMSController) Info() {
	this.TplName = "info.html"
}

// @router /list.html [get]
func (this *CMSController) List() {
	this.TplName = "list.html"
}
func run() {
	cmd := exec.Command("/bin/sh", "-c", "ping 127.0.0.1")
	_, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	}

	if err := cmd.Start(); err != nil {
		panic(err.Error())
	}

	if err := cmd.Wait(); err != nil {
		panic(err.Error())
	}

}

// @router /gitpull [get]
func (this *CMSController) Gitpull() {
	//res := system("who ")
	//ll 不能用会直接退出换成ls
	//res := system("ls $GOPATH/src/medexam")
	res := system("cd $GOPATH/src/medexam && git pull")
	this.Ctx.WriteString(res)
	return
}

// @router /shell [get]
func (this *CMSController) Shell() {

	go run()
	time.Sleep(1e9)

	cmd := exec.Command("/bin/sh", "-c", `ps -ef | grep -v "grep" | grep "ping"`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("StderrPipe: ", err.Error())
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Start: ", err.Error())
		return
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println("ReadAll stderr: ", err.Error())
		return
	}

	if len(bytesErr) != 0 {
		fmt.Printf("stderr is not nil: %s", bytesErr)
		return
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll stdout: ", err.Error())
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait: ", err.Error())
		return
	}

	fmt.Printf("stdout: %s", bytes)
	jsoninfo := this.GetString("jsoninfo")
	if jsoninfo == "" {
		this.Ctx.WriteString("jsoninfo is empty")
		return
	}
}

//调用系统指令的方法，参数s 就是调用的shell命令
func system(s string) string {
	cmd := exec.Command("/bin/sh", "-c", s) //调用Command函数
	var out bytes.Buffer                    //缓冲字节

	cmd.Stdout = &out //标准输出
	err := cmd.Run()  //运行指令 ，做判断
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String()) //输出执行结果
	return out.String()
}
