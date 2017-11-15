// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/jackielii/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// ReferentialConstraint represents a row from 'information_schema.referential_constraints'.
type ReferentialConstraint struct {
	ConstraintCatalog       pgtypes.SQLIdentifier `json:"constraint_catalog"`        // constraint_catalog
	ConstraintSchema        pgtypes.SQLIdentifier `json:"constraint_schema"`         // constraint_schema
	ConstraintName          pgtypes.SQLIdentifier `json:"constraint_name"`           // constraint_name
	UniqueConstraintCatalog pgtypes.SQLIdentifier `json:"unique_constraint_catalog"` // unique_constraint_catalog
	UniqueConstraintSchema  pgtypes.SQLIdentifier `json:"unique_constraint_schema"`  // unique_constraint_schema
	UniqueConstraintName    pgtypes.SQLIdentifier `json:"unique_constraint_name"`    // unique_constraint_name
	MatchOption             pgtypes.CharacterData `json:"match_option"`              // match_option
	UpdateRule              pgtypes.CharacterData `json:"update_rule"`               // update_rule
	DeleteRule              pgtypes.CharacterData `json:"delete_rule"`               // delete_rule
}
