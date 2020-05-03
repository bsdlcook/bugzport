package config

func Get() (Config, error) {
	cfg, err := ParseConfig(filePath())

	if err != nil {
		return Config{}, err
	}

	return cfg, err
}
