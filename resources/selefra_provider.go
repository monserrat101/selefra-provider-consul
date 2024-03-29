package resources

import (
	"context"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/terraform/bridge"
)

func GetSelefraProvider() *provider.Provider {
	diagnostics := schema.NewDiagnostics()
	selefraProvider, d := GetSelefraTerraformProvider().ToSelefraProvider(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) *bridge.TerraformBridge {
		return taskClient.(*Client).TerraformBridge
	})
	if diagnostics.AddDiagnostics(d).HasError() {
		panic(diagnostics.ToString())
	}

	selefraProvider.TableList = GetSelefraTables()

	return selefraProvider
}

func GetSelefraTables() []*schema.Table {

	diagnostics := schema.NewDiagnostics()
	tables := make([]*schema.Table, 0)
	var table *schema.Table
	var d *schema.Diagnostics

	table, d = TableSchemaGenerator_consul_prepared_query()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_config_entry()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_agent_service()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_catalog_entry()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_admin_partition()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_keys()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_acl_token()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_license()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_acl_policy()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_namespace_role_attachment()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_peering()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_certificate_authority()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_acl_role()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_acl_token_role_attachment()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_key_prefix()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_node()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_intention()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_service()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_acl_binding_rule()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_namespace()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_acl_auth_method()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_network_area()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_namespace_policy_attachment()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_peering_token()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_acl_token_policy_attachment()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	table, d = TableSchemaGenerator_consul_autopilot_config()
	if !diagnostics.AddDiagnostics(d).HasError() {
		tables = append(tables, table)
	}

	if diagnostics.HasError() {
		panic(diagnostics.ToString())
	}

	return tables
}
