### Getting Started
Golang starter REST API with Gin Gonic

running the application
`go run main.go`

generate swagger 
`swaggo init -g main.go`

### Structure
|  Folder  |  Sub / File  |  Description  |
| ------------ | ------------ | ------------ |
| api  | db  |  ORM setup database , models, and seeder  |
|   | logs   | Generated logs file   |
|   | modules   | All bussines process with modularity method and dependecy injection   |
|   | utils   | library / utility application   |
|   | routes.go   | Initialized all route from modules   |
|   | server.go   | Initilalized server  framework |
|   | wire.go   | Initialized google wire build |
|   | wire_gen.go   | generated google wire |

### Prequisite
Required Dependencies in your local computer

| Dependency  | Source   |
| ------------ | ------------ |
|  Google Wire Cmd |  github.com/google/wire/cmd/wire  |

Core Dependencies

| Dependency  | Source   | Desc   |
| ------------ | ------------ | ------------ |
|  Google Wire |  https://github.com/google/wire  | Dependecy Injection
|  Gin Gonic |  https://github.com/gin-gonic/gin  | HTTP web framework
|  Gorm |  https://github.com/jinzhu/gorm  | ORM Library
|  Swagger |  https://github.com/swaggo/swag  | API documentation
| Zerolog | https://github.com/rs/zerolog | Logger
