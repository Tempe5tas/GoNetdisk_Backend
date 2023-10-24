package preference

import (
	"github.com/gin-gonic/gin"
	"go-netdisk/internal/db"
	"go-netdisk/internal/db/models"
	"go-netdisk/pkg/utils"
)

// curl -X POST http://localhost:5000/api/preference/fetch/
func FetchHandler(c *gin.Context) {
	var p models.Preference
	if err := db.DB.First(&p).Error; err != nil {
		utils.Error(c, err)
	}
	utils.Ok(c, p)
}
