package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main()  {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user/:name", GetUserS)
	r.POST("/upload", uploadFile)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func GetUserS(c *gin.Context)  {
	name := c.Param("name")
	c.String(http.StatusOK, name)
}
func uploadFile(c * gin.Context)  {

	file,_ := c.FormFile("file")
	dst := "./Operation/"+file.Filename
	log.Println(file.Filename)
	// 上传到指定路劲
	c.SaveUploadedFile(file, dst)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

