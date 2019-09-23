// 后台认证API
package auth

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"leonzhangxf-api/util"
	"net/http"
	"os"
	"time"
)

var mgrConfig managerConfig

type managerConfig struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
	Issuer   string `toml:"issuer"`
	TokenKey string `toml:"tokenKey"`
	Expires  int64  `toml:"expires"`
}

func init() {
	// load the manager config
	const managerFileName = "mgr.toml"
	dir, _ := os.Getwd()
	mgrFileLocation := fmt.Sprintf("%v%vconfig%v%v", dir, string(os.PathSeparator),
		string(os.PathSeparator), managerFileName)
	util.Log.Infoln("Current mgr config dir is ", mgrFileLocation)
	if _, err := toml.DecodeFile(mgrFileLocation, &mgrConfig); nil != err {
		util.Log.Fatalln("Parse the mgr config fail.", err)
	}
	util.Log.Infoln("The mgr config parsed success.")
}

type Api struct {
}

// @Summary Login
// @Description Login and get the token and refresh token
// @Tags auth
// @Param username query string true "login username"
// @Param password query string true "login password"
// @Produce json
// @Success 201 {obj} auth.TokenResponse
// @Failure 400 {string} Login fail
// @Failure 500 {string} Generate token fail
// @Router /api/v1/auth/login [post]
func (api *Api) Login(context *gin.Context) {
	username := context.Query("username")
	password := context.Query("password")
	util.Log.Debugln("Username is ", username, " and password is ", password)
	if username != mgrConfig.Username || password != mgrConfig.Password {
		context.String(http.StatusBadRequest, "Login fail")
		return
	}

	token, err := getTokenString(username)
	if nil != err {
		context.String(http.StatusInternalServerError, "Generate token fail")
		return
	}
	refreshToken, err := getRefreshTokenString(username)
	if nil != err {
		context.String(http.StatusInternalServerError, "Generate token fail")
		return
	}
	context.JSON(http.StatusOK, &TokenResponse{
		Token:        token,
		RefreshToken: refreshToken,
	})
}

// @Summary Refresh and get the new token
// @Description Refresh and get the new token with refresh token
// @Tags auth
// @Param refresh_token query string true "refresh token"
// @Produce json
// @Success 201 {obj} auth.TokenResponse
// @Failure 400 {string} Params error or validate fail
// @Failure 500 {string} Generate token fail
// @Router /api/v1/auth/refresh_token [post]
func (api *Api) RefreshToken(context *gin.Context) {
	refreshToken := context.Query("refresh_token")
	token, err := parseToken(refreshToken)
	if nil != err {
		context.String(http.StatusBadRequest, "Validate fail")
		return
	}
	standardClaims := token.Claims.(*jwt.StandardClaims)
	tokenString, err := getTokenString(standardClaims.Subject)
	refreshTokenString, err := getRefreshTokenString(standardClaims.Subject)
	if nil != err {
		context.String(http.StatusInternalServerError, "Generate token fail")
		return
	}
	context.JSON(http.StatusOK, &TokenResponse{
		Token:        tokenString,
		RefreshToken: refreshTokenString,
	})
}

func parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(mgrConfig.TokenKey), nil
	})
	return token, err
}

// Token check Middleware
func (api *Api) TokenCheck(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	if "" == tokenString {
		ctx.String(http.StatusUnauthorized, "Authorization fail")
		return
	}
	_, err := parseToken(tokenString)
	if nil != err {
		ctx.String(http.StatusUnauthorized, "Authorization fail")
		return
	}
	ctx.Next()
}

type TokenResponse struct {
	Token        string `json:"token"`         // 用于调用接口传入Authorization，默认300s过期
	RefreshToken string `json:"refresh_token"` // 用于刷新Token，默认30min过期
}

func getTokenString(username string) (tokenString string, err error) {
	tokenString, err = generateToken(mgrConfig.Expires, username)
	return
}

func getRefreshTokenString(username string) (tokenString string, err error) {
	const refreshTokenExpires = 1800
	tokenString, err = generateToken(refreshTokenExpires, username)
	return
}

func generateToken(tokenExpires int64, username string) (string, error) {
	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(tokenExpires * time.Second.Nanoseconds())).Unix(),
		Issuer:    mgrConfig.Issuer,
		Subject:   username,
	}
	signKey := []byte(mgrConfig.TokenKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signKey)
	return tokenString, err
}
