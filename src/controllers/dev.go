package controllers

import (
	"dev-crud/src/database"
	"dev-crud/src/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Create a dev
// @Description create a dev
// @Tags devs
// @Accept json
// @Produce json
// @Param dev body models.Dev true "Dev"
// @Success 201 {object} models.Dev
// @Router /dev [post]
func CreateDev(context *gin.Context){
	var developer models.Dev

	if err := context.ShouldBindJSON(&developer); err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := database.Instance.Create(&developer); result.Error != nil {
		context.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	context.JSON(201, developer)
}

// @Summary Get all devs
// @Description get devs
// @Tags devs
// @Produce json
// @Success 200 {array} models.Dev
// @Router /dev [get]
func GetDevs(context *gin.Context){
	var devs []models.Dev
	
	if result := database.Instance.Find(&devs); result.Error != nil {
		context.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	context.JSON(200, devs)
}

// @Summary Get dev by id
// @Description get dev by id
// @Tags devs
// @Produce json
// @Success 200 {object} models.Dev
// @Router /dev/:id [get]
func GetDevById(context *gin.Context) {
    id := context.Param("id")
    var dev models.Dev
    var result *models.Dev
    var err error

    if result, err = dev.FindById(database.Instance, id); err != nil {
        context.JSON(404, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(200, result)
}

// @Summary Remove dev by id
// @Description remove dev by id
// @Tags devs
// @Produce json
// @Success 204
// @Router /dev/:id [delete]
func RemoveDev(context *gin.Context){
	id := context.Param("id")
	var dev models.Dev
	var result *models.Dev
    var err error
	
   if result, err = dev.FindById(database.Instance, id); err != nil {
        context.JSON(404, gin.H{
            "error": err.Error(),
        })
        return
    }

	if update := database.Instance.Model(&result).UpdateColumn("deleted_at", gorm.DeletedAt{}); update.Error != nil {
		context.JSON(500, gin.H{
			"error": update.Error.Error(),
		})
		return
	}

	context.JSON(204, nil)
}