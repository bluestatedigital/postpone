package common

// Base set of options that are common to all programs in this repo.  Should be
// embedded in a struct specific to the program.
type Options struct {
    Debug   bool   `env:"DEBUG"    long:"debug"    description:"enable debug"`
    LogFile string `env:"LOG_FILE" long:"log-file" description:"path to JSON log file"`
    
    RedisHost string `env:"REDIS_HOST" long:"redis-host" description:"redis hostname" default:"localhost"`
    RedisPort int    `env:"REDIS_PORT" long:"redis-port" description:"redis port"     default:"6379"`
    RedisDb   int    `env:"REDIS_DB"   long:"redis-db"   description:"redis database" default:"0"`
}
