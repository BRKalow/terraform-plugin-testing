// generated by stringer -type=GraphNodeConfigType graph_config_node_type.go; DO NOT EDIT

package terraform

import "fmt"

const _GraphNodeConfigType_name = "GraphNodeConfigTypeInvalidGraphNodeConfigTypeResourceGraphNodeConfigTypeProviderGraphNodeConfigTypeModuleGraphNodeConfigTypeOutput"

var _GraphNodeConfigType_index = [...]uint8{0, 26, 53, 80, 105, 130}

func (i GraphNodeConfigType) String() string {
	if i < 0 || i+1 >= GraphNodeConfigType(len(_GraphNodeConfigType_index)) {
		return fmt.Sprintf("GraphNodeConfigType(%d)", i)
	}
	return _GraphNodeConfigType_name[_GraphNodeConfigType_index[i]:_GraphNodeConfigType_index[i+1]]
}
