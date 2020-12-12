package Tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func success(context *gin.Context , v interface{}) {
	context.JSON(200,gin.H{
		"status" : 1 ,
		"massage" : "成功",
		"data" : v,
	})

}
func failed(context *gin.Context , v interface{}){
	context.JSON(http.StatusOK,gin.H{
		"status" : 0 ,
		"maasage" : v,
	})
	
}
