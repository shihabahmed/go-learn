package utils

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DbConfig DbConfig `yaml:"database"`
	LogPath  string   `yaml:"log_path"`
}

// Parse unmarshals the provided byte array into a Config struct from YAML
func (c *Config) Parse(data []byte) error {
	c.LogPath = "."
	return yaml.Unmarshal(data, c)
}

// FromFile creates an instance of the configuration
func (c *Config) FromFile(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal().Err(err).Msgf("Could not read config file %s", path)
		return fmt.Errorf("could not read file %s: %w", path, err)
	}

	if err = c.Parse(b); err != nil {
		log.Fatal().Err(err).Msgf("Unable to parse config file %s", path)
		return fmt.Errorf("could not parse file %s: %w", path, err)
	}

	return err
}

// DbConfig is the struct that holds the configuration details for connecting to the DB
type DbConfig struct {
	Host     string   `yaml:"host"`
	Port     int32    `yaml:"port"`
	User     string   `yaml:"user"`
	Pass     string   `yaml:"pass"`
	Name     string   `yaml:"name"`
	MaxConns int      `yaml:"max_connections"`
	Conn     *gorm.DB `yaml:"-"`
}

// getConnectString returns a connection string for the database driver
func (c *DbConfig) getConnectString() string {
	return os.Getenv("CONN_STRING")

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	c.Host, c.Port, c.User, c.Pass, c.Name)
	// return psqlInfo
}

// Connect connects to the DB and returns the DB connection object
func (c *DbConfig) Connect() (db *gorm.DB, err error) {
	if c.Conn == nil {
		// c.Conn, err = gorm.Open(postgres.New(postgres.Config{
		// 	DSN:                  c.getConnectString(),
		// 	PreferSimpleProtocol: true,
		// }), &gorm.Config{})
		c.Conn, err = gorm.Open(postgres.Open(c.getConnectString()))

		if err != nil {
			log.Panic().AnErr("error", err).Msg("Unable to connect to DB")
		}
	}

	return c.Conn, err
}
