package main

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"reconciliation/internal/delivery/http/route"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic().Msg("please set config.json file")
	}

	if viper.GetBool(`debug`) {
		log.Info().Msg("Service RUN on DEBUG mode")
	}
}

func main() {
	host := viper.GetString(`server.host`)
	port := viper.GetString(`server.port`)
	router := route.NewRouteConfig()
	router.Listen(host + ":" + port)
}
