package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type tfPlan struct {
	Changes []tfChangeResource `json:"resource_changes"`
}

type tfChangeResource struct {
	Address  string   `json:"address"`
	Type     string   `json:"type"`
	Provider string   `json:"provider_name"`
	Change   tfChange `json:"change"`
}

type tfChange struct {
	Actions    []string               `json:"actions"`
	Properties map[string]interface{} `json:"after"`
}

func generateTerraformPlan() *tfPlan {
	file, err := ioutil.TempFile("", "tfplan")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	cmd := exec.Command("terraform", "plan", "-out="+file.Name())
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatalf("generating terraform plan failed with %s\n", err)
	}

	return parsePlan(file.Name())
}

func parsePlan(planFile string) *tfPlan {
	cmd := exec.Command("terraform", "show", "-json", planFile)
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("getting terraform plan json failed with %s\n", err)
	}

	var plan tfPlan
	json.Unmarshal(out, &plan)

	return &plan
}

func getNewResources(plan *tfPlan) *[]NewResource {
	var newResources = make([]NewResource, 0)
	for _, change := range plan.Changes {
		if len(change.Change.Actions) != 1 {
			continue
		}

		if change.Change.Actions[0] != "create" {
			continue
		}

		newResource := NewResource{
			Address:    change.Address,
			Type:       change.Type,
			Provider:   change.Provider,
			Properties: change.Change.Properties,
		}
		newResources = append(newResources, newResource)
	}
	return &newResources
}

func GetNewResourcesFromFile(planFile string) *[]NewResource {
	return getNewResources(parsePlan(planFile))
}

func GetNewResourcesFromTerraform() *[]NewResource {
	return getNewResources(generateTerraformPlan())
}
