package file_service_config

type Config struct {
	MaxSizeMb int
}

func (h Config) IsNotValidSize(size int) bool {
	return (h.MaxSizeMb << 20) < size
}
