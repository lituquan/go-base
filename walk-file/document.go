package main

import (
	"bytes"
	"encoding/json"
	ioutil "io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const SUFFIX = ".go"
const FROM_CACHE = false

var FILE_INDEX = "data/index.json"
var FILE_TOKEN = "data/token.json"

var ops = 0                 //记录文件总数
var codeIndex = make(index) //倒排索引：token-->ids
var files = []File{}        //文档

// 遍历文档，插入索引表
func ListDir(dirPth string) (err error) {
	//遍历目录
	return filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(fi.Name(), SUFFIX) {
			readDocument(filename)
		}
		return nil
	})
}

// 这里相当于做了索引和文档存储
func loadCache(dir string) {
	indexs, err := ioutil.ReadFile(FILE_INDEX)
	if err == nil && FROM_CACHE {
		json.Unmarshal(indexs, &files)
		tokens, _ := ioutil.ReadFile(FILE_TOKEN)
		json.Unmarshal(tokens, &codeIndex)
		return
	}
	ListDir(dir)
	writeFile(codeIndex, FILE_TOKEN)
	writeFile(files, FILE_INDEX)
}

// 写json文件
func writeFile(obj interface{}, file string) (err error) {
	tokens, _ := json.Marshal(obj)
	var out bytes.Buffer
	json.Indent(&out, tokens, "", "\t")
	err = ioutil.WriteFile(file, out.Bytes(), 0666)
	if err != nil {
		panic(err)
	}
	return err
}

type File struct {
	ID      int
	Path    string
	Content string
}

// 插入一个文档和它的索引
func readDocument(s string) (err error) {
	log.Println(s)
	bytes, err := ioutil.ReadFile(s)
	if err != nil {
		log.Println(err)
		return
	}
	file := File{ID: ops, Path: s, Content: string(bytes)}
	files = append(files, file)
	codeIndex.add(file)
	ops += 1
	if ops%500 == 0 {
		log.Printf("ops %s\n", ops)
	}
	return
}
