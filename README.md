# Overview

It's an Clean Architecture Skeleton project based on Chi framework.
Our aim is reducing development time on default features that you can meet very often when your work on API.
There is a useful set of tools that described below. Feel free to contribute!


## What's inside:

- CRUD API
- Migrations
- Request validation
- Swagger docs
- Environment configuration
- Docker development environment


## Usage

1. Copy .env.dist to .env and set the environment variables. In the .env file set these variables:
2. Browse to {HTTP_HOST}:{HTTP_PORT}/swagger/index.html. You will see Swagger 2.0 API documents.


## Directories

1. **main.go**: contains the application's main entry point(s) or command-line interfaces (CLIs). Each subdirectory represents a different executable within the project
2. **/internal**: houses the internal components of your application that are not intended to be imported by external projects. This directory typically contains packages/modules related to business logic, domain models, repositories, services, and configuration.
3. **/internal/app**: this section may include any initialization code that needs to be executed before the application starts. For example, setting up configuration, connecting to databases, or initializing logging.
4. **/internal/cache**: directory allows for the separation of caching concerns from other parts of the application, promoting modularity and maintainability. By isolating caching-related code, it becomes easier to manage and test caching functionality independently. However, the specific directory structure and organization may vary based on the project's needs and preferences.
5. **/internal/config**: holds the configuration-related code and files. It includes the logic to read and parse configuration files, environment variables, or other sources of configuration data. It provides a centralized way to manage and access application configuration throughout the codebase.
6. **/internal/domain**: directory, you separate the core business logic from infrastructure-specific or framework-specific code. This separation helps keep your code clean, maintainable, and easier to test. It also allows for better reusability and modularity, as the domain layer can be used independently of the specific infrastructure or framework being used.
7. **/internal/handler**: contains the HTTP or RPC handlers for the application. These handlers are responsible for receiving incoming requests, parsing them, invoking the necessary business logic, and returning the appropriate responses. Each handler typically corresponds to a specific endpoint or operation in the application's API.
8. **/internal/repository**: contains the implementation of data access and persistence logic. It provides an abstraction over the data storage layer, allowing the application to interact with databases, or other external systems. Repositories handle the CRUD operations and data querying required by the application.
9. **/internal/service**: contains the implementation of the application's business logic. It encapsulates the core functionality of the application and provides high-level operations that the handlers can use to accomplish specific tasks. Services interact with data repositories, external APIs, or other dependencies to fulfill the application's requirements.
10. **/migrations**: contains database migration scripts, which are used to manage database schema changes over time.
11. **/pkg**: contains packages that can be imported and used by external projects. These packages are typically utilities, libraries, or modules that have potential for reuse across different projects.


## Libraries

1. Router: https://github.com/go-chi/chi
2. Migrations: https://github.com/golang-migrate/migrate
3. Swagger: https://github.com/swaggo/http-swagger


### Swagger getting started

1. Add comments to your API source code, See [Declarative Comments Format](#declarative-comments-format).

2. Download swag by using:
```sh
go install github.com/swaggo/swag/cmd/swag@latest
```
To build from source you need [Go](https://golang.org/dl/) (1.17 or newer).

Or download a pre-compiled binary from the [release page](https://github.com/swaggo/swag/releases).

3. Run `swag init` in the project's root folder which contains the `main.go` file. This will parse your comments and generate the required files (`docs` folder and `docs/docs.go`).
```sh
swag init
```

  Make sure to import the generated `docs/docs.go` so that your specific configuration gets `init`'ed. If your General API annotations do not live in `main.go`, you can let swag know with `-g` flag.
  ```sh
  swag init -g internal/handler/handler.go
  ```

4. (optional) Use `swag fmt` format the SWAG comment. (Please upgrade to the latest version)

  ```sh
  swag fmt
  ```


## License

The project is developed by [NIX Solutions](http://nixsolutions.com) Go team and distributed under [MIT LICENSE](https://github.com/nixsolutions/golang-echo-boilerplate/blob/master/LICENSE)