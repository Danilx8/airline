# airline
## Intro
Это базовая ветка для бэкеров. Тут находится "общий" докер компоуз.
## Project tree
```
├── app
│   │── api
│   │   │── controller
│   │   └─ route
│   │── bootstrap
│   │── cmd
│   │── domain
│   │── repository
│   └── usecase
├── cron
│   ├── crontab
│   └── scripts
│       └── utils
├── README.md
├── docker-compose.yml
├── sql_scripts
│   └── sql files...
└── source_database
    └── mysql files
```
## HOWTO
Тут будет описываться как запустить проект и другие детали для разработчиков.
### Generate docs
Чтобы сгенерировать swagger, нужно установить все зависимости запустить данную команду:
```
swag init -g cmd/main.go -d app/ --pd true
```
Далее переходим [сюда](http://localhost:8080/docs/index.html) и наслаждаемся
