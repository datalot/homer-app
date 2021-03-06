package migration

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/sipcapture/homer-app/migration/jsonschema"
	"github.com/sipcapture/homer-app/model"
	"github.com/sipcapture/homer-app/utils/heputils"
	"github.com/sirupsen/logrus"
)

type RollesTable struct {
	Username   string `json:"Username"`
	Attributes string `json:"Attributes"`
}

// getSession creates a new root session and panics if connection error occurs
func GetDataRootDBSession(user *string, password *string, dbname *string, host *string, port *int) (*gorm.DB, error) {

	//connectString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s ssldmode=disable password=%s", *host, *port, *user, *dbname, *password)
	connectString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", *host, *user, *dbname)

	if *port != 0 {
		connectString += fmt.Sprintf(" port=%d", *port)
	}

	if len(*password) != 0 {
		connectString += fmt.Sprintf(" password=%s", *password)
	}

	heputils.Colorize(heputils.ColorYellow, (fmt.Sprintf("\nCONNECT to DB ROOT STRING: [%s]\n", connectString)))

	db, err := gorm.Open("postgres", connectString)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	logrus.Println("----------------------------------- ")
	logrus.Println("*** Database Data Root Session created *** ")
	logrus.Println("----------------------------------- ")
	return db, nil
}

func CreateNewUser(dataRootDBSession *gorm.DB, user *string, password *string) {

	createString := fmt.Sprintf("\r\nHOMER - creating user [user=%s password=%s]", *user, *password)

	heputils.Colorize(heputils.ColorRed, createString)

	sql := fmt.Sprintf("CREATE USER %s WITH PASSWORD  '%s'", *user, *password)

	dataRootDBSession.Debug().Exec(sql)

	heputils.Colorize(heputils.ColorYellow, "\r\nDONE")

}

func DeleteNewUser(dataRootDBSession *gorm.DB, user *string) {

	createString := fmt.Sprintf("\r\nHOMER - delete user [user=%s]", *user)

	heputils.Colorize(heputils.ColorRed, createString)

	sql := fmt.Sprintf("DROP ROLE IF EXISTS %s", *user)

	dataRootDBSession.Debug().Exec(sql)

	heputils.Colorize(heputils.ColorYellow, "\r\nDONE")
}

func CreateHomerDB(dataRootDBSession *gorm.DB, dbname *string, user *string) {

	createString := fmt.Sprintf("\r\nHOMER - create db [%s] with [name=%s]", *dbname, *user)

	heputils.Colorize(heputils.ColorRed, createString)

	sql := fmt.Sprintf("CREATE DATABASE %s OWNER %s", *dbname, *user)

	dataRootDBSession.Debug().Exec(sql)

	heputils.Colorize(heputils.ColorYellow, "\r\nDONE")
}

func CreateHomerRole(dataRootDBSession *gorm.DB, user *string, homerDBconfig *string, homerDBdata *string) {

	createString := fmt.Sprintf("\r\nHOMER - creating role for user [user=%s dbconfig=%s, dbdata=%s]", *user, *homerDBconfig, *homerDBdata)

	heputils.Colorize(heputils.ColorRed, createString)

	sql := fmt.Sprintf("GRANT ALL PRIVILEGES ON DATABASE %s to %s;", *homerDBconfig, *user)
	dataRootDBSession.Debug().Exec(sql)

	sql = fmt.Sprintf("GRANT ALL PRIVILEGES ON DATABASE %s to %s;", *homerDBdata, *user)
	dataRootDBSession.Debug().Exec(sql)

	sql = fmt.Sprintf("ALTER DATABASE %s OWNER TO %s;", *homerDBconfig, *user)
	dataRootDBSession.Debug().Exec(sql)

	sql = fmt.Sprintf("ALTER DATABASE %s OWNER TO %s;", *homerDBdata, *user)
	dataRootDBSession.Debug().Exec(sql)

	heputils.Colorize(heputils.ColorYellow, "\r\nDONE")
}

func RevokeHomerRole(dataRootDBSession *gorm.DB, user *string, homerDBconfig *string, homerDBdata *string) {

	createString := fmt.Sprintf("\r\nHOMER - revoke role for user [user=%s dbconfig=%s, dbdata=%s]", *user, *homerDBconfig, *homerDBdata)

	heputils.Colorize(heputils.ColorRed, createString)

	sql := fmt.Sprintf("REVOKE ALL PRIVILEGES ON DATABASE %s FROM %s;", *homerDBconfig, *user)
	dataRootDBSession.Debug().Exec(sql)

	sql = fmt.Sprintf("REVOKE ALL PRIVILEGES ON DATABASE %s FROM %s;", *homerDBdata, *user)
	dataRootDBSession.Debug().Exec(sql)

	sql = fmt.Sprintf("ALTER DATABASE %s OWNER TO postgres;", *homerDBconfig)
	dataRootDBSession.Debug().Exec(sql)

	sql = fmt.Sprintf("ALTER DATABASE %s OWNER TO postgres;", *homerDBdata)
	dataRootDBSession.Debug().Exec(sql)

	heputils.Colorize(heputils.ColorYellow, "\r\nDONE")
}

func ShowUsers(dataRootDBSession *gorm.DB) {

	createString := fmt.Sprintf("\r\nHOMER - show users")

	heputils.Colorize(heputils.ColorRed, createString)

	sql2 := "SELECT u.usename AS \"Username\", CASE WHEN u.usesuper AND u.usecreatedb THEN CAST('superuser, create database' AS pg_catalog.text)" +
		"WHEN u.usesuper THEN CAST('superuser' AS pg_catalog.text) WHEN u.usecreatedb THEN CAST('create database' AS  pg_catalog.text)" +
		"ELSE CAST('' AS pg_catalog.text) END AS \"Attributes\" FROM pg_catalog.pg_user u ORDER BY 1;"

	fmt.Println("\tRole name\t|\tAttributes")
	fmt.Println("------------------------------------------------")

	var Username, Attributes string
	rows, _ := dataRootDBSession.Raw(sql2).Rows() // (*sql.Rows, error)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Username, &Attributes)
		if err == nil {
			fmt.Println(fmt.Sprintf("\t%s\t|\t%s\t", Username, Attributes))
		}

	}
	heputils.Colorize(heputils.ColorYellow, "\r\nDONE")
}

func CreateHomerConfigTables(configDBSession *gorm.DB, homerDBconfig string) {

	createString := fmt.Sprintf("\r\nHOMER - creating table for the config DB [dbname=%s]", homerDBconfig)

	heputils.Colorize(heputils.ColorRed, createString)

	configDBSession.AutoMigrate(&model.TableAlias{},
		&model.TableGlobalSettings{},
		&model.TableMappingSchema{},
		&model.TableUserSettings{},
		&model.TableHepsubSchema{},
		&model.TableUser{},
		&model.TableAgentLocationSession{},
		&model.TableVersions{})

	heputils.Colorize(heputils.ColorYellow, "\r\nDONE")
}

func checkHomerConfigTables(configDBSession *gorm.DB) map[string]bool {

	data := []model.TableVersions{}
	createTables := map[string]bool{}
	if err := configDBSession.Debug().
		Table("versions").
		Find(&data).Error; err != nil {
		return createTables
	}

	for _, row := range data {
		createTables[row.NameTable] = false
		if jsonschema.TableVersion[row.NameTable] < row.VersionTable {
			fmt.Println("Found older version:", row.NameTable)
			createTables[row.NameTable] = true
		}
	}

	return createTables
}

func PopulateHomerConfigTables(configDBSession *gorm.DB, homerDBconfig string, force bool) {

	createString := fmt.Sprintf("\r\nHOMER - filling tables for the config DB [dbname=%s]", homerDBconfig)
	usersData := []model.TableUser{
		model.TableUser{
			Id:         1,
			UserName:   "admin",
			PartId:     10,
			Email:      "root@localhost",
			FirstName:  "Homer",
			LastName:   "Admin",
			Department: "Develop",
			UserGroup:  "admin",
			Hash:       string(jsonschema.DefaultAdminPassword),
			GUID:       uuid.NewV4().String(),
		},
		model.TableUser{
			Id:         2,
			UserName:   "support",
			PartId:     10,
			Email:      "support@localhost",
			FirstName:  "Homer",
			LastName:   "Support",
			Department: "Develop",
			UserGroup:  "admin",
			Hash:       string(jsonschema.DefaultSupportPassword),
			GUID:       uuid.NewV4().String(),
		},
	}

	globalSettingData := []model.TableGlobalSettings{
		model.TableGlobalSettings{
			GUID:     uuid.NewV4().String(),
			PartId:   1,
			Category: "search",
			Param:    "lokiserver",
			Data:     jsonschema.LokiConfig,
		},
		model.TableGlobalSettings{
			GUID:     uuid.NewV4().String(),
			PartId:   1,
			Category: "search",
			Param:    "promserver",
			Data:     jsonschema.PrometheusConfig,
		},
		model.TableGlobalSettings{
			GUID:     uuid.NewV4().String(),
			PartId:   1,
			Category: "search",
			Param:    "grafana",
			Data:     jsonschema.GrafanaConfig,
		},
	}

	agentLocationSession := []model.TableAgentLocationSession{
		model.TableAgentLocationSession{
			GUID:       uuid.NewV4().String(),
			Gid:        10,
			Host:       "127.0.0.1",
			Port:       8080,
			Protocol:   "rtp",
			Path:       "/api/search",
			Node:       "rtpnode01",
			Type:       "cdr",
			CreateDate: time.Now(),
			ExpireDate: time.Now(),
			Active:     1,
		},
	}

	hepsubSchema := []model.TableHepsubSchema{
		model.TableHepsubSchema{
			GUID:       uuid.NewV4().String(),
			Profile:    "default",
			Hepid:      1,
			HepAlias:   "SIP",
			Version:    1,
			Mapping:    jsonschema.CorrelationMappingdefault,
			CreateDate: time.Now(),
		},
		model.TableHepsubSchema{
			GUID:       uuid.NewV4().String(),
			Profile:    "call",
			Hepid:      1,
			HepAlias:   "SIP",
			Version:    1,
			Mapping:    jsonschema.CorrelationMappingdefault,
			CreateDate: time.Now(),
		},
		model.TableHepsubSchema{
			GUID:       uuid.NewV4().String(),
			Profile:    "registration",
			Hepid:      1,
			HepAlias:   "SIP",
			Version:    1,
			Mapping:    jsonschema.CorrelationMappingdefault,
			CreateDate: time.Now(),
		},
	}

	dashboardUsers := []model.TableUserSettings{
		model.TableUserSettings{
			GUID:       uuid.NewV4().String(),
			UserName:   "admin",
			Param:      "home",
			PartId:     10,
			Category:   "dashboard",
			Data:       jsonschema.DashboardHome,
			CreateDate: time.Now(),
		},
		model.TableUserSettings{
			GUID:       uuid.NewV4().String(),
			UserName:   "support",
			Param:      "home",
			PartId:     10,
			Category:   "dashboard",
			Data:       jsonschema.DashboardHome,
			CreateDate: time.Now(),
		},
	}

	tableVersions := []model.TableVersions{
		model.TableVersions{
			NameTable:    "versions",
			VersionTable: jsonschema.TableVersion["versions"],
		},
		model.TableVersions{
			NameTable:    "agent_location_session",
			VersionTable: jsonschema.TableVersion["agent_location_session"],
		},
		model.TableVersions{
			NameTable:    "alias",
			VersionTable: jsonschema.TableVersion["alias"],
		},
		model.TableVersions{
			NameTable:    "global_settings",
			VersionTable: jsonschema.TableVersion["global_settings"],
		},
		model.TableVersions{
			NameTable:    "hepsub_mapping_schema",
			VersionTable: jsonschema.TableVersion["hepsub_mapping_schema"],
		},
		model.TableVersions{
			NameTable:    "mapping_schema",
			VersionTable: jsonschema.TableVersion["mapping_schema"],
		},
		model.TableVersions{
			NameTable:    "users",
			VersionTable: jsonschema.TableVersion["users"],
		},
		model.TableVersions{
			NameTable:    "user_settings",
			VersionTable: jsonschema.TableVersion["user_settings"],
		},
	}

	mappingSchema := []model.TableMappingSchema{
		model.TableMappingSchema{
			GUID:               uuid.NewV4().String(),
			Profile:            "default",
			Hepid:              1,
			HepAlias:           "SIP",
			PartId:             10,
			Version:            1,
			Retention:          10,
			PartitionStep:      10,
			CreateIndex:        jsonschema.EmptyJson,
			CreateTable:        "CREATE TABLE test(id integer, data text);",
			FieldsMapping:      jsonschema.FieldsMapping1default,
			CorrelationMapping: jsonschema.CorrelationMapping1default,
			MappingSettings:    jsonschema.EmptyJson,
			SchemaMapping:      jsonschema.EmptyJson,
			SchemaSettings:     jsonschema.EmptyJson,
			CreateDate:         time.Now(),
		},
		model.TableMappingSchema{
			GUID:               uuid.NewV4().String(),
			Profile:            "call",
			Hepid:              1,
			HepAlias:           "SIP",
			PartId:             10,
			Version:            1,
			Retention:          10,
			PartitionStep:      10,
			CreateIndex:        jsonschema.EmptyJson,
			CreateTable:        "CREATE TABLE test(id integer, data text);",
			FieldsMapping:      jsonschema.FieldsMapping1default,
			CorrelationMapping: jsonschema.CorrelationMapping1default,
			MappingSettings:    jsonschema.EmptyJson,
			SchemaMapping:      jsonschema.EmptyJson,
			SchemaSettings:     jsonschema.EmptyJson,
			CreateDate:         time.Now(),
		},
		model.TableMappingSchema{
			GUID:               uuid.NewV4().String(),
			Profile:            "registration",
			Hepid:              1,
			HepAlias:           "SIP",
			PartId:             10,
			Version:            1,
			Retention:          10,
			PartitionStep:      10,
			CreateIndex:        jsonschema.EmptyJson,
			CreateTable:        "CREATE TABLE test(id integer, data text);",
			FieldsMapping:      jsonschema.FieldsMapping1default,
			CorrelationMapping: jsonschema.CorrelationMapping1default,
			MappingSettings:    jsonschema.EmptyJson,
			SchemaMapping:      jsonschema.EmptyJson,
			SchemaSettings:     jsonschema.EmptyJson,
			CreateDate:         time.Now(),
		},
		model.TableMappingSchema{
			GUID:               uuid.NewV4().String(),
			Profile:            "default",
			Hepid:              100,
			HepAlias:           "LOG",
			PartId:             10,
			Version:            1,
			Retention:          10,
			PartitionStep:      10,
			CreateIndex:        jsonschema.EmptyJson,
			CreateTable:        "CREATE TABLE test(id integer, data text);",
			FieldsMapping:      jsonschema.FieldsMapping100default,
			CorrelationMapping: jsonschema.CorrelationMapping100default,
			MappingSettings:    jsonschema.EmptyJson,
			SchemaMapping:      jsonschema.EmptyJson,
			SchemaSettings:     jsonschema.EmptyJson,
			CreateDate:         time.Now(),
		},
		model.TableMappingSchema{
			GUID:               uuid.NewV4().String(),
			Profile:            "default",
			Hepid:              34,
			HepAlias:           "RTP-FULL-REPORT",
			PartId:             10,
			Version:            1,
			Retention:          10,
			PartitionStep:      10,
			CreateIndex:        jsonschema.EmptyJson,
			CreateTable:        "CREATE TABLE test(id integer, data text);",
			FieldsMapping:      jsonschema.FieldsMapping34default,
			CorrelationMapping: jsonschema.CorrelationMapping34default,
			MappingSettings:    jsonschema.EmptyJson,
			SchemaMapping:      jsonschema.EmptyJson,
			SchemaSettings:     jsonschema.EmptyJson,
			CreateDate:         time.Now(),
		},
		model.TableMappingSchema{
			GUID:               uuid.NewV4().String(),
			Profile:            "default",
			Hepid:              1000,
			HepAlias:           "JANUS",
			PartId:             10,
			Version:            1,
			Retention:          10,
			PartitionStep:      10,
			CreateIndex:        jsonschema.EmptyJson,
			CreateTable:        "CREATE TABLE test(id integer, data text);",
			FieldsMapping:      jsonschema.FieldsMapping1000default,
			CorrelationMapping: jsonschema.CorrelationMapping1000default,
			MappingSettings:    jsonschema.EmptyJson,
			SchemaMapping:      jsonschema.EmptyJson,
			SchemaSettings:     jsonschema.EmptyJson,
			CreateDate:         time.Now(),
		},
		model.TableMappingSchema{
			GUID:               uuid.NewV4().String(),
			Profile:            "default",
			Hepid:              2000,
			HepAlias:           "LOKI",
			PartId:             10,
			Version:            1,
			Retention:          10,
			PartitionStep:      10,
			CreateIndex:        jsonschema.EmptyJson,
			CreateTable:        "CREATE TABLE test(id integer, data text);",
			CorrelationMapping: jsonschema.EmptyJson,
			FieldsMapping:      jsonschema.FieldsMapping2000loki,
			MappingSettings:    jsonschema.EmptyJson,
			SchemaMapping:      jsonschema.EmptyJson,
			SchemaSettings:     jsonschema.EmptyJson,
			CreateDate:         time.Now(),
		},
	}

	createTables := checkHomerConfigTables(configDBSession)

	/*********************************************/
	heputils.Colorize(heputils.ColorRed, createString)

	if val, ok := createTables["users"]; !ok || ok && val || force {
		/* User data */
		heputils.Colorize(heputils.ColorRed, "reinstalling users")
		configDBSession.Exec("TRUNCATE TABLE users")
		for _, el := range usersData {
			configDBSession.Save(&el)
		}
	}

	if val, ok := createTables["global_settings"]; !ok || ok && val || force {
		/* globalSettingData data */
		heputils.Colorize(heputils.ColorRed, "reinstalling global_settings")
		configDBSession.Exec("TRUNCATE TABLE global_settings")
		for _, el := range globalSettingData {
			configDBSession.Save(&el)
		}
	}

	if val, ok := createTables["agent_location_session"]; !ok || ok && val || force {
		/* agentLocationSession data */
		heputils.Colorize(heputils.ColorRed, "reinstalling agent_location_session")
		configDBSession.Exec("TRUNCATE TABLE agent_location_session")
		for _, el := range agentLocationSession {
			configDBSession.Save(&el)
		}
	}

	if val, ok := createTables["hepsub_mapping_schema"]; !ok || ok && val || force {
		/* hepsubSchema data */
		heputils.Colorize(heputils.ColorRed, "reinstalling hepsub_mapping_schema")
		configDBSession.Exec("TRUNCATE TABLE hepsub_mapping_schema")
		for _, el := range hepsubSchema {
			configDBSession.Save(&el)
		}
	}

	if val, ok := createTables["user_settings"]; !ok || ok && val || force {
		/* dashboardUsers data */
		heputils.Colorize(heputils.ColorRed, "reinstalling user_settings")
		configDBSession.Exec("TRUNCATE TABLE user_settings")
		for _, el := range dashboardUsers {
			configDBSession.Save(&el)
		}
	}

	if val, ok := createTables["mapping_schema"]; !ok || ok && val || force {
		/* mappingSchema data */
		heputils.Colorize(heputils.ColorRed, "reinstalling mapping_schema")
		configDBSession.Exec("TRUNCATE TABLE mapping_schema")
		for _, el := range mappingSchema {
			configDBSession.Save(&el)
		}
	}

	/* tableVersions data */
	heputils.Colorize(heputils.ColorRed, "reinstalling versions")
	configDBSession.Exec("TRUNCATE TABLE versions")
	for _, el := range tableVersions {
		configDBSession.Save(&el)
	}

	heputils.Colorize(heputils.ColorYellow, "\r\nDONE")
}
