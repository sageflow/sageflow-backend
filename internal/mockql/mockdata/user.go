package mockdata

import (
	"github.com/sageflow/sageapi/internal/mockql/model"
)

// GetSessionUser ...
func GetSessionUser() *model.User {
	return &model.User{
		ID:           "4e6c8b75-4a74-42fd-84b3-f319a8f153c6",
		Profile:      getUserProfile(),
		Session:      getSession(),
		Workspaces:   getWorkspaces(),
		Integrations: getIntegrations(),
	}
}
