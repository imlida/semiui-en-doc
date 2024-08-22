package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	sourceDir := "./semi-design"
	targetDir := "./output"

	// 遍历源目录
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 检查文件是否以 US.md 结尾
		if !info.IsDir() && strings.HasSuffix(info.Name(), "US.md") {
			// 构建目标路径
			relPath, err := filepath.Rel(sourceDir, path)
			if err != nil {
				return err
			}
			targetPath := filepath.Join(targetDir, relPath)

			// 创建目标目录
			if err := os.MkdirAll(filepath.Dir(targetPath), os.ModePerm); err != nil {
				return err
			}

			// 复制文件
			if err := copyFile(path, targetPath); err != nil {
				return err
			}

			fmt.Printf("Copied: %s -> %s\n", path, targetPath)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// 复制文件
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	targetFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	_, err = io.Copy(targetFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}