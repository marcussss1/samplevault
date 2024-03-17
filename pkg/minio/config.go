package minio

// TODO без дефолта
type Config struct {
	DSN       string `env:"MINIO_DSN,required" envDefault:"minio:9000"`
	AccessKey string `env:"MINIO_ACCESS_KEY,required" envDefault:"your_access_key"`
	SecretKey string `env:"MINIO_SECRET_KEY,required" envDefault:"your_secret_key"`
}
