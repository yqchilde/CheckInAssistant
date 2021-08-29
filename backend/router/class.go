package router

import (
	"github.com/gin-gonic/gin"

	"CheckInAssistant/model"
	"CheckInAssistant/pkg/app"
	"CheckInAssistant/service"
)

// CreateClass create class ...
func CreateClass(ctx *gin.Context) {
	class := new(model.Class)
	errDecode := ctx.ShouldBindJSON(class)
	if errDecode != nil {
		app.NewErrorBodyJSON(ctx.Writer)
		return
	}

	if class.ClassName == "" {
		app.NewErrorMissField(ctx.Writer, "class_name")
		return
	}
	if class.ClassCapacity == 0 {
		app.NewErrorInvalidField(ctx.Writer, "class_capacity")
		return
	}
	if class.Creator == "" {
		app.NewErrorInvalidField(ctx.Writer, "creator")
		return
	}

	err := service.NewClassService(ctx).CreateClass(class)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", nil)
}

// GetClassList return class list info
func GetClassList(ctx *gin.Context) {
	tabIndex := ctx.Query("tabIndex")
	userID := ctx.Param("userID")

	data, err := service.NewClassService(ctx).GetClassList(userID, tabIndex)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", data)
}

// UpdateClass update class info
func UpdateClass(ctx *gin.Context) {
	class := new(model.Class)
	errDecode := ctx.ShouldBindJSON(class)
	if errDecode != nil {
		app.NewErrorBodyJSON(ctx.Writer)
		return
	}

	if class.ClassName == "" {
		app.NewErrorMissField(ctx.Writer, "class_name")
		return
	}
	if class.ClassCapacity == 0 {
		app.NewErrorInvalidField(ctx.Writer, "class_capacity")
		return
	}

	class.ClassID = ctx.Param("classID")

	err := service.NewClassService(ctx).UpdateClass(class)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", nil)
}

// DeleteClass delete class
func DeleteClass(ctx *gin.Context) {
	classID := ctx.Param("classID")

	err := service.NewClassService(ctx).DeleteClass(classID)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", nil)
}

// OutClass out class
func OutClass(ctx *gin.Context) {
	classID := ctx.Param("classID")
	userID := ctx.Param("userID")

	err := service.NewClassService(ctx).OutClass(classID, userID)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", nil)
}

// JoinClass join class by user id and class id
func JoinClass(ctx *gin.Context) {
	conn := new(model.Connection)
	errDecode := ctx.ShouldBindJSON(conn)
	if errDecode != nil {
		app.NewErrorBodyJSON(ctx.Writer)
		return
	}

	if conn.UserID == "" {
		app.NewErrorInvalidField(ctx.Writer, "user_id")
		return
	}
	if conn.ClassID == "" {
		app.NewErrorInvalidField(ctx.Writer, "class_id")
		return
	}

	err := service.NewClassService(ctx).JoinClass(conn)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", nil)
}

// GetClassUserList get class user list by class id
func GetClassUserList(ctx *gin.Context) {
	classID := ctx.Param("classID")

	data, err := service.NewClassService(ctx).GetClassUserList(classID)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", data)
}
