package config

import (
    "fmt"
    "github.com/spf13/viper"
)

type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    JWT      JWTConfig     `mapstructure:"jwt"`
}

type ServerConfig struct {
    Port int    `mapstructure:"port"`
    Mode string `mapstructure:"mode"`
}

type DatabaseConfig struct {
    Driver    string `mapstructure:"driver"`
    Host      string `mapstructure:"host"`
    Port      int    `mapstructure:"port"`
    Username  string `mapstructure:"username"`
    Password  string `mapstructure:"password"`
    DBName    string `mapstructure:"dbname"`
    Charset   string `mapstructure:"charset"`
    ParseTime bool   `mapstructure:"parse_time"`
    Loc       string `mapstructure:"loc"`
}

type JWTConfig struct {
    Secret string `mapstructure:"secret"`
    Expire int    `mapstructure:"expire"`
}

func (d *DatabaseConfig) DSN() string {
    return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
        d.Username, d.Password, d.Host, d.Port, d.DBName, d.Charset, d.ParseTime, d.Loc)
}

func LoadConfig() (*Config, error) {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("/Users/systom/go/src/robot-go/configs")

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, err
    }

    return &config, nil
}