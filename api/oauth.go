package api

import (
	"io"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var conf = &oauth2.Config{
	ClientID:     "xxxxxxxx.apps.googleusercontent.com",
	ClientSecret: "xxxxxxxxxxxxxxxx",
	RedirectURL:  "http://localhost:8080/callback",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}

func GoogleOAuth(c *fiber.Ctx) error {
	// 這邊應該用加密字串，讓callback去驗證，這邊是方便測試而已
	// 拿到的url用重新導向讓，會跑去google的登入畫面
	url := conf.AuthCodeURL("random-string")
	return c.Redirect(url)
}

func CallBack(c *fiber.Ctx) error {
	// 驗證是否合規
	state := c.Query("state")
	if state != "random-string" {
		return c.Status(400).SendString("Not valid")
	}
	ctx := c.Context()
	code := c.Query("code")
	if code == "" {
		return c.Status(400).SendString("Code not found")
	}
	token, err := conf.Exchange(ctx, code)
	if err != nil {
		return c.Status(500).SendString("Failed to exchange token." + err.Error())
	}

	// google允許的oauth api拿資料
	// 需要完整功能可以讀、寫個人信息等請使用:google people api
	// 基本的驗證功能請使用: OAuth2 api
	apiURL := "https://www.googleapis.com/oauth2/v2/userinfo"
	client := conf.Client(ctx, token)
	resp, err := client.Get(apiURL)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendString(string(body))
}
