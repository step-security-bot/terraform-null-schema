package generated

import (
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-null-schema/v3/generated/data"
	"github.com/lonegunmanb/terraform-null-schema/v3/generated/resource"
)

var Resources map[string]*tfjson.Schema
var DataSources map[string]*tfjson.Schema

func init() {
	resources := make(map[string]*tfjson.Schema, 0)
	dataSources := make(map[string]*tfjson.Schema, 0)  
	resources["null_resource"] = resource.NullResourceSchema()  
	dataSources["null_data_source"] = data.NullDataSourceSchema()  
	Resources = resources
	DataSources = dataSources
}