package router

import (
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

	userApi := r.Group("/user")
	//userApi.Use(jwt.JWT())
	{
		userApi.POST("/login", UserLogin)
		userApi.POST("/register", UserRegister)
		userApi.PUT("/info", UpdateUserInfo)
		userApi.GET("/info", GetUserInfo)
	}

	classApi := r.Group("/class")
	{
		classApi.POST("", CreateClass)
		classApi.GET("/user/:userID", GetClassList)
		classApi.PUT("/:classID", UpdateClass)
		classApi.DELETE("/del/:classID", DeleteClass)
		classApi.POST("/join", JoinClass)
		classApi.DELETE("/out/:classID/user/:userID", OutClass)
		classApi.GET("/users/:classID", GetClassUserList)
	}

	classRoomApi := r.Group("/classRoom")
	{
		classRoomApi.POST("", CreateClassRoom)
		classRoomApi.GET("/user/:userID", GetClassRoomList)
		classRoomApi.GET("/room/:roomID", GetClassRoomDetail)
		classRoomApi.DELETE("/:roomID", DeleteClassRoom)
	}

	checkInApi := r.Group("/checkIn")
	{
		checkInApi.POST("", CreateCheckIn)
		checkInApi.GET("/detail/:checkInID/user/:userID", GetCheckInDetail)
		checkInApi.GET("/user/:userID", GetCheckInList)
		checkInApi.POST("/startCheckIn", UserCheckIn)
		checkInApi.GET("/detail/:checkInID/checkUsers", GetCheckInAllUsers)
	}

	appApi := r.Group("/app")
	{
		appApi.GET("/share", GetShareUrl)
	}

	return r
}
