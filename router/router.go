package router

import (
	"CheckInAssistant/middleware/jwt"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	//r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	//r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	//r.POST("/auth", api.GetAuth)
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{



		//获取标签列表
		//apiv1.GET("/tags", v1.GetTags)
		////新建标签
		//apiv1.POST("/tags", v1.AddTag)
		////更新指定标签
		//apiv1.PUT("/tags/:id", v1.EditTag)
		////删除指定标签
		//apiv1.DELETE("/tags/:id", v1.DeleteTag)
		////导出标签
		//r.POST("/tags/export", v1.ExportTag)
		////导入标签
		//r.POST("/tags/import", v1.ImportTag)
	}

	return r
}
