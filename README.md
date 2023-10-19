# Structure Golang Clean Architecture

Clean Architecture is a design philosophy that can be applied to build maintainable, scalable, and testable applications. It follows the same principles as Clean Architecture in general, but it is adapted to the specific features and characteristics of the Go programming language.

## Structure Directory & PKG

Minimum versi Go adalah v.1.15 dengan go mod sebagai dependency injectionnya.

```bash
<root>
  ├── bin
  |   ├── <service_name_executable>
  |   └── ...
  |
  ├── cmd
  |   ├── <service_name>_main.go
  |   └── ...
  |
  ├── deploy
  |   └── <service_name>
  |       ├── deployment.yml
  |       ├── Dockerfile
  |       └── service.yml
  |
  ├── domain
  |   ├── entity
  │   |   ├── <domain_name_entity>.go
  |   |   └── <domain_name_entity>_test.go
  |   |
  |   ├── value_object
  │   |   ├── <domain_value_object>.go
  |   |   └── <domain_value_object>_test.go
  |   |
  |   ├── repository
  │   |   └── <domain_repository>_interface.go
  |   |
  |   ├── use_case
  │   |   └── <domain_use_case>_interface.go
  |   |
  |   ├── service
  │   |   └── <domain_service>_interface.go
  |   |
  |   └── application
  │       └── <domain_application>_interface.go
  |
  ├── internal  
  |   | 
  │   ├── delivery
  │   │   └── http
  |   |       ├── <domain_name>_init.go
  |   |       ├── <domain_name>_init_test.go
  |   |       ├── <domain_name>_handler.go
  |   |       └── <domain_name>_handler_test.go
  |   |   
  |   ├── repository
  |   |   └── mocks
  |   |   |   └── <repo_implementation>Mock.go
  |   |   └── <repo_implementation>.go
  |   |
  |   ├── service
  |   |   └── mocks
  |   |   |   └── <service_implementation>Mock.go
  |   |   └── <service_implementation>.go
  |   |
  |   ├── application
  |   |   └── mocks
  |   |   |   └── <application_implementation>Mock.go
  |   |   └── <application_implementation>.go
  |   |
  |   └── usecase
  |       ├── mocks
  |       |  └── <usecase_name>UseCaseMock.go
  |       ├── <usecase_name>.go
  |       └── <usecase_name>_test.go
  |
  ├── pkg
  |   └── <pkg_name>
  |       ├── <pkg_implementation>.go
  |       └── <pkg_implementation>_test.go
  | 
  ├── vendor
  ├── README.txt
  └── MAKEFILE
```

| Name                                                 | Purpose |
|------------------------------------------------------| ------- |
| bin                                                  | Golang executable file per service name        |
| cmd                                                  | Golang main program per service name       |
| deploy                                               | Configuration file for deployment per service name       |
| domain/*                                             | Golang file that holds definiton for data model, usecase interface, and repository interface for each domain       |
| pkg/*                                                | Functions, configs used by microsevices       |
| internal/delivery/http                               | As per port adapter pattern, this the port definition that interfacing with outside world       |
| internal/repository/<repo_implmentation>.go          | Encapsulated Implementation of Repository Interface        |
| internal/repository/mocks/*                          | Mock file for testing purpose         |
| internal/usecase/<usecase_implementation>.go         | Encapsulated Implementation of UseCase Interface        |
| internal/usecase/mocks/*                             | Mock file for testing purpose         |
| internal/service/<service_implmentation>.go          | Encapsulated Implementation of service Interface        |
| internal/service/mocks/*                             | Mock file for testing purpose         |
| internal/application/<application_implementation>.go | Encapsulated Implementation of application Interface        |
| internal/application/mocks/*                         | Mock file for testing purpose         |
| vendor                                               | Golang external package dependencies       |
| README.md                                            | Manual handbook to use the service        |
| makefile                                             | Script definition to build, test, lint, and deploy the application       |

## Recommendation HTTP ROUTER GO

Gorilla-Mux : https://github.com/gorilla/mux

Echo - Router : https://echo.labstack.com/docs/routing

Gin-Gonic : https://github.com/gin-gonic/gin


