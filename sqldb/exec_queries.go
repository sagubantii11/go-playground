package sqldb

import "database/sql"

func ExecuteSelect(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func ExecuteInserts(query string, args ...interface{}) (int64, error) {
	inserQuery, err := DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	result, err := inserQuery.Exec(args...)
	defer inserQuery.Close()
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}