package unit_testing

// TerraModule is the object to instance terraform modules to testing for
type TerraModule struct {
	RootFolderPath      string
	TerraformModulePath string
	Variables           map[string]interface{}
}

// TerraModuleStatus enumerates the values for module status validation.
type TerraModuleStatus string

const (
	// Successful .. when TerraModule meet with health check validation
	Successful TerraModuleStatus = "Successful"
	// Failed .. when TerraModule not meet with health check validation
	Failed TerraModuleStatus = "Failed"
)
