package config

func Get() (Config, error) {
	cfg, err := ParseConfig()

	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
