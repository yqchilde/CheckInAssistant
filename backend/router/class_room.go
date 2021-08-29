package router

import (
	"github.com/gin-gonic/gin"

	"CheckInAssistant/pkg/app"
	"CheckInAssistant/service"
	"CheckInAssistant/viewmodel"
)

// CreateClassRoom create class room
func CreateClassRoom(ctx *gin.Context) {
	classRoom := new(viewmodel.ClassRoom)
	errDecode := ctx.ShouldBindJSON(classRoom)
	if errDecode != nil {
		app.NewErrorBodyJSON(ctx.Writer)
		return
	}

	if classRoom.RoomName == "" {
		app.NewErrorMissField(ctx.Writer, "room_name")
		return
	}
	if len(classRoom.ClassIDList) == 0 {
		app.NewErrorInvalidField(ctx.Writer, "class_list")
		return
	}

	err := service.NewClassRoomService(ctx).CreateClassRoom(classRoom)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", nil)
}

// GetClassRoomList get class room list
func GetClassRoomList(ctx *gin.Context) {
	tabIndex := ctx.Query("tabIndex")
	userID := ctx.Param("userID")

	data, err := service.NewClassRoomService(ctx).GetClassRoomList(userID, tabIndex)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", data)
}

// DeleteClassRoom delete class room
func DeleteClassRoom(ctx *gin.Context) {
	roomID := ctx.Param("roomID")

	err := service.NewClassRoomService(ctx).DeleteClassRoom(roomID)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", nil)
}

// GetClassRoomDetail get class room detail
func GetClassRoomDetail(ctx *gin.Context) {
	roomID := ctx.Param("roomID")

	data, err := service.NewClassRoomService(ctx).GetClassRoomDetail(roomID)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", data)
}
