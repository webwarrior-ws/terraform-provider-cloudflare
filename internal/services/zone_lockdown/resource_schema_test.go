// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_lockdown_test

import (
	"context"
	"testing"

	"github.com/webwarrior-ws/terraform-provider-cloudflare/internal/services/zone_lockdown"
	"github.com/webwarrior-ws/terraform-provider-cloudflare/internal/test_helpers"
)

func TestZoneLockdownModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zone_lockdown.ZoneLockdownModel)(nil)
	schema := zone_lockdown.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
