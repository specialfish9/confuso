# confuso
A simple go project made just for fun

## How to use it

1. Write your struct

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

2. Write your config

```
Db.Host=${DB_HOSTNAME}
Db.Port=5432
Db.Username=admin
Db.Password=${DB_PASSWORD}
Db.use_ssl=true

Http.website_hostname=${HTTP_HOSTNAME}
Http.Port=${HTTP_PORT}
```

3. Load the config
```go
	var config = Config{}
	err := goconf.LoadConf("myconf.whatev", &config)

	if err != nil {
		log.Fatal(err)
	}
```

4. Enjoy!
```go
fmt.Printf("My website is: %s!\n", config.Http.Port)
```
