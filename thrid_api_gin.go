package ThirdApi_Gin

import (
	"net/http"
)

type ThridInteraction interface {
	First(c *gin.Context)
}

type ThirdSx struct {
}

type FirstReq struct {
	Content string `json:"content"`
	Id      string `json:"id"`
}

type AA struct {
}

func (t *ThirdSx) First(c *gin.Context) {
	tt := FirstReq{
		Content: "123",
		Id:      "123",
	}
	AA.M(c, tt)
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}
func ThirdApi(r *gin.Engine, t ThridInteraction) {
	r.GET("/First-API", t.First)
}
