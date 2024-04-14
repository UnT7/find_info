package info

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

func ExecuteCommand(command string) error {

	outputFile := fmt.Sprintf("output.txt")

	// 创建命令
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("bash", "-c", command)
	}

	// 执行命令
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("执行命令时出错: %v\n输出: %s", err, output)
	}

	// 打开文件并追加结果
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("打开文件时出错: %v", err)
	}
	defer file.Close()

	// 将结果追加到文件中
	if _, err := file.WriteString(fmt.Sprintf("=== %s ===\r", command)); err != nil {
		return fmt.Errorf("写入标题时出错: %v", err)
	}
	if _, err := file.Write(output); err != nil {
		return fmt.Errorf("写入结果时出错: %v", err)
	}

	//fmt.Println("结果", outputFile, "文件中")
	return nil
}

// ReadFile 读取指定文件的内容并将结果追加到输出文件中
func ReadFile(inputFile string) error {
	// 读取输入文件的内容
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("读取文件 %s 出错: %v", inputFile, err)
	}

	// 输出文件路径
	outputFile := "output.txt"

	// 打开输出文件，如果文件不存在则创建，以追加模式写入
	output, err := os.OpenFile(outputFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("打开文件 %s 出错: %v", outputFile, err)
	}
	defer output.Close()

	// 将文件内容写入到输出文件中
	_, err = output.Write(data)
	if err != nil {
		return fmt.Errorf("写入文件 %s 出错: %v", outputFile, err)
	}

	//fmt.Println("文件内容已追加到", outputFile)
	return nil
}

func AppendToFile(content string) error {
	// 文件路径
	filePath := "output.txt"

	// 打开文件，如果文件不存在则创建，以追加模式写入
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("打开文件 %s 出错: %v", filePath, err)
	}
	defer file.Close()

	// 写入内容到文件中
	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("写入文件 %s 出错: %v", filePath, err)
	}

	//fmt.Println("内容已成功追加到", filePath)
	return nil
}

func win_info() {

	ExecuteCommand("whoami")
	ExecuteCommand("ipconfig /all")
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

func Main() {

	if runtime.GOOS == "windows" {
		win_info()

	} else {
		linux_info()
	}
}
