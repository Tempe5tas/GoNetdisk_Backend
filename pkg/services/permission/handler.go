package permission

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-netdisk/internal/db"
	models2 "go-netdisk/internal/db/models"
	"go-netdisk/pkg/utils"

	"gorm.io/gorm"
)

// curl http://localhost:5000/api/permission/self_permissions/
func SelfPermissionsHandler(c *gin.Context) {
	var p *models2.Permission

	username := c.GetString("username")
	p, err := models2.GetPermissionByUsername(username)

	if err != nil {
		// No permission item found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Ok(c, gin.H{})
			return
		}
		utils.Error(c, err)
	}

	utils.Ok(c, p)
}

// curl http://localhost:5000/api/permission/get_my_project/
func MyProjectHandler(c *gin.Context) {
	var perm *models2.Permission
	var project *models2.Project

	username := c.GetString("username")
	perm, err := models2.GetPermissionByUsername(username)

	if err != nil {
		// No permission item found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Ok(c, gin.H{})
			return
		}
		utils.Error(c, err)
	}

	if perm.Role == models2.ADMINISTRATOR {
		err = db.DB.Order("-created_at").First(&project).Error
	} else {
		err = db.DB.First(&project, "uuid = ?", perm.ProjectUUID).Error
	}

	if !errors.Is(err, nil) {
		utils.Error(c, err)
		return
	}

	utils.Ok(c, project)
}
