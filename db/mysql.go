package db

import (
	"fmt"
)

func getMysqlConnectionString(params InitDBParams) string {
	switch params.Protocol {
	case "tcp", "tcp6":
		return fmt.Sprintf("%s:%s@%s([%s]:%d)/%s?parseTime=true&timeout=10s",
			params.User, params.Password, params.Protocol, params.Host, params.Port, params.Db)
	case "unix":
		return fmt.Sprintf("%s:%s@%s(%s)/%s?parseTime=true&timeout=10s",
			params.User, params.Password, params.Protocol, params.Path, params.Db)
	}

	panic(fmt.Errorf("This protocol is not supported: %v", params.Protocol))
	return ""
}
