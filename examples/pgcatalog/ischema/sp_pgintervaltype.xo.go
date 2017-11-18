// Package ischema contains the types for schema 'information_schema'.
package ischema

import "gitlab.com/tesgo/core/pkg/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// PgIntervalType calls the stored procedure 'information_schema._pg_interval_type(oid, integer) text' on db.
func PgIntervalType(db XODB, v0 pgtypes.Oid, v1 int) (string, error) {
	var err error

	// sql query
	const sqlstr = `SELECT information_schema._pg_interval_type($1, $2)`

	// run query
	var ret string
	XOLog(sqlstr, v0, v1)
	err = db.QueryRow(sqlstr, v0, v1).Scan(&ret)
	if err != nil {
		return "", err
	}

	return ret, nil
}
