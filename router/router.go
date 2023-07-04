package router

import (
	"CRM-test/database"
	"CRM-test/structures"
	"CRM-test/utils"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func Initialization() *gin.Engine {
	router := gin.Default()

	store := sessions.NewCookieStore([]byte("secretWord"))
	router.Use(sessions.Sessions("session", store))

	router.LoadHTMLGlob("template/*.html")
	router.Static("assets/", "assets")

	router.GET("/", indexHandler)
	router.GET("/get-events", getEvents)

	routerApplication := router.Group("/application")
	routerApplication.GET("/", applicationHandler)
	routerApplication.GET("/create-:text", applicationCreateHandler)
	routerApplication.PUT("/create", applicationCreate)
	routerApplication.PUT("/create-alternative", applicationCreateAlternative)
	routerApplication.GET("/view-:id", applicationViewHandler)
	routerApplication.GET("/get-abonents", getAbonents)
	routerApplication.GET("/get-all", getApplications)
	routerApplication.POST("/processing-:type/:id", switchTypeProcessingApplication)
	routerApplication.POST("/switch-:obj/:id", switchObjStatus)

	routerHouses := router.Group("/houses")
	routerHouses.GET("/", housesHandler)
	routerHouses.GET("/get-all", getHouses)
	routerHouses.GET("/view/:id", houseViewHandler)
	routerHouses.GET("/create", houseCreateHandler)
	routerHouses.PUT("/create", houseCreate)
	routerHouses.GET("/edit/:id", houseEditHandler)
	routerHouses.PUT("/edit/:id", houseEdit)

	routerUser := router.Group("/users")
	routerUser.GET("/", usersHandler)
	routerUser.GET("/create", usersCreateHandler)
	routerUser.PUT("/create", usersAdd)
	routerUser.GET("/login", loginHandler)
	routerUser.POST("/login", loginCheck)
	routerUser.GET("/edit/:id", usersEditHandler)
	routerUser.PUT("/update/:id", userUpdate)
	//routerUser.GET("/get-user/:id", getUser)
	routerUser.DELETE("/exit", exit)
	routerUser.POST("/user-blocked-switch", userBlockedSwitch)

	routerAbonent := router.Group("/abonents")
	routerAbonent.GET("/", abonentsHandler)
	routerAbonent.GET("/create", abonentsCreateHandler)
	routerAbonent.PUT("/create", abonentsCreate)
	routerAbonent.GET("/edit/:id", abonentEditHandler)
	routerAbonent.PUT("/update/:id", abonentUpdate)
	routerAbonent.GET("/view/:id", abonentViewHandler)

	return router
}

func houseEdit(c *gin.Context) {
	session := getSession(c)

	if session.User.Role.ID == 1 || session.User.Role.ID == 6 || session.User.Role.ID == 2 {
		id := c.Param("id")
		var house structures.House
		e := c.BindJSON(&house)
		if e != nil {
			utils.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		e = database.UpdateHouse(house, id)
		if e != nil {
			utils.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		c.JSON(200, true)
	} else {
		c.JSON(403, false)
	}
}

func houseEditHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.Role.ID == 1 || session.User.Role.ID == 6 || session.User.Role.ID == 2 {
		id := c.Param("id")
		house, e := database.GetHouse(id)
		if e != nil {
			utils.Logger.Println(e)
			c.HTML(200, "houses-edit", gin.H{
				"Theme":       getTheme(c),
				"Active":      "houses",
				"SessionUser": session.User,
			})
			return
		}

		splitAddress := strings.Split(house.Name, " ")
		number := splitAddress[len(splitAddress)-1]
		var street string

		if len(splitAddress) > 2 {
			for i := 0; i < len(splitAddress)-1; i++ {
				street = street + splitAddress[i] + " "
			}
		} else {
			street = splitAddress[0]
		}

		c.HTML(200, "houses-edit", gin.H{
			"Theme":       getTheme(c),
			"Active":      "houses",
			"House":       house,
			"Number":      number,
			"Street":      street,
			"SessionUser": session.User,
		})
	} else {
		c.Redirect(301, "/")
	}
}

func houseCreate(c *gin.Context) {
	session := getSession(c)

	if session.User.Role.ID == 1 || session.User.Role.ID == 6 || session.User.Role.ID == 2 {

		var house structures.House
		e := c.BindJSON(&house)
		if e != nil {
			utils.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		e = database.CreateHouse(house)
		if e != nil {
			utils.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		c.JSON(200, true)
	} else {
		c.JSON(403, false)
	}
}

func houseCreateHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.Role.ID == 1 || session.User.Role.ID == 6 || session.User.Role.ID == 2 {
		c.HTML(200, "houses-create", gin.H{
			"Theme":       getTheme(c),
			"Active":      "houses",
			"SessionUser": session.User,
		})
	} else {
		c.Redirect(301, "/")
	}
}

func houseViewHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		id := c.Param("id")
		house, e := database.GetHouse(id)
		if e != nil {
			utils.Logger.Println(e)
			c.HTML(200, "houses-view", gin.H{
				"Theme":       getTheme(c),
				"Active":      "houses",
				"SessionUser": session.User,
			})
			return
		}

		abonents, e := database.GetAbonentsByAddress(house.Name)
		if e != nil {
			utils.Logger.Println(e)
			c.HTML(200, "houses-view", gin.H{
				"Theme":       getTheme(c),
				"Active":      "houses",
				"House":       house,
				"SessionUser": session.User,
			})
			return
		}

		applications, e := database.GetApplicationsByAddress(house.Name)
		if e != nil {
			utils.Logger.Println(e)
			c.HTML(200, "houses-view", gin.H{
				"Theme":       getTheme(c),
				"Active":      "houses",
				"House":       house,
				"Abonents":    abonents,
				"SessionUser": session.User,
			})
			return
		}

		c.HTML(200, "houses-view", gin.H{
			"Theme":        getTheme(c),
			"Active":       "houses",
			"House":        house,
			"Abonents":     abonents,
			"Applications": applications,
			"SessionUser":  session.User,
		})
	} else {
		c.Redirect(301, "/users/login")
	}
}

func getHouses(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		houses, e := database.GetHouses()
		if e != nil {
			utils.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		c.JSON(200, houses)
	} else {
		c.JSON(403, false)
	}
}

func abonentViewHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		id := c.Param("id")
		abonent, e := database.GetAbonent(id)
		if e != nil {
			utils.Logger.Println(e)
			c.HTML(200, "abonent-view", gin.H{
				"Theme":       getTheme(c),
				"Active":      "abonents",
				"SessionUser": session.User,
			})
			return
		}

		c.HTML(200, "abonent-view", gin.H{
			"Theme":       getTheme(c),
			"Active":      "abonents",
			"Abonent":     abonent,
			"SessionUser": session.User,
		})
	} else {
		c.Redirect(301, "/users/login")
	}
}

func abonentUpdate(c *gin.Context) {
	session := getSession(c)

	if session.User.Role.ID == 1 || session.User.Role.ID == 6 || session.User.Role.ID == 3 {
		id := c.Param("id")
		var abonent structures.Abonent
		e := c.BindJSON(&abonent)
		if e != nil {
			utils.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		e = database.UpdateAbonent(abonent, id)
		if e != nil {
			utils.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		c.JSON(200, true)
	} else {
		c.JSON(403, false)
	}
}

func abonentEditHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.Role.ID == 1 || session.User.Role.ID == 6 || session.User.Role.ID == 3 {
		id := c.Param("id")
		abonent, e := database.GetAbonent(id)
		if e != nil {
			utils.Logger.Println(e)
			c.HTML(200, "abonent-edit", gin.H{
				"Theme":       getTheme(c),
				"Active":      "abonents",
				"SessionUser": session.User,
			})
			return
		}

		c.HTML(200, "abonent-edit", gin.H{
			"Theme":       getTheme(c),
			"Active":      "abonents",
			"Abonent":     abonent,
			"SessionUser": session.User,
		})
	} else {
		c.Redirect(301, "/")
	}
}

func switchObjStatus(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		applicationID := c.Param("id")
		obj := c.Param("obj")
		var event structures.Event
		e := c.BindJSON(&event)
		if e != nil {
			utils.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		if database.SwitchObjStatus(applicationID, obj, event, session.User.ID) {
			c.JSON(200, true)
			return
		}

		c.JSON(400, false)
	} else {
		c.JSON(403, false)
	}
}

func switchTypeProcessingApplication(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		applicationID := c.Param("id")
		typeProcessing := c.Param("type")
		switch typeProcessing {
		case "start":
			e := database.SwitchProcessingApplication(applicationID, session.User.ID, 3)
			if e != nil {
				utils.Logger.Println(e)
				c.JSON(400, false)
				return
			}

			if database.CreateEvent(applicationID, session.User.ID, "Принял в обработку", "") {
				c.JSON(200, true)
				return
			}
			break
		case "stop":
			var event structures.Event
			e := c.BindJSON(&event)
			if e != nil {
				utils.Logger.Println(e)
				c.JSON(400, false)
				return
			}
			e = database.SwitchProcessingApplication(applicationID, session.User.ID, 2)
			if e != nil {
				utils.Logger.Println(e)
				c.JSON(400, false)
				return
			}

			if database.CreateEvent(applicationID, session.User.ID, "Закрыл заявку", event.Comment) {
				c.JSON(200, true)
				return
			}
			break
		case "reopen":
			e := database.SwitchProcessingApplication(applicationID, session.User.ID, 3)
			if e != nil {
				utils.Logger.Println(e)
				c.JSON(400, false)
				return
			}

			if database.CreateEvent(applicationID, session.User.ID, "Переоткрыл заявку", "") {
				c.JSON(200, true)
				return
			}
			break
		}

		c.JSON(400, false)
	} else {
		c.JSON(403, false)
	}
}

func getEvents(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		events, e := database.GetAllEvents()
		if e != nil {
			utils.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		c.JSON(200, events)
	} else {
		c.JSON(403, false)
	}
}

func abonentsCreate(c *gin.Context) {
	session := getSession(c)

	if session.User.Role.ID == 1 || session.User.Role.ID == 6 || session.User.Role.ID == 3 {
		var abonent structures.Abonent
		e := c.BindJSON(&abonent)
		if e != nil {
			utils.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		e = database.CreateAbonent(abonent)
		if e != nil {
			utils.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		c.JSON(200, true)
	} else {
		c.JSON(403, false)
	}
}

func abonentsCreateHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.Role.ID == 1 || session.User.Role.ID == 6 || session.User.Role.ID == 3 {
		c.HTML(200, "abonents-create", gin.H{
			"Theme":       getTheme(c),
			"Active":      "abonents",
			"SessionUser": session.User,
		})
	} else {
		c.Redirect(301, "/")
	}
}

func abonentsHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		abonents, e := database.GetAbonents()
		if e != nil {
			utils.Logger.Println(e)
			c.HTML(200, "abonents", gin.H{
				"Theme":       getTheme(c),
				"SessionUser": session.User,
				"Active":      "abonents",
			})
			return
		}

		c.HTML(200, "abonents", gin.H{
			"Abonents":    abonents,
			"Theme":       getTheme(c),
			"SessionUser": session.User,
			"Active":      "abonents",
		})
	} else {
		c.Redirect(301, "/users/login")
	}
}

func applicationViewHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		id := c.Param("id")
		application, e := database.GetApplication(id)
		if e != nil {
			utils.Logger.Println(e)
			return
		}

		//application.Abonent.Phone = "+7 (" + application.Abonent.Phone[0:3] + ") " + application.Abonent.Phone[3:6] + " " + application.Abonent.Phone[6:8] + "-" + application.Abonent.Phone[8:10]

		events, e := database.GetEventsByApplicationID(application.ID)
		if e != nil {
			utils.Logger.Println(e)
			return
		}

		c.HTML(200, "application-view", gin.H{
			"Application": application,
			"Events":      events,
			"Theme":       getTheme(c),
			"Active":      "application",
			"SessionUser": session.User,
		})
	} else {
		c.Redirect(301, "/users/login")
	}
}

func getApplications(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		applications, e := database.GetAllApplications()
		if e != nil {
			utils.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		c.JSON(200, applications)
	} else {
		c.JSON(403, false)
	}
}

func applicationCreateAlternative(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		var application structures.Application

		e := c.BindJSON(&application)
		if e != nil {
			utils.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		abonentID, e := database.GetAbonentAlternative(application.Abonent)
		if e != nil {
			utils.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		application.Abonent.ID = abonentID

		if database.CreateApplication(&application, session) {
			if database.CreateEvent(strconv.Itoa(application.ID), session.User.ID, "Создал заявку", "") {
				c.JSON(200, true)
				return
			}
		}

		c.JSON(400, false)
	} else {
		c.JSON(403, false)
	}
}

func applicationCreate(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		var application structures.Application
		e := c.BindJSON(&application)
		if e != nil {
			utils.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		if database.CreateApplication(&application, session) {
			if database.CreateEvent(strconv.Itoa(application.ID), session.User.ID, "Создал заявку", "") {
				c.JSON(200, true)
				return
			}
		}

		c.JSON(400, false)
	} else {
		c.JSON(403, false)
	}
}

func getAbonents(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		abonents, e := database.GetAbonents()
		if e != nil {
			utils.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		c.JSON(200, abonents)
	} else {
		c.JSON(403, false)
	}
}

func userBlockedSwitch(c *gin.Context) {
	session := getSession(c)

	if session.User.Role.ID == 1 || session.User.Role.ID == 6 {
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
	} else {
		c.JSON(403, false)
	}
}

func userUpdate(c *gin.Context) {
	session := getSession(c)

	if session.User.Role.ID == 1 || session.User.Role.ID == 6 {
		id := c.Param("id")
		var user structures.User
		e := c.BindJSON(&user)
		if e != nil {
			utils.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		e = database.UpdateUser(user, id)
		if e != nil {
			utils.Logger.Println(e)
			c.JSON(400, false)
			return
		}

		if user.Password != "" {
			user.Password, e = utils.Encrypt(user.Password)
			if e != nil {
				utils.Logger.Println(e)
				c.JSON(400, false)
				return
			}

			e = database.NewPasswordUser(user.Password, id)
			if e != nil {
				utils.Logger.Println(e)
				c.JSON(400, false)
				return
			}
		}

		c.JSON(200, true)
	} else {
		c.JSON(403, false)
	}
}

func usersCreateHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.Role.ID == 1 || session.User.Role.ID == 6 {
		c.HTML(200, "users-create", gin.H{
			"Theme":       getTheme(c),
			"Active":      "users",
			"SessionUser": session.User,
		})
	} else {
		c.Redirect(301, "/")
	}
}

func usersEditHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.Role.ID == 1 || session.User.Role.ID == 6 {
		userID := c.Param("id")
		user, e := database.GetUserByID(userID)
		if e != nil {
			utils.Logger.Println(e)
			return
		}
		c.HTML(200, "users-edit", gin.H{
			"User":        user,
			"Theme":       getTheme(c),
			"Active":      "users",
			"SessionUser": session.User,
		})
	} else {
		c.Redirect(301, "/")
	}
}

func usersHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.Role.ID == 1 || session.User.Role.ID == 6 {
		users, e := database.GetAllUsers()
		if e != nil {
			utils.Logger.Println(e)
			return
		}

		c.HTML(200, "users", gin.H{
			"Users":       users,
			"Theme":       getTheme(c),
			"Active":      "users",
			"SessionUser": session.User,
		})
	} else {
		c.Redirect(301, "/")
	}

}

func housesHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		c.HTML(200, "houses", gin.H{
			"Theme":       getTheme(c),
			"Active":      "houses",
			"SessionUser": session.User,
		})
	} else {
		c.Redirect(301, "/users/login")
	}

}

func applicationCreateHandler(c *gin.Context) {
	session := getSession(c)

	description := c.Param("text")
	var departmentID int
	var departmentName string

	switch description {
	case "none":
		description = ""
		departmentID = 0
		departmentName = ""
		break
	case "Нет интернета":
		departmentID = 3
		departmentName = "Техники"
		break
	case "Смена ТП":
		departmentID = 1
		departmentName = "Админы"
		break
	case "ТВ":
		departmentID = 4
		departmentName = "ТВ"
		break
	}

	if session.User.ID > 0 {
		c.HTML(200, "application-create", gin.H{
			"Theme":           getTheme(c),
			"Active":          "applicationCreate",
			"SessionUser":     session.User,
			"DescriptionText": description,
			"DepartmentID":    departmentID,
			"DepartmentName":  departmentName,
		})
	} else {
		c.Redirect(301, "/users/login")
	}

}

func applicationHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		c.HTML(200, "application", gin.H{
			"Theme":       getTheme(c),
			"Active":      "application",
			"SessionUser": session.User,
		})
	} else {
		c.Redirect(301, "/users/login")
	}
}

func indexHandler(c *gin.Context) {
	session := getSession(c)

	if session.User.ID > 0 {
		c.HTML(200, "index", gin.H{
			"Theme":       getTheme(c),
			"Active":      "index",
			"SessionUser": session.User,
		})
	} else {
		c.HTML(200, "login", gin.H{
			"Theme": getTheme(c),
		})
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
	session := getSession(c)

	if session.User.Role.ID == 1 || session.User.Role.ID == 6 {
		var userData structures.User
		e := c.BindJSON(&userData)
		if e != nil {
			utils.Logger.Println(e)
			return
		}

		c.JSON(200, database.UsersAdd(&userData))
	} else {
		c.JSON(403, false)
	}
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
		return "light"
	}

	return cookie
}
