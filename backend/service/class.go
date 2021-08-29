package service

import (
	"github.com/gin-gonic/gin"

	"CheckInAssistant/core"
	"CheckInAssistant/model"
	"CheckInAssistant/store"
	"CheckInAssistant/viewmodel"
)

type ClassService struct {
	Ctx *gin.Context
}

func NewClassService(ctx *gin.Context) *ClassService {
	return &ClassService{Ctx: ctx}
}

// CreateClass create class ...
func (c *ClassService) CreateClass(classInfo *model.Class) error {
	var classStore = store.NewDBClass()

	// Insert class
	class := &model.Class{
		ClassID:       core.NewUUID(),
		ClassName:     classInfo.ClassName,
		ClassCapacity: classInfo.ClassCapacity,
		Creator:       classInfo.Creator,
	}
	err := classStore.Insert(class)
	if err != nil {
		return err
	}

	return err
}

// GetClassList get class list by user id and type
func (c *ClassService) GetClassList(userID, tabIndex string) ([]*model.Class, error) {
	classList, err := store.NewDBClass().GetByUserIDAndType(userID, tabIndex)
	if err != nil {
		return nil, err
	}

	for i := range classList {
		cnt, err := store.NewDBConnection().GetByClassID(classList[i].ClassID)
		if err != nil {
			return nil, err
		}

		classList[i].CurrentPeopleNumber = cnt
	}

	return classList, err
}

// UpdateClass update class info
func (c *ClassService) UpdateClass(class *model.Class) error {
	return store.NewDBClass().UpdateClass(class)
}

// DeleteClass delete class
func (c *ClassService) DeleteClass(classID string) error {
	return store.NewDBClass().DeleteClass(classID)
}

// OutClass out class
func (c *ClassService) OutClass(classID, userID string) error {
	return store.NewDBConnection().DeleteByClassIDAndUserID(classID, userID)
}

// JoinClass join class by user id and class id
func (c *ClassService) JoinClass(connection *model.Connection) error {
	connStore := store.NewDBConnection()

	conn := &model.Connection{
		ClassID: connection.ClassID,
		UserID:  connection.UserID,
	}

	err := connStore.Insert(conn)
	if err != nil {
		return err
	}

	return err
}

// GetClassUserList get class user list by class id
func (c *ClassService) GetClassUserList(classID string) ([]*viewmodel.User, error) {
	return store.NewDBClass().GetClassUserList(classID)
}
