package main

import (
	"Goup/commands"
	"Goup/constants"
	"Goup/service"
	"Goup/utils"
	"fmt"
	"log"
	"os"
	"regexp"
	// "time"
	"github.com/urfave/cli/v2"
	// "time"
	// "github.com/fatih/color"
)

const VERSION = "v1.0"
const TITLE = "goup"
const DESC = "a tool for upload blog"

func imgReg() {
	// imgPattern := regexp.MustCompile("(!\\[\\w+\\]\\((.*?\\.jpg|jpeg|png)\\))")
	imgPattern := regexp.MustCompile(`(!\[\w*\]\((.*?\.jpg|jpeg|png)\))`)
	// imgPattern := regexp2.MustCompile(`g(?=])`, 0)
	var buf string = `
	> [世界坐标系，相机坐标系，图像坐标系，像素坐标系转换](https://blog.csdn.net/qq_28448117/article/details/79526431)
> [世界坐标系、相机坐标系、图像坐标系、像素坐标系之间的关系](https://blog.csdn.net/u011574296/article/details/73658560)

![](../assert/coordnate.jpg)

### 世界坐标系：

 客观三维世界的绝对坐标系，也称客观坐标系。**因为数码相机安放在三维空间中，我们需要世界坐标系这个基准坐标系来描述数码相机的位置**，并且**用它来描述安放在此三维环境中的其它任何物体的位置**，用（X, Y, Z）表示其坐标值

### 相机坐标系（光心坐标系）：
以相机的光心为坐标原点，X 轴和Y 轴分别平行于图像坐标系的 X 轴和Y 轴，相机的光轴为Z 轴，用（Xc, Yc, Zc）表示其坐标值。
### 图像坐标系：
以CCD 图像平面的中心为坐标原点，X轴和Y 轴分别平行于图像平面的两条垂直边，用( x , y )表示其坐标值。图像坐标系是用物理单位（例如毫米）表示像素在图像中的位置。
### 像素坐标系：
以 CCD 图像平面的左上角顶点为原点，X 轴和Y 轴分别平行于图像坐标系的 X 轴和Y 轴，用(u , v )表示其坐标值。**数码相机采集的图像首先是形成标准电信号的形式，然后再通过模数转换变换为数字图像**。每幅图像的存储形式是M × N的数组，M 行 N 列的图像中的每一个元素的数值代表的是图像点的灰度。这样的每个元素叫像素，像素坐标系就是以像素为单位的图像坐标系

总转换公式:

![img](https://img-blog.csdn.net/20180312144108329)

像素坐标系与世界坐标系的关系
![](../assert/coord_cv.jpg)


	`
	if imgPattern == nil { //解释失败，返回nil
		fmt.Println("regexp err")
		return
	}
	//根据规则提取关键信息
	result1 := imgPattern.FindAllStringSubmatch(buf, -1)
	// if err != nil{
	// 	fmt.Println("sf", err.Error())
	// }
	// result1 := imgPattern.Fin
	// fmt.Println(result1.Groups()[2].Capture.String())
	fmt.Println(result1)
	for _, res := range result1 {
		if len(res) >= 2 {
			fmt.Println("res = ", res[len(res)-1])
		}
	}
}

func fix(s *string){
	*s = "img"
}

func main() {
	// var s [2]string
	// go func (s [2]string)  {
	// 	s[0] = "sf"
	// 	s[1] = "im"
	// }(s)
	// time.Sleep(time.Duration(3)*time.Second)
	// go func (s [2]string)  {
	// 	fmt.Println(s[1])
	// }(s)
	// fmt.Print("s ",s)
	// imgReg()
	
		var cnb service.CnBlog
		cnb.Init(constants.USERNAME, constants.PASSWD, constants.BLOGADDR, constants.URL)
		app := commands.GetApp(TITLE)

		app.Action = func(ctx *cli.Context) error {

			if isDir := ctx.Bool("dir"); isDir {
				fmt.Println("sorry, we not'support apply dir currently, coming soon!")
				return nil
			}
			md := ctx.String("file")

			if md == "" {
				fmt.Println("file can't be null")
				return nil
			}
			if isFile, err := utils.CheckFile(md); err == nil {
				if !isFile {
					fmt.Println("the param file should be the path onf a file")
					return nil
				}
				tmpMd := []string{md}
				nPost := utils.ParseMd(&cnb, tmpMd...)
				cnb.UpBlog(nPost)
			} else {
				fmt.Println("the param file is invalid, please check it")
				return nil
			}

			return nil
		}

		err := app.Run(os.Args)
		if err != nil {
			log.Fatal(err)
		}

		// md := []string{`E:\Type\Documents\deepin\编程珠玑\语言系列\Go专栏\1Go常用.md`}
		//

		// cnb.GetUserBlogs()
	
}
