// Declare the package name as 'tf_provider'
package tf_provider

// Import the 'schema' package from the Terraform Plugin SDK
import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Define the Provider function which returns a pointer to a schema.Provider
func Provider() *schema.Provider {
	// Return a pointer to a new schema.Provider struct
	return &schema.Provider{
		// Initialize the ResourcesMap field with a map of resource names to resource definitions
		ResourcesMap: map[string]*schema.Resource{
			// Map the resource name 'example_resource' to its resource schema and CRUD functions
			"example_resource": resourceExample(),
		},
	}
}
