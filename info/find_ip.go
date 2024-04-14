package info

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"regexp"
	"strings"
)

func find_ip_dm() {
	inputFilename := "output.txt"
	outputFilename := "output_ip_dm.txt"

	// 读取文件内容
	content, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		fmt.Printf("读取文件时出错: %v", err)
		return
	}

	// 正则表达式匹配IP地址和域名
	ipPattern := `\b(?:\d{1,3}\.){3}\d{1,3}\b`
	domainPattern := `(?i)\b((?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z]{2,6})\b`

	ipRegex := regexp.MustCompile(ipPattern)
	domainRegex := regexp.MustCompile(domainPattern)

	ipMatches := ipRegex.FindAllString(string(content), -1)
	domainMatches := domainRegex.FindAllString(string(content), -1)

	// 分类IP地址
	internalIPs := make(map[string]bool)
	externalIPs := make(map[string]bool)

	for _, ip := range ipMatches {
		if isInternalIP(ip) {
			ip = addSubnetMask(ip)
			internalIPs[ip] = true
		} else {
			externalIPs[ip] = true
		}
	}

	// 将分类后的IP地址写入文件
	err = writeIPClassificationToFile(outputFilename, internalIPs, externalIPs)
	if err != nil {
		fmt.Printf("写入文件时出错: %v", err)
		return
	}

	// 将域名写入文件
	err = writeDomainsToFile(outputFilename, domainMatches)
	if err != nil {
		fmt.Printf("写入域名到文件时出错: %v", err)
		return
	}

}

// 判断IP地址是否为内网地址
func isInternalIP(ip string) bool {
	// 内网地址段
	internalIPRanges := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}

	for _, ipRange := range internalIPRanges {
		if isInSubnet(ip, ipRange) {
			return true
		}
	}
	return false
}

// 将IP地址加上/24子网掩码
func addSubnetMask(ip string) string {
	// 如果IP地址已经带有子网掩码，则不修改
	if strings.Contains(ip, "/") {
		return ip
	}
	return ip
}

// 判断IP地址是否属于指定网段
func isInSubnet(ip, subnet string) bool {
	_, ipNet, err := net.ParseCIDR(subnet)
	if err != nil {
		return false
	}

	ipAddr := net.ParseIP(ip)
	if ipAddr == nil {
		return false
	}

	return ipNet.Contains(ipAddr)
}

// 将分类后的IP地址写入文件
func writeIPClassificationToFile(filename string, internalIPs, externalIPs map[string]bool) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString("内网地址:\r")
	for ip := range internalIPs {
		file.WriteString(ip + "\r")
	}

	file.WriteString("\r内网IP段:\r")
	for ip := range internalIPs {
		ipWithoutSubnet := strings.Split(ip, "/")[0]
		file.WriteString(ipWithoutSubnet + "/24\r")
	}

	file.WriteString("\r外网地址:\r")
	for ip := range externalIPs {
		file.WriteString(ip + "\r")
	}

	// 将内网IP地址段写入文件

	return nil
}

// 将域名写入文件
func writeDomainsToFile(filename string, domains []string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString("\r域名:\r")
	for _, domain := range domains {
		file.WriteString(domain + "\r")
	}

	return nil
}
