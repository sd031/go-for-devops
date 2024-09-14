// Declare the package name as 'tf_provider'
package tf_provider

// Import the 'schema' package from the Terraform Plugin SDK
import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Define the 'resourceExample' function that returns a pointer to a 'schema.Resource'
func resourceExample() *schema.Resource {
	// Return a pointer to a new 'schema.Resource' struct
	return &schema.Resource{
		// Assign the 'Create' function for resource creation
		Create: resourceExampleCreate,
		// Assign the 'Read' function for resource reading
		Read: resourceExampleRead,
		// Assign the 'Update' function for resource updating
		Update: resourceExampleUpdate,
		// Assign the 'Delete' function for resource deletion
		Delete: resourceExampleDelete,

		// Define the schema of the resource
		Schema: map[string]*schema.Schema{
			// Define 'example_value' as a required string attribute
			"example_value": &schema.Schema{
				// Specify the type of the attribute as a string
				Type: schema.TypeString,
				// Mark the attribute as required
				Required: true,
			},
		},
	}
}

// Define the 'resourceExampleCreate' function to handle resource creation
func resourceExampleCreate(d *schema.ResourceData, m interface{}) error {
	// Set the resource ID to the value provided by the user in 'example_value'
	d.SetId(d.Get("example_value").(string))
	// Call the 'Read' function to refresh the state after creation
	return resourceExampleRead(d, m)
}

// Define the 'resourceExampleRead' function to handle resource reading
func resourceExampleRead(d *schema.ResourceData, m interface{}) error {
	// Logic for reading the resource state (typically involves fetching data from an API)
	return nil
}

// Define the 'resourceExampleUpdate' function to handle resource updating
func resourceExampleUpdate(d *schema.ResourceData, m interface{}) error {
	// Logic for updating the resource
	// After updating, refresh the resource state
	return resourceExampleRead(d, m)
}

// Define the 'resourceExampleDelete' function to handle resource deletion
func resourceExampleDelete(d *schema.ResourceData, m interface{}) error {
	// Remove the resource ID to indicate that the resource has been deleted
	d.SetId("")
	return nil
}
