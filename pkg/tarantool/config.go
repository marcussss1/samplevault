package tarantool

// TODO без дефолта.
type Config struct {
	DSN      string `env:"TARANTOOL_DSN,required" envDefault:"tarantool:3301"`
	User     string `env:"TARANTOOL_USER,required" envDefault:"tarantool"`
	Password string `env:"TARANTOOL_PASSWORD,required" envDefault:"tarantool"`
}
