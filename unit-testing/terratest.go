package unit_testing

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

// TerratestExecution test only one workload module isolated. The test stages include: init, create or update, provisioning validate and destroy.
// Terraform commands wrapped by Terratest
// Parameters:
// t - current test state
// Return:
// result - TerraModuleStatus
func (module *TerraModule) TerratestExecution(t *testing.T) (result TerraModuleStatus) {
	if module.RootFolderPath == "" {
		module.RootFolderPath = "../" // Set Default root folder path assuming that test file is located in ./test
	}
	if module.TerraformModulePath == "" {
		module.TerraformModulePath = "." // Set Default tf module folder path assuming that .tf files are located in the root
	}
	terraformOptions := &terraform.Options{
		TerraformDir: test_structure.CopyTerraformFolderToTemp(t, module.RootFolderPath, module.TerraformModulePath),
		Vars:         module.Variables,
	}

	// Run `terraform init` and `terraform apply` and failed if en error in founded
	terraform.InitAndApply(t, terraformOptions)

	// at the end of the test the resources are destroyed
	defer terraform.Destroy(t, terraformOptions)
	return
}
