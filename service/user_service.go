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
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

// UserService 用户服务
type UserService struct {
	userDao  *dao.UserDao
	jwt      *xjwt.JwtHelper
	enforcer *casbin.Enforcer
}

// NewUserService 创建服务
func NewUserService(
	userDao *dao.UserDao,
	jwt *xjwt.JwtHelper,
	enforcer *casbin.Enforcer,
) *UserService {
	return &UserService{
		userDao:  userDao,
		jwt:      jwt,
		enforcer: enforcer,
	}
}

// ListPage 用户列表
func (a *UserService) ListPage(ctx context.Context, req model.UserListReq) *result.Result[model.PageData[model.UserModel]] {
	data := a.userDao.ListPage(ctx, req)
	return result.Ok[model.PageData[model.UserModel]](data)
}

// CreateByEmail 根据邮箱创建用户
func (a *UserService) CreateByEmail(ctx context.Context, email model.UserEmailRegister) *result.Result[model.UserModel] {
	user := a.userDao.GetByEmail(ctx, email.Email)
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
	user, err = a.userDao.Add(ctx, user)
	if err != nil {
		logger.Errorf("创建用户出错：%s", err.Error())
		return result.Failure[model.UserModel]("创建用户出错")
	}
	var userModel model.UserModel
	copier.Copy(&userModel, &user)
	return result.Ok[model.UserModel](userModel)
}

// LoginByEmail 邮箱登录
func (a *UserService) LoginByEmail(ctx context.Context, loginModel model.UserLoginModel) *result.Result[string] {
	user := a.userDao.GetByEmail(ctx, loginModel.Email)
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
	token, err := a.jwt.CreateToken(userClaims)
	if err != nil {
		logger.Errorf("生成token错误：%s", err.Error())
		return result.Error[string](exception.SysError)
	}
	return result.Ok[string](token.String())
}

// AssignRole 给用户分配角色
func (a *UserService) AssignRole(ctx context.Context, userRole model.AssignRoleModel) *result.Result[any] {
	err := a.userDao.SetRole(ctx, userRole.UserId, userRole.RoleCode)
	if err != nil {
		logger.Errorf("db update error: %s", err.Error())
		return result.Failure[any]("分配角色出错")
	}
	// 更新casbin中的用户与角色关系
	uid := strconv.FormatInt(userRole.UserId, 10)
	_, _ = a.enforcer.DeleteRolesForUser(uid)
	// 角色为空，表示清除此用户的角色,无需添加
	if userRole.RoleCode != "" {
		_, _ = a.enforcer.AddGroupingPolicy(uid, userRole.RoleCode)
	}
	return result.Success[any]()
}

// DeleteById 根据ID删除
func (a *UserService) DeleteById(ctx context.Context, id int64) *result.Result[any] {
	err := a.userDao.DeleteById(ctx, id)
	if err != nil {
		logger.Errorf("delete error: %s", err.Error())
		return result.Failure[any]("刪除出错")
	}
	// 清除 casbin 中用户信息
	_, err = a.enforcer.DeleteRolesForUser(strconv.FormatInt(id, 10))
	if err != nil {
		logger.Errorf("Enforcer.DeleteRolesForUser error: %s", err)
	}
	return result.Success[any]()
}
