package main

import (
	"github.com/gin-gonic/gin"
	"io"
)

func initGroupHandlers(r *gin.Engine) {
	r.GET("/group", GetAllGroup)
	r.POST("/group", PostGroup)
	r.GET("/group/:id", GetGroup)
	r.GET("/group/:id/avatar", GetGroupAvatar)
	r.PUT("/group/:id/avatar", SetGroupAvatar)
	r.PUT("/group/:id", UpdateGroup)
	r.DELETE("/group/:id", DeleteGroup)
}

func GetAllGroup(c *gin.Context) {
	var groups []Group
	db.Model(&Group{}).Find(&groups)
	c.JSON(200, groups)
}

func PostGroup(c *gin.Context) {
	var group Group
	err := c.BindJSON(&group)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Create(&group)
	c.JSON(200, group)
}

func GetGroup(c *gin.Context) {
	var group Group
	id := c.Params.ByName("id")
	notFound := db.First(&group, id).RecordNotFound()
	if notFound {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, group)
}

func UpdateGroup(c *gin.Context) {
	var group Group
	id := c.Params.ByName("id")
	db.First(&group, id)
	err := c.BindJSON(&group)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Save(&group)
	c.JSON(200, group)
}

func DeleteGroup(c *gin.Context) {
	var group Group
	id := c.Params.ByName("id")
	db.First(&group, id)
	db.Delete(&group)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func GetGroupAvatar(c *gin.Context) {
	var group Group
	id := c.Params.ByName("id")
	db.First(&group, id)
	c.Data(200, "image/png", group.Avatar)
}

func SetGroupAvatar(c *gin.Context) {
	var group Group
	id := c.Params.ByName("id")
	db.First(&group, id)
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	group.Avatar, err = io.ReadAll(file)
	db.Save(&group)
	c.JSON(200, group)
}
