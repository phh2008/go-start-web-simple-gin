package service

import (
	"com.gientech/selection/dao"
	"com.gientech/selection/entity"
	"com.gientech/selection/model"
	"com.gientech/selection/model/result"
	"com.gientech/selection/pkg/exception"
	"com.gientech/selection/pkg/logger"
	"com.gientech/selection/pkg/xjwt"
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/cristalhq/jwt/v5"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

var UserSet = wire.NewSet(wire.Struct(new(UserService), "*"))

type UserService struct {
	UserDao  *dao.UserDao
	Jwt      *xjwt.JwtHelper
	Enforcer *casbin.Enforcer
}

func (a *UserService) CreateByEmail(ctx context.Context, email model.UserEmailRegister) *result.Result[model.UserModel] {
	user := a.UserDao.GetByEmail(ctx, email.Email)
	if user.Id > 0 {
		return result.Failure[model.UserModel]("email 已存在")
	}
	pwd, err := bcrypt.GenerateFromPassword([]byte(email.Password), 1)
	if err != nil {
		logger.Errorf("生成密码出错：%s", err.Error())
		return result.Error[model.UserModel](err)
	}

	user = entity.UserEntity{
		Email:    email.Email,
		RealName: email.Email,
		UserName: email.Email,
		Password: string(pwd),
		Status:   1,
		RoleCode: "",
	}
	user, err = a.UserDao.Add(ctx, user)
	if err != nil {
		logger.Errorf("创建用户出错：%s", err.Error())
		return result.Failure[model.UserModel]("创建用户出错")
	}
	var userModel model.UserModel
	copier.Copy(&userModel, &user)
	return result.Ok[model.UserModel](userModel)
}

func (a *UserService) LoginByEmail(ctx context.Context, loginModel model.UserLoginModel) *result.Result[string] {
	user := a.UserDao.GetByEmail(ctx, loginModel.Email)
	if user.Id == 0 {
		return result.Failure[string]("用户或密码错误")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginModel.Password))
	if err != nil {
		return result.Failure[string]("用户或密码错误")
	}
	// 生成token
	userClaims := xjwt.UserClaims{}
	userClaims.ID = strconv.FormatInt(user.Id, 10)
	userClaims.Role = user.RoleCode
	userClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7))
	token, err := a.Jwt.CreateToken(userClaims)
	if err != nil {
		logger.Errorf("生成token错误：%s", err.Error())
		return result.Error[string](exception.SysError)
	}
	return result.Ok[string](token.String())
}

// AssignRole 给用户分配角色
func (a *UserService) AssignRole(ctx context.Context, userRole model.AssignRoleModel) *result.Result[any] {
	err := a.UserDao.SetRole(ctx, userRole.UserId, userRole.RoleCode)
	if err != nil {
		logger.Errorf("db update error: %s", err.Error())
		return result.Failure[any]("分配角色出错")
	}
	// 更新casbin中的用户与角色关系
	uid := strconv.FormatInt(userRole.UserId, 10)
	_, _ = a.Enforcer.DeleteRolesForUser(uid)
	_, _ = a.Enforcer.AddGroupingPolicy(uid, userRole.RoleCode)
	return result.Success[any]()
}

func (a *UserService) DeleteById(ctx context.Context, id int64) *result.Result[any] {
	err := a.UserDao.DeleteById(ctx, id)
	if err != nil {
		logger.Errorf("delete error: %s", err.Error())
		return result.Failure[any]("刪除出错")
	}
	return result.Success[any]()
}
