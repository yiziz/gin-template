package db

import (
	"github.com/yiziz/gin-template/path"
	"github.com/yiziz/gin-template/services/yml"
)

func configFilename() string {
	return path.DBConfigFilename()
}

func dbAdapter(s string) (adapter string) {
	if s == "mysql" {
		adapter = "mysql"
	}
	return
}

func dbParameters(adapter string, m map[interface{}]interface{}) (params string) {
	dbName := m["database"].(string)
	username := m["username"].(string)
	password, pwOK := m["password"].(string)
	if !pwOK {
		password = ""
	}
	if adapter == "mysql" {
		params = username
		if password != "" {
			params += ":" + password
		}
		params += "@/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	}
	return
}

func dbConfigFromYML(env string, name string) (adapter, params string) {
	m := yml.ConfigYML(name)
	dbConfig := m[env].(map[interface{}]interface{})
	adapter = dbAdapter(dbConfig["adapter"].(string))
	params = dbParameters(adapter, dbConfig)
	return
}

// Config returns the database config strings necessary to open a DB conn
func Config(env string, path ...string) (db, parameters string) {
	if env == "" {
		// test db should be wiped after tests so there shouldn't be
		// data loss worries over accidently using the test db
		env = "test"
	}
	var configName string
	if len(path) == 0 {
		configName = configFilename()
	}
	return dbConfigFromYML(env, configName)
}
