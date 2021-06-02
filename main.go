package main

import (
	"bufio"
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"strings"

	//"os"
)

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func GetIp() (string, error){

	is, err := net.Interfaces()
	if err != nil{
		return "", fmt.Errorf("get interfaces error, %v", err)
	}

	for _, v := range is{
		if v.Name == "Netease UU PPP Connection" {
			adds, err := v.Addrs()
			if err != nil{
				panic(err)
			}
			fmt.Printf("i:%v, add=%v\n", v, adds[0])
			return strings.Split(adds[0].String(),"/")[0], nil
		}
	}

	return "", fmt.Errorf("找不到 uu加速器")
}

func main()  {
	ip, err := GetIp()
	fmt.Println(ip, err)

	//if err == nil {
	//	setRouteCmd := exec.Command("netsh", "interface", "ip", "show");
	//	stdout, err := setRouteCmd.StdoutPipe()
	//	if err != nil{
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	err = setRouteCmd.Start()
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	if opBytes, err := ioutil.ReadAll(stdout); err != nil{
	//		fmt.Println(err)
	//		return
	//	}else{
	//		str := ConvertToString(string(opBytes), "gbk", "utf-8")
	//		fmt.Println(str)
	//	}
	//}

	if err == nil{
		setRouteCmd := exec.Command("route","add", "0.0.0.0", "mask", "0.0.0.0" , ip);
		stdout, err := setRouteCmd.StdoutPipe()
		if err != nil{
			fmt.Println(err)
			return
		}
		// 执行
		err = setRouteCmd.Start()
		if err != nil{
			fmt.Println(err)
			//return
		}
		// 读取返回结果
		if opBytes, err := ioutil.ReadAll(stdout); err != nil{
			fmt.Println(err)
			return
		}else{
			str := ConvertToString(string(opBytes), "gbk", "utf-8")
			fmt.Println("route xx")
			fmt.Println(str)
		}
		setRouteCmd.Wait()

		// 设置dns
		dnsCmd := exec.Command("netsh","interface", "ip", "set","dns", "Netease UU PPP Connection", "static","119.29.29.29");
		stdout, err = dnsCmd.StdoutPipe()
		if err != nil{
			fmt.Println(err)
			return
		}
		// 执行
		err = dnsCmd.Start()
		if err != nil{
			fmt.Println(err)
			//return
		}
		// 读取返回结果
		if opBytes, err := ioutil.ReadAll(stdout); err != nil{
			fmt.Println(err)
			return
		}else{
			str := ConvertToString(string(opBytes), "gbk", "utf-8")
			fmt.Println("set dns1")
			fmt.Println(str)
		}
		dnsCmd.Wait()

		// 设置dns
		//dns2Cmd := exec.Command("netsh","interface", "ip", "add", "dns", "Netease UU PPP Connection", "182.254.116.116", "index=2");
		//stdout, err = dns2Cmd.StdoutPipe()
		//if err != nil{
		//	fmt.Println(err)
		//	return
		//}
		//// 执行
		//err = dns2Cmd.Start()
		//if err != nil{
		//	fmt.Println(err)
		//	//return
		//}
		//// 读取返回结果
		//if opBytes, err := ioutil.ReadAll(stdout); err != nil{
		//	fmt.Println(err)
		//	return
		//}else{
		//	str := ConvertToString(string(opBytes), "gbk", "utf-8")
		//	fmt.Println("set dns2")
		//	fmt.Println(str)
		//}
		//dns2Cmd.Wait()

	}

	stdi := bufio.NewReader(os.Stdin)

	fmt.Print("设置成功，敲入任意按键退出：")
	wmsg, _ := stdi.ReadString('\n')
	fmt.Println(wmsg)
	return

}
