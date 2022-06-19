package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"path/filepath"
	"regexp"
	"Goup/service"
)

var MdCategories = []string{"[Markdown]"}

// TODO: 改成map来传
func ParseMd(cb *service.CnBlog, fileinfo ...string) map[string]interface{} {
	if len(fileinfo) < 1 {
		fmt.Println("file info can't be null")
		panic("file info can't be null")
	}
	var title string
	var tag string = ""
	var filePath = fileinfo[0]
	var lenFileinfo = len(fileinfo)
	var tmpFile = strings.ReplaceAll(filePath, "\\", "/")
	var curdir = filepath.Dir(tmpFile)
	
	// print("here")
	switch lenFileinfo {
	case 1:
		var tmpFileSlice = strings.Split(tmpFile, "/") // 分割
		var suffix = strings.Split(tmpFileSlice[len(tmpFileSlice)-1], ".")
		title = suffix[0]
	case 2:
		title = fileinfo[1]
	case 3:
		title = fileinfo[1]
		tag = fileinfo[2]
	}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		panic("file read fail")
	}
	content := string(data)
	imgRejg(&content, cb, curdir)
	var nPost = map[string]interface{}{
		// "time":        time.Now().Format("2022-06-05 19:15:19"),
		"title":       title,
		"description": content,
		"categories":  MdCategories,
		"mt_keywords": tag,
	}
	return nPost
}

func imgRejg(content *string, cb *service.CnBlog, curdir string) {
	// imgPattern := regexp.MustCompile("(!\\[\\w+\\]\\((.*?\\.jpg|jpeg|png)\\))")
	imgPattern := regexp.MustCompile(`(!\[\w*\]\((.*?\.jpg|jpeg|png)\))`)
	if imgPattern == nil { //解释失败，返回nil
		fmt.Println("regexp err")
		return
	}
	
	var localImgs []string
	//根据规则提取关键信息
	result1 := imgPattern.FindAllStringSubmatch(*content, -1)
	for _, res := range result1 {
		if len(res)>=2{
			// fmt.Println("res = ", res[len(res)-1]," ",filepath.Join(curdir,res[2]))
			localImgs = append(localImgs, filepath.Join(curdir,res[2]))
		}
	}
	var remoteImgs = make(chan map[int]string, len(localImgs))

	if len(localImgs) > 0{
		for idx, path := range localImgs{
			fmt.Println("go path: ", path)
			go func (idx int, path string) {
				isFile, _ := CheckFile(path)
				if !isFile {
					fmt.Println("the media object path is not a file: ",path)
					return
				}
				fp, err := os.OpenFile(path, os.O_RDONLY, 0) // 只读
				if err != nil {
					fmt.Println("the media object  is invalid: ", path)
					return
				}
				defer fp.Close()
				remoteUrl, err :=cb.NewMediaObj(fp)
				if remoteUrl == "" || err != nil{
					fmt.Println("new media object fail, please check")
					// remoteImgs <- map[int]string{-1:""}
					return
				}
				remoteImgs <- map[int]string{idx:remoteUrl}
			}(idx, path)
		}
		remoteUrls := make(map[int]string)
		for i := 0; i < len(localImgs); i++ {
			IdxUrl := <- remoteImgs
			for k,v := range IdxUrl{
				if k==-1{
					break
				}
				remoteUrls[k] = v
			}
			// strings.ReplaceAll(*content, localImgs[])
		}
		for idx, rUrl := range remoteUrls{
			fmt.Println("sfjj: ",result1[idx][2])
			*content = strings.ReplaceAll(*content, result1[idx][2], rUrl)
		}
	} else {
		fmt.Println("not need upload img!")
	}
}

func CheckFile(file string) (bool,error){
	s, err := os.Stat(file)
	if err != nil{
		return false, err
	}
	if s.IsDir(){
		return false, nil
	}
	return true, nil//如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}