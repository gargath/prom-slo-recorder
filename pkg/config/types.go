package config

import()

type Config struct {
        Services map[string]Service
}

type Service struct {
        Name string
        PromAddr string
        Query string
}
