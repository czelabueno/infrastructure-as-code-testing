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
func (module *TerraModule) TerratestExecution(t *testing.T, useStaticAnalysis bool) (result TerraModuleStatus) {
	terraformOptions, err := CreateTerraformOptions(t, module)

	if err != nil {
		t.Fatal()
		return Failed
	}

	if useStaticAnalysis {
		module.RunStaticAnalysis(t, terraformOptions)
		// Run `terraform apply` and failed if en error in founded
		terraform.Apply(t, terraformOptions)
	} else {
		// Run `terraform init` and `terraform apply` and failed if en error in founded
		terraform.InitAndApply(t, terraformOptions)
	}
	// at the end of the test the resources are destroyed
	defer terraform.Destroy(t, terraformOptions)
	return Successful
}

// RunStaticAnalysis ... something comment
func (module *TerraModule) RunStaticAnalysis(t *testing.T, terraformOptions *terraform.Options) (result TerraModuleStatus) {
	if terraformOptions == nil {
		tfOptions, err := CreateTerraformOptions(t, module)
		terraformOptions = tfOptions
		if err != nil {
			t.Fatal()
			return Failed
		}
	}
	// terraform init ...
	terraform.Init(t, terraformOptions)
	// terraform validate ...
	args := []string{"validate"}
	args = append(args, terraform.FormatTerraformBackendConfigAsArgs(terraformOptions.BackendConfig)...)
	terraform.RunTerraformCommand(t, terraformOptions, args...)
	// terraform plan ...
	defer terraform.Plan(t, terraformOptions)
	return Successful
}

// CreateTerraformOptions ... create terraform options with terraform module
func CreateTerraformOptions(t *testing.T, module *TerraModule) (*terraform.Options, error) {
	if module.RootFolderPath == "" {
		module.RootFolderPath = "../" // Set Default root folder path assuming that test file is located in ./test
	}
	if module.TerraformModulePath == "" {
		module.TerraformModulePath = "." // Set Default tf module folder path assuming that .tf files are located in the root
	}

	return &terraform.Options{
		TerraformDir: test_structure.CopyTerraformFolderToTemp(t, module.RootFolderPath, module.TerraformModulePath),
		Vars:         module.Variables,
	}, nil
}
