package service

import (
	"fmt"
	"io"

	// "reflect"
	// "time"
	// "encoding/base64"
	"io/ioutil"

	"github.com/mattn/go-xmlrpc"
	// "github.com/kolo/xmlrpc"
)

// type methods struct {
// 	newPost string
// }

type services interface {
	Init(userName string, passWd string, blogAddr string, url string)
	UpBlog(nPost map[string]interface{}) string
	// TODO: 其他service
}

type CnBlog struct {
	blogId   string
	userName string
	passWd   string
	blogAddr string
	url      string
	client   *xmlrpc.Client
}

func (cb *CnBlog) Init(userName string, passWd string, blogAddr string, url string) {
	cb.userName = userName
	cb.passWd = passWd
	cb.blogAddr = blogAddr
	cb.url = url
	cb.client = xmlrpc.NewClient(cb.url)
	cb.GetUserBlogs()
}

func (cb *CnBlog) NewMediaObj(fp io.Reader) (string, error) {
	fbytes, err := ioutil.ReadAll(fp)
	if err != nil {
		return "", err
	}
	// fEncBytes := base64.StdEncoding.EncodeToString(fbytes)
	// newByte := new(bytes.Buffer)
	params := []interface{}{cb.blogId, cb.userName, cb.passWd, map[string]interface{}{
		"bits": fbytes,
		"name": "abc.jpg",
		"type": "image/jpg"}}
	res, err := cb.client.Call("metaWeblog.newMediaObject", params...)
	if err != nil {
		fmt.Println("NewMediaObj error happen")
		fmt.Println(err.Error())
		return "", nil
	} else {
		url := res.(xmlrpc.Struct)
		fmt.Println("new media object success!: ",url)
		return url["url"].(string), nil
	}
}

func (cb *CnBlog) UpBlog(nPost map[string]interface{}) string {
	params := []interface{}{cb.blogId, cb.userName, cb.passWd, nPost, true}
	res, err := cb.client.Call("metaWeblog.newPost", params...)
	if err != nil {
		fmt.Println("UpBlog error happen")
		fmt.Println(err.Error())
	} else {
		fmt.Println("publish blog: success! ", res.(string))
	}
	return "sf"
}

/*
func (cb *CnBlog) UpBlog(blogId string) string {
	categories := []string{"[Markdown]"}
	fmt.Println(time.Now())
	md := "E:\\Type\\Documents\\deepin\\编程珠玑\\语言系列\\Go专栏\\1Go常用.md"
	data, err := ioutil.ReadFile(md)
	if err!=nil{
		fmt.Println(err)
		return ""
	}
	content := string(data)
	var nPost = map[string]interface{}{
		// "time":        time.Now().Format("2022-06-05 19:15:19"),
		"title":       "Go常用",
		"description": content,
		"categories":  categories,
		"mt_keywords": "go",
	}
	// var res string
	// parms := struct {
	// 	BlogId    string
	// 	UserName  string
	// 	PassWord  string
	// 	Npost     Post
	// 	IsPublish bool
	// }{"740438",
	// 	cb.userName,
	// 	cb.passWd,
	// 	nPost,
	// 	true}
	// parms := make([]interface{}, 5)
	// parms[0] = "740438"
	// parms[1] = cb.userName
	// parms[2] = cb.passWd
	// parms[3] = nPost
	// parms[4] = true
	params := []interface{}{"740438", cb.userName, cb.passWd, nPost, true}
	res, err := cb.client.Call("metaWeblog.newPost", params...)
	if err != nil {
		fmt.Println("error happen")
		fmt.Println(err.Error())
	} else {
		fmt.Println("publish blog: success! ", res.(string))
	}

	return "sf"
}
*/

func (cb *CnBlog) GetUserBlogs() {
	params := []interface{}{cb.blogAddr, cb.userName, cb.passWd}
	fmt.Println("call blogger.getUserBlogs")
	v, err := cb.client.Call("blogger.getUsersBlogs", params...)
	if err != nil {
		fmt.Println("GetUserBlogs error happen")
		fmt.Println(err.Error())
		return
	} else {
		tar := v.(xmlrpc.Array)[0].(xmlrpc.Struct)
		/*
			fmt.Println("tar: ",reflect.TypeOf(tar)) // 实际运行时候知道
			for k, v := range tar.(xmlrpc.Struct) { // 编译时候不知道
				fmt.Printf("%s=%v\n", k, v)
			}*/
		fmt.Println("apply id: ", tar["blogid"])
		cb.blogId = tar["blogid"].(string)
	}
}
