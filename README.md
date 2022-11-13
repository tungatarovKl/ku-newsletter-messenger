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

## Инициализировать библиотеки
```shell
go mod vendor
```
## Указать пароль в docker-compose.yml и config/local.toml
## Создать базу данных MySQL из коренной папки (Docker), запустив ее 
```shell
docker-compose up
```
## Запустить проект
```shell
 go run cmd/api/main.go -config=config/local.toml
```
