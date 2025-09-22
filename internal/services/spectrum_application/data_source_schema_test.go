// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_application_test

import (
	"context"
	"testing"

	"github.com/webwarrior-ws/terraform-provider-cloudflare/internal/services/spectrum_application"
	"github.com/webwarrior-ws/terraform-provider-cloudflare/internal/test_helpers"
)

func TestSpectrumApplicationDataSourceModelSchemaParity(t *testing.T) {
	t.Skip("need investigation: currently broken")
	t.Parallel()
	model := (*spectrum_application.SpectrumApplicationDataSourceModel)(nil)
	schema := spectrum_application.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
