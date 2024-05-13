package helpers

import (
	"log"
	"time"

	"github.com/cozkul/umai/server/config"
	"github.com/cozkul/umai/server/database"
	"github.com/cozkul/umai/server/models"
	"gorm.io/gorm/clause"
)

func SetUpServerTick() error {
	ticker := time.NewTicker(time.Duration(config.TickInterval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		serverTick()
	}

	return nil
}

func serverTick() error {
	var planets []models.Planet
	database.DB.Where("user_id IS NOT NULL").Preload(clause.Associations).Find(&planets)
	for _, p := range planets {
		for _, b := range p.Buildings {
			p.Resource.AddValue(b.GetProduction())
			log.Printf("Crystal product: %f\n", p.Resource.Crystal)
		}
		p.Resource.Clamp(p.GetMaxCapacity())
		log.Printf("Crystal product clamped: %f\n", p.Resource.Crystal)
		database.DB.Save(&p.Resource)
	}
	return nil
}
