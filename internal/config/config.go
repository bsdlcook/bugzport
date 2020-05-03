package config

func Get() (Config, error) {
	cfg, err := ParseConfig(File())

	if err != nil {
		return Config{}, err
	}

	return cfg, err
}
