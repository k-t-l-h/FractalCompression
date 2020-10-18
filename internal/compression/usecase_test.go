package compression

import (
	"FractalCompression/internal"
	"testing"
)

func TestTable_Compress(t1 *testing.T) {
	type fields struct {
		K              uint64
		Name           string
		Database       internal.IDatabase
		Columns        []*Column
		Compressible   []int
		Incompressible []int
		Domens         []int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{},
	}
	for _, _ = range tests {

	}
}

func TestTable_getCompressible(t1 *testing.T) {
	type fields struct {
		K              uint64
		Name           string
		Database       internal.IDatabase
		Columns        []*Column
		Compressible   []int
		Incompressible []int
		Domens         []int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{},
	}
	for _, _ = range tests {

	}
}

func TestTable_getConstrains(t1 *testing.T) {
	type fields struct {
		K              uint64
		Name           string
		Database       internal.IDatabase
		Columns        []*Column
		Compressible   []int
		Incompressible []int
		Domens         []int
	}
	tests := []struct {
		name   string
		fields fields
		result bool
	}{
		{},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Table{
				K:              tt.fields.K,
				Name:           tt.fields.Name,
				Database:       tt.fields.Database,
				Columns:        tt.fields.Columns,
				Compressible:   tt.fields.Compressible,
				Incompressible: tt.fields.Incompressible,
				Domens:         tt.fields.Domens,
			}
			if err := t.getConstrains(); (err != nil) != tt.result {
				t1.Errorf("getConstrains() error = %v, wantErr %v", err, tt.result)
			}
		})
	}
}

func TestTable_getDomens(t1 *testing.T) {
	type fields struct {
		K              uint64
		Name           string
		Database       internal.IDatabase
		Columns        []*Column
		Compressible   []int
		Incompressible []int
		Domens         []int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{},
	}
	for _, _ = range tests {

	}
}

func TestTable_getMeta(t1 *testing.T) {
	type fields struct {
		K              uint64
		Name           string
		Database       internal.IDatabase
		Columns        []*Column
		Compressible   []int
		Incompressible []int
		Domens         []int
	}
	tests := []struct {
		name   string
		fields fields
		result bool
	}{
		{},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Table{
				K:              tt.fields.K,
				Name:           tt.fields.Name,
				Database:       tt.fields.Database,
				Columns:        tt.fields.Columns,
				Compressible:   tt.fields.Compressible,
				Incompressible: tt.fields.Incompressible,
				Domens:         tt.fields.Domens,
			}
			if err := t.getMeta(); (err != nil) != tt.result {
				t1.Errorf("getMeta() error = %v, wantErr %v", err, tt.result)
			}
		})
	}
}

func TestTable_getPriorities(t1 *testing.T) {
	type fields struct {
		K              uint64
		Name           string
		Database       internal.IDatabase
		Columns        []*Column
		Compressible   []int
		Incompressible []int
		Domens         []int
	}
	tests := []struct {
		name   string
		fields fields
		result bool
	}{
		{},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Table{
				K:              tt.fields.K,
				Name:           tt.fields.Name,
				Database:       tt.fields.Database,
				Columns:        tt.fields.Columns,
				Compressible:   tt.fields.Compressible,
				Incompressible: tt.fields.Incompressible,
				Domens:         tt.fields.Domens,
			}
			if err := t.getPriorities(); (err != nil) != tt.result {
				t1.Errorf("getPriorities() error = %v, wantErr %v", err, tt.result)
			}
		})
	}
}

func TestTable_getValue(t1 *testing.T) {
	type fields struct {
		K              uint64
		Name           string
		Database       internal.IDatabase
		Columns        []*Column
		Compressible   []int
		Incompressible []int
		Domens         []int
	}
	tests := []struct {
		name   string
		fields fields
		result bool
	}{
		{},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Table{
				K:              tt.fields.K,
				Name:           tt.fields.Name,
				Database:       tt.fields.Database,
				Columns:        tt.fields.Columns,
				Compressible:   tt.fields.Compressible,
				Incompressible: tt.fields.Incompressible,
				Domens:         tt.fields.Domens,
			}
			if err := t.getValue(); (err != nil) != tt.result {
				t1.Errorf("getValue() error = %v, wantErr %v", err, tt.result)
			}
		})
	}
}

func TestTable_getValueFactor(t1 *testing.T) {
	type fields struct {
		K              uint64
		Name           string
		Database       internal.IDatabase
		Columns        []*Column
		Compressible   []int
		Incompressible []int
		Domens         []int
	}
	tests := []struct {
		name   string
		fields fields
		result bool
	}{
		{},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Table{
				K:              tt.fields.K,
				Name:           tt.fields.Name,
				Database:       tt.fields.Database,
				Columns:        tt.fields.Columns,
				Compressible:   tt.fields.Compressible,
				Incompressible: tt.fields.Incompressible,
				Domens:         tt.fields.Domens,
			}
			if err := t.getValueFactor(); (err != nil) != tt.result {
				t1.Errorf("getValueFactor() error = %v, wantErr %v", err, tt.result)
			}
		})
	}
}
