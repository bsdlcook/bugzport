package config

func Get() (Config, error) {
	cfg, err := ParseConfig(ConfigFile())
	if err != nil {
		return Config{}, err
	}

	return cfg, err
}
