package initial

import (
	"go-netdisk/internal/db"
	models2 "go-netdisk/internal/db/models"
	"go-netdisk/internal/settings"
	"go-netdisk/pkg/utils/misc"
	"log"
)

func InitData() {
	if settings.ENV.NeedMigrate {
		log.Printf("Start migrate database\n")
		_ = db.DB.AutoMigrate(&models2.Project{}, &models2.User{}, &models2.Permission{}, &models2.Matter{}, models2.Preference{})
	}

	log.Printf("Create superuser: %s", settings.ENV.SuperUser)
	if _, err := models2.GetOrCreateUser(settings.ENV.SuperUser, true); err != nil {
		panic(err)
	}

	perm := &models2.Permission{}
	db.DB.Where(models2.Permission{UserName: settings.ENV.SuperUser}).Attrs(models2.Permission{
		Role: models2.ADMINISTRATOR,
	}).FirstOrCreate(&perm)
	log.Printf("GetOrCreate permission: %s\n", misc.PrettyJson(perm))

	prefer := &models2.Preference{}
	db.DB.Where(models2.Preference{Name: "netdisk"}).Attrs(models2.Preference{
		AllowRegister: true,
	}).FirstOrCreate(&prefer)
	log.Printf("GetOrCreate preference: %s\n", misc.PrettyJson(prefer))

	if db.DB.First(&models2.Project{}).RowsAffected == 0 {
		log.Printf("Create default project")
		project := models2.Project{
			Name:        "DEMO",
			Description: "DEMO",
		}
		if err := db.DB.Create(&project).Error; err != nil {
			panic(err)
		}
	}

}
