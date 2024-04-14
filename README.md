简单的基于命令行和读取文件敏感字符搜索与敏感文件搜索并且打包


<img width="1574" alt="image" src="https://github.com/hackerxj007/find_info/assets/23031720/9186b351-b87f-4ad7-b885-ca183b4ece3d">


效果如下：

<img width="991" alt="image" src="https://github.com/hackerxj007/find_info/assets/23031720/de014073-f0a4-45d2-8cba-6c3ef2e4c983">


正则匹配的规则：find/find.go

regexp.MustCompile(`(?i).*账.*|.*密码.*|.*录.*|.*构.*|.*组.*|.*VPN.*|.*简.*|.*拓扑.*|.*历.*|.*记.*|.*说.*|.*服.*|.*地址.*|.*资产.*|.*表.*|.*部.*|.*使用.*|.*单.*|.*表.*|.*配.*|.*介绍.*|.*工.*|.*图.*|.*人.*|.*备份.*|.*.册.*|.*桌面.*|.*交接.*|.*维.*`)


info搜索基于如下

func win_info() {

	ExecuteCommand("whoami")
	ExecuteCommand("ipconfig /all")
	ExecuteCommand("arp -a")
	ExecuteCommand("nslookup")
	ExecuteCommand("systeminfo")
	ExecuteCommand("netstat -an")
	AppendToFile("=== Read HOSTS FILE ===\n")
	ReadFile("C:\\Windows\\System32\\drivers\\etc\\hosts")
	find_ip_dm()
}

func linux_info() {

	ExecuteCommand("id")
	ExecuteCommand("ifconfig")
	ExecuteCommand("uname -a")
	ExecuteCommand("arp -a")
	ExecuteCommand("w")
	ExecuteCommand("last")
	ExecuteCommand("netstat -an")
	AppendToFile("=== Read HOSTS FILE ===\n")
	ReadFile("/etc/hosts")
	AppendToFile("=== Read passwd FILE ===\n")
	ReadFile("/etc/passwd")
	AppendToFile("=== Read version FILE ===\n")
	ReadFile("/proc/version")
	find_ip_dm()

}

