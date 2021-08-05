// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// NamespacesColumns holds the columns for the "namespaces" table.
	NamespacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 64},
		{Name: "created", Type: field.TypeTime},
	}
	// NamespacesTable holds the schema information for the "namespaces" table.
	NamespacesTable = &schema.Table{
		Name:       "namespaces",
		Columns:    NamespacesColumns,
		PrimaryKey: []*schema.Column{NamespacesColumns[0]},
	}
	// WorkflowsColumns holds the columns for the "workflows" table.
	WorkflowsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "created", Type: field.TypeTime},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 1024, Default: ""},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "revision", Type: field.TypeInt, Default: 0},
		{Name: "workflow", Type: field.TypeBytes},
		{Name: "log_to_events", Type: field.TypeString, Nullable: true},
		{Name: "namespace_workflows", Type: field.TypeString, Nullable: true, Size: 64},
	}
	// WorkflowsTable holds the schema information for the "workflows" table.
	WorkflowsTable = &schema.Table{
		Name:       "workflows",
		Columns:    WorkflowsColumns,
		PrimaryKey: []*schema.Column{WorkflowsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workflows_namespaces_workflows",
				Columns:    []*schema.Column{WorkflowsColumns[8]},
				RefColumns: []*schema.Column{NamespacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "workflow_name_namespace_workflows",
				Unique:  true,
				Columns: []*schema.Column{WorkflowsColumns[1], WorkflowsColumns[8]},
			},
		},
	}
	// WorkflowEventsColumns holds the columns for the "workflow_events" table.
	WorkflowEventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "events", Type: field.TypeJSON},
		{Name: "correlations", Type: field.TypeJSON},
		{Name: "signature", Type: field.TypeBytes, Nullable: true},
		{Name: "count", Type: field.TypeInt},
		{Name: "workflow_wfevents", Type: field.TypeUUID, Nullable: true},
		{Name: "workflow_instance_instance", Type: field.TypeInt, Nullable: true},
	}
	// WorkflowEventsTable holds the schema information for the "workflow_events" table.
	WorkflowEventsTable = &schema.Table{
		Name:       "workflow_events",
		Columns:    WorkflowEventsColumns,
		PrimaryKey: []*schema.Column{WorkflowEventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workflow_events_workflows_wfevents",
				Columns:    []*schema.Column{WorkflowEventsColumns[5]},
				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "workflow_events_workflow_instances_instance",
				Columns:    []*schema.Column{WorkflowEventsColumns[6]},
				RefColumns: []*schema.Column{WorkflowInstancesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// WorkflowEventsWaitsColumns holds the columns for the "workflow_events_waits" table.
	WorkflowEventsWaitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "events", Type: field.TypeJSON},
		{Name: "workflow_events_wfeventswait", Type: field.TypeInt, Nullable: true},
	}
	// WorkflowEventsWaitsTable holds the schema information for the "workflow_events_waits" table.
	WorkflowEventsWaitsTable = &schema.Table{
		Name:       "workflow_events_waits",
		Columns:    WorkflowEventsWaitsColumns,
		PrimaryKey: []*schema.Column{WorkflowEventsWaitsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workflow_events_waits_workflow_events_wfeventswait",
				Columns:    []*schema.Column{WorkflowEventsWaitsColumns[2]},
				RefColumns: []*schema.Column{WorkflowEventsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// WorkflowInstancesColumns holds the columns for the "workflow_instances" table.
	WorkflowInstancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "instance_id", Type: field.TypeString, Unique: true},
		{Name: "invoked_by", Type: field.TypeString},
		{Name: "status", Type: field.TypeString},
		{Name: "revision", Type: field.TypeInt},
		{Name: "begin_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime, Nullable: true},
		{Name: "flow", Type: field.TypeJSON, Nullable: true},
		{Name: "input", Type: field.TypeString},
		{Name: "output", Type: field.TypeString, Nullable: true},
		{Name: "state_data", Type: field.TypeString, Nullable: true},
		{Name: "memory", Type: field.TypeString, Nullable: true},
		{Name: "deadline", Type: field.TypeTime, Nullable: true},
		{Name: "attempts", Type: field.TypeInt, Nullable: true},
		{Name: "error_code", Type: field.TypeString, Nullable: true},
		{Name: "error_message", Type: field.TypeString, Nullable: true},
		{Name: "state_begin_time", Type: field.TypeTime, Nullable: true},
		{Name: "controller", Type: field.TypeString, Nullable: true},
		{Name: "workflow_instances", Type: field.TypeUUID, Nullable: true},
	}
	// WorkflowInstancesTable holds the schema information for the "workflow_instances" table.
	WorkflowInstancesTable = &schema.Table{
		Name:       "workflow_instances",
		Columns:    WorkflowInstancesColumns,
		PrimaryKey: []*schema.Column{WorkflowInstancesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workflow_instances_workflows_instances",
				Columns:    []*schema.Column{WorkflowInstancesColumns[18]},
				RefColumns: []*schema.Column{WorkflowsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		NamespacesTable,
		WorkflowsTable,
		WorkflowEventsTable,
		WorkflowEventsWaitsTable,
		WorkflowInstancesTable,
	}
)

func init() {
	WorkflowsTable.ForeignKeys[0].RefTable = NamespacesTable
	WorkflowEventsTable.ForeignKeys[0].RefTable = WorkflowsTable
	WorkflowEventsTable.ForeignKeys[1].RefTable = WorkflowInstancesTable
	WorkflowEventsWaitsTable.ForeignKeys[0].RefTable = WorkflowEventsTable
	WorkflowInstancesTable.ForeignKeys[0].RefTable = WorkflowsTable
}
