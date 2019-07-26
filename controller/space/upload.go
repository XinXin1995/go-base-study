package space

import (
	"blog/pkg/e"
	"blog/pkg/util"
	"fmt"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//单文件上传
func UploadSingle(c *gin.Context) {
	file, _ := c.FormFile("file")
	// 1 图片 2 文件 3 视频 4 音频
	fileType, _ := com.StrTo(c.PostForm("type")).Int()
	code := e.INVALID_PARAMS
	valid := validation.Validation{}
	valid.Required(file, "file").Message("未获取到上传文件")
	valid.Required(fileType, "fileType").Message("上传类型为空值")
	res := &util.Res{}
	if !valid.HasErrors() {
		filePath := "static/"
		switch fileType {
		case 1:
			filePath += "img/"
			break
		case 2:
			filePath += "file/"
			break
		case 3:
			filePath += "video/"
			break
		case 4:
			filePath += "audio/"
			break
		}
		temp := com.ToStr(time.Now().Unix())
		filePath = fmt.Sprintf("%s%s_%s", filePath, temp, file.Filename)
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			log.Println("saveFile err :", err)
		}
		code = e.SUCCESS
		res.Code = code
		res.Msg = e.MsgUser[code]
		res.Data = filePath
	} else {
		for _, value := range valid.Errors {
			log.Println(value)
		}
		res.Msg = e.MsgUser[code]
	}
	c.JSON(http.StatusOK, res)
}
