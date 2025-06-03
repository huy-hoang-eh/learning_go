package sqlite

import "database/sql"

func Init() {
	lite := sqlite{}
	lite.init()
}

func Execute(query string, args ...interface{}) (sql.Result, error) {
	lite := sqlite{}
	return lite.execute(query, args...)
}
