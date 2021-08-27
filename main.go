package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

type Config struct {
	ServerUrl string
	Endpoint  string
}

func configServer(config *Config) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Can't read config file: %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Can't unmarshal config file: %v", err)
	}
}

func main() {
	app := iris.New()
	var config Config
	configServer(&config)
	fmt.Println(config)
	app.RegisterView(iris.HTML("./views", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello World!")

		ctx.View("hello.html")
	})

	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
		userID, _ := ctx.Params().GetUint64("id")
		ctx.Writef("User ID: %d", userID)
	})

	res, err := http.Get("https://bec142c164838d04d3725cc336cce26d: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@testingapi-mineloop.myshopify.com/admin/api/2021-07/products.json")

	if err != nil {
		log.Fatalf(err.Error())
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(string(body))
	app.Run(iris.Addr(":8080"))
}
