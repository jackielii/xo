// Package ischema contains the types for schema 'information_schema'.
package ischema

import "gitlab.com/tesgo/core/pkg/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// Sequence represents a row from 'information_schema.sequences'.
type Sequence struct {
	SequenceCatalog       pgtypes.SQLIdentifier  `json:"sequence_catalog"`        // sequence_catalog
	SequenceSchema        pgtypes.SQLIdentifier  `json:"sequence_schema"`         // sequence_schema
	SequenceName          pgtypes.SQLIdentifier  `json:"sequence_name"`           // sequence_name
	DataType              pgtypes.CharacterData  `json:"data_type"`               // data_type
	NumericPrecision      pgtypes.CardinalNumber `json:"numeric_precision"`       // numeric_precision
	NumericPrecisionRadix pgtypes.CardinalNumber `json:"numeric_precision_radix"` // numeric_precision_radix
	NumericScale          pgtypes.CardinalNumber `json:"numeric_scale"`           // numeric_scale
	StartValue            pgtypes.CharacterData  `json:"start_value"`             // start_value
	MinimumValue          pgtypes.CharacterData  `json:"minimum_value"`           // minimum_value
	MaximumValue          pgtypes.CharacterData  `json:"maximum_value"`           // maximum_value
	Increment             pgtypes.CharacterData  `json:"increment"`               // increment
	CycleOption           pgtypes.YesOrNo        `json:"cycle_option"`            // cycle_option
}
