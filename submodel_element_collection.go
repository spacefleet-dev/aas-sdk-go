package aassdkgo

import (
	"github.com/spacefleet-dev/aas-sdk-go/api"
)

type SubmodelElementCollection struct {
	SubmodelElement
	Items []api.ISubmodelElement
}

func NewSubmodelElementCollection(sme api.ISubmodelElement) (*SubmodelElementCollection, bool) {
	if sme.SubmodelElementCollection == nil {
		return nil, false
	}

	return &SubmodelElementCollection{
		SubmodelElement: SubmodelElement{
			Extensions:                 sme.SubmodelElementCollection.Extensions,
			Category:                   strFromPtr(sme.SubmodelElementCollection.Category),
			Checksum:                   strFromPtr(sme.SubmodelElementCollection.Checksum),
			Description:                sme.SubmodelElementCollection.Description,
			DisplayName:                sme.SubmodelElementCollection.DisplayName,
			IDShort:                    strFromPtr(sme.SubmodelElementCollection.IdShort),
			Kind:                       sme.SubmodelElementCollection.Kind,
			SemanticID:                 NewReference(sme.SubmodelElementCollection.SemanticId),
			ModelType:                  sme.SubmodelElementCollection.ModelType,
			SupplementalSemanticIDs:    apiRefSliceToRefSlice(sme.SubmodelElementCollection.SupplementalSemanticIds),
			Qualifiers:                 sme.SubmodelElementCollection.Qualifiers,
			EmbeddedDataSpecifications: sme.SubmodelElementCollection.EmbeddedDataSpecifications,
		},
		Items: sme.SubmodelElementCollection.Value,
	}, sme.SubmodelElementCollection != nil
}

func (smec SubmodelElementCollection) ISubmodelElement() api.ISubmodelElement {
	return api.ISubmodelElement{
		SubmodelElementList: &api.SubmodelElementList{
			Extensions:                 smec.Extensions,
			Category:                   strPtr(smec.Category),
			Checksum:                   strPtr(smec.Checksum),
			Description:                smec.Description,
			DisplayName:                smec.DisplayName,
			IdShort:                    strPtr(smec.IDShort),
			Kind:                       smec.Kind,
			SemanticId:                 smec.SemanticID.ptr(),
			ModelType:                  smec.ModelType,
			SupplementalSemanticIds:    refSliceToAPIRefSlice(smec.SupplementalSemanticIDs),
			Qualifiers:                 smec.Qualifiers,
			EmbeddedDataSpecifications: smec.EmbeddedDataSpecifications,
			Value:                      smec.Items,
		},
	}
}

func (smec SubmodelElementCollection) Find(idShort string) AnySubmodelElement {
	return findSubmodelElement(smec.Items, idShort)
}

func (smec SubmodelElementCollection) FindAll(idShort string) []AnySubmodelElement {
	return findAllSubmodelElements(smec.Items, idShort)
}

func (smec SubmodelElementCollection) FindAllByModelType(modelType api.ModelType) []AnySubmodelElement {
	return findAllSubmodelElementsByModelType(smec.Items, modelType)
}
