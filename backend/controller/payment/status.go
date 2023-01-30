package controller

import (
	"github.com/sasiprapa45/test18/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)


//ไว้สำหรับดึงข้อมูล--------------------------------------
// GET /Status/:id
func GetStatus(c *gin.Context) {
	var status entity.Status
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM statuses WHERE id = ?", id).Scan(&status).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": status})
}

// GET /Payee
func ListStatus(c *gin.Context) {
	var status []entity.Status
	if err := entity.DB().Raw("SELECT * FROM statuses").Scan(&status).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": status})
}

