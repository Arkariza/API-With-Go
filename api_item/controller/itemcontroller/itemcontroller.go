package itemcontroller

import (
	"net/http"
	"strconv"

	"github.com/user/api_item/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var item []models.Item

	models.DB.Find(&item)
	c.JSON(http.StatusOK, gin.H{"items":item})
}
func Show(c *gin.Context) {
	var item models.Item
	id := c.Param("id")
	if err := models.DB.First(&item, id).Error; err != nil{
		switch err{
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message":"data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message":err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"item": item})
}
func Create(c *gin.Context) {
	var item models.Item
	
	if err := c.ShouldBindJSON(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":err.Error()})
		return
	}
	models.DB.Create(&item)
	c.JSON(http.StatusOK, gin.H{"item": item})
}
func Update(c *gin.Context) {
	var item models.Item
	id := c.Param("id")
	
	if err := c.ShouldBindJSON(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":err.Error()})
		return
	}
	if models.DB.Model(&item).Where("id = ?", id ).Updates(&item).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat Memperbarui"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di Update"})
}

func Delete(c *gin.Context) {
	var item models.Item
	input := map[string]string{"id": "0"}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":err.Error()})
		return
	}
	id, _ := strconv.ParseInt(input["id"], 10, 64)
	if models.DB.Delete(&item, id).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":"Tidak dapat Menghapus Item"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di Hapus"})
}