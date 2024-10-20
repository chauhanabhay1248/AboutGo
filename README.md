# AboutGo
``` go run app.go ```
or 
``` go run . ``` if you have the go module.

**go build**
```
- it wants a mod file with say the module to build.
- package main says that this module is an executable module and can have mutiple main packages.
- main func says this is the entry point for the application, we can't have mutiple main function under in one application.
```

**Variables:**
```
- var newVariable <type> = 1000 -
    ` here var is use to define new variable and newVariable should be in camelCase format.
    ` type [ which is OPTIONAL ] can be: int, float64, string bool.
- newVariable := 1000 - 
    a way to declare ans assign a new variable
- const newVariable = 100 -
    here we can't change the newVariable value further.
- we have use `` for "", where `` can be use for multiline string but "" can't.
- interface {} -
    it represents the empty interface, it can hold value of any type.
```