package unit_testing

import (
	"testing"

	ct "github.com/daviddengcn/go-colortext"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

// TerratestExecution test only one module isolated. The test stages include: Parser, dry run, init, create or update, provisioning validate and destroy.
// Terraform commands wrapped by Terratest
// Parameters:
// t - current test state
// useStaticAnalysis - flag to indicate if execution will include static analysis
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
		ct.Foreground(ct.Blue, true)
		logger.Logf(t, "Creating or changing module...")
		ct.Foreground(ct.White, false)
		// Run `terraform apply` and failed if en error in founded
		terraform.Apply(t, terraformOptions)
	} else {
		ct.Foreground(ct.Blue, true)
		logger.Logf(t, "Creating or changing module...")
		ct.Foreground(ct.White, false)
		// Run `terraform init` and `terraform apply` and failed if en error in founded
		terraform.InitAndApply(t, terraformOptions)
	}

	// at the end of the test the resources are destroyed
	ct.Foreground(ct.Blue, true)
	logger.Logf(t, "Destroying module...")
	ct.Foreground(ct.White, false)
	defer terraform.Destroy(t, terraformOptions)
	return Successful
}

// RunStaticAnalysis test just static analysis. Include Parser, dry run: init, validate and plan.
// Terraform commands wrapped by Terratest
// Parameters:
// t - current test state
// terraformOptions - structure that contains module path and args.
// Return:
// result - TerraModuleStatus
func (module *TerraModule) RunStaticAnalysis(t *testing.T, terraformOptions *terraform.Options) (result TerraModuleStatus) {
	ct.Foreground(ct.Blue, true)
	logger.Logf(t, "Running static analysis...")
	ct.Foreground(ct.White, false)

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

// CreateTerraformOptions ... create terraform options using TerraModule fields
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
