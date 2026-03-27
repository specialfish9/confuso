package confuso_test

import "github.com/specialfish9/confuso/v2"

type Config struct {
	This struct {
		Is struct {
			A struct {
				String string `confuso:"string"`
				Bool   bool   `confuso:"bool"`
				Number int    `confuso:"number"`
			} `confuso:"a"`
		} `confuso:"is"`
	} `confuso:"this"`
	Other struct {
		Object string `confuso:"object"`
	} `confuso:"other"`
}

type ConfigWithOptional struct {
	This struct {
		Is struct {
			A struct {
				String    confuso.Optional[string] `confuso:"string"`
				Bool      confuso.Optional[bool]   `confuso:"bool"`
				Number    confuso.Optional[int]    `confuso:"number"`
				OptString confuso.Optional[string] `confuso:"non_existent_string"`
				OptBool   confuso.Optional[bool]   `confuso:"non_existent_bool"`
				OptNumber confuso.Optional[int]    `confuso:"non_existent_number"`
			} `confuso:"a"`
		} `confuso:"is"`
	} `confuso:"this"`
}
