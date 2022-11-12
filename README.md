# Steps 

## Создать систему папок
Для начала написания кода нужна система папок, внутри который мы будем хранить будущий проект

```


├───cmd
│       ├───api
│       │       └───main.go
│       │               
│       │               
│       └───bot
│               └───main.go            
├───config
│       
└───internal
        └───models

```

**Цели каждой папки**:
- cmd = хранит пакеты, которые доступны извне
- config = программа хранит внутри данной папки основную информацию о DSN баз данных, BotToken телеграмм-бота и т.д.
- internal = хранит пакеты , которые доступны только внутри программы, недоступные извне
- cmd/api = хранит наш каркас Api, благодаря работает наш веб-сервис
- cmd/bot = хранит наш каркас Бота
- internal/models = хранит структуры объектов (entities) и подключения нашей базы данных к самому проекту.

>В проекте 2 main файла в 2 директориях. Делается это для того, чтобы запустить 2 сервера в рамках 1 проекта

## Создать модуль

```shell
go mod init upgrade
```

## Создать подключение к пакету файла конфигурации

```shell
go get github.com/BurntSushi/toml@latest
```

## Создаем базу данных users.db
## Создаем файл local.toml , который содержит информацию о названии нашей базы данных
config/local.toml
```toml
Env="local"
BotToken="5667090127:AAHj1OTBYvj2KJ_MA0-cL1Ys9oa019npuPw"
Dsn="upgrade.db"
```

## Подключаем пакеты GORM 
```shell
go get gorm.io/gorm  
go get gorm.io/driver/sqlite
```

## Дополняем main.go

``` go
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	BotToken string
	Dsn      string
}

func handleNewsLetter(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

func main() {

	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	cfg := &Config{}
	_, err := toml.DecodeFile(*configPath, cfg)

	if err != nil {
		log.Fatalf("Ошибка декодирования файла конфигов %v", err)
	}

	db, err := gorm.Open(sqlite.Open("../../"+cfg.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Ошибка подключения к БД %v", err)
	}
	fmt.Printf(db.Name())

	http.HandleFunc("/newsletter", handleNewsLetter)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### Отвечает за запуск вебсервиса по заданной функции
``` go
    func handleNewsLetter(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
    }


	http.HandleFunc("/newsletter", handleNewsLetter)
	log.Fatal(http.ListenAndServe(":8080", nil))
```



### Отвечает за открытие соединения к базе данных users.db  
``` go
	db, err := gorm.Open(sqlite.Open("../../"+cfg.Dsn), &gorm.Config{})
```
