# coding: utf-8

# flake8: noqa

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


from __future__ import absolute_import

# import apis into sdk package
from flyteadmin.api.admin_service_api import AdminServiceApi

# import ApiClient
from flyteadmin.api_client import ApiClient
from flyteadmin.configuration import Configuration
# import models into sdk package
from flyteadmin.models.admin_abort_metadata import AdminAbortMetadata
from flyteadmin.models.admin_annotations import AdminAnnotations
from flyteadmin.models.admin_auth import AdminAuth
from flyteadmin.models.admin_auth_role import AdminAuthRole
from flyteadmin.models.admin_cluster_assignment import AdminClusterAssignment
from flyteadmin.models.admin_cluster_resource_attributes import AdminClusterResourceAttributes
from flyteadmin.models.admin_cron_schedule import AdminCronSchedule
from flyteadmin.models.admin_description import AdminDescription
from flyteadmin.models.admin_description_entity import AdminDescriptionEntity
from flyteadmin.models.admin_description_entity_list import AdminDescriptionEntityList
from flyteadmin.models.admin_description_format import AdminDescriptionFormat
from flyteadmin.models.admin_domain import AdminDomain
from flyteadmin.models.admin_email_notification import AdminEmailNotification
from flyteadmin.models.admin_envs import AdminEnvs
from flyteadmin.models.admin_execution import AdminExecution
from flyteadmin.models.admin_execution_closure import AdminExecutionClosure
from flyteadmin.models.admin_execution_cluster_label import AdminExecutionClusterLabel
from flyteadmin.models.admin_execution_create_request import AdminExecutionCreateRequest
from flyteadmin.models.admin_execution_create_response import AdminExecutionCreateResponse
from flyteadmin.models.admin_execution_list import AdminExecutionList
from flyteadmin.models.admin_execution_metadata import AdminExecutionMetadata
from flyteadmin.models.admin_execution_queue_attributes import AdminExecutionQueueAttributes
from flyteadmin.models.admin_execution_recover_request import AdminExecutionRecoverRequest
from flyteadmin.models.admin_execution_relaunch_request import AdminExecutionRelaunchRequest
from flyteadmin.models.admin_execution_spec import AdminExecutionSpec
from flyteadmin.models.admin_execution_state import AdminExecutionState
from flyteadmin.models.admin_execution_state_change_details import AdminExecutionStateChangeDetails
from flyteadmin.models.admin_execution_terminate_request import AdminExecutionTerminateRequest
from flyteadmin.models.admin_execution_terminate_response import AdminExecutionTerminateResponse
from flyteadmin.models.admin_execution_update_request import AdminExecutionUpdateRequest
from flyteadmin.models.admin_execution_update_response import AdminExecutionUpdateResponse
from flyteadmin.models.admin_fixed_rate import AdminFixedRate
from flyteadmin.models.admin_fixed_rate_unit import AdminFixedRateUnit
from flyteadmin.models.admin_flyte_ur_ls import AdminFlyteURLs
from flyteadmin.models.admin_get_version_response import AdminGetVersionResponse
from flyteadmin.models.admin_labels import AdminLabels
from flyteadmin.models.admin_launch_plan import AdminLaunchPlan
from flyteadmin.models.admin_launch_plan_closure import AdminLaunchPlanClosure
from flyteadmin.models.admin_launch_plan_create_request import AdminLaunchPlanCreateRequest
from flyteadmin.models.admin_launch_plan_create_response import AdminLaunchPlanCreateResponse
from flyteadmin.models.admin_launch_plan_list import AdminLaunchPlanList
from flyteadmin.models.admin_launch_plan_metadata import AdminLaunchPlanMetadata
from flyteadmin.models.admin_launch_plan_spec import AdminLaunchPlanSpec
from flyteadmin.models.admin_launch_plan_state import AdminLaunchPlanState
from flyteadmin.models.admin_launch_plan_update_request import AdminLaunchPlanUpdateRequest
from flyteadmin.models.admin_launch_plan_update_response import AdminLaunchPlanUpdateResponse
from flyteadmin.models.admin_list_matchable_attributes_response import AdminListMatchableAttributesResponse
from flyteadmin.models.admin_literal_map_blob import AdminLiteralMapBlob
from flyteadmin.models.admin_matchable_attributes_configuration import AdminMatchableAttributesConfiguration
from flyteadmin.models.admin_matchable_resource import AdminMatchableResource
from flyteadmin.models.admin_matching_attributes import AdminMatchingAttributes
from flyteadmin.models.admin_named_entity import AdminNamedEntity
from flyteadmin.models.admin_named_entity_identifier import AdminNamedEntityIdentifier
from flyteadmin.models.admin_named_entity_identifier_list import AdminNamedEntityIdentifierList
from flyteadmin.models.admin_named_entity_list import AdminNamedEntityList
from flyteadmin.models.admin_named_entity_metadata import AdminNamedEntityMetadata
from flyteadmin.models.admin_named_entity_state import AdminNamedEntityState
from flyteadmin.models.admin_named_entity_update_request import AdminNamedEntityUpdateRequest
from flyteadmin.models.admin_named_entity_update_response import AdminNamedEntityUpdateResponse
from flyteadmin.models.admin_node_execution_closure import AdminNodeExecutionClosure
from flyteadmin.models.admin_node_execution_event_request import AdminNodeExecutionEventRequest
from flyteadmin.models.admin_node_execution_event_response import AdminNodeExecutionEventResponse
from flyteadmin.models.admin_node_execution_get_data_response import AdminNodeExecutionGetDataResponse
from flyteadmin.models.admin_node_execution_list import AdminNodeExecutionList
from flyteadmin.models.admin_node_execution_meta_data import AdminNodeExecutionMetaData
from flyteadmin.models.admin_notification import AdminNotification
from flyteadmin.models.admin_notification_list import AdminNotificationList
from flyteadmin.models.admin_pager_duty_notification import AdminPagerDutyNotification
from flyteadmin.models.admin_plugin_override import AdminPluginOverride
from flyteadmin.models.admin_plugin_overrides import AdminPluginOverrides
from flyteadmin.models.admin_project import AdminProject
from flyteadmin.models.admin_project_attributes import AdminProjectAttributes
from flyteadmin.models.admin_project_attributes_delete_request import AdminProjectAttributesDeleteRequest
from flyteadmin.models.admin_project_attributes_delete_response import AdminProjectAttributesDeleteResponse
from flyteadmin.models.admin_project_attributes_get_response import AdminProjectAttributesGetResponse
from flyteadmin.models.admin_project_attributes_update_request import AdminProjectAttributesUpdateRequest
from flyteadmin.models.admin_project_attributes_update_response import AdminProjectAttributesUpdateResponse
from flyteadmin.models.admin_project_domain_attributes import AdminProjectDomainAttributes
from flyteadmin.models.admin_project_domain_attributes_delete_request import AdminProjectDomainAttributesDeleteRequest
from flyteadmin.models.admin_project_domain_attributes_delete_response import AdminProjectDomainAttributesDeleteResponse
from flyteadmin.models.admin_project_domain_attributes_get_response import AdminProjectDomainAttributesGetResponse
from flyteadmin.models.admin_project_domain_attributes_update_request import AdminProjectDomainAttributesUpdateRequest
from flyteadmin.models.admin_project_domain_attributes_update_response import AdminProjectDomainAttributesUpdateResponse
from flyteadmin.models.admin_project_identifier import AdminProjectIdentifier
from flyteadmin.models.admin_project_register_request import AdminProjectRegisterRequest
from flyteadmin.models.admin_project_register_response import AdminProjectRegisterResponse
from flyteadmin.models.admin_project_update_response import AdminProjectUpdateResponse
from flyteadmin.models.admin_projects import AdminProjects
from flyteadmin.models.admin_raw_output_data_config import AdminRawOutputDataConfig
from flyteadmin.models.admin_reason import AdminReason
from flyteadmin.models.admin_schedule import AdminSchedule
from flyteadmin.models.admin_slack_notification import AdminSlackNotification
from flyteadmin.models.admin_sort import AdminSort
from flyteadmin.models.admin_source_code import AdminSourceCode
from flyteadmin.models.admin_system_metadata import AdminSystemMetadata
from flyteadmin.models.admin_task import AdminTask
from flyteadmin.models.admin_task_closure import AdminTaskClosure
from flyteadmin.models.admin_task_execution_closure import AdminTaskExecutionClosure
from flyteadmin.models.admin_task_execution_event_request import AdminTaskExecutionEventRequest
from flyteadmin.models.admin_task_execution_event_response import AdminTaskExecutionEventResponse
from flyteadmin.models.admin_task_execution_get_data_response import AdminTaskExecutionGetDataResponse
from flyteadmin.models.admin_task_execution_list import AdminTaskExecutionList
from flyteadmin.models.admin_task_list import AdminTaskList
from flyteadmin.models.admin_task_resource_attributes import AdminTaskResourceAttributes
from flyteadmin.models.admin_task_resource_spec import AdminTaskResourceSpec
from flyteadmin.models.admin_task_spec import AdminTaskSpec
from flyteadmin.models.admin_url_blob import AdminUrlBlob
from flyteadmin.models.admin_version import AdminVersion
from flyteadmin.models.admin_workflow import AdminWorkflow
from flyteadmin.models.admin_workflow_attributes import AdminWorkflowAttributes
from flyteadmin.models.admin_workflow_attributes_delete_request import AdminWorkflowAttributesDeleteRequest
from flyteadmin.models.admin_workflow_attributes_delete_response import AdminWorkflowAttributesDeleteResponse
from flyteadmin.models.admin_workflow_attributes_get_response import AdminWorkflowAttributesGetResponse
from flyteadmin.models.admin_workflow_attributes_update_request import AdminWorkflowAttributesUpdateRequest
from flyteadmin.models.admin_workflow_attributes_update_response import AdminWorkflowAttributesUpdateResponse
from flyteadmin.models.admin_workflow_closure import AdminWorkflowClosure
from flyteadmin.models.admin_workflow_create_request import AdminWorkflowCreateRequest
from flyteadmin.models.admin_workflow_create_response import AdminWorkflowCreateResponse
from flyteadmin.models.admin_workflow_execution_config import AdminWorkflowExecutionConfig
from flyteadmin.models.admin_workflow_execution_event_request import AdminWorkflowExecutionEventRequest
from flyteadmin.models.admin_workflow_execution_event_response import AdminWorkflowExecutionEventResponse
from flyteadmin.models.admin_workflow_execution_get_data_response import AdminWorkflowExecutionGetDataResponse
from flyteadmin.models.admin_workflow_execution_get_metrics_response import AdminWorkflowExecutionGetMetricsResponse
from flyteadmin.models.admin_workflow_list import AdminWorkflowList
from flyteadmin.models.admin_workflow_spec import AdminWorkflowSpec
from flyteadmin.models.blob_type_blob_dimensionality import BlobTypeBlobDimensionality
from flyteadmin.models.catalog_reservation_status import CatalogReservationStatus
from flyteadmin.models.comparison_expression_operator import ComparisonExpressionOperator
from flyteadmin.models.conjunction_expression_logical_operator import ConjunctionExpressionLogicalOperator
from flyteadmin.models.connection_set_id_list import ConnectionSetIdList
from flyteadmin.models.container_architecture import ContainerArchitecture
from flyteadmin.models.core_alias import CoreAlias
from flyteadmin.models.core_approve_condition import CoreApproveCondition
from flyteadmin.models.core_array_node import CoreArrayNode
from flyteadmin.models.core_artifact_binding_data import CoreArtifactBindingData
from flyteadmin.models.core_artifact_id import CoreArtifactID
from flyteadmin.models.core_artifact_key import CoreArtifactKey
from flyteadmin.models.core_artifact_query import CoreArtifactQuery
from flyteadmin.models.core_artifact_tag import CoreArtifactTag
from flyteadmin.models.core_binary import CoreBinary
from flyteadmin.models.core_binding import CoreBinding
from flyteadmin.models.core_binding_data import CoreBindingData
from flyteadmin.models.core_binding_data_collection import CoreBindingDataCollection
from flyteadmin.models.core_binding_data_map import CoreBindingDataMap
from flyteadmin.models.core_blob import CoreBlob
from flyteadmin.models.core_blob_metadata import CoreBlobMetadata
from flyteadmin.models.core_blob_type import CoreBlobType
from flyteadmin.models.core_boolean_expression import CoreBooleanExpression
from flyteadmin.models.core_branch_node import CoreBranchNode
from flyteadmin.models.core_catalog_artifact_tag import CoreCatalogArtifactTag
from flyteadmin.models.core_catalog_cache_status import CoreCatalogCacheStatus
from flyteadmin.models.core_catalog_metadata import CoreCatalogMetadata
from flyteadmin.models.core_comparison_expression import CoreComparisonExpression
from flyteadmin.models.core_compiled_task import CoreCompiledTask
from flyteadmin.models.core_compiled_workflow import CoreCompiledWorkflow
from flyteadmin.models.core_compiled_workflow_closure import CoreCompiledWorkflowClosure
from flyteadmin.models.core_conjunction_expression import CoreConjunctionExpression
from flyteadmin.models.core_connection_set import CoreConnectionSet
from flyteadmin.models.core_container import CoreContainer
from flyteadmin.models.core_container_port import CoreContainerPort
from flyteadmin.models.core_data_loading_config import CoreDataLoadingConfig
from flyteadmin.models.core_enum_type import CoreEnumType
from flyteadmin.models.core_error import CoreError
from flyteadmin.models.core_execution_error import CoreExecutionError
from flyteadmin.models.core_extended_resources import CoreExtendedResources
from flyteadmin.models.core_gpu_accelerator import CoreGPUAccelerator
from flyteadmin.models.core_gate_node import CoreGateNode
from flyteadmin.models.core_io_strategy import CoreIOStrategy
from flyteadmin.models.core_identifier import CoreIdentifier
from flyteadmin.models.core_identity import CoreIdentity
from flyteadmin.models.core_if_block import CoreIfBlock
from flyteadmin.models.core_if_else_block import CoreIfElseBlock
from flyteadmin.models.core_input_binding_data import CoreInputBindingData
from flyteadmin.models.core_k8s_object_metadata import CoreK8sObjectMetadata
from flyteadmin.models.core_k8s_pod import CoreK8sPod
from flyteadmin.models.core_key_value_pair import CoreKeyValuePair
from flyteadmin.models.core_label_value import CoreLabelValue
from flyteadmin.models.core_literal import CoreLiteral
from flyteadmin.models.core_literal_collection import CoreLiteralCollection
from flyteadmin.models.core_literal_map import CoreLiteralMap
from flyteadmin.models.core_literal_type import CoreLiteralType
from flyteadmin.models.core_node import CoreNode
from flyteadmin.models.core_node_execution_identifier import CoreNodeExecutionIdentifier
from flyteadmin.models.core_node_execution_phase import CoreNodeExecutionPhase
from flyteadmin.models.core_node_metadata import CoreNodeMetadata
from flyteadmin.models.core_o_auth2_client import CoreOAuth2Client
from flyteadmin.models.core_o_auth2_token_request import CoreOAuth2TokenRequest
from flyteadmin.models.core_o_auth2_token_request_type import CoreOAuth2TokenRequestType
from flyteadmin.models.core_operand import CoreOperand
from flyteadmin.models.core_output_reference import CoreOutputReference
from flyteadmin.models.core_parameter import CoreParameter
from flyteadmin.models.core_parameter_map import CoreParameterMap
from flyteadmin.models.core_partitions import CorePartitions
from flyteadmin.models.core_primitive import CorePrimitive
from flyteadmin.models.core_promise_attribute import CorePromiseAttribute
from flyteadmin.models.core_quality_of_service import CoreQualityOfService
from flyteadmin.models.core_quality_of_service_spec import CoreQualityOfServiceSpec
from flyteadmin.models.core_resource_type import CoreResourceType
from flyteadmin.models.core_resources import CoreResources
from flyteadmin.models.core_retry_strategy import CoreRetryStrategy
from flyteadmin.models.core_runtime_metadata import CoreRuntimeMetadata
from flyteadmin.models.core_scalar import CoreScalar
from flyteadmin.models.core_schema import CoreSchema
from flyteadmin.models.core_schema_type import CoreSchemaType
from flyteadmin.models.core_secret import CoreSecret
from flyteadmin.models.core_security_context import CoreSecurityContext
from flyteadmin.models.core_signal_condition import CoreSignalCondition
from flyteadmin.models.core_simple_type import CoreSimpleType
from flyteadmin.models.core_sleep_condition import CoreSleepCondition
from flyteadmin.models.core_span import CoreSpan
from flyteadmin.models.core_sql import CoreSql
from flyteadmin.models.core_structured_dataset import CoreStructuredDataset
from flyteadmin.models.core_structured_dataset_metadata import CoreStructuredDatasetMetadata
from flyteadmin.models.core_structured_dataset_type import CoreStructuredDatasetType
from flyteadmin.models.core_task_execution_identifier import CoreTaskExecutionIdentifier
from flyteadmin.models.core_task_execution_phase import CoreTaskExecutionPhase
from flyteadmin.models.core_task_log import CoreTaskLog
from flyteadmin.models.core_task_metadata import CoreTaskMetadata
from flyteadmin.models.core_task_node import CoreTaskNode
from flyteadmin.models.core_task_node_overrides import CoreTaskNodeOverrides
from flyteadmin.models.core_task_template import CoreTaskTemplate
from flyteadmin.models.core_type_annotation import CoreTypeAnnotation
from flyteadmin.models.core_type_structure import CoreTypeStructure
from flyteadmin.models.core_typed_interface import CoreTypedInterface
from flyteadmin.models.core_union import CoreUnion
from flyteadmin.models.core_union_info import CoreUnionInfo
from flyteadmin.models.core_union_type import CoreUnionType
from flyteadmin.models.core_variable import CoreVariable
from flyteadmin.models.core_variable_map import CoreVariableMap
from flyteadmin.models.core_void import CoreVoid
from flyteadmin.models.core_workflow_execution_identifier import CoreWorkflowExecutionIdentifier
from flyteadmin.models.core_workflow_execution_phase import CoreWorkflowExecutionPhase
from flyteadmin.models.core_workflow_metadata import CoreWorkflowMetadata
from flyteadmin.models.core_workflow_metadata_defaults import CoreWorkflowMetadataDefaults
from flyteadmin.models.core_workflow_node import CoreWorkflowNode
from flyteadmin.models.core_workflow_template import CoreWorkflowTemplate
from flyteadmin.models.data_loading_config_literal_map_format import DataLoadingConfigLiteralMapFormat
from flyteadmin.models.event_event_reason import EventEventReason
from flyteadmin.models.event_external_resource_info import EventExternalResourceInfo
from flyteadmin.models.event_node_execution_event import EventNodeExecutionEvent
from flyteadmin.models.event_parent_node_execution_metadata import EventParentNodeExecutionMetadata
from flyteadmin.models.event_parent_task_execution_metadata import EventParentTaskExecutionMetadata
from flyteadmin.models.event_resource_pool_info import EventResourcePoolInfo
from flyteadmin.models.event_task_execution_event import EventTaskExecutionEvent
from flyteadmin.models.event_workflow_execution_event import EventWorkflowExecutionEvent
from flyteadmin.models.execution_error_error_kind import ExecutionErrorErrorKind
from flyteadmin.models.execution_metadata_execution_mode import ExecutionMetadataExecutionMode
from flyteadmin.models.flyteidladmin_dynamic_workflow_node_metadata import FlyteidladminDynamicWorkflowNodeMetadata
from flyteadmin.models.flyteidladmin_node_execution import FlyteidladminNodeExecution
from flyteadmin.models.flyteidladmin_task_create_request import FlyteidladminTaskCreateRequest
from flyteadmin.models.flyteidladmin_task_create_response import FlyteidladminTaskCreateResponse
from flyteadmin.models.flyteidladmin_task_execution import FlyteidladminTaskExecution
from flyteadmin.models.flyteidladmin_task_node_metadata import FlyteidladminTaskNodeMetadata
from flyteadmin.models.flyteidladmin_workflow_node_metadata import FlyteidladminWorkflowNodeMetadata
from flyteadmin.models.flyteidlevent_dynamic_workflow_node_metadata import FlyteidleventDynamicWorkflowNodeMetadata
from flyteadmin.models.flyteidlevent_task_execution_metadata import FlyteidleventTaskExecutionMetadata
from flyteadmin.models.flyteidlevent_task_node_metadata import FlyteidleventTaskNodeMetadata
from flyteadmin.models.flyteidlevent_workflow_node_metadata import FlyteidleventWorkflowNodeMetadata
from flyteadmin.models.io_strategy_download_mode import IOStrategyDownloadMode
from flyteadmin.models.io_strategy_upload_mode import IOStrategyUploadMode
from flyteadmin.models.plugin_override_missing_plugin_behavior import PluginOverrideMissingPluginBehavior
from flyteadmin.models.project_project_state import ProjectProjectState
from flyteadmin.models.protobuf_any import ProtobufAny
from flyteadmin.models.protobuf_list_value import ProtobufListValue
from flyteadmin.models.protobuf_null_value import ProtobufNullValue
from flyteadmin.models.protobuf_struct import ProtobufStruct
from flyteadmin.models.protobuf_value import ProtobufValue
from flyteadmin.models.quality_of_service_tier import QualityOfServiceTier
from flyteadmin.models.resources_resource_entry import ResourcesResourceEntry
from flyteadmin.models.resources_resource_name import ResourcesResourceName
from flyteadmin.models.runtime_metadata_runtime_type import RuntimeMetadataRuntimeType
from flyteadmin.models.schema_column_schema_column_type import SchemaColumnSchemaColumnType
from flyteadmin.models.schema_type_schema_column import SchemaTypeSchemaColumn
from flyteadmin.models.secret_mount_type import SecretMountType
from flyteadmin.models.sort_direction import SortDirection
from flyteadmin.models.sql_dialect import SqlDialect
from flyteadmin.models.structured_dataset_type_dataset_column import StructuredDatasetTypeDatasetColumn
from flyteadmin.models.task_execution_metadata_instance_class import TaskExecutionMetadataInstanceClass
from flyteadmin.models.task_log_message_format import TaskLogMessageFormat
from flyteadmin.models.workflow_metadata_on_failure_policy import WorkflowMetadataOnFailurePolicy
