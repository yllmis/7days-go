package app

import (
	"github.com/vote_demo/app/model"
	"github.com/vote_demo/app/router"
)

func Strat() {
	model.NewMysql()
	defer func() {
		model.Close()
	}()

	router.NewRouter()
}
