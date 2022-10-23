package model

import (
	"gorm.io/gorm"
)

/*
DB Table Details
-------------------------------------


Table: migration_schema
[ 0] version                                        INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
[ 1] dirty                                          BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []


JSON Sample
-------------------------------------
{    "version": 16,    "dirty": false}



*/

// MigrationSchema struct is a row record of the migration_schema table in the go_template_test database
type MigrationSchema struct {
	//[ 0] version                                        INT8                 null: false  primary: true   isArray: false  auto: false  col: INT8            len: -1      default: []
	Version int64 `gorm:"primary_key;column:version;type:INT8;" db:"version"`
	//[ 1] dirty                                          BOOL                 null: false  primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	Dirty bool `gorm:"column:dirty;type:BOOL;" db:"dirty"`
}

var migration_schemaTableInfo = &TableInfo{
	Name: "migration_schema",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "version",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "Version",
			GoFieldType:        "int64",
			JSONFieldName:      "version",
			ProtobufFieldName:  "version",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "dirty",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "Dirty",
			GoFieldType:        "bool",
			JSONFieldName:      "dirty",
			ProtobufFieldName:  "dirty",
			ProtobufType:       "bool",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *MigrationSchema) TableName() string {
	return "migration_schema"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *MigrationSchema) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *MigrationSchema) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *MigrationSchema) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *MigrationSchema) TableInfo() *TableInfo {
	return migration_schemaTableInfo
}
