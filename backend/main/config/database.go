package config

import (
	"main/main/system/types"
)

/**
 * A database config for the app.
 *
 * @param	UseDatabase		True if this application use a database, false if not.
 * @param	Driver			Database driver name. If you use other databases like MongoDB,
 *							you must import the driver in system/database/database.go.
 * @param	Username		Username used to connect to the database.
 * @param	Password		Password used to connect to the database.
 * @param	Host			The database host.
 * @param	Port			The port of database server.const.
 * @param	Database		The database name used for this app.
 */
var DBConfig types.DatabaseConfig = types.DatabaseConfig{
	UseDatabase: false,
	Driver:      "mysql",
	Username:    "root",
	Password:    "",
	Host:        "127.0.0.1",
	Port:        3306,
	Database:    "test",
	ParseTime:   true,
}
