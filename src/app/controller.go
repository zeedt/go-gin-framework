package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func registerRoute() *gin.Engine {
	route := gin.Default()

	route.LoadHTMLGlob("templates/**/*.html")

	route.GET("/test", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello from %s", "Gin" )
	})
	route.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil )
	})

	route.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", nil)
	})


	route.GET("/employees/:id/vacation", func(context *gin.Context) {
		id := context.Param("id")
		timesOff, ok := TimesOff[id]
		if !ok {
			context.String(http.StatusNotFound, "Page not found", nil)
			return
		}
		context.HTML(http.StatusOK, "vacation-overview.html", gin.H{
			"TimesOff" : timesOff,
		})
	})

	route.POST("/employees/:id/vacation/new", func(context *gin.Context) {
		var timeOff TimeOff
		err := context.BindJSON(&timeOff)
		if err != nil {
			context.String(http.StatusBadRequest, err.Error())
			return
		}

		id := context.Param("id")
		timesOff, ok := TimesOff[id]

		if !ok {
			TimesOff[id] = []TimeOff{}
		}
		TimesOff[id] = append(timesOff, timeOff)

		context.JSON(http.StatusCreated, timeOff)
	})

	route.GET("/load-file", func(context *gin.Context) {
		context.File("/Users/saheedyususf/Downloads/ref.pdf")
	})

	admin := route.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin" : "admin",
	}))

	admin.GET("/employees/:id", func(context *gin.Context) {
		id := context.Param("id")
		if id == "add" {
			context.HTML(http.StatusOK, "admin-employee-add.html", nil)
			return
		}

		employee, ok := employees[id]
		if !ok {
			context.String(http.StatusNotFound, "Employee not found")
			return
		}

		context.HTML(http.StatusOK, "admin-employee-edit.html", gin.H{
			"Employee" : employee,
		})
	})

	admin.POST("/employees/:id", func(context *gin.Context) {
		id := context.Param("id")
		if id == "add" {
			startDate, err := time.Parse("2006-01-02", context.PostForm("startDate"))
			if err != nil {
				context.String(http.StatusBadRequest, err.Error())
				return
			}

			var employee Employee
			context.Bind(&employee)
			if err != nil {
				context.String(http.StatusBadRequest, err.Error())
				return
			}
			employee.ID = 42
			employee.StartDate = startDate
			employee.Status = "Active"

			employees["42"] = employee

			context.Redirect(http.StatusMovedPermanently, "/admin/employees/42")
		}
	})

	admin.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK,"admin-overview.html", nil)
	})


	route.Static("public", "./public")

	return route
}
