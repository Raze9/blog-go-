package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func NewRoute() *gin.Engine {
	r := gin.Default()
	r.POST("/api/re", func(c *gin.Context) {
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")
		if len(telephone) != 11 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 422,
				"msg":  "手机号必须为11位",
			})
		}
		if len(password) != 6 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 422,
				"msg":  "密码长度不足",
			})
		}
		if len(name) == 0 {
			name = fmt.Sprintf("admin" + r2and(10))
			fmt.Println(name)
		}
		log.Println(name, telephone, password)
		c.JSON(http.StatusOK, "注册成功")
	})
	return r
}
func r2and(n int) string {
	rand.Seed(time.Now().Unix())
	ran := rand.Intn(n)
	return strconv.Itoa(ran)
}
