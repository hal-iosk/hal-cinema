package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/controller"
	"github.com/hal-iosk/hal-cinema/model"
	"github.com/hal-iosk/hal-cinema/service"
	"github.com/satori/go.uuid"
)

var db = model.GetDBConn()
var CustomerTokenKey = "halCinemaUser"

func UserLoginHandle(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	user, ok := service.Admin.Logincheck(email, password)
	if !ok {
		controller.Batequest("ログイン失敗", c)
		return
	}
	SetCookie(user.ID, c)
}

func SetCookie(userID uint, c *gin.Context) {
	uid := uuid.NewV4().String()
	token := model.Token{
		AdminFlag: false,
		UserID:    userID,
		Token:     uid,
	}
	err := db.Create(&token).Error
	if err != nil {
		panic(err)
	}
	http.SetCookie(c.Writer, &http.Cookie{Name: CustomerTokenKey, Value: uid})
}

func AuthViewMiddleware(c *gin.Context) {
	token, err := c.Cookie(CustomerTokenKey)
	if err != nil {
		c.Redirect(300, "/login")
		c.Abort()
		return
	}
	user, ok := tokenCheck(token)
	if !ok {
		c.Redirect(300, "/login")
		c.Abort()
		return
	}
	c.Set("userID", user.UserID)
}
func AuthApiMiddleware(c *gin.Context) {
	token, err := c.Cookie(CustomerTokenKey)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"err": "権限がありまてん＾＾;",
		})
		c.Abort()
		return
	}
	user, ok := tokenCheck(token)
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{
			"err": "権限がありまてん＾＾;",
		})
		c.Abort()
		return
	}
	c.Set("userID", user.UserID)
}

func tokenCheck(token string) (*model.Token, bool) {
	dbtoken := model.Token{}
	db.Where("token = ? AND admin_flag = 0", token).First(&dbtoken)
	return &dbtoken, dbtoken.Token != ""

}
