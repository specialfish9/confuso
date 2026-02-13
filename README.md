# Confuso
A simple Go project for reading configuration files.

## Features
- Read your configuration from a YAML file
- Read your configuration from environment variables
- Built-in `Optional[T]` type

## How to use it

0. Add dependency
```bash
go get github.com/specialfish9/confuso/v2
```

1. Write your custom config struct, for example:

```go
type Config struct {
    Database   DB     `confuso:"db"`
    Server     HTTP   `confuso:"http"`
}

type DB struct {
    Host     string  					`confuso:"hostname"`
    Port     int     					`confuso:"port"`
    Username string  					`confuso:"username"`
    Password string  					`confuso:"password"`
	SSL 	 confuso.Optional[bool]		`confuso:"ssl"`
}

type HTTP struct {
	Hostname confuso.Optional[string] 	`confuso:"website_hostname"`
    Port     int    					`confuso:"port"`
}

```

2. Write your matching config

```yaml
# whatev.yaml
db:
    hostname: localhost
    port: 5432
    username: admin
    password: secret
http:
    website_hostname: example.com
    port: 8080
```

3. Load it
```go
var config = Config{}

err := confuso.Do("myconf.whatev", &config)
if err != nil {
	log.Fatal(err)
}
```

4. Enjoy!
```go
fmt.Printf("Server listening at: %d!\n", config.Server.Port)
```

## About the `Optional[T]` type
The `Optional[T]` type is a simple wrapper around a value of type `T` that allows
you to distinguish between a set, unset and zero value. No more messing up with 
pointers or default values.

You can declare it with:
```go
type Config struct {
    Port confuso.Optional[int] `confuso:"port"`
}
```

and then use it:
```go
// Check if the port is set
port, ok := config.Port.Val()
if !ok {
    fmt.Println("Port is not set, using default 8080")
}

// Or use the `Or` method to provide a default value
server.Listen(config.Port.Or(8080)) 

// If you are 100% sure that the port is set, you can use `MustVal` 
// which will panic if the value is not set
im100PercentSureThisPortIsNotNill := config.Port.MustVal()
```

## Integrations

### Validator

You can integrate `confuso` with the [validator package](https://github.com/go-validator/validator).


```go
type UserConfig struct {
    Username string `confuso:"username" validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
    Name string     `confuso:"name" validate:"nonzero"`
    Password string `confuso:"password" validate:"min=8"`
}

var userConfig = UserConfig{}

err := confuso.Do("whatev.yaml", &userConfig)
if err != nil {
	log.Fatal(err)
}

if errs := validator.Validate(userConfig); errs != nil {
	// values not valid, deal with errors here
}
```
