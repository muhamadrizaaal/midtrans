package config

// Config menyimpan konfigurasi aplikasi
type Config struct {
	EnvConfig
}

// LoadConfig digunakan untuk memuat konfigurasi aplikasi
func LoadConfig() Config {
	env := LoadEnvConfig()
	return Config{
		EnvConfig: env,
	}
}
