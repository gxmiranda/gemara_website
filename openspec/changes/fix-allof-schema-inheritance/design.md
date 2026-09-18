## Context

The website build converts OpenAPI schemas into Markdown reference pages. Embedded CUE definitions will be emitted as local OpenAPI `allOf` references, while the current renderer primarily reads direct `properties` and `required` values. The fix must remain in the generator, preserve the existing Markdown output conventions, and avoid changing upstream schemas or hand-editing generated pages.

## Goals / Non-Goals

**Goals:**

- Resolve local component-schema references used by `allOf` during rendering.
- Flatten inherited properties and required fields into the documented schema view.
- Support recursive inheritance, direct-property precedence, wrapper descriptions, and cycle termination.
- Keep behavior deterministic and covered by focused fixtures and regression tests.

**Non-Goals:**

- Resolving remote references or introducing network access during generation.
- General-purpose JSON Schema evaluation beyond the composition needed by the renderer.
- Changing the OpenAPI generator in the Gemara specification repository.
- Editing committed generated schema pages as the source of truth.

## Decisions

### Resolve only local component references

The renderer will resolve `$ref` values that point to schemas in the loaded document's local components map. Remote references and unsupported reference forms remain unchanged or are reported through existing generator behavior. This avoids network-dependent builds and keeps resolution bounded to known input.

### Flatten composition at the schema-rendering boundary

Composition will be resolved immediately before property and required-field rendering, using a recursive resolver that carries a visited-reference set. This keeps the loaded OpenAPI model intact and limits the behavior change to generated documentation.

### Merge in declaration order with derived properties winning

Inherited `allOf` members will be merged first, followed by the current schema's direct properties. Required fields will be unioned while preserving stable input order. A direct property with the same name overrides the inherited property, matching the acceptance criteria.

### Treat description-plus-allOf as one property

When a property wrapper contains a description and a single `allOf` reference, the renderer will resolve the referenced schema for that property while retaining the wrapper description and property identity. It will not emit a duplicate nested property.

### Detect cycles with an active recursion set

A reference encountered while already active will stop recursive expansion at that edge. The generator will continue rendering the non-cyclic portions rather than recurse indefinitely or fail the whole page.

## Risks / Trade-offs

- [Risk] Flattening may expose more fields than existing pages and alter generated Markdown broadly. → Mitigate with fixture snapshots or focused assertions and regenerate only after tests pass.
- [Risk] Cyclic or malformed schemas may produce incomplete documentation. → Mitigate by terminating cycles deterministically and preserving direct declarations.
- [Risk] OpenAPI library types may represent `$ref` wrappers differently across versions. → Mitigate by testing both schema-level and property-level composition using the repository's actual parser types.
- [Risk] Required-field ordering could become unstable. → Mitigate by using ordered de-duplication rather than an unordered set in output construction.

## Migration Plan

1. Add resolver logic and fixtures alongside the existing generator tests.
2. Run the generator test suite and the site link/build checks.
3. Regenerate schema pages through the normal build when the upstream OpenAPI change is available.
4. If regressions appear, revert the renderer change; generated output can be regenerated from the unchanged source.

## Open Questions

- Confirm the exact OpenAPI parser representation of a description-plus-`allOf` property in the current toolchain before implementation.
