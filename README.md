# Modular Go Boilerplate

This repository is a codebase for Golang Project with Modular Implementation. One module didn't dependence to other modules. It's suitable for you who want to implement monorepo microservices.

## Use Case
Our company wants to develop new product. We decided to not use Microservice. But we want to get the benefit of microservice where each service will not have interdependance to other service. So, we decide to create one Go Boilerplate with Modular Monorepo implementation.

1. Each Module is dedicated.
2. You can make each module to git submodule. Where each developer has access to specific module
3. Easy to migrate to microservice in near future

## Feature
- Modular Module
- GraphQL Server
- Rest API
- Integration Test & Unit Testing
- Translations
- Data Seeders
- Docker

## How to usage

### Before Run
Execute the service dependency via Docker.

Since we use Air for our development support. Install the Air Live Reload
If you got issue with the installation please follow official docs: https://github.com/cosmtrek/air

```
make run-setup 
```

Then run up the Docker Compose
```
docker-compose up -d
```

### Start Dev Server

to run command on development:

```
make run-dev
```

### Seeder

to create dummy data execute

```
make run-seeds
```

### Create & Register new Module

Just duplicate sample module and replace module name. Then register those module to the module registry

```
moduleName.Init(options.GinEngine, options.Repository)
```

## GraphQL Server & Playground
You can access the graphql playground via

```
http://localhost:3000/graph
```

and the graphql server depends on the module you have

Root Graphql Server

**http://localhost:3000/graph**

Module Graphql Server

**http://localhost:3000/moduleName/graph/query**

Sample

```
// Auth Module
http://localhost:3000/auth/graph/query

// Other Module
http://localhost:3000/other/graph/query
```

## Library

1. Web Server : Gin Framework
2. Database : PostgreSQL
3. Cache Storage : Redis
4. GraphQL Server: GQLGen

## Commit Prefixes

- feat: (new feature for the user, not a new feature for build script)
- fix: (bug fix for the user, not a fix to a build script)
- docs: (changes to the documentation)
- style: (formatting, missing semi colons, etc; no production code change)
- refactor: (refactoring production code, eg. renaming a variable)
- test: (adding missing tests, refactoring tests; no production code change)
- chore: (updating grunt tasks etc; no production code change)

> https://www.conventionalcommits.org
>
> https://gist.github.com/joshbuchea/6f47e86d2510bce28f8e7f42ae84c716
