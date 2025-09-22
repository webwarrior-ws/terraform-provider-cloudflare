// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum_application_test

import (
	"context"
	"testing"

	"github.com/webwarrior-ws/terraform-provider-cloudflare/internal/services/spectrum_application"
	"github.com/webwarrior-ws/terraform-provider-cloudflare/internal/test_helpers"
)

func TestSpectrumApplicationsDataSourceModelSchemaParity(t *testing.T) {
	t.Skip("need investigation: currently broken")
	t.Parallel()
	model := (*spectrum_application.SpectrumApplicationsDataSourceModel)(nil)
	schema := spectrum_application.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
