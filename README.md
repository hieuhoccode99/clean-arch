## Clean architecture
```
├── cmd
│   └── main.go # entry point
├── config
│   ├── config.go # config struct
│   └── config.yml # config file
├── domain 
│   ├── entity # define entity
│   │   └── user.go
|       ├── base.go
├── infrastructure
│   ├── database.go # connect to database
├── pkg
│   ├── user # define user feature
│   │   ├── repository.go # define repository
│   │   ├── handler.go # define handler
│   │   ├── usecase.go # define usecase
│   │   └── dto.go # define dto
├── route 
│   └── route.go # define route
├── utils
│   ├── constants
│   │   └── constants.go # define constants
│   ├── data.go # transform data
```

