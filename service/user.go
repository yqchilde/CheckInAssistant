package service

import (
	"errors"

	"github.com/gin-gonic/gin"

	"CheckInAssistant/core"
	"CheckInAssistant/model"
	"CheckInAssistant/store"
	"CheckInAssistant/viewmodel"
)

type UserService struct {
	Ctx *gin.Context
}

func NewUserService(ctx *gin.Context) *UserService {
	return &UserService{Ctx: ctx}
}

// UserLogin user login ...
func (u *UserService) UserLogin(userLogin *viewmodel.User) (*model.User, error) {
	// Check user is exist
	userStore := store.NewDBUser()

	has, user, err := userStore.GetByPhone(userLogin.Phone)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("该用户不存在")
	}

	// Check user password is it right or not
	// todo password rules
	salt := core.GetSalt(user.Password)
	encPwd := core.EncryptPassword(userLogin.Password, salt)
	if encPwd != user.Password {
		return nil, errors.New("账号或密码错误")
	}

	return user, err
}

// UserRegister user register ...
func (u *UserService) UserRegister(userRegister *viewmodel.User) error {
	// Check user is exist
	userStore := store.NewDBUser()
	has, _, err := userStore.GetByPhone(userRegister.Phone)
	if err != nil {
		return err
	}
	if has {
		return errors.New("该用户已注册")
	}

	// todo password rule
	userRegister.Password = encryptPassword(userRegister.Password)

	user := &model.User{
		UserID:   core.NewUUID(),
		Phone:    userRegister.Phone,
		Password: userRegister.Password,
		Role:     model.RoleStudent,
	}
	err = userStore.Insert(user)
	if err != nil {
		return err
	}

	return err
}

// UpdateUserInfo update user info
func (u *UserService) UpdateUserInfo(typ string, userInfo *model.User) error {
	userStore := store.NewDBUser()

	has, user, err := userStore.GetByID(userInfo.UserID)
	if err != nil {
		return err
	}
	if !has {
		return errors.New("The user not exists. ")
	}

	switch typ {
	case "realName":
		user.RealName = userInfo.RealName
	case "email":
		user.Email = userInfo.Email
	case "class":
		user.ClassID = "11111"
	case "password":
		user.Password = encryptPassword(userInfo.Pwd)
	}
	err = userStore.UpdateUserInfo(user)
	if err != nil {
		return err
	}

	return err
}

// GetUserInfo get user info
func (u *UserService) GetUserInfo(userID string) (*model.User, error) {
	has, user, err := store.NewDBUser().GetByID(userID)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("The user not exists. ")
	}

	return user, err
}

func encryptPassword(password string) string {
	salt := core.GenerateSalt(6)
	encPwd := core.EncryptPassword(password, salt)
	return encPwd
}
