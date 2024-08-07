package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func fix() {
	dir := "/Users/xxxx/xxxx/go_study/package/main"
	files, err := ioutil.ReadDir(dir)
	if nil != err {
		log.Fatalln(err)
	}
	paths := []string{}
	for _, f := range files {
		fName := f.Name()
		if f.IsDir() && !strings.Contains(fName, ".") {
			subDir := dir + "/" + fName
			subFiles, err := ioutil.ReadDir(subDir)
			if err != nil {
				log.Println("Error reading subdirectory:", err)
				continue
			}
			for _, subFile := range subFiles {
				if !subFile.IsDir() {
					subfileName := subFile.Name()
					absFileName := dir + "/" + fName + "/" + subfileName
					changeContent(subfileName, absFileName, fName)
				}
			}
		}
	}
	for _, v := range paths {
		log.Println(v)
	}
}

func changeContent(simpleFileName string, absFileName string, pkg string) {
	if !strings.HasSuffix(simpleFileName, "go") {
		return
	}
	if strings.Contains(simpleFileName, "Template") {
		return
	}
	content, err := ioutil.ReadFile(absFileName)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	// 将内容转换为字符串
	text := string(content)
	replaceName := strings.ReplaceAll(simpleFileName, ".go", "")
	// 修改内容
	oldWord := "main"
	newWord := replaceName + "main"
	pkg = strings.ToLower(pkg)

	newText := strings.ReplaceAll(text, "package main", "package "+pkg)
	newText = strings.ReplaceAll(newText, oldWord, newWord)
	oldWord = "test"
	newWord = replaceName + "test"
	newText = strings.ReplaceAll(newText, oldWord, newWord)

	// 将修改后的内容写回文件
	err = ioutil.WriteFile(absFileName, []byte(newText), 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %s", err)
	}

	log.Println("File updated successfully")
}
