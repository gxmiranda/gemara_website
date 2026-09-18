## Why

The schema reference generator will receive OpenAPI `allOf` output for embedded CUE definitions, but it currently documents only fields declared directly on each schema. This will make generated Gemara reference pages incomplete and can hide inherited required fields when the upstream model change lands.

## What Changes

- Resolve local OpenAPI `allOf` `$ref` entries while rendering schema reference pages.
- Include inherited properties and required fields in generated documentation.
- Support multi-level inheritance and direct-property override precedence.
- Preserve descriptions on properties represented by a description-plus-`allOf` wrapper.
- Detect cyclic schema references so generation terminates safely.
- Add fixtures and regression tests covering inheritance, overrides, cycles, and wrapped references.

## Capabilities

### New Capabilities

- `schema-inheritance-resolution`: Resolve OpenAPI schema composition when generating complete schema reference documentation.

### Modified Capabilities

<!-- No existing OpenSpec capabilities are defined in this repository. -->

## Impact

- OpenAPI-to-Markdown generation tooling and its schema/property rendering logic.
- Generator fixtures and automated tests.
- Generated `schema/*.md` output after the next documentation build; no generated files should be edited directly.
- No public runtime API or deployment configuration changes.
