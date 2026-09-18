package cmd

import (
	"os"
	"strings"
	"testing"

	"github.com/goccy/go-yaml"
)

func loadAllOfFixture(t *testing.T) OpenAPISpec {
	t.Helper()
	data, err := os.ReadFile("testdata/allof.yaml")
	if err != nil {
		t.Fatal(err)
	}
	var spec OpenAPISpec
	if err := yaml.Unmarshal(data, &spec); err != nil {
		t.Fatal(err)
	}
	return spec
}

func TestResolveSchemaComposition(t *testing.T) {
	spec := loadAllOfFixture(t)
	schema, err := resolveSchemaByName("Derived", spec, make(map[string]bool))
	if err != nil {
		t.Fatal(err)
	}

	if got := schema.Properties["base"].(map[string]interface{})["type"]; got != "integer" {
		t.Fatalf("base override type = %v, want integer", got)
	}
	for _, name := range []string{"base", "middle", "own"} {
		if _, ok := schema.Properties[name]; !ok {
			t.Errorf("missing inherited property %q", name)
		}
	}
	for _, name := range []string{"base", "middle", "own"} {
		if !contains(schema.Required, name) {
			t.Errorf("missing required field %q", name)
		}
	}
}

func TestResolveSchemaCompositionCycle(t *testing.T) {
	spec := loadAllOfFixture(t)
	schema, err := resolveSchemaByName("CycleA", spec, make(map[string]bool))
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"a", "b"} {
		if _, ok := schema.Properties[name]; !ok {
			t.Errorf("missing cyclic property %q", name)
		}
	}
}

func TestGenerateRootSectionPreservesWrappedDescription(t *testing.T) {
	spec := loadAllOfFixture(t)
	schema, err := resolveSchemaByName("Wrapped", spec, make(map[string]bool))
	if err != nil {
		t.Fatal(err)
	}
	output := generateRootSection("Wrapped", schema, spec, map[string]string{})
	if strings.Count(output, "`metadata`") != 1 {
		t.Fatalf("metadata rendered more than once:\n%s", output)
	}
	if !strings.Contains(output, "metadata description") {
		t.Fatalf("wrapped description missing:\n%s", output)
	}
}

func contains(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}
