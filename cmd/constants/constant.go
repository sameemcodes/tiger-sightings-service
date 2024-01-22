package constants

import (
	"context"
	"flag"
	durable "tigerhall-kittens/cmd/durables"
	"time"
)

const XPlatformVersion string = "x-platform-version"
const XPlatform string = "x-platform"
const SessionUserKey string = "X-USER-ID"

type JsonType map[string]interface{}

var (
	Env                = flag.String("env", "local-dev", "")
	DbUsername         = flag.String("DB_USERNAME", "root", "")
	DbPassword         = flag.String("DB_PASSWORD", "root@0830", "")
	RedisPassword      = flag.String("REDIS_PASSWORD", "password", "")
	Ctx                = context.Background()
	DevConfigFile      = "src/config/dev.yml"
	ProdConfigFile     = "src/config/prod.yml"
	ReleaseConfigFile  = "src/config/release.yml"
	StageConfigFile    = "src/config/stage.yml"
	LocalDevConfigFile = "src/config/local.yml"
	PreferenceIndex    = "user_preference"
	StandardPrefMap    map[string]interface{}
	Logger             *durable.Logger
	ProductConfigList  []string
)

// mysql_db constants
const WhereUserId string = "user_id = ?"
const WhereTigerId string = "tiger_id = ?"
const WhereSightingId string = "sighting_id = ?"
const WhereEmail string = "email = ?"

// context key MYSQLDB

type MysqlDbKeyType string

const MysqlDbKey MysqlDbKeyType = "mysqlDb"

// cookie Expiry
const CookieExpiryTime int = 300 // 5 minute

// Token Expiry
const TokenExpiryTime time.Duration = 300 // 5 minute
