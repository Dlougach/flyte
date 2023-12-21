package validation

import (
	"context"

	"google.golang.org/grpc/codes"
	"k8s.io/apimachinery/pkg/util/validation"

	"github.com/flyteorg/flyte/flyteadmin/pkg/errors"
	"github.com/flyteorg/flyte/flyteadmin/pkg/manager/impl/shared"
	repositoryInterfaces "github.com/flyteorg/flyte/flyteadmin/pkg/repositories/interfaces"
	runtimeInterfaces "github.com/flyteorg/flyte/flyteadmin/pkg/runtime/interfaces"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"
)

const projectID = "project_id"
const projectName = "project_name"
const projectDescription = "project_description"
const maxNameLength = 64
const maxDescriptionLength = 300
const maxLabelArrayLength = 16

func ValidateProjectRegisterRequest(request admin.ProjectRegisterRequest) error {
	if request.Project == nil {
		return shared.GetMissingArgumentError(shared.Project)
	}
	project := *request.Project
	if err := ValidateEmptyStringField(project.Name, projectName); err != nil {
		return err
	}
	return ValidateProject(project)
}

func ValidateProject(project admin.Project) error {
	if err := ValidateEmptyStringField(project.Id, projectID); err != nil {
		return err
	}
	if err := validateLabels(project.Labels); err != nil {
		return err
	}
	if errs := validation.IsDNS1123Label(project.Id); len(errs) > 0 {
		return errors.NewFlyteAdminErrorf(codes.InvalidArgument, "invalid project id [%s]: %v", project.Id, errs)
	}
	if err := ValidateMaxLengthStringField(project.Name, projectName, maxNameLength); err != nil {
		return err
	}
	if err := ValidateMaxLengthStringField(project.Description, projectDescription, maxDescriptionLength); err != nil {
		return err
	}
	if project.Domains != nil {
		return errors.NewFlyteAdminError(codes.InvalidArgument,
			"Domains are currently only set system wide. Please retry without domains included in your request.")
	}
	return nil
}

type WithProjectAndDomain interface {
	GetProject() string
	GetDomain() string
	GetOrg() string
}

// Validates that a specified project and domain combination has been registered and exists in the db.
func ValidateProjectAndDomain(
	ctx context.Context, db repositoryInterfaces.Repository, config runtimeInterfaces.ApplicationConfiguration, message WithProjectAndDomain) error {
	project, err := db.ProjectRepo().Get(ctx, &admin.ProjectIdentifier{
		Id:  message.GetProject(),
		Org: message.GetOrg(),
	})
	if err != nil {
		return errors.NewFlyteAdminErrorf(codes.InvalidArgument,
			"failed to validate that project [%s] and domain [%s] are registered, err: [%+v]",
			message.GetProject(), message.GetDomain(), err)
	}
	if *project.State != int32(admin.Project_ACTIVE) {
		return errors.NewFlyteAdminErrorf(codes.InvalidArgument,
			"project [%s] is not active", message.GetProject())
	}
	var validDomain bool
	domains := config.GetDomainsConfig()
	for _, domain := range *domains {
		if domain.ID == message.GetDomain() {
			validDomain = true
			break
		}
	}
	if !validDomain {
		return errors.NewFlyteAdminErrorf(codes.InvalidArgument, "domain [%s] is unrecognized by system", message.GetDomain())
	}
	return nil
}

func ValidateProjectForUpdate(
	ctx context.Context, db repositoryInterfaces.Repository, id *admin.ProjectIdentifier) error {

	project, err := db.ProjectRepo().Get(ctx, id)
	if err != nil {
		return errors.NewFlyteAdminErrorf(codes.InvalidArgument,
			"failed to validate that project [%s] is registered, err: [%+v]",
			projectID, err)
	}
	if *project.State != int32(admin.Project_ACTIVE) {
		return errors.NewFlyteAdminErrorf(codes.InvalidArgument,
			"project [%s] is not active", projectID)
	}
	return nil
}

// ValidateProjectExists doesn't check that the project is active. This is used to get Project level attributes, which you should
// be able to do even for inactive projects.
func ValidateProjectExists(
	ctx context.Context, db repositoryInterfaces.Repository, id *admin.ProjectIdentifier) error {

	_, err := db.ProjectRepo().Get(ctx, id)
	if err != nil {
		return errors.NewFlyteAdminErrorf(codes.InvalidArgument,
			"failed to validate that project [%s] exists, err: [%+v]",
			projectID, err)
	}
	return nil
}
