package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	todoapp "todo-app"
	"todo-app/db"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/service"
)

func Start()  {
	/***************** Инициализация config-ов *****************/
	if err := initConfigs(); err != nil {
		log.Fatalf("Error while initializing configs. Error is: %s", err.Error())
	}
	/**********************************************************/

	/***************** Инициализация базы данных *****************/
	database, err := repository.NewSqliteDB(viper.GetString("db.dbname"))
	if err != nil {
		log.Fatalf("Error while opening DB. Error is: %s", err.Error())
	}
	db.Init(database)
	/**********************************************************/

	/***************** Внедрение зависимостей *****************/
	repos := repository.NewRepository(database)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	/**********************************************************/

	/***************** Starting App *****************/
	MainServer := new(todoapp.Server)
	go func() {
		if err := MainServer.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("Error while running http server. Error is %s", err.Error())
		}
	}()
	fmt.Println("TodoApp Started its work")
	fmt.Printf("Server is listening port: %s\n", viper.GetString("port"))
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit
	/**********************************************************/

	/***************** Shutting App Down *****************/
	fmt.Println("TodoApp Shutting Down")
	if err := MainServer.Shutdown(context.Background()); err != nil  {
		log.Fatalf("error while shutting server down. Error is: %s", err.Error())
	}
	if err := database.Close(); err != nil {
		log.Fatalf("error while closing DB. Error is: %s", err.Error())
	}
	/**********************************************************/
}

func main() {
	Start()
}

//Функция инициализации config-ов
func initConfigs() error {
	viper.AddConfigPath("configs") //адрес директории
	viper.SetConfigName("config") //имя файла
	return viper.ReadInConfig() //считывает config и сохраняет данные во внутренний объект viper
}