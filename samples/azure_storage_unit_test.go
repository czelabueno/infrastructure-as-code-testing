package samples

import (
	"testing"

	unit_testing "github.com/czelabueno/infrastructure-as-code-testing/unit-testing"
	"github.com/stretchr/testify/assert"
)

func Test_storageMod(t *testing.T) {
	t.Parallel()

	vars := make(map[string]interface{})
	vars["account_tier"] = "Standard"
	vars["account_replication_type"] = "LRS"

	myunitmodule := unit_testing.TerraModule{
		RootFolderPath:      "../",
		TerraformModulePath: "examples/azure/tf-storage",
		Variables:           vars,
	}

	result := myunitmodule.TerratestExecution(t, false)

	assert.Equal(t, unit_testing.Successful, result)
}

func Test_WithStaticAnalysis(t *testing.T) {
	t.Parallel()

	vars := make(map[string]interface{})
	vars["account_tier"] = "Standard"
	vars["account_replication_type"] = "LRS"

	myunitmodule := unit_testing.TerraModule{
		RootFolderPath:      "../",
		TerraformModulePath: "examples/azure/tf-storage",
		Variables:           vars,
	}

	result := myunitmodule.TerratestExecution(t, true) // enable true to include check static analysis

	assert.Equal(t, unit_testing.Successful, result)
}
