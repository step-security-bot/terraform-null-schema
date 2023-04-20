package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const nullDataSource = `{
  "block": {
    "attributes": {
      "has_computed_default": {
        "computed": true,
        "description": "If set, its literal value will be stored and returned. If not, its value defaults to ` + "`" + `\"default\"` + "`" + `. This argument exists primarily for testing and has little practical use.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "deprecated": true,
        "description": "This attribute is only present for some legacy compatibility issues and should not be used. It will be removed in a future version.",
        "description_kind": "plain",
        "type": "string"
      },
      "inputs": {
        "description": "A map of arbitrary strings that is copied into the ` + "`" + `outputs` + "`" + ` attribute, and accessible directly for interpolation.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "outputs": {
        "computed": true,
        "description": "After the data source is \"read\", a copy of the ` + "`" + `inputs` + "`" + ` map.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "random": {
        "computed": true,
        "description": "A random value. This is primarily for testing and has little practical use; prefer the [hashicorp/random provider](https://registry.terraform.io/providers/hashicorp/random) for more practical random number use-cases.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "deprecated": true,
    "description": "The ` + "`" + `null_data_source` + "`" + ` data source implements the standard data source lifecycle but does not\ninteract with any external APIs.\n\nHistorically, the ` + "`" + `null_data_source` + "`" + ` was typically used to construct intermediate values to re-use elsewhere in configuration. The\nsame can now be achieved using [locals](https://www.terraform.io/docs/language/values/locals.html).\n",
    "description_kind": "plain"
  },
  "version": 0
}`

func NullDataSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(nullDataSource), &result)
	return &result
}
