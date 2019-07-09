package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config 配置参数
type Config struct {
	Web     Web
	Gorm    Gorm
	MySQL   MySQL
	Sqlite3 Sqlite3
}

// 站点配置参数
type Web struct {
	Domain       string
	StaticPath   string
	Port         int
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

// Gorm gorm配置参数
type Gorm struct {
	Debug        bool
	DBType       string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	TablePrefix  string
	DSN          string
}

// MySQL mysql配置参数
type MySQL struct {
	Host       string
	Port       int
	User       string
	Password   string
	DBName     string
	Parameters string
}

// MySQL 数据库连接串
func (a MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}

// Sqlite3 配置参数
type Sqlite3 struct {
	Path string
}

// Sqlite3 数据库连接串
func (a Sqlite3) DSN() string {
	return a.Path
}

// 加载配置
func LoadConfig(fpath string) (c *Config, err error) {
	v := viper.New()
	v.SetConfigFile(fpath)
	v.SetConfigType("yaml")
	if err1 := v.ReadInConfig(); err1 != nil {
		err = err1
		return
	}
	c = &Config{}
	c.Web.StaticPath = v.GetString("web.static_path")
	c.Web.Domain = v.GetString("web.domain")
	c.Web.Port = v.GetInt("web.port")
	c.Web.ReadTimeout = v.GetInt("web.read_timeout")
	c.Web.WriteTimeout = v.GetInt("web.write_timeout")
	c.Web.IdleTimeout = v.GetInt("web.idle_timeout")
	c.MySQL.Host = v.GetString("mysql.host")
	c.MySQL.Port = v.GetInt("mysql.port")
	c.MySQL.User = v.GetString("mysql.user")
	c.MySQL.Password = v.GetString("mysql.password")
	c.MySQL.DBName = v.GetString("mysql.db_name")
	c.MySQL.Parameters = v.GetString("mysql.parameters")
	c.Sqlite3.Path = v.GetString("sqlite3.path")
	c.Gorm.Debug = v.GetBool("gorm.debug")
	c.Gorm.DBType = v.GetString("gorm.db_type")
	c.Gorm.MaxLifetime = v.GetInt("gorm.max_lifetime")
	c.Gorm.MaxOpenConns = v.GetInt("gorm.max_open_conns")
	c.Gorm.MaxIdleConns = v.GetInt("gorm.max_idle_conns")
	c.Gorm.TablePrefix = v.GetString("gorm.table_prefix")
	return
}
