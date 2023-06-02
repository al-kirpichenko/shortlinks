package app

import (
	"context"
	"github.com/al-kirpichenko/shortlinks/internal/database"
	"log"
	"net/http"
	"time"
)

func (a *App) Ping(w http.ResponseWriter, r *http.Request) {

	db, err := database.Connect(a.cfg.DataBase)
	if err != nil {
		log.Println("don't connect Database")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		log.Println("don't ping Database")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
