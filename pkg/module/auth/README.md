# Authentication Module

## Getting started

This is repo for our backend module, Authentication Module. Authentication Module responsible to handle
user authentication.

## Implementation
Make sure you already have the source in your local by executing this command below
```
cd $PROJECT_HOME
git submodule sync
```

Then, you need to register this module by inserting this code below to `reg_mod.go`

```go
// ===> START REGISTER AUTH MODULE <=== //
authModule := auth.NewModule(module.ModuleConfig{
    Repository: *repo,
})
authModule.Router.Register(r.Group("/auth"))
// <=== END REGISTER AUTH MODULE ===> //
```

## Group routes
- Rest API Server: `localhost:3000/moduleName`
- GraphQL Server: `localhost:3000/moduleName/graph`

replace `moduleName` with this module name.
