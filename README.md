# confuso
A simple go project for reading configuration files

## Features
- Read your configuration from a file
- Automatically resolve environment variables
- Custom name for config fields

## How to use it

0. Add dependency
```bash
go get github.com/specialfish9/confuso
```

1. Write your custom config struct, for example:

```go
type Config struct {
	Db   Db
	Http Http
}

type Db struct {
	Host     string
	Port     int
	Username string
	Password string
	UseSsl   bool `confuso:"use_ssl"`
}

type Http struct {
	Hostname string `confuso:"website_hostname"`
	Port     int
}

```

2. Write your matching config

```
# Database
Db.Host=${DB_HOSTNAME}
Db.Port=5432
Db.Username=admin
Db.Password=${DB_PASSWORD}
Db.use_ssl=true

# Http
Http.website_hostname=${HTTP_HOSTNAME}
Http.Port=${HTTP_PORT}
```

3. Load it
```go
var config = Config{}
err := confuso.LoadConf("myconf.whatev", &config)

if err != nil {
	log.Fatal(err)
}
```

4. Enjoy!
```go
fmt.Printf("My website is: %s!\n", config.Http.Port)
```

## Integrations

### Validator

You can integrate `confuso` with the [validator package](https://github.com/go-validator/validator).


```go
type UserConfig struct {
	Username string `validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
	Name string     `validate:"nonzero"`
	Password string `validate:"min=8"`
}

var userConfig = UserConfig{}
err := confuso.LoadConf("myconf.whatev", &userConfig)

if err != nil {
	log.Fatal(err)
}

if errs := validator.Validate(userConfig); errs != nil {
	// values not valid, deal with errors here
}
```
