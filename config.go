package dockertbls

import (
	"fmt"
	"strconv"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type Database struct {
	Name     string `env:"TBLS_DATABASE_NAME"`
	Schema   string `env:"TBLS_DATABASE_SCHEMA"`
	Username string
	Password string
	Host     string
	Port     string
}

func (db Database) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", db.Username, db.Password, db.Host, db.Port, db.Name, db.Schema)
}

func (db Database) DSNDefaultDBName() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable", db.Username, db.Password, db.Host, db.Port)
}

func (db Database) QuotedName() string {
	return "`" + db.Name + "`"
}

func (db Database) GetPort() uint32 {
	u, _ := strconv.ParseUint(db.Port, 10, 32)
	return uint32(u)
}

type Config struct {
	Database     Database
	MigrationDir string `env:"TBLS_MIGRATION_DIR"`
	TblsCfgFile  string `env:"TBLS_CONFIG_FILE,default=.tbls.yml"`
}

// NewConfig creates an instance of Config.
// It needs the path of the env file to be used.
func NewConfig(env string) (Config, error) {
	err := godotenv.Load(env)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
