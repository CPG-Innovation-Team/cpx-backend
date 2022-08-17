package biz

import (
	"context"
	"cpx-backend/app/mall/portal/internal/conf"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/copier"

	v1 "cpx-backend/api/mall/portal/v1"
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrPasswordInvalid = errors.New("password invalid")
	ErrUsernameInvalid = errors.New("username invalid")
	ErrUpdateUser      = errors.New("update user failed")
)

type User struct {
	Uid      int64
	Name     string
	Nickname string

	// 类型。1-用户；2-供应商；3-运营
	Category uint64
	Email    string
	Password string
	Avatar   string

	// 性别:0-保密;1-男;2-女;
	Gender    uint64
	Introduce string
}

func checkUserParams(name, password string) error {
	// check username
	if len(name) <= 0 {
		return ErrUsernameInvalid
	}
	// check password
	if len(password) <= 0 {
		return ErrPasswordInvalid
	}

	return nil
}

// UserRepo is a retrieve user information repository.
type UserRepo interface {
	FindByEmail(ctx context.Context, email string) (user *User, err error, uid int64)
	VerifyPassword(ctx context.Context, email, password string) (uid int64, err error)
	GetUserProfile(ctx context.Context, uid int64) (user *User, err error)
	GetUserList(ctx context.Context, uid []int64, typ int64) (users *[]User, err error)

	Create(ctx context.Context, user *User) (uid int64, err error)
	Update(ctx context.Context, user *User) error
}

// UserUseCase is a user use case.
type UserUseCase struct {
	key string
	ur  UserRepo
	log *log.Helper
}

// NewUserUseCase new a user use case.
func NewUserUseCase(ur UserRepo, conf *conf.Auth, logger log.Logger) *UserUseCase {
	return &UserUseCase{key: conf.ApiKey, ur: ur, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) generateToken(uid int64, email string) (token string, err error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": uid,
		"email":   email,
	})

	token, err = claims.SignedString([]byte(uc.key))
	if err != nil {
		return "", v1.ErrorLoginFailed("generate token failed: %s", err.Error())
	}
	return
}

// RegisterUser creates a User, and returns the new User.
func (uc *UserUseCase) RegisterUser(ctx context.Context, request *v1.RegisterRequest) (uid int64, token string, err error) {
	// check email
	_, err, _ = uc.ur.FindByEmail(ctx, request.User.Email)

	if !errors.Is(err, ErrUserNotFound) {
		return 0, "", v1.ErrorRegisterFailed("username already exists")
	}

	// check user params
	err = checkUserParams(request.GetUser().Name, request.GetUser().Password)

	// create user
	user := User{}
	_ = copier.Copy(&user, request.User)
	uid, err = uc.ur.Create(ctx, &user)
	if err != nil {
		return 0, "", err
	}

	// generate token
	token, err = uc.generateToken(uid, request.User.Email)
	if err != nil {
		return uid, "", err
	}

	return uid, token, nil
}

func (uc *UserUseCase) Login(ctx context.Context, email, password string) (uid int64, token string, err error) {
	// check params

	// verify password
	uid, err = uc.ur.VerifyPassword(ctx, email, password)
	if err != nil {
		return 0, "", err
	}

	// generate token
	token, err = uc.generateToken(uid, email)
	if err != nil {
		return uid, "", err
	}

	return uid, token, nil
}

func (uc UserUseCase) Update(ctx context.Context, request *v1.UpdateRequest) (uid int64, token string, err error) {
	// check user params

	// update user
	user := User{}
	_ = copier.Copy(&user, request.User)
	err = uc.ur.Update(ctx, &user)
	if err == nil {
		return request.User.Uid, "", err
	}

	// generate token
	token, err = uc.generateToken(request.User.Uid, request.User.Email)
	if err != nil {
		return request.User.Uid, "", err
	}

	return request.User.Uid, token, nil
}

func (uc UserUseCase) GetUserProfile(ctx context.Context, uid int64) (user *User, err error) {
	return uc.ur.GetUserProfile(ctx, uid)
}

func (uc UserUseCase) GetUserList(ctx context.Context, uid []int64, typ int64) (user *[]User, err error) {
	return uc.ur.GetUserList(ctx, uid, typ)
}
