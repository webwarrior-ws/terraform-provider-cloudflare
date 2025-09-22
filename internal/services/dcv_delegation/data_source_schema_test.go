// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dcv_delegation_test

import (
	"context"
	"testing"

	"github.com/webwarrior-ws/terraform-provider-cloudflare/internal/services/dcv_delegation"
	"github.com/webwarrior-ws/terraform-provider-cloudflare/internal/test_helpers"
)

func TestDCVDelegationDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*dcv_delegation.DCVDelegationDataSourceModel)(nil)
	schema := dcv_delegation.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
