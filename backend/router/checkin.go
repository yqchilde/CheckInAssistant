package router

import (
	"github.com/gin-gonic/gin"

	"CheckInAssistant/model"
	"CheckInAssistant/pkg/app"
	"CheckInAssistant/service"
)

// CreateCheckIn ...
func CreateCheckIn(ctx *gin.Context) {
	checkIn := new(model.CheckIn)
	errDecode := ctx.ShouldBindJSON(checkIn)
	if errDecode != nil {
		app.NewErrorBodyJSON(ctx.Writer)
		return
	}

	if checkIn.Creator == "" {
		app.NewErrorMissField(ctx.Writer, "creator")
		return
	}
	if checkIn.RoomID == "" {
		app.NewErrorMissField(ctx.Writer, "room_id")
		return
	}
	if checkIn.CheckInWay != model.CheckInTypeCode && checkIn.CheckInWay != model.CheckInTypeGPS {
		app.NewErrorInvalidField(ctx.Writer, "check_in_way")
		return
	}
	if checkIn.Duration == 0 {
		app.NewErrorInvalidField(ctx.Writer, "duration")
		return
	}

	data, err := service.NewCheckInService(ctx).CreateCheckIn(checkIn)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", data)
}

// GetCheckInDetail get check in detail
func GetCheckInDetail(ctx *gin.Context) {
	checkInID := ctx.Param("checkInID")
	userID := ctx.Param("userID")

	data, err := service.NewCheckInService(ctx).GetCheckInDetail(checkInID, userID)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", data)
}

// GetCheckInList return check in list
func GetCheckInList(ctx *gin.Context) {
	userID := ctx.Param("userID")
	tabIndex := ctx.Query("tabIndex")

	data, err := service.NewCheckInService(ctx).GetCheckInList(userID, tabIndex)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", data)
}

// UserCheckIn user check in
func UserCheckIn(ctx *gin.Context) {
	checkInItem := new(model.CheckInItem)
	errDecode := ctx.ShouldBindJSON(checkInItem)
	if errDecode != nil {
		app.NewErrorBodyJSON(ctx.Writer)
		return
	}

	if checkInItem.CheckInID == "" {
		app.NewErrorInvalidField(ctx.Writer, "check_in_id")
		return
	}
	if checkInItem.UserID == "" {
		app.NewErrorInvalidField(ctx.Writer, "user_id")
		return
	}

	err := service.NewCheckInService(ctx).UserCheckIn(checkInItem)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", nil)
}

// GetCheckInAllUsers get check in all users
func GetCheckInAllUsers(ctx *gin.Context) {
	checkInID := ctx.Param("checkInID")

	data, err := service.NewCheckInService(ctx).GetCheckInAllUsers(checkInID)
	if err != nil {
		app.NewErrorResponse(ctx.Writer, err)
		return
	}

	app.NewSuccessResponse(ctx.Writer, "", data)
}
