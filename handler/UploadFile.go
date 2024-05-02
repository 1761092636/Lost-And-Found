package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func Upload(c *gin.Context) {

	file, _ := c.FormFile("file")
	fileExt := strings.ToLower(path.Ext(file.Filename))
	if fileExt != ".jpg" {
		c.JSON(http.StatusOK, gin.H{"uploading": "done", "message": "上传图片不是jpg格式文件"})
		return
	}
	in, _ := file.Open()
	defer in.Close()
	dst := path.Join("./static/back/upload", file.Filename)
	out, _ := os.Create(dst)
	defer out.Close()
	io.Copy(out, in)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Filename))
	c.File(dst)
	//c.JSON(http.StatusOK, gin.H{"uploading": "done", "message": "success", "url": "http://" + c.Request.Host + "./static/back/upload" + file.Filename})

}
