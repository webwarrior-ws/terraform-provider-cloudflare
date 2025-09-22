// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schema_validation_schemas_test

import (
	"context"
	"testing"

	"github.com/webwarrior-ws/terraform-provider-cloudflare/internal/services/schema_validation_schemas"
	"github.com/webwarrior-ws/terraform-provider-cloudflare/internal/test_helpers"
)

func TestSchemaValidationSchemasModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*schema_validation_schemas.SchemaValidationSchemasModel)(nil)
	schema := schema_validation_schemas.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
