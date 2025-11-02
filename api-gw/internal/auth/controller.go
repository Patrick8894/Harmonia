package auth

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	users        *UserRepo
	sess         SessionStore
	cookieName   string
	cookieDom    string
	cookieSec    bool
	cookieMaxAge int
}

func NewController(users *UserRepo, sess SessionStore, cookieName, cookieDomain string, cookieSecure bool, cookieMaxAge int) *Controller {
	return &Controller{
		users:        users,
		sess:         sess,
		cookieName:   cookieName,
		cookieDom:    cookieDomain,
		cookieSec:    cookieSecure,
		cookieMaxAge: cookieMaxAge,
	}
}

type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a *Controller) Register(rg *gin.RouterGroup) {
	g := rg.Group("/auth")
	g.POST("/register", a.RegisterUser)
	g.POST("/login", a.Login)
	g.POST("/logout", a.Logout)
	g.GET("/me", a.Me)
}

// RegisterUser godoc
// @Summary      Register a new user
// @Description  Creates a user with unique username; logs user in by setting the session cookie
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        payload  body  registerReq  true  "New user"
// @Success      201      {object}  map[string]string
// @Failure      400      {object}  map[string]string
// @Failure      409      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /auth/register [post]
func (a *Controller) RegisterUser(c *gin.Context) {
	var req registerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	req.Username = strings.TrimSpace(req.Username)
	if len(req.Username) < 3 || len(req.Username) > 64 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username must be 3-64 chars"})
		return
	}
	if len(req.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password must be at least 6 chars"})
		return
	}

	u, err := a.users.Create(c, req.Username, req.Password)
	if err != nil {
		if err == ErrUserExists {
			c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	// Auto-login: create session and set cookie
	token, err := a.sess.Create(u.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create session"})
		return
	}

	maxAge := a.cookieMaxAge
	c.SetCookie(a.cookieName, token, maxAge, "/", a.cookieDom, a.cookieSec, true)
	c.Header("Set-Cookie", (&http.Cookie{
		Name:     a.cookieName,
		Value:    token,
		Path:     "/",
		Domain:   a.cookieDom,
		Secure:   a.cookieSec,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   maxAge,
		Expires:  time.Now().Add(time.Duration(maxAge) * time.Second),
	}).String())

	c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

// @Summary      Login
// @Description  Create a session and set an HTTP-only cookie
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        payload  body  loginReq  true  "Credentials"
// @Success      200      {object}  map[string]string
// @Failure      401      {object}  map[string]string
// @Router       /auth/login [post]
func (a *Controller) Login(c *gin.Context) {
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil || req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	u, err := a.users.GetByUsername(c, req.Username)
	if err != nil || u == nil || !CheckPassword(u.PasswordHash, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := a.sess.Create(u.Username) // <— uses interface
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create session"})
		return
	}

	// Set cookie (SameSite=Lax via explicit header)
	maxAge := a.cookieMaxAge
	c.SetCookie(a.cookieName, token, maxAge, "/", a.cookieDom, a.cookieSec, true)
	c.Header("Set-Cookie", (&http.Cookie{
		Name:     a.cookieName,
		Value:    token,
		Path:     "/",
		Domain:   a.cookieDom,
		Secure:   a.cookieSec,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   maxAge,
		Expires:  time.Now().Add(time.Duration(maxAge) * time.Second),
	}).String())

	c.JSON(http.StatusOK, gin.H{"message": "logged in"})
}

// @Summary      Logout
// @Description  Invalidate the session and clear the cookie
// @Tags         auth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /auth/logout [post]
func (a *Controller) Logout(c *gin.Context) {
	token, _ := c.Cookie(a.cookieName)
	if token != "" {
		a.sess.Delete(token)
	}
	// Clear cookie
	c.SetCookie(a.cookieName, "", -1, "/", a.cookieDom, a.cookieSec, true)
	c.Header("Set-Cookie", (&http.Cookie{
		Name:     a.cookieName,
		Value:    "",
		Path:     "/",
		Domain:   a.cookieDom,
		Secure:   a.cookieSec,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	}).String())

	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

// @Summary      Current user
// @Description  Returns the logged-in user if any
// @Tags         auth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /auth/me [get]
func (a *Controller) Me(c *gin.Context) {
	if token, _ := c.Cookie(a.cookieName); token != "" {
		if user, ok := a.sess.Get(token); ok { // <— uses interface
			c.JSON(http.StatusOK, gin.H{"user": user})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"user": ""})
}
