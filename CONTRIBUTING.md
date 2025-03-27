# Contributing

## Pre-Requisites

Install https://openapi-generator.tech/#try

## Regenerating the SDK

**Update the OpenAPI spec in `openapi/conductor.yaml` from the [official spec](https://docs.conductor.is/api-ref/openapi) if needed**
> Note: Comment out the entire block for `/end-users/{id}/passthrough/{integrationSlug}`: since `openapi-generator` doesn't parse it correctly.

```bash
# With local installation (with brew)
openapi-generator generate -i ./openapi/conductor.yaml -g go -o . -c gen-config.yml

go mod tidy
```