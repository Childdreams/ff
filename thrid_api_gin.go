package ThirdApi_Gin

import (
	"github.com/gin-gonic/gin"
)

type ThridInteraction interface {
	First(c *gin.Context)
}

//type ThirdSx struct {
//}
//
//func (t *ThirdSx) First(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"msg": "OK",
//	})
//}
func ThirdApi(r *gin.Engine, t ThridInteraction) {
	r.GET("/First-API", t.First)
}
