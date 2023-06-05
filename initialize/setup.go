package initialize

import (
	"github.com/chaperone/server/config"
	"github.com/chaperone/server/router"
)

func Setup() {
	config.SetupEnv()
	config.SetupMysql()

	router.SetupRouter()
}

func init() {

}

func init() {

}
