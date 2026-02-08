# Confuso
A simple Go project for reading configuration files.

## Features
- Read your configuration from a YAML file
- Built-in `Optional[T]` type

## How to use it

0. Add dependency
```bash
go get github.com/specialfish9/confuso
```

1. Write your custom config struct, for example:

```go
type Config struct {
    Database   DB     `confuso:"db"`
    Server     HTTP   `confuso:"http"`
}

type DB struct {
    Host     string  `confuso:"hostname"`
    Port     int     `confuso:"port"`
    Username string  `confuso:"username"`
    Password string  `confuso:"password"`
}

type HTTP struct {
	Hostname string `confuso:"website_hostname"`
    Port     int    `confuso:"port"`
}

```

2. Write your matching config

```
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
