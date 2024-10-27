package service

import (
	"ddCode-server/dao"
	"ddCode-server/global"
	"ddCode-server/models"
	"ddCode-server/models/dto"
	"ddCode-server/models/vo"
	"ddCode-server/pkg/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
	Service
}

// LoginByUserNameAndPassword 登录信息
func (s *UserService) LoginByUserNameAndPassword(param *dto.UserLoginParam) (*vo.UserLogin, error) {
	userDao := dao.UserDao{}
	deptDao := dao.DeptDao{}
	// 用户基本信息
	userInfo, err := userDao.GetUserInfo(param.Username)
	if err != nil {
		global.Logger.Errorf("获取用户信息失败: %v", err)
		return nil, err
	}
	if userInfo == nil {
		return nil, fmt.Errorf("用户名或密码错误")
	}
	if utils.VerifyPassword(userInfo.Password, param.Password) {
		return nil, fmt.Errorf("用户名或密码错误")
	}
	// 部门
	deptName, err := deptDao.GetDeptNameById(userInfo.DeptId)
	if err != nil {
		global.Logger.Errorf("获取部门信息失败: %v", err)
		return nil, err
	}
	// 角色
	roles, err := global.Rbac.GetRoleByUserName(userInfo.UserName)
	if err != nil {
		global.Logger.Errorf("获取角色信息失败: %v", err)
		return nil, err
	}
	// 菜单权限
	roleMenuService := &RoleMenuService{}
	menus, err := roleMenuService.FindRoleMenuByRoleNames(roles)
	if err != nil {
		global.Logger.Errorf("获取角色菜单信息失败: %v", err)
		return nil, err
	}

	userTokenInfo := &vo.UserTokenInfo{
		UserName: userInfo.UserName,
		Phone:    userInfo.Phone,
		Email:    userInfo.Email,
		DeptId:   userInfo.DeptId,
		Roles:    roles,
	}
	userTokenMap := utils.StructTagToMap(userTokenInfo, "json")
	// Token
	token, _, err := global.JwtMaker.CreateToken(userTokenMap)
	if err != nil {
		global.Logger.Errorf("生成token失败: %v", err)
		return nil, err
	}
	// token 存Redis
	userLogin := &vo.UserLogin{
		Info: vo.UserLoginInfo{
			UserName: userInfo.UserName,
			NickName: userInfo.NickName,
			Email:    userInfo.Email,
			Phone:    userInfo.Phone,
			Sex:      userInfo.Sex,
			Avatar:   userInfo.Avatar,
			DeptName: deptName,
			Roles:    roles,
			Menus:    menus,
		},
		Token: token,
	}
	//	Redis操作
	return userLogin, nil
}

func (s *UserService) CreateUser(param *dto.UserRegisterParam) error {
	userDao := dao.UserDao{}
	userInfo, err := userDao.GetUserInfo(param.Username)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if userInfo.UserName != "" {
		return fmt.Errorf("用户名已被使用")
	}
	password, _ := utils.GeneratePasswordHash(userInfo.Password)
	err = userDao.CreateUser(&models.User{
		UserName:  param.Username,
		Password:  password,
		NickName:  param.Nickname,
		Email:     param.Email,
		Phone:     param.Phone,
		Sex:       param.Sex,
		Avatar:    param.Avatar,
		DeptId:    param.DeptId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
		Comment:   param.Comment,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateUser(param *dto.UserInfoUpdateParam) error {
	userDao := dao.UserDao{}
	err := userDao.UpdateUserInfo(&models.User{
		Id:        param.Id,
		UserName:  param.Username,
		NickName:  param.Nickname,
		Email:     param.Email,
		Phone:     param.Phone,
		Sex:       param.Sex,
		Avatar:    param.Avatar,
		DeptId:    param.DeptId,
		UpdatedAt: time.Now(),
		DeletedAt: &param.DeletedAt,
		Comment:   param.Comment,
	})
	if param.Password != "" {
		password, _ := utils.GeneratePasswordHash(param.Password)
		param.Password = password
	}
	if err != nil {
		return err
	}
	oldUserInfo, err := userDao.GetUserInfoById(param.Id)
	if err != nil {
		return err
	}
	if oldUserInfo.UserName != param.Username {
		err := global.Rbac.UpdateUser(oldUserInfo.UserName, param.Username)
		if err != nil {
			global.Logger.Errorf("UpdateRole err:%v", err)
			return err
		}
	}
	return nil
}

func (s *UserService) FindUserByPage(p *dto.UserListParam) ([]dto.UserListDto, int64, error) {
	userDao := dao.UserDao{}
	users, count, err := userDao.FindUserByParamAndPage(p)
	list := make([]dto.UserListDto, 0, len(users))
	for _, v := range users {
		list = append(list, dto.UserListDto{
			Id:       v.Id,
			UserName: v.UserName,
			NickName: v.NickName,
			Email:    v.Email,
			Phone:    v.Phone,
			Sex:      v.Sex,
			Avatar:   v.Avatar,
			Comment:  v.Comment,
		})
	}
	return list, count, err
}
