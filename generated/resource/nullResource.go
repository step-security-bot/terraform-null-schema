package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const nullResource = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "This is set to a random value at create time.",
        "description_kind": "plain",
        "type": "string"
      },
      "triggers": {
        "description": "A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "The ` + "`" + `null_resource` + "`" + ` resource implements the standard resource lifecycle but takes no further action.\n\nThe ` + "`" + `triggers` + "`" + ` argument allows specifying an arbitrary set of values that, when changed, will cause the resource to be replaced.",
    "description_kind": "plain"
  },
  "version": 0
}`

func NullResourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(nullResource), &result)
	return &result
}
