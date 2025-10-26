package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	store        *Store
	cookieName   string
	cookieDom    string
	cookieSec    bool
	cookieMaxAge int
}

// In a real system, replace this with DB/LDAP/OAuth etc.
func validateUserPass(user, pass string) bool {
	// demo: single local user
	return (user == "admin" && pass == "admin") || (user == "patrick" && pass == "patrick")
}

func NewController(store *Store, cookieName, cookieDomain string, cookieSecure bool, cookieMaxAge int) *Controller {
	return &Controller{
		store:        store,
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

func (a *Controller) Register(rg *gin.RouterGroup) {
	g := rg.Group("/auth")
	g.POST("/login", a.Login)
	g.POST("/logout", a.Logout)
	g.GET("/me", a.Me)
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
	if !validateUserPass(req.Username, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	token, _, err := a.store.Create(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create session"})
		return
	}

	maxAge := a.cookieMaxAge
	// Set cookie (HTTPOnly, SameSite=Lax)
	c.SetCookie(
		a.cookieName,
		token,
		maxAge,      // seconds
		"/",         // path
		a.cookieDom, // domain (can be empty)
		a.cookieSec, // secure
		true,        // httpOnly
	)

	// Also set SameSite explicitly (Gin SetCookie doesn’t expose it directly prior to v1.9)
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
		a.store.Delete(token)
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
	// Best-effort read (no 401)
	token, _ := c.Cookie(a.cookieName)
	if token == "" {
		c.JSON(http.StatusOK, gin.H{"user": ""})
		return
	}
	if sess, ok := a.store.Get(token); ok {
		c.JSON(http.StatusOK, gin.H{"user": sess.User})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": ""})
}
