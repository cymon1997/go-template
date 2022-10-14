package db

const (
	MySQL    Driver = "mysql"
	Postgres Driver = "postgres"
)

type Driver string

type Config struct {
	Driver          Driver
	DSN             string
	Timezone        string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxIdleTime int
	ConnMaxLifeTime int
}
