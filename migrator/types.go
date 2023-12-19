package migrator

import (
	"database/sql"
	"fmt"
	"strings"
)

type Config struct {
	DB                    *sql.DB
	PriorityLpad          int
	VersionLpad           int
	MigrationFilesPath    string
	DbObjectPath          string
	CreateObjectsFilename string
	DropObjectsFilename   string
}

type DbObject struct {
	Name     string
	Priority int
	Versions []int
}

type DbObjectVersion struct {
	DbObject *DbObject
	Version  int
}

func (o *DbObject) FileNameForStep(step int, config *Config) string {
	version := 0
	for _, vs := range o.Versions {
		if version < vs && step <= vs {
			version = vs
		}
	}
	if version == 0 {
		return ""
	}
	return o.FileName(version, config)

}

func (o *DbObject) FileName(version int, config *Config) string {
	return config.DbObjectPath + "/" + LPAD(o.Priority, config.PriorityLpad) + "_" + o.Name + "/" + LPAD(o.Priority, config.PriorityLpad) + "_" + o.Name + ".sql"
}

func LPAD(number int, length int) string {
	numStr := fmt.Sprintf("%d", number)
	if len(numStr) >= length {
		return numStr
	}
	return strings.Repeat("0", length-len(numStr)) + numStr
}
