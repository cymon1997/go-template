package model

import (
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

/*
DB Table Details
-------------------------------------


Table: sample
[ 0] uuid                                           UUID                 null: false  primary: true   isArray: false  auto: false  col: UUID            len: -1      default: [md5(((random())]
[ 1] data                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 2] create_time                                    TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [now()]
[ 3] create_by                                      VARCHAR(20)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 20      default: []
[ 4] update_time                                    TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 5] update_by                                      VARCHAR(20)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 20      default: []


JSON Sample
-------------------------------------
{    "uuid": "OQbjvNcAaffNSVrsfTHmAxcuO",    "data": "VBjPjoPqNeGwFquXiknadXiTX",    "create_time": "2181-12-11T16:05:52.394847175+07:00",    "create_by": "eyXWlrnrRQoDIgOorYmmAywyT",    "update_time": "2269-08-19T08:59:54.815061515+07:00",    "update_by": "jMMFiVSpwIZSqFDudIaDrVRLl"}



*/

// Sample struct is a row record of the sample table in the go_template_test database
type Sample struct {
	//[ 0] uuid                                           UUID                 null: false  primary: true   isArray: false  auto: true   col: UUID            len: -1      default: [md5(((random())]
	UUID string `gorm:"primary_key;AUTO_INCREMENT;column:uuid;type:UUID;" db:"uuid"`
	//[ 1] data                                           TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Data null.String `gorm:"column:data;type:TEXT;" db:"data"`
	//[ 2] create_time                                    TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: [now()]
	CreateTime null.Time `gorm:"column:create_time;type:TIMESTAMP;" db:"create_time"`
	//[ 3] create_by                                      VARCHAR(20)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 20      default: []
	CreateBy null.String `gorm:"column:create_by;type:VARCHAR;size:20;" db:"create_by"`
	//[ 4] update_time                                    TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
	UpdateTime null.Time `gorm:"column:update_time;type:TIMESTAMP;" db:"update_time"`
	//[ 5] update_by                                      VARCHAR(20)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 20      default: []
	UpdateBy null.String `gorm:"column:update_by;type:VARCHAR;size:20;" db:"update_by"`
}

var sampleTableInfo = &TableInfo{
	Name: "sample",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "uuid",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "UUID",
			DatabaseTypePretty: "UUID",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "UUID",
			ColumnLength:       -1,
			GoFieldName:        "UUID",
			GoFieldType:        "string",
			JSONFieldName:      "uuid",
			ProtobufFieldName:  "uuid",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "data",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Data",
			GoFieldType:        "null.String",
			JSONFieldName:      "data",
			ProtobufFieldName:  "data",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "create_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "CreateTime",
			GoFieldType:        "null.Time",
			JSONFieldName:      "create_time",
			ProtobufFieldName:  "create_time",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "create_by",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       20,
			GoFieldName:        "CreateBy",
			GoFieldType:        "null.String",
			JSONFieldName:      "create_by",
			ProtobufFieldName:  "create_by",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "update_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "UpdateTime",
			GoFieldType:        "null.Time",
			JSONFieldName:      "update_time",
			ProtobufFieldName:  "update_time",
			ProtobufType:       "uint64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "update_by",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       20,
			GoFieldName:        "UpdateBy",
			GoFieldType:        "null.String",
			JSONFieldName:      "update_by",
			ProtobufFieldName:  "update_by",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *Sample) TableName() string {
	return "sample"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *Sample) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *Sample) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *Sample) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *Sample) TableInfo() *TableInfo {
	return sampleTableInfo
}
