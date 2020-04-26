package core

type NewResource struct {
	Address    string
	Type       string
	Provider   string
	Properties map[string]interface{}
}

type RunConfig struct {
	TerraformPlanPath string
}

func Run(c *RunConfig) error {
	var newResources *[]NewResource
	if c.TerraformPlanPath != "" {
		newResources = GetNewResourcesFromFile(c.TerraformPlanPath)
	} else {
		newResources = GetNewResourcesFromTerraform()
	}

	for _, newResource := range *newResources {
		println(newResource.Address)
	}

	return nil
}
