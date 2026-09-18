## ADDED Requirements

### Requirement: Resolve inherited schema properties

The schema documentation generator SHALL resolve local OpenAPI `allOf` references and include inherited properties in the generated page for the derived schema.

#### Scenario: Derived schema includes inherited properties
- **WHEN** a schema contains an `allOf` reference to a local component schema with properties
- **THEN** the generated page includes those inherited properties along with the derived schema's direct properties

### Requirement: Preserve inherited required fields

The generator SHALL include required fields inherited through local `allOf` references and SHALL preserve direct required fields.

#### Scenario: Required fields are merged
- **WHEN** a base schema and derived schema each declare required fields
- **THEN** the generated page marks the union of those fields as required

### Requirement: Apply derived-property precedence

The generator SHALL use a derived schema's direct property definition when it declares the same property name as an inherited schema.

#### Scenario: Direct property overrides inherited property
- **WHEN** an inherited schema and its derived schema define a property with the same name
- **THEN** the generated page renders the derived property's definition and does not duplicate the property

### Requirement: Resolve multi-level inheritance

The generator SHALL resolve local `allOf` inheritance across more than one schema level.

#### Scenario: Grandparent properties are rendered
- **WHEN** a schema inherits from a schema that itself inherits from another local schema
- **THEN** the generated page includes properties and required fields from all reachable non-cyclic ancestors

### Requirement: Terminate cyclic references safely

The generator SHALL terminate local schema resolution when an `allOf` reference forms a cycle and SHALL still render the non-cyclic declarations available at the point of traversal.

#### Scenario: Cyclic inheritance does not fail generation
- **WHEN** local schemas reference one another through a cycle
- **THEN** generation completes without infinite recursion and renders each reachable direct declaration at most once

### Requirement: Preserve wrapped property descriptions

The generator SHALL render a property represented by a description-plus-`allOf` wrapper as one property and retain the wrapper description.

#### Scenario: Described reference property remains one property
- **WHEN** a property has a description and an `allOf` reference to a local schema
- **THEN** the generated page contains one property entry with the wrapper description and the referenced property's resolved type or fields

### Requirement: Cover composition with fixtures

The generator test suite SHALL include fixtures and assertions for single-level inheritance, required-field merging, overrides, multi-level inheritance, cycles, and described reference wrappers.

#### Scenario: Regression fixtures validate all composition cases
- **WHEN** the composition fixture suite runs
- **THEN** each listed composition case is verified against the expected generated schema documentation
