package migrate

import (
	"net/url"
)

func addSchemaTableName(dsn string, tableName string) (string, error) {
	raw, err := url.Parse(dsn)
	if err != nil {
		return dsn, err
	}
	q := raw.Query()
	q.Set(SchemaTableHeader, tableName)
	raw.RawQuery = q.Encode()
	return raw.String(), nil
}
