package service

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"

	"CheckInAssistant/core"
	"CheckInAssistant/model"
	"CheckInAssistant/store"
	"CheckInAssistant/viewmodel"
)

type ClassRoomService struct {
	Ctx *gin.Context
}

func NewClassRoomService(ctx *gin.Context) *ClassRoomService {
	return &ClassRoomService{Ctx: ctx}
}

// CreateClassRoom create class room
func (c *ClassRoomService) CreateClassRoom(classObj *viewmodel.ClassRoom) error {
	tr := store.GetDBTransaction()
	err := tr.Begin()
	if err != nil {
		log.Printf("ClassRoomService CreateClassRoom tr.Begin, error:%s \n", err.Error())
		return err
	}

	defer func() {
		if err == nil {
			_ = tr.Commit()
		} else {
			log.Printf("ClassRoomService CreateClassRoom failed, error:%s \n", err.Error())
			_ = tr.Rollback()
		}
		_ = tr.Close()
	}()

	var (
		roomStore      = store.NewDBRoom(tr)
		classRoomStore = store.NewDBClassRoom(tr)
	)
	// Create Room
	room := &model.Room{
		RoomID:   core.NewUUID(),
		RoomName: classObj.RoomName,
		Creator:  classObj.Creator,
	}
	err = roomStore.Insert(room)
	if err != nil {
		return err
	}

	// Create Class Room
	for _, classID := range classObj.ClassIDList {
		classRoom := &model.ClassRoom{
			ClassRoomID: core.NewUUID(),
			ClassID:     classID,
			RoomID:      room.RoomID,
		}
		err := classRoomStore.Insert(classRoom)
		if err != nil {
			return err
		}
	}

	return err
}

// GetClassRoomList return class room list
func (c *ClassRoomService) GetClassRoomList(userID, tabIndex string) ([]*model.Room, error) {
	return store.NewDBRoom().GetByUserIDAndType(userID, tabIndex)
}

// DeleteClassRoom delete class room by room id
func (c *ClassRoomService) DeleteClassRoom(roomID string) error {
	return store.NewDBRoom().DeleteRoom(roomID)
}

// GetClassRoomDetail return class room detail by room id
func (c *ClassRoomService) GetClassRoomDetail(roomID string) (*viewmodel.ClassRoom, error) {
	var (
		roomStore  = store.NewDBRoom()
		userStore  = store.NewDBUser()
		classStore = store.NewDBClass()
	)
	// Get room info
	has, room, err := roomStore.GetByID(roomID)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("The room not exist. ")
	}

	// classRoom
	classRoom := &viewmodel.ClassRoom{
		RoomID:    room.RoomID,
		RoomName:  room.RoomName,
		Creator:   room.Creator,
		CreatedAt: room.CreatedAt,
	}

	// Get class room
	cmList, err := classStore.GetClassRoomList(roomID)
	if err != nil {
		return nil, err
	}
	classRoom.ClassList = cmList

	// Get class room member
	var classIDList []string
	for i := range cmList {
		classIDList = append(classIDList, cmList[i].ClassID)
	}
	userList, err := userStore.GetByClassIDList(classIDList)
	if err != nil {
		return nil, err
	}
	classRoom.ClassMember = userList

	return classRoom, err
}
