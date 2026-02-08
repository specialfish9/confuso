package confuso

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
				String    Optional[string] `confuso:"string"`
				Bool      Optional[bool]   `confuso:"bool"`
				Number    Optional[int]    `confuso:"number"`
				OptString Optional[string] `confuso:"non_existent_string"`
				OptBool   Optional[bool]   `confuso:"non_existent_bool"`
				OptNumber Optional[int]    `confuso:"non_existent_number"`
			} `confuso:"a"`
		} `confuso:"is"`
	} `confuso:"this"`
}
