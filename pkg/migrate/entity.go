package migrate

const (
	SchemaTableHeader      = "x-migrations-table"
	DefaultSchemaTableName = "migration_schema"
)

type Config struct {
	// DSN database connection
	DSN string
	// FilePath location of the migration files
	FilePath string
	// SchemaTableName table name used for maintain migration versions
	SchemaTableName string
}
