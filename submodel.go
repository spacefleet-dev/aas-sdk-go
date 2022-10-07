package aassdkgo

import (
	"github.com/spacefleet-dev/aas-sdk-go/api"
)

type Submodel struct {
	Submodel *api.Submodel
}

func (s Submodel) Elements() []api.ISubmodelElement {
	return s.Submodel.SubmodelElements
}

func (s Submodel) IDShort() string {
	if s.Submodel.IdShort == nil {
		return ""
	}

	return *s.Submodel.IdShort
}

func (s Submodel) SemanticID() *api.Reference {
	return s.Submodel.SemanticId
}
