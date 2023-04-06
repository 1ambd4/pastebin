package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	registerRouter(r)
	Init()

	r.Run()
}

func registerRouter(e *gin.Engine) {
	new(RouteManager).Router(e)
}

type RouteManager struct{}

func (m *RouteManager) Router(e *gin.Engine) {
	v1 := e.Group("/v1")
	{
		v1.GET("/s/:id", m.GetText)
		v1.PUT("/s/:id", m.PutText)
		v1.POST("/s", m.PostText)
		v1.DELETE("/s/:id", m.DeleteText)
	}
}

// xxx.yyy.zzz/v1/s/:id
func (m *RouteManager) GetText(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	context, err := GetContext(id)
	if err != nil {
		panic(err)
	}
	if len(context) == 0 {
		context = "invalid id"
	}

	c.JSON(200, map[string]any{
		"id":      id,
		"context": context,
	})
}

func (m *RouteManager) PutText(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	context := c.PostForm("context")

	if err := UpdateContext(id, context); err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"context": context,
	})
}

// xxx.yyy.zzz/v1/s
func (m *RouteManager) PostText(c *gin.Context) {
	context := c.DefaultPostForm("context", "empty data")
	if id, err := InsertContext(context); err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	}
}

// xxx.yyy.zzz/v1/s/:id
func (m *RouteManager) DeleteText(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if msg, err := DeleteContext(id); err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":  id,
			"msg": msg,
		})
	}
}
