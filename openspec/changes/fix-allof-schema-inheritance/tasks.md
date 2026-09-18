## 1. Understand Existing Generator

- [x] 1.1 Confirm the OpenAPI `Schema` representation and current property/required rendering flow in `tools/internal/cmd/openapi_types.go` and `tools/internal/cmd/openapi2md.go`.
- [x] 1.2 Confirm how the active OpenAPI parser represents schema-level and property-level description-plus-`allOf` wrappers.

## 2. Implement Resolution

- [x] 2.1 Add a local component-reference resolver that rejects unsupported remote references and reports missing local schemas through existing error handling.
- [x] 2.2 Add recursive `allOf` flattening with an active-reference set so multi-level inheritance resolves and cycles terminate safely.
- [x] 2.3 Merge inherited properties before direct properties, with direct definitions overriding duplicate names while preserving deterministic order.
- [x] 2.4 Merge inherited and direct required fields with stable de-duplication.
- [x] 2.5 Integrate resolved schemas at the schema-rendering boundary without mutating the loaded OpenAPI document.
- [x] 2.6 Preserve wrapper descriptions and render description-plus-`allOf` properties as one property entry.

## 3. Add Regression Coverage

- [x] 3.1 Add OpenAPI fixtures for single-level inheritance and inherited required fields.
- [x] 3.2 Add fixtures covering multi-level inheritance and direct-property override precedence.
- [x] 3.3 Add fixtures covering cyclic references and assert generation completes without infinite recursion.
- [x] 3.4 Add a fixture for a described property wrapper containing `allOf` and assert the description and single-property output.
- [x] 3.5 Add focused Go tests asserting generated Markdown contains the expected inherited, required, overridden, and wrapped-property output.

## 4. Verify Generated Documentation

- [x] 4.1 Run the Go formatter and generator test suite.
- [x] 4.2 Run the site build or relevant `make` checks, including link validation where available. (Go checks pass; `make test-links` requires unavailable `bundle`.)
- [x] 4.3 Regenerate schema documentation against the updated upstream OpenAPI input and inspect representative catalog and log pages.
- [x] 4.4 Run cleanup and verify no unintended generated artifacts or term-link changes remain in the diff.
