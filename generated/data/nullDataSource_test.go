package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-null-schema/v3/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestNullDataSourceSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.NullDataSourceSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
