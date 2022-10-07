package aassdkgo

import (
	"github.com/spacefleet-dev/aas-sdk-go/api"
)

type SubmodelElementList struct {
	SubmodelElement

	OrderRelevant         bool
	SemanticIDListElement *Reference
	TypeValueListElement  api.AasSubmodelElements
	Items                 []api.ISubmodelElement
	ItemTypeListElement   *api.DataTypeDefXsd
}

func NewSubmodelElementList(sme api.ISubmodelElement) (*SubmodelElementList, bool) {
	if sme.SubmodelElementList == nil {
		return nil, false
	}

	var orderRelevant bool
	if sme.SubmodelElementList.OrderRelevant != nil {
		orderRelevant = *sme.SubmodelElementList.OrderRelevant
	}

	return &SubmodelElementList{
		SubmodelElement: SubmodelElement{
			Extensions:                 sme.SubmodelElementList.Extensions,
			Category:                   strFromPtr(sme.SubmodelElementList.Category),
			Checksum:                   strFromPtr(sme.SubmodelElementList.Checksum),
			Description:                sme.SubmodelElementList.Description,
			DisplayName:                sme.SubmodelElementList.DisplayName,
			IDShort:                    strFromPtr(sme.SubmodelElementList.IdShort),
			Kind:                       sme.SubmodelElementList.Kind,
			SemanticID:                 NewReference(sme.SubmodelElementList.SemanticId),
			ModelType:                  sme.SubmodelElementList.ModelType,
			SupplementalSemanticIDs:    apiRefSliceToRefSlice(sme.SubmodelElementList.SupplementalSemanticIds),
			Qualifiers:                 sme.SubmodelElementList.Qualifiers,
			EmbeddedDataSpecifications: sme.SubmodelElementList.EmbeddedDataSpecifications,
		},
		OrderRelevant:         orderRelevant,
		SemanticIDListElement: NewReference(sme.SubmodelElementList.SemanticIdListElement),
		TypeValueListElement:  sme.SubmodelElementList.TypeValueListElement,
		Items:                 sme.SubmodelElementList.Value,
		ItemTypeListElement:   sme.SubmodelElementList.ValueTypeListElement,
	}, sme.SubmodelElementList != nil
}

func (smel SubmodelElementList) ISubmodelElement() api.ISubmodelElement {
	return api.ISubmodelElement{
		SubmodelElementList: &api.SubmodelElementList{
			Extensions:                 smel.Extensions,
			Category:                   strPtr(smel.Category),
			Checksum:                   strPtr(smel.Checksum),
			Description:                smel.Description,
			DisplayName:                smel.DisplayName,
			IdShort:                    strPtr(smel.IDShort),
			Kind:                       smel.Kind,
			SemanticId:                 smel.SemanticID.ptr(),
			ModelType:                  smel.ModelType,
			SupplementalSemanticIds:    refSliceToAPIRefSlice(smel.SupplementalSemanticIDs),
			Qualifiers:                 smel.Qualifiers,
			EmbeddedDataSpecifications: smel.EmbeddedDataSpecifications,
			OrderRelevant:              &smel.OrderRelevant,
			SemanticIdListElement:      smel.SemanticIDListElement.ptr(),
			TypeValueListElement:       smel.TypeValueListElement,
			Value:                      smel.Items,
			ValueTypeListElement:       smel.ItemTypeListElement,
		},
	}
}

func (smel SubmodelElementList) Find(idShort string) AnySubmodelElement {
	return findSubmodelElement(smel.Items, idShort)
}

func (smel SubmodelElementList) FindAll(idShort string) []AnySubmodelElement {
	return findAllSubmodelElements(smel.Items, idShort)
}

func (smel SubmodelElementList) FindAllByModelType(modelType api.ModelType) []AnySubmodelElement {
	return findAllSubmodelElementsByModelType(smel.Items, modelType)
}
