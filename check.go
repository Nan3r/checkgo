package main

import (
  "fmt"
  "runtime"
  "time"
  "os"
  "os/user"
  "github.com/shirou/gopsutil/mem"
  "github.com/shirou/gopsutil/disk"
  "github.com/shirou/gopsutil/process"
)

//判断磁盘信息
func GetDisk()bool {
	diskPart,err := disk.Partitions(false)
	if err != nil {
		return false
	}else{
		if(len(diskPart) == 1){
			diskUsed,_ := disk.Usage(diskPart[0].Mountpoint)
			res := diskUsed.Total/1024/1024/1024
			//fmt.Printf("分区总大小: %d GB \n",res)
			if(res > 450){
				//fmt.Println("GetDisk fine")
				return true
			}else{
				return false
			}
		}else{
			//fmt.Println("GetDisk fine")
			return true
		}
	}

}

//判断内存大小，大于2G
func GetMemPercent()bool {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.Total > 2000000000
}


//判断CPU核数，大于2
func GetCpuCount()bool {
	c := runtime.GOMAXPROCS(0)
	return c >= 2
}

//判断机器名是否是沙箱的机器名
func NoBlockComputerName()bool {
	known := []string{
			"SANDBOX",
			"7SILVIA",
			"HANSPETER-PC",
			"JOHN-PC",
			"MUELLER-PC",
			"WIN7-TRAPS",
			"FORTINET",
			"TEQUILABOOMBOOM",
			"VBCCSC-PC",
			"DESKTOP-SVONXYD",
			"WIN-2HBXSRKWCRY",
			"WIN-2HBXSRKWCRY",
			"WIN-IVE99JTTEQ6",
			"WIN-HHQMQDCBT7E",
			"0CC47AC83803",
			"AMAZING-AVOCADO",
			"rbmhuwvcing",
			"STACAS84",
			"SDJ-FFD0FEB05DC"}
	name, _ := os.Hostname()

	for _,v :=range known{
		if (v == name){
			return false
		}
	}
	//fmt.Println("NoBlockComputerName fine")
	return true
	
}

//判断用户名是否是沙箱的用户名
func NoBlockUserName()bool {
	known := []string{
		"CurrentUser",
		"Sandbox",
		"Emily",
		"HAPUBWS",
		"Hong Lee",
		"IT-ADMIN",
		"Johnson",
		"Miller",
		"milozs",
		"Peter Wilson",
		"timmy",
		"sand box",
		"malware",
		"maltest",
		"test user",
		"virus",
		"John Doe",
		"vbccsb",
		"jason",
		"jojo",
		"lichao"}
	name, _ := user.Current()

	for _,v :=range known{
		if (v == name.Username){
			return false
		}
	}
	//fmt.Println("NoBlockUserName fine")
	return true
	
}

//判断是否有反调试进程和包含某些必要进程
func NoBlockUserProcess()bool {
	known := []string{
		"ollydbg.exe",
		"ProcessHacker.exe",
		"tcpview.exe",
		"autoruns.exe",
		"autorunsc.exe",
		"filemon.exe",
		"procmon.exe",
		"regmon.exe",
		"procexp.exe",
		"idaq.exe",
		"idaq64.exe",
		"ImmunityDebugger.exe",
		"Wireshark.exe",
		"dumpcap.exe",
		"HookExplorer.exe",
		"ImportREC.exe",
		"PETools.exe",
		"LordPE.exe",
		"SysInspector.exe",
		"proc_analyzer.exe",
		"sysAnalyzer.exe",
		"sniff_hit.exe",
		"windbg.exe",
		"joeboxcontrol.exe",
		"joeboxserver.exe",
		"joeboxserver.exe",
		"ResourceHacker.exe",
		"x32dbg.exe",
		"x64dbg.exe",
		"Fiddler.exe",
		"httpdebugger.exe"}
	pids,_ := process.Pids()
	pname := []string{}
	for _, pid := range pids {
		pn,_ := process.NewProcess(pid)
		pName,_ :=pn.Name()
		pname = append(pname,pName)
	}

	for _,v :=range pname{
		for _,v1 :=range known{
			if (v1 == v){
				return false
			}
		}
	}
	
	//fmt.Println("NoBlockUserProcess fine")
	return true
}


//判断是否有必要的系统文件
func HaveFile()bool {
	return true
}


//判断程序名称及是否在某几个目录下运行
func RunPath()bool {
	checkName := "run.exe"
	known := []string{
		"C:\\Users\\Public",
		"C:\\Programdata"}

	path := os.Args[0]
	Name := path[len(path)-len(checkName):len(path)]
	spath := path[0:len(path)-len(checkName)-1]
	//fmt.Println(spath)
	if(Name != checkName){
		return false
	}

	for _,v :=range known{
		if (spath == v){
			return true
		}
	}
	
	return false
}



func runbefore()bool {
	//休眠多少秒
	time.Sleep(time.Duration(1)*time.Second)
	res := RunPath() && NoBlockUserProcess() && GetMemPercent() && GetCpuCount() && NoBlockUserName() && NoBlockComputerName() && GetDisk() && HaveFile()
	return res
}

func main() {
	if(runbefore()){
		fmt.Println("all fine")
	}
}
