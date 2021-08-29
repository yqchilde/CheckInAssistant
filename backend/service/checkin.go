package service

import (
	"time"

	"github.com/gin-gonic/gin"

	"CheckInAssistant/core"
	"CheckInAssistant/model"
	"CheckInAssistant/store"
	"CheckInAssistant/viewmodel"
)

type CheckInService struct {
	Ctx *gin.Context
}

func NewCheckInService(ctx *gin.Context) *CheckInService {
	return &CheckInService{Ctx: ctx}
}

// CreateCheckIn ...
func (c *CheckInService) CreateCheckIn(checkInObj *model.CheckIn) (*model.CheckIn, error) {
	checkInStore := store.NewDBCheckIn()

	checkIn := &model.CheckIn{
		CheckInID:  core.NewUUID(),
		Creator:    checkInObj.Creator,
		RoomID:     checkInObj.RoomID,
		CheckInWay: checkInObj.CheckInWay,
		BeginTime:  time.Now().UTC(),
		EndTime:    time.Now().UTC().Add(time.Duration(checkInObj.Duration) * time.Minute),
	}
	if checkIn.CheckInWay == model.CheckInTypeCode {
		checkIn.CheckInCode = checkInObj.CheckInCode
	} else if checkIn.CheckInWay == model.CheckInTypeGPS {
		checkIn.Longitude = checkInObj.Longitude
		checkIn.Latitude = checkInObj.Latitude
	}

	err := checkInStore.Insert(checkIn)
	if err != nil {
		return nil, err
	}

	return checkIn, err
}

// GetCheckInDetail get check in detail
func (c *CheckInService) GetCheckInDetail(checkInID, userID string) (*viewmodel.CheckIn, error) {
	var (
		checkInStore = store.NewDBCheckIn()
	)
	// Get check in info
	checkIn, err := checkInStore.GetCheckInDetail(checkInID)
	if err != nil {
		return nil, err
	}

	// Get check in item
	hasItem, _, err := store.NewDBCheckInItem().GetByCheckInIDAndUserID(checkInID, userID)
	if err != nil {
		return nil, err
	}

	dur := checkIn.EndTime.Sub(time.Now().UTC())
	if dur.Seconds() > 0 {
		if hasItem {
			checkIn.Status = model.CheckInStatusFinish
		} else {
			checkIn.Status = model.CheckInStatusProcessing
		}
	} else {
		if hasItem {
			checkIn.Status = model.CheckInStatusEndAndFinish
		} else {
			checkIn.Status = model.CheckInStatusEnd
		}
	}

	return checkIn, err
}

// GetCheckInList return check in list
func (c *CheckInService) GetCheckInList(userID, tabIndex string) ([]*viewmodel.CheckIn, error) {
	checkInList, err := store.NewDBCheckIn().GetByUserIDAndType(userID, tabIndex)
	if err != nil {
		return nil, err
	}

	var cnt = 0

	for _, checkIn := range checkInList {
		// Get check in item
		hasItem, _, err := store.NewDBCheckInItem().GetByCheckInIDAndUserID(checkIn.CheckInID, userID)
		if err != nil {
			return nil, err
		}

		dur := checkIn.EndTime.Sub(time.Now().UTC())
		if dur.Seconds() > 0 {
			if hasItem {
				cnt++
				checkIn.Status = model.CheckInStatusFinish
			} else {
				checkIn.Status = model.CheckInStatusProcessing
			}
		} else {
			if hasItem {
				cnt++
				checkIn.Status = model.CheckInStatusEndAndFinish
			} else {
				checkIn.Status = model.CheckInStatusEnd
			}
		}

		checkIn.CheckInTotal = 1
		checkIn.PeopleTotal = 1
	}

	return checkInList, err
}

// UserCheckIn user check in
func (c *CheckInService) UserCheckIn(checkInItem *model.CheckInItem) error {
	return store.NewDBCheckInItem().Insert(checkInItem)
}

// GetCheckInAllUsers return check in all users
func (c *CheckInService) GetCheckInAllUsers(checkInID string) (map[string]interface{}, error) {
	var (
		checkInStore     = store.NewDBCheckIn()
		checkInItemStore = store.NewDBCheckInItem()
		cnt              = 0
	)

	_, checkIn, err := checkInStore.GetByID(checkInID)
	if err != nil {
		return nil, err
	}

	detail, err := NewClassRoomService(c.Ctx).GetClassRoomDetail(checkIn.RoomID)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{}, 3)

	for _, cm := range detail.ClassMember {
		// select check in
		has, ci, err := checkInItemStore.GetByCheckInIDAndUserID(checkInID, cm.UserID)
		if err != nil {
			return nil, err
		}
		if has {
			cnt++
			cm.CheckInStatus = model.CheckInUserComplete
			cm.CreatedAt = ci.CreatedAt
		} else {
			cm.CheckInStatus = model.CheckInUserUndone
			cm.CreatedAt = time.Time{}
		}
	}

	m["people_total"] = len(detail.ClassMember)
	m["check_in_total"] = cnt
	m["user_list"] = detail.ClassMember

	return m, err
}
