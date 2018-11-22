package apiv0

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type api struct {
	DB *gorm.DB
}

func NewAPIv0(r *mux.Router, db *gorm.DB) {
	apiHandler := &api{
		DB: db,
	}

	db.AutoMigrate(&Stand{})

	r.HandleFunc("/stand", apiHandler.getStands)
}