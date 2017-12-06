package session

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/controller"
	"github.com/hal-iosk/hal-cinema/model"
	"github.com/hal-iosk/hal-cinema/service"
	"github.com/satori/go.uuid"
)

var AdminTokenKey = "halCinemaAdmin"

func AdminLoginHandle(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")
	user, ok := service.Customer.Logincheck(email, password)
	if !ok {
		controller.Batequest("ログイン失敗", c)
		return
	}
	SetCookie(user.ID, c)
}

func AdminSetCookie(userID uint, c *gin.Context) {
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
	http.SetCookie(c.Writer, &http.Cookie{Name: AdminTokenKey, Value: uid})
}

func AdminAuthViewMiddleware(c *gin.Context) {
	token, err := c.Cookie(AdminTokenKey)
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
func Admin(c *gin.Context) {
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

func AdmintokenCheck(token string) (*model.Token, bool) {
	dbtoken := model.Token{}
	db.Where("token = ? AND admin_flag = 1", token).First(&dbtoken)
	return &dbtoken, dbtoken.Token != ""
}