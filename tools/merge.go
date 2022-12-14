package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"sigs.k8s.io/yaml"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("usage: merge.go path/to/models.yaml path/to/api.yaml")
	}

	err := run(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
}

func run(modelsFile, apiFile string) error {
	modelsYAML, err := os.ReadFile(modelsFile)
	if err != nil {
		return err
	}

	apiYAML, err := os.ReadFile(apiFile)
	if err != nil {
		return err
	}

	models := map[string]any{}
	api := map[string]any{}

	err = yaml.Unmarshal(modelsYAML, &models)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(apiYAML, &api)
	if err != nil {
		return err
	}

	for komp, schema := range mustGet[map[string]any](models, "components.schemas") {
		if komp == "Identifiable" || komp == "AssetAdministrationShellDescriptor" {
			continue
		}

		submodelElemetArray, ok := get[string](schema, "properties.value.items.$ref")
		if ok && submodelElemetArray == "#/components/schemas/SubmodelElement" {
			set(schema, "properties.value.items.$ref", "#/components/schemas/ISubmodelElement")
		}

		submodelElemetArray, ok = get[string](schema, "allOf[1].properties.value.items.$ref")
		if ok && submodelElemetArray == "#/components/schemas/SubmodelElement" {
			set(schema, "allOf[1].properties.value.items.$ref", "#/components/schemas/ISubmodelElement")
		}

		submodelElemetArray, ok = get[string](schema, "allOf[5].properties.submodelElements.items.$ref")
		if ok && submodelElemetArray == "#/components/schemas/SubmodelElement" {
			set(schema, "allOf[5].properties.submodelElements.items.$ref", "#/components/schemas/ISubmodelElement")
		}

		renameDataElementToSubmodelElement, ok := get[string](schema, "allOf[1].properties.annotations.items.$ref")
		if ok && renameDataElementToSubmodelElement == "#/components/schemas/DataElement" {
			set(schema, "allOf[1].properties.annotations.items.$ref", "#/components/schemas/ISubmodelElement")
		}

		renameSubmodelElementToISubmodelElement, ok := get[string](schema, "properties.value.$ref")
		if ok && renameSubmodelElementToISubmodelElement == "#/components/schemas/SubmodelElement" {
			set(schema, "properties.value.$ref", "#/components/schemas/ISubmodelElement")
		}

		if komp == "AssetAdministrationShell" {
			props, ok := get[map[string]any](schema, "allOf[2].properties")
			if ok {
				props["asset"] = map[string]any{"$ref": "#/components/schemas/Reference"}
				set(schema, "allOf[2].properties", props)
			}
		}

		set(api, "components.schemas."+komp, schema)
	}

	// override the incorrect return type
	set(api, "paths./shells/{aasId}/aas/submodels.get.responses.200.content.application/json.schema", map[string]any{"type": "array", "items": map[string]any{"$ref": "#/components/schemas/Submodel"}})

	y, err := yaml.Marshal(api)
	if err != nil {
		return err
	}

	fmt.Println(string(y))

	return nil
}

func mustGet[T any](obj any, path string) T {
	r, ok := get[T](obj, path)
	if !ok {
		panic("can't get value " + path)
	}
	return r
}

func get[T any](obj any, path string) (T, bool) {
	pathParts := strings.Split(strings.ReplaceAll(strings.ReplaceAll(path, "]", ""), "[", "."), ".")

	last := obj

	for _, p := range pathParts {
		index, err := strconv.Atoi(p)
		if err == nil {
			a, ok := last.([]any)
			if !ok || index >= len(a) {
				var r T
				return r, false
			}

			last = a[index]
			continue
		}

		m, ok := last.(map[string]any)
		if !ok {
			var r T
			return r, false
		}

		last = m[p]
	}

	if last == nil {
		var r T
		return r, false
	}

	return last.(T), true
}

func set(obj any, path string, value any) bool {
	pathParts := strings.Split(strings.ReplaceAll(strings.ReplaceAll(path, "]", ""), "[", "."), ".")

	last := obj

	for i := 0; i < len(pathParts); i++ {
		isLast := i == len(pathParts)-1
		p := pathParts[i]
		index, err := strconv.Atoi(p)
		if err == nil {
			a, ok := last.([]any)
			if ok && index < len(a) {
				if isLast {
					a[index] = value
					break
				}

				last = a[index]
				continue
			}
		}

		m, ok := last.(map[string]any)
		if !ok {
			return false
		}

		if isLast {
			m[p] = value
			break
		}

		last = m[p]

	}

	return true
}
