/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Metadata for the entire workflow. To be used in the future.
type CoreWorkflowMetadata struct {
	// Total wait time a workflow can be delayed by queueing.
	QueuingBudget string `json:"queuing_budget,omitempty"`
}
