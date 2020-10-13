package samples

import (
	"testing"

	unit_testing "github.com/czelabueno/infrastructure-as-code-testing/unit-testing"
	"github.com/stretchr/testify/assert"
)

func Test_storageMod(t *testing.T) {
	t.Parallel()

	// Define yout terraform input variables
	vars := make(map[string]interface{})
	vars["account_tier"] = "Standard"
	vars["account_replication_type"] = "LRS"

	//Initialize TerraModule struct entering the data to locate individual module that you want test.
	myunitmodule := unit_testing.TerraModule{
		RootFolderPath:      "../",                       //Put root path from this .go file is located
		TerraformModulePath: "examples/azure/tf-storage", // Put module path where is located .tf files.
		Variables:           vars,
	}

	//Call to TerratestExecution to run unit test function: deploy, validate and undeploy.
	result := myunitmodule.TerratestExecution(t, false) // change to (t, true) to include static analysis

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
