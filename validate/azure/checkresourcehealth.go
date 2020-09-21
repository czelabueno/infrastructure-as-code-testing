package azure

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/resourcehealth/mgmt/2017-07-01/resourcehealth"
	unit_testing "github.com/czelabueno/infrastructure-as-code-testing/unit-testing"
	ct "github.com/daviddengcn/go-colortext"
	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

// ValidateModule ... Check resource health status into azure rest api
// Oficial documentation here: https://docs.microsoft.com/en-us/rest/api/resourcehealth/
// Parameters:
// t - current test state
// terraformOptions - structure that contains module path and args.
// Return:
// result - unit_testing.TerraModuleStatus
func ValidateModule(t *testing.T, terraformOptions *terraform.Options) (result unit_testing.TerraModuleStatus) {
	ct.Foreground(ct.Blue, true)
	logger.Logf(t, "Validating provisioned resource ...")
	ct.Foreground(ct.White, false)

	// Validate terraformOptions param because is mandatory for validation
	if terraformOptions != nil {
		subscriptionID := terraform.OutputRequired(t, terraformOptions, "subscriptionId")
		resourceID := terraform.OutputRequired(t, terraformOptions, "resourceId")
		resourceName := terraform.OutputRequired(t, terraformOptions, "resourceName")

		// Create availabilityStatusesClient
		availabilityStatusesClient := resourcehealth.NewAvailabilityStatusesClient(subscriptionID)

		// Get azure token access from Environment(Client credential or client certificate) or File json or Az CLI
		authorizer, err := azure.NewAuthorizer()
		if err == nil {
			availabilityStatusesClient.Authorizer = *authorizer
		} else {
			t.Fatalf("Authorization is Failed: %s", err.Error())
			result = unit_testing.Failed
			t.Fail()
		}

		// Get health availability status of the given resource
		ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
		defer cancel()
		availabilityStatus, err := availabilityStatusesClient.GetByResource(ctx, resourceID, "", "")
		if err != nil {
			t.Fatalf("Cant connect with azure resourcehealth api service: %s", err.Error())
			result = unit_testing.Failed
			t.Fail() // So if error is not null the test must be fail
		}

		// Checking that resource validated is healthy available
		if availabilityStatus.Properties.AvailabilityState != resourcehealth.Available {
			ct.Foreground(ct.Red, false)
			logger.Logf(t, "Resource %s is unhealthy status: \t%s", resourceName, fmt.Sprint(availabilityStatus.Properties.AvailabilityState))
			ct.Foreground(ct.White, false)
			t.Error("Resource " + resourceName + " is unhealthy :( . Please check resource config")
			result = unit_testing.Failed
			t.Fail() // So if resource is unhealthy the test should be fail
		} else {
			ct.Foreground(ct.Green, true)
			result = unit_testing.Successful
			logger.Logf(t, "Validation complete! Resource "+resourceName+" is: "+fmt.Sprint(resourcehealth.Available))
			ct.Foreground(ct.White, false)
		}

	} else {
		ct.Foreground(ct.Red, false)
		logger.Logf(t, "ERROR: terraform.Options cant be nill")
		ct.Foreground(ct.White, false)
		result = unit_testing.Failed
		t.Fatal("terraform.Options cant be nil :( . Please check your test code")
		t.Fail()
	}
	return result

}
