package router

import (
	"CRM-test/database"
	"CRM-test/structures"
	"CRM-test/utils"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Initialization() *gin.Engine {
	router := gin.Default()

	store := sessions.NewCookieStore([]byte("secretWord"))
	router.Use(sessions.Sessions("session", store))

	router.LoadHTMLGlob("template/*.html")
	router.Static("assets/", "assets")

	router.GET("/", indexHandler)
	routerApplication := router.Group("/application")
	routerApplication.GET("/", applicationHandler)
	routerApplication.GET("/create", applicationCreateHandler)
	routerApplication.PUT("/create", applicationCreate)
	routerApplication.GET("/get-abonents", getAbonents)
	routerApplication.GET("/get-data", getApplications)
	router.GET("/houses", housesHandler)

	routerUser := router.Group("/users")
	routerUser.GET("/", usersHandler)
	routerUser.GET("/create", usersCreateHandler)
	routerUser.PUT("/create", usersAdd)
	routerUser.GET("/login", loginHandler)
	routerUser.POST("/login", loginCheck)
	routerUser.GET("/edit/:id", usersEditHandler)
	routerUser.GET("/get-user/:id", getUser)
	routerUser.DELETE("/exit", exit)
	routerUser.POST("/user-blocked-switch", userBlockedSwitch)

	return router
}

func getApplications(c *gin.Context) {
	applications, e := database.GetAllApplications()
	if e != nil {
		utils.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	c.JSON(200, applications)
}

func applicationCreate(c *gin.Context) {
	var application structures.Application
	e := c.BindJSON(&application)
	if e != nil {
		utils.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	if database.CreateApplication(application) {
		c.JSON(200, true)
		return
	}

	c.JSON(400, false)
}

func getAbonents(c *gin.Context) {
	abonents, e := database.GetAbonents()
	if e != nil {
		utils.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	c.JSON(200, abonents)
}

func userBlockedSwitch(c *gin.Context) {
	var user structures.User
	e := c.BindJSON(&user)
	if e != nil {
		utils.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	user, e = database.GetUserByID(strconv.Itoa(user.ID))
	if e != nil {
		utils.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	var res bool
	res, e = database.UserBlockedSwitch(user)
	if e != nil {
		utils.Logger.Println(e)
		c.JSON(400, false)
		return
	}

	c.JSON(200, res)
}

func getUser(c *gin.Context) {
	userID := c.Param("id")
	user, e := database.GetUserByID(userID)
	if e != nil {
		utils.Logger.Println(e)
		c.JSON(400, nil)
		return
	}
	c.JSON(200, user)
}

func usersCreateHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		c.HTML(200, "users-create", gin.H{
			"Theme": getTheme(c),
		})
	} else {
		c.Redirect(301, "/users/login")
	}
}

func usersEditHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		userID := c.Param("id")
		user, e := database.GetUserByID(userID)
		if e != nil {
			utils.Logger.Println(e)
			return
		}
		c.HTML(200, "users-edit", gin.H{
			"User":  user,
			"Theme": getTheme(c),
		})
	} else {
		c.Redirect(301, "/users/login")
	}
}

func usersHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		users, e := database.GetAllUsers()
		if e != nil {
			utils.Logger.Println(e)
			return
		}

		c.HTML(200, "users", gin.H{
			"Users": users,
			"Theme": getTheme(c),
		})
	} else {
		c.Redirect(301, "/users/login")
	}

}

func housesHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		c.HTML(200, "houses", gin.H{
			"Theme": getTheme(c),
		})
	} else {
		c.Redirect(301, "/users/login")
	}

}

func applicationCreateHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		c.HTML(200, "application-create", gin.H{
			"Theme": getTheme(c),
		})
	} else {
		c.Redirect(301, "/users/login")
	}

}

func applicationHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		c.HTML(200, "application", gin.H{
			"Theme": getTheme(c),
		})
	} else {
		c.Redirect(301, "/users/login")
	}
}

func indexHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		c.HTML(200, "index", gin.H{
			"Theme": getTheme(c),
		})
	} else {
		c.Redirect(301, "/users/login")
	}
}

func loginCheck(c *gin.Context) {
	session := sessions.Default(c)

	var user structures.User
	e := c.BindJSON(&user)
	if e != nil {
		utils.Logger.Println(e)
		c.Status(400)
		return
	}

	user.Password, e = utils.Encrypt(user.Password)
	if e != nil {
		utils.Logger.Println(e)
		c.Status(400)
		return
	}

	if database.LoginCheck(&user) {
		hash, ok := database.CreateSession(&user)
		if ok {
			session.Set("SessionSecretKey", hash)
			e = session.Save()
			if e != nil {
				utils.Logger.Println(e)
				return
			}

			c.JSON(200, true)

		}
	}

	c.Status(400)
}

func loginHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		c.Redirect(301, "/")
	} else {
		c.HTML(200, "login", gin.H{
			"Theme": getTheme(c),
		})
	}
}

func usersAdd(c *gin.Context) {
	var userData structures.User
	e := c.BindJSON(&userData)
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	c.JSON(200, database.UsersAdd(&userData))
}

func exit(c *gin.Context) {
	session := sessions.Default(c)
	_session := getSession(c)

	_, ok := session.Get("SessionSecretKey").(string)
	if ok {
		session.Clear()
		_ = session.Save()
		c.SetCookie("hello", "", -1, "/", c.Request.URL.Hostname(), false, true)
		session.Delete("SessionSecretKey")
	}

	database.DeleteSession(_session)

	c.JSON(301, true)
}

func getSession(c *gin.Context) *structures.Session {
	_session := sessions.Default(c)

	sessionHash, ok := _session.Get("SessionSecretKey").(string)
	if ok {
		session := database.GetSession(sessionHash)
		if session != nil {
			session.Exists = true
			return session
		}
	}

	return &structures.Session{
		Exists: false,
	}
}

func getTheme(c *gin.Context) string {
	cookie, e := c.Cookie("crm-theme")
	if e != nil {
		utils.Logger.Println(400, "Куки не найден")
		return "light"
	}

	return cookie
}
