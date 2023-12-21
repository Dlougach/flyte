// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: flyteidl/service/admin.proto

#include "flyteidl/service/admin.pb.h"

#include <algorithm>

#include <google/protobuf/stubs/common.h>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/wire_format_lite_inl.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>

namespace flyteidl {
namespace service {
}  // namespace service
}  // namespace flyteidl
void InitDefaults_flyteidl_2fservice_2fadmin_2eproto() {
}

constexpr ::google::protobuf::Metadata* file_level_metadata_flyteidl_2fservice_2fadmin_2eproto = nullptr;
constexpr ::google::protobuf::EnumDescriptor const** file_level_enum_descriptors_flyteidl_2fservice_2fadmin_2eproto = nullptr;
constexpr ::google::protobuf::ServiceDescriptor const** file_level_service_descriptors_flyteidl_2fservice_2fadmin_2eproto = nullptr;
const ::google::protobuf::uint32 TableStruct_flyteidl_2fservice_2fadmin_2eproto::offsets[1] = {};
static constexpr ::google::protobuf::internal::MigrationSchema* schemas = nullptr;
static constexpr ::google::protobuf::Message* const* file_default_instances = nullptr;

::google::protobuf::internal::AssignDescriptorsTable assign_descriptors_table_flyteidl_2fservice_2fadmin_2eproto = {
  {}, AddDescriptors_flyteidl_2fservice_2fadmin_2eproto, "flyteidl/service/admin.proto", schemas,
  file_default_instances, TableStruct_flyteidl_2fservice_2fadmin_2eproto::offsets,
  file_level_metadata_flyteidl_2fservice_2fadmin_2eproto, 0, file_level_enum_descriptors_flyteidl_2fservice_2fadmin_2eproto, file_level_service_descriptors_flyteidl_2fservice_2fadmin_2eproto,
};

const char descriptor_table_protodef_flyteidl_2fservice_2fadmin_2eproto[] =
  "\n\034flyteidl/service/admin.proto\022\020flyteidl"
  ".service\032\034google/api/annotations.proto\032\034"
  "flyteidl/admin/project.proto\032.flyteidl/a"
  "dmin/project_domain_attributes.proto\032\'fl"
  "yteidl/admin/project_attributes.proto\032\031f"
  "lyteidl/admin/task.proto\032\035flyteidl/admin"
  "/workflow.proto\032(flyteidl/admin/workflow"
  "_attributes.proto\032 flyteidl/admin/launch"
  "_plan.proto\032\032flyteidl/admin/event.proto\032"
  "\036flyteidl/admin/execution.proto\032\'flyteid"
  "l/admin/matchable_resource.proto\032#flytei"
  "dl/admin/node_execution.proto\032#flyteidl/"
  "admin/task_execution.proto\032\034flyteidl/adm"
  "in/version.proto\032\033flyteidl/admin/common."
  "proto\032\'flyteidl/admin/description_entity"
  ".proto2\361p\n\014AdminService\022m\n\nCreateTask\022!."
  "flyteidl.admin.TaskCreateRequest\032\".flyte"
  "idl.admin.TaskCreateResponse\"\030\202\323\344\223\002\022\"\r/a"
  "pi/v1/tasks:\001*\022\324\001\n\007GetTask\022 .flyteidl.ad"
  "min.ObjectGetRequest\032\024.flyteidl.admin.Ta"
  "sk\"\220\001\202\323\344\223\002\211\001\022=/api/v1/tasks/{id.project}"
  "/{id.domain}/{id.name}/{id.version}ZH\022F/"
  "api/v1/tasks/{id.org}/{id.project}/{id.d"
  "omain}/{id.name}/{id.version}\022\301\001\n\013ListTa"
  "skIds\0220.flyteidl.admin.NamedEntityIdenti"
  "fierListRequest\032).flyteidl.admin.NamedEn"
  "tityIdentifierList\"U\202\323\344\223\002O\022#/api/v1/task"
  "_ids/{project}/{domain}Z(\022&/api/v1/tasks"
  "/{org}/{project}/{domain}\022\240\002\n\tListTasks\022"
  "#.flyteidl.admin.ResourceListRequest\032\030.f"
  "lyteidl.admin.TaskList\"\323\001\202\323\344\223\002\314\001\0220/api/v"
  "1/tasks/{id.project}/{id.domain}/{id.nam"
  "e}Z;\0229/api/v1/tasks/{id.org}/{id.project"
  "}/{id.domain}/{id.name}Z(\022&/api/v1/tasks"
  "/{id.project}/{id.domain}Z1\022//api/v1/tas"
  "ks/{id.org}/{id.project}/{id.domain}\022}\n\016"
  "CreateWorkflow\022%.flyteidl.admin.Workflow"
  "CreateRequest\032&.flyteidl.admin.WorkflowC"
  "reateResponse\"\034\202\323\344\223\002\026\"\021/api/v1/workflows"
  ":\001*\022\344\001\n\013GetWorkflow\022 .flyteidl.admin.Obj"
  "ectGetRequest\032\030.flyteidl.admin.Workflow\""
  "\230\001\202\323\344\223\002\221\001\022A/api/v1/workflows/{id.project"
  "}/{id.domain}/{id.name}/{id.version}ZL\022J"
  "/api/v1/workflows/{id.org}/{id.project}/"
  "{id.domain}/{id.name}/{id.version}\022\315\001\n\017L"
  "istWorkflowIds\0220.flyteidl.admin.NamedEnt"
  "ityIdentifierListRequest\032).flyteidl.admi"
  "n.NamedEntityIdentifierList\"]\202\323\344\223\002W\022\'/ap"
  "i/v1/workflow_ids/{project}/{domain}Z,\022*"
  "/api/v1/workflows/{org}/{project}/{domai"
  "n}\022\270\002\n\rListWorkflows\022#.flyteidl.admin.Re"
  "sourceListRequest\032\034.flyteidl.admin.Workf"
  "lowList\"\343\001\202\323\344\223\002\334\001\0224/api/v1/workflows/{id"
  ".project}/{id.domain}/{id.name}Z\?\022=/api/"
  "v1/workflows/{id.org}/{id.project}/{id.d"
  "omain}/{id.name}Z,\022*/api/v1/workflows/{i"
  "d.project}/{id.domain}Z5\0223/api/v1/workfl"
  "ows/{id.org}/{id.project}/{id.domain}\022\206\001"
  "\n\020CreateLaunchPlan\022\'.flyteidl.admin.Laun"
  "chPlanCreateRequest\032(.flyteidl.admin.Lau"
  "nchPlanCreateResponse\"\037\202\323\344\223\002\031\"\024/api/v1/l"
  "aunch_plans:\001*\022\356\001\n\rGetLaunchPlan\022 .flyte"
  "idl.admin.ObjectGetRequest\032\032.flyteidl.ad"
  "min.LaunchPlan\"\236\001\202\323\344\223\002\227\001\022D/api/v1/launch"
  "_plans/{id.project}/{id.domain}/{id.name"
  "}/{id.version}ZO\022M/api/v1/launch_plans/{"
  "id.org}/{id.project}/{id.domain}/{id.nam"
  "e}/{id.version}\022\357\001\n\023GetActiveLaunchPlan\022"
  "\'.flyteidl.admin.ActiveLaunchPlanRequest"
  "\032\032.flyteidl.admin.LaunchPlan\"\222\001\202\323\344\223\002\213\001\022>"
  "/api/v1/active_launch_plans/{id.project}"
  "/{id.domain}/{id.name}ZI\022G/api/v1/active"
  "_launch_plans/{id.org}/{id.project}/{id."
  "domain}/{id.name}\022\234\001\n\025ListActiveLaunchPl"
  "ans\022+.flyteidl.admin.ActiveLaunchPlanLis"
  "tRequest\032\036.flyteidl.admin.LaunchPlanList"
  "\"6\202\323\344\223\0020\022./api/v1/active_launch_plans/{p"
  "roject}/{domain}\022\330\001\n\021ListLaunchPlanIds\0220"
  ".flyteidl.admin.NamedEntityIdentifierLis"
  "tRequest\032).flyteidl.admin.NamedEntityIde"
  "ntifierList\"f\202\323\344\223\002`\022*/api/v1/launch_plan"
  "_ids/{project}/{domain}Z2\0220/api/v1/launc"
  "h_plan_ids/{org}/{project}/{domain}\022\310\002\n\017"
  "ListLaunchPlans\022#.flyteidl.admin.Resourc"
  "eListRequest\032\036.flyteidl.admin.LaunchPlan"
  "List\"\357\001\202\323\344\223\002\350\001\0227/api/v1/launch_plans/{id"
  ".project}/{id.domain}/{id.name}ZB\022@/api/"
  "v1/launch_plans/{id.org}/{id.project}/{i"
  "d.domain}/{id.name}Z/\022-/api/v1/launch_pl"
  "ans/{id.project}/{id.domain}Z8\0226/api/v1/"
  "launch_plans/{id.org}/{id.project}/{id.d"
  "omain}\022\211\002\n\020UpdateLaunchPlan\022\'.flyteidl.a"
  "dmin.LaunchPlanUpdateRequest\032(.flyteidl."
  "admin.LaunchPlanUpdateResponse\"\241\001\202\323\344\223\002\232\001"
  "\032D/api/v1/launch_plans/{id.project}/{id."
  "domain}/{id.name}/{id.version}:\001*ZO\032M/ap"
  "i/v1/launch_plans/{id.org}/{id.project}/"
  "{id.domain}/{id.name}/{id.version}\022\201\001\n\017C"
  "reateExecution\022&.flyteidl.admin.Executio"
  "nCreateRequest\032\'.flyteidl.admin.Executio"
  "nCreateResponse\"\035\202\323\344\223\002\027\"\022/api/v1/executi"
  "ons:\001*\022\216\001\n\021RelaunchExecution\022(.flyteidl."
  "admin.ExecutionRelaunchRequest\032\'.flyteid"
  "l.admin.ExecutionCreateResponse\"&\202\323\344\223\002 \""
  "\033/api/v1/executions/relaunch:\001*\022\213\001\n\020Reco"
  "verExecution\022\'.flyteidl.admin.ExecutionR"
  "ecoverRequest\032\'.flyteidl.admin.Execution"
  "CreateResponse\"%\202\323\344\223\002\037\"\032/api/v1/executio"
  "ns/recover:\001*\022\327\001\n\014GetExecution\022+.flyteid"
  "l.admin.WorkflowExecutionGetRequest\032\031.fl"
  "yteidl.admin.Execution\"\177\202\323\344\223\002y\0225/api/v1/"
  "executions/{id.project}/{id.domain}/{id."
  "name}Z@\022>/api/v1/executions/{id.org}/{id"
  ".project}/{id.domain}/{id.name}\022\352\001\n\017Upda"
  "teExecution\022&.flyteidl.admin.ExecutionUp"
  "dateRequest\032\'.flyteidl.admin.ExecutionUp"
  "dateResponse\"\205\001\202\323\344\223\002\177\0325/api/v1/execution"
  "s/{id.project}/{id.domain}/{id.name}:\001*Z"
  "C\032>/api/v1/executions/{id.org}/{id.proje"
  "ct}/{id.domain}/{id.name}:\001*\022\202\002\n\020GetExec"
  "utionData\022/.flyteidl.admin.WorkflowExecu"
  "tionGetDataRequest\0320.flyteidl.admin.Work"
  "flowExecutionGetDataResponse\"\212\001\202\323\344\223\002\203\001\022:"
  "/api/v1/data/executions/{id.project}/{id"
  ".domain}/{id.name}ZE\022C/api/v1/data/execu"
  "tions/{id.org}/{id.project}/{id.domain}/"
  "{id.name}\022\301\001\n\016ListExecutions\022#.flyteidl."
  "admin.ResourceListRequest\032\035.flyteidl.adm"
  "in.ExecutionList\"k\202\323\344\223\002e\022+/api/v1/execut"
  "ions/{id.project}/{id.domain}Z6\0224/api/v1"
  "/executions/{id.org}/{id.project}/{id.do"
  "main}\022\371\001\n\022TerminateExecution\022).flyteidl."
  "admin.ExecutionTerminateRequest\032*.flytei"
  "dl.admin.ExecutionTerminateResponse\"\213\001\202\323"
  "\344\223\002\204\001*5/api/v1/executions/{id.project}/{"
  "id.domain}/{id.name}:\001*ZH*C/api/v1/data/"
  "executions/{id.org}/{id.project}/{id.dom"
  "ain}/{id.name}:\001*\022\336\002\n\020GetNodeExecution\022\'"
  ".flyteidl.admin.NodeExecutionGetRequest\032"
  "\035.flyteidl.admin.NodeExecution\"\201\002\202\323\344\223\002\372\001"
  "\022n/api/v1/node_executions/{id.execution_"
  "id.project}/{id.execution_id.domain}/{id"
  ".execution_id.name}/{id.node_id}Z\207\001\022\204\001/a"
  "pi/v1/node_executions/{id.execution_id.o"
  "rg}/{id.execution_id.project}/{id.execut"
  "ion_id.domain}/{id.execution_id.name}/{i"
  "d.node_id}\022\365\002\n\022ListNodeExecutions\022(.flyt"
  "eidl.admin.NodeExecutionListRequest\032!.fl"
  "yteidl.admin.NodeExecutionList\"\221\002\202\323\344\223\002\212\002"
  "\022s/api/v1/node_executions/{workflow_exec"
  "ution_id.project}/{workflow_execution_id"
  ".domain}/{workflow_execution_id.name}Z\222\001"
  "\022\217\001/api/v1/node_executions/{workflow_exe"
  "cution_id.org}/{workflow_execution_id.pr"
  "oject}/{workflow_execution_id.domain}/{w"
  "orkflow_execution_id.name}\022\213\010\n\031ListNodeE"
  "xecutionsForTask\022/.flyteidl.admin.NodeEx"
  "ecutionForTaskListRequest\032!.flyteidl.adm"
  "in.NodeExecutionList\"\231\007\202\323\344\223\002\222\007\022\251\003/api/v1"
  "/children/task_executions/{task_executio"
  "n_id.node_execution_id.execution_id.proj"
  "ect}/{task_execution_id.node_execution_i"
  "d.execution_id.domain}/{task_execution_i"
  "d.node_execution_id.execution_id.name}/{"
  "task_execution_id.node_execution_id.node"
  "_id}/{task_execution_id.task_id.project}"
  "/{task_execution_id.task_id.domain}/{tas"
  "k_execution_id.task_id.name}/{task_execu"
  "tion_id.task_id.version}/{task_execution"
  "_id.retry_attempt}Z\343\003\022\340\003/api/v1/children"
  "/task_executions/{task_execution_id.node"
  "_execution_id.execution_id.org}/{task_ex"
  "ecution_id.node_execution_id.execution_i"
  "d.project}/{task_execution_id.node_execu"
  "tion_id.execution_id.domain}/{task_execu"
  "tion_id.node_execution_id.execution_id.n"
  "ame}/{task_execution_id.node_execution_i"
  "d.node_id}/{task_execution_id.task_id.pr"
  "oject}/{task_execution_id.task_id.domain"
  "}/{task_execution_id.task_id.name}/{task"
  "_execution_id.task_id.version}/{task_exe"
  "cution_id.retry_attempt}\022\377\002\n\024GetNodeExec"
  "utionData\022+.flyteidl.admin.NodeExecution"
  "GetDataRequest\032,.flyteidl.admin.NodeExec"
  "utionGetDataResponse\"\213\002\202\323\344\223\002\204\002\022s/api/v1/"
  "data/node_executions/{id.execution_id.pr"
  "oject}/{id.execution_id.domain}/{id.exec"
  "ution_id.name}/{id.node_id}Z\214\001\022\211\001/api/v1"
  "/data/node_executions/{id.execution_id.o"
  "rg}/{id.execution_id.project}/{id.execut"
  "ion_id.domain}/{id.execution_id.name}/{i"
  "d.node_id}\022\177\n\017RegisterProject\022&.flyteidl"
  ".admin.ProjectRegisterRequest\032\'.flyteidl"
  ".admin.ProjectRegisterResponse\"\033\202\323\344\223\002\025\"\020"
  "/api/v1/projects:\001*\022\377\001\n\rUpdateProject\022\027."
  "flyteidl.admin.Project\032%.flyteidl.admin."
  "ProjectUpdateResponse\"\255\001\202\323\344\223\002\246\001\032\025/api/v1"
  "/projects/{id}:\001*Z%\032 /api/v1/projects/{i"
  "dentifier.id}:\001*Z+\032&/api/v1/projects/{id"
  "entifier.org}/{id}:\001*Z6\0321/api/v1/project"
  "s/{identifier.org}/{identifier.id}:\001*\022f\n"
  "\014ListProjects\022\".flyteidl.admin.ProjectLi"
  "stRequest\032\030.flyteidl.admin.Projects\"\030\202\323\344"
  "\223\002\022\022\020/api/v1/projects\022\231\001\n\023CreateWorkflow"
  "Event\022-.flyteidl.admin.WorkflowExecution"
  "EventRequest\032..flyteidl.admin.WorkflowEx"
  "ecutionEventResponse\"#\202\323\344\223\002\035\"\030/api/v1/ev"
  "ents/workflows:\001*\022\211\001\n\017CreateNodeEvent\022)."
  "flyteidl.admin.NodeExecutionEventRequest"
  "\032*.flyteidl.admin.NodeExecutionEventResp"
  "onse\"\037\202\323\344\223\002\031\"\024/api/v1/events/nodes:\001*\022\211\001"
  "\n\017CreateTaskEvent\022).flyteidl.admin.TaskE"
  "xecutionEventRequest\032*.flyteidl.admin.Ta"
  "skExecutionEventResponse\"\037\202\323\344\223\002\031\"\024/api/v"
  "1/events/tasks:\001*\022\307\005\n\020GetTaskExecution\022\'"
  ".flyteidl.admin.TaskExecutionGetRequest\032"
  "\035.flyteidl.admin.TaskExecution\"\352\004\202\323\344\223\002\343\004"
  "\022\231\002/api/v1/task_executions/{id.node_exec"
  "ution_id.execution_id.project}/{id.node_"
  "execution_id.execution_id.domain}/{id.no"
  "de_execution_id.execution_id.name}/{id.n"
  "ode_execution_id.node_id}/{id.task_id.pr"
  "oject}/{id.task_id.domain}/{id.task_id.n"
  "ame}/{id.task_id.version}/{id.retry_atte"
  "mpt}Z\304\002\022\301\002/api/v1/task_executions/{id.no"
  "de_execution_id.execution_id.org}/{id.no"
  "de_execution_id.execution_id.project}/{i"
  "d.node_execution_id.execution_id.domain}"
  "/{id.node_execution_id.execution_id.name"
  "}/{id.node_execution_id.node_id}/{id.tas"
  "k_id.project}/{id.task_id.domain}/{id.ta"
  "sk_id.name}/{id.task_id.version}/{id.ret"
  "ry_attempt}\022\355\003\n\022ListTaskExecutions\022(.fly"
  "teidl.admin.TaskExecutionListRequest\032!.f"
  "lyteidl.admin.TaskExecutionList\"\211\003\202\323\344\223\002\202"
  "\003\022\252\001/api/v1/task_executions/{node_execut"
  "ion_id.execution_id.project}/{node_execu"
  "tion_id.execution_id.domain}/{node_execu"
  "tion_id.execution_id.name}/{node_executi"
  "on_id.node_id}Z\322\001\022\317\001/api/v1/task_executi"
  "ons/{node_execution_id.execution_id.org}"
  "/{node_execution_id.execution_id.project"
  "}/{node_execution_id.execution_id.domain"
  "}/{node_execution_id.execution_id.name}/"
  "{node_execution_id.node_id}\022\350\005\n\024GetTaskE"
  "xecutionData\022+.flyteidl.admin.TaskExecut"
  "ionGetDataRequest\032,.flyteidl.admin.TaskE"
  "xecutionGetDataResponse\"\364\004\202\323\344\223\002\355\004\022\236\002/api"
  "/v1/data/task_executions/{id.node_execut"
  "ion_id.execution_id.project}/{id.node_ex"
  "ecution_id.execution_id.domain}/{id.node"
  "_execution_id.execution_id.name}/{id.nod"
  "e_execution_id.node_id}/{id.task_id.proj"
  "ect}/{id.task_id.domain}/{id.task_id.nam"
  "e}/{id.task_id.version}/{id.retry_attemp"
  "t}Z\311\002\022\306\002/api/v1/data/task_executions/{id"
  ".node_execution_id.execution_id.org}/{id"
  ".node_execution_id.execution_id.project}"
  "/{id.node_execution_id.execution_id.doma"
  "in}/{id.node_execution_id.execution_id.n"
  "ame}/{id.node_execution_id.node_id}/{id."
  "task_id.project}/{id.task_id.domain}/{id"
  ".task_id.name}/{id.task_id.version}/{id."
  "retry_attempt}\022\307\002\n\035UpdateProjectDomainAt"
  "tributes\0224.flyteidl.admin.ProjectDomainA"
  "ttributesUpdateRequest\0325.flyteidl.admin."
  "ProjectDomainAttributesUpdateResponse\"\270\001"
  "\202\323\344\223\002\261\001\032J/api/v1/project_domain_attribut"
  "es/{attributes.project}/{attributes.doma"
  "in}:\001*Z`\032[/api/v1/project_domain_attribu"
  "tes/{attributes.org}/{attributes.project"
  "}/{attributes.domain}:\001*\022\377\001\n\032GetProjectD"
  "omainAttributes\0221.flyteidl.admin.Project"
  "DomainAttributesGetRequest\0322.flyteidl.ad"
  "min.ProjectDomainAttributesGetResponse\"z"
  "\202\323\344\223\002t\0224/api/v1/project_domain_attribute"
  "s/{project}/{domain}Z<\022:/api/v1/project_"
  "domain_attributes/{org}/{project}/{domai"
  "n}\022\217\002\n\035DeleteProjectDomainAttributes\0224.f"
  "lyteidl.admin.ProjectDomainAttributesDel"
  "eteRequest\0325.flyteidl.admin.ProjectDomai"
  "nAttributesDeleteResponse\"\200\001\202\323\344\223\002z*4/api"
  "/v1/project_domain_attributes/{project}/"
  "{domain}:\001*Z\?*:/api/v1/project_domain_at"
  "tributes/{org}/{project}/{domain}:\001*\022\206\002\n"
  "\027UpdateProjectAttributes\022..flyteidl.admi"
  "n.ProjectAttributesUpdateRequest\032/.flyte"
  "idl.admin.ProjectAttributesUpdateRespons"
  "e\"\211\001\202\323\344\223\002\202\001\032//api/v1/project_attributes/"
  "{attributes.project}:\001*ZL\032G/api/v1/proje"
  "ct_domain_attributes/{attributes.org}/{a"
  "ttributes.project}:\001*\022\324\001\n\024GetProjectAttr"
  "ibutes\022+.flyteidl.admin.ProjectAttribute"
  "sGetRequest\032,.flyteidl.admin.ProjectAttr"
  "ibutesGetResponse\"a\202\323\344\223\002[\022$/api/v1/proje"
  "ct_attributes/{project}Z3\0221/api/v1/proje"
  "ct_domain_attributes/{org}/{project}\022\343\001\n"
  "\027DeleteProjectAttributes\022..flyteidl.admi"
  "n.ProjectAttributesDeleteRequest\032/.flyte"
  "idl.admin.ProjectAttributesDeleteRespons"
  "e\"g\202\323\344\223\002a*$/api/v1/project_attributes/{p"
  "roject}:\001*Z6*1/api/v1/project_domain_att"
  "ributes/{org}/{project}:\001*\022\330\002\n\030UpdateWor"
  "kflowAttributes\022/.flyteidl.admin.Workflo"
  "wAttributesUpdateRequest\0320.flyteidl.admi"
  "n.WorkflowAttributesUpdateResponse\"\330\001\202\323\344"
  "\223\002\321\001\032Z/api/v1/workflow_attributes/{attri"
  "butes.project}/{attributes.domain}/{attr"
  "ibutes.workflow}:\001*Zp\032k/api/v1/workflow_"
  "attributes/{attributes.org}/{attributes."
  "project}/{attributes.domain}/{attributes"
  ".workflow}:\001*\022\373\001\n\025GetWorkflowAttributes\022"
  ",.flyteidl.admin.WorkflowAttributesGetRe"
  "quest\032-.flyteidl.admin.WorkflowAttribute"
  "sGetResponse\"\204\001\202\323\344\223\002~\0229/api/v1/workflow_"
  "attributes/{project}/{domain}/{workflow}"
  "ZA\022\?/api/v1/workflow_attributes/{org}/{p"
  "roject}/{domain}/{workflow}\022\213\002\n\030DeleteWo"
  "rkflowAttributes\022/.flyteidl.admin.Workfl"
  "owAttributesDeleteRequest\0320.flyteidl.adm"
  "in.WorkflowAttributesDeleteResponse\"\213\001\202\323"
  "\344\223\002\204\001*9/api/v1/workflow_attributes/{proj"
  "ect}/{domain}/{workflow}:\001*ZD*\?/api/v1/w"
  "orkflow_attributes/{org}/{project}/{doma"
  "in}/{workflow}:\001*\022\240\001\n\027ListMatchableAttri"
  "butes\022..flyteidl.admin.ListMatchableAttr"
  "ibutesRequest\032/.flyteidl.admin.ListMatch"
  "ableAttributesResponse\"$\202\323\344\223\002\036\022\034/api/v1/"
  "matchable_attributes\022\343\001\n\021ListNamedEntiti"
  "es\022&.flyteidl.admin.NamedEntityListReque"
  "st\032\037.flyteidl.admin.NamedEntityList\"\204\001\202\323"
  "\344\223\002~\0229/api/v1/named_entities/{resource_t"
  "ype}/{project}/{domain}ZA\022\?/api/v1/named"
  "_entities/{resource_type}/{org}/{project"
  "}/{domain}\022\377\001\n\016GetNamedEntity\022%.flyteidl"
  ".admin.NamedEntityGetRequest\032\033.flyteidl."
  "admin.NamedEntity\"\250\001\202\323\344\223\002\241\001\022I/api/v1/nam"
  "ed_entities/{resource_type}/{id.project}"
  "/{id.domain}/{id.name}ZT\022R/api/v1/named_"
  "entities/{resource_type}/{id.org}/{id.pr"
  "oject}/{id.domain}/{id.name}\022\231\002\n\021UpdateN"
  "amedEntity\022(.flyteidl.admin.NamedEntityU"
  "pdateRequest\032).flyteidl.admin.NamedEntit"
  "yUpdateResponse\"\256\001\202\323\344\223\002\247\001\032I/api/v1/named"
  "_entities/{resource_type}/{id.project}/{"
  "id.domain}/{id.name}:\001*ZW\032R/api/v1/named"
  "_entities/{resource_type}/{id.org}/{id.p"
  "roject}/{id.domain}/{id.name}:\001*\022l\n\nGetV"
  "ersion\022!.flyteidl.admin.GetVersionReques"
  "t\032\".flyteidl.admin.GetVersionResponse\"\027\202"
  "\323\344\223\002\021\022\017/api/v1/version\022\262\002\n\024GetDescriptio"
  "nEntity\022 .flyteidl.admin.ObjectGetReques"
  "t\032!.flyteidl.admin.DescriptionEntity\"\324\001\202"
  "\323\344\223\002\315\001\022_/api/v1/description_entities/{id"
  ".resource_type}/{id.project}/{id.domain}"
  "/{id.name}/{id.version}Zj\022h/api/v1/descr"
  "iption_entities/{id.org}/{id.resource_ty"
  "pe}/{id.project}/{id.domain}/{id.name}/{"
  "id.version}\022\300\003\n\027ListDescriptionEntities\022"
  ",.flyteidl.admin.DescriptionEntityListRe"
  "quest\032%.flyteidl.admin.DescriptionEntity"
  "List\"\317\002\202\323\344\223\002\310\002\022O/api/v1/description_enti"
  "ties/{resource_type}/{id.project}/{id.do"
  "main}/{id.name}ZZ\022X/api/v1/description_e"
  "ntities/{resource_type}/{id.org}/{id.pro"
  "ject}/{id.domain}/{id.name}ZG\022E/api/v1/d"
  "escription_entities/{resource_type}/{id."
  "project}/{id.domain}ZP\022N/api/v1/descript"
  "ion_entities/{resource_type}/{id.org}/{i"
  "d.project}/{id.domain}\022\221\002\n\023GetExecutionM"
  "etrics\0222.flyteidl.admin.WorkflowExecutio"
  "nGetMetricsRequest\0323.flyteidl.admin.Work"
  "flowExecutionGetMetricsResponse\"\220\001\202\323\344\223\002\211"
  "\001\022=/api/v1/metrics/executions/{id.projec"
  "t}/{id.domain}/{id.name}ZH\022F/api/v1/metr"
  "ics/executions/{id.org}/{id.project}/{id"
  ".domain}/{id.name}B\?Z=github.com/flyteor"
  "g/flyte/flyteidl/gen/pb-go/flyteidl/serv"
  "iceb\006proto3"
  ;
::google::protobuf::internal::DescriptorTable descriptor_table_flyteidl_2fservice_2fadmin_2eproto = {
  false, InitDefaults_flyteidl_2fservice_2fadmin_2eproto, 
  descriptor_table_protodef_flyteidl_2fservice_2fadmin_2eproto,
  "flyteidl/service/admin.proto", &assign_descriptors_table_flyteidl_2fservice_2fadmin_2eproto, 15131,
};

void AddDescriptors_flyteidl_2fservice_2fadmin_2eproto() {
  static constexpr ::google::protobuf::internal::InitFunc deps[16] =
  {
    ::AddDescriptors_google_2fapi_2fannotations_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fproject_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fproject_5fdomain_5fattributes_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fproject_5fattributes_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2ftask_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fworkflow_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fworkflow_5fattributes_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2flaunch_5fplan_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fevent_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fexecution_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fmatchable_5fresource_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fnode_5fexecution_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2ftask_5fexecution_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fversion_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fcommon_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fdescription_5fentity_2eproto,
  };
 ::google::protobuf::internal::AddDescriptors(&descriptor_table_flyteidl_2fservice_2fadmin_2eproto, deps, 16);
}

// Force running AddDescriptors() at dynamic initialization time.
static bool dynamic_init_dummy_flyteidl_2fservice_2fadmin_2eproto = []() { AddDescriptors_flyteidl_2fservice_2fadmin_2eproto(); return true; }();
namespace flyteidl {
namespace service {

// @@protoc_insertion_point(namespace_scope)
}  // namespace service
}  // namespace flyteidl
namespace google {
namespace protobuf {
}  // namespace protobuf
}  // namespace google

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
