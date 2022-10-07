package aassdkgo

import (
	"github.com/spacefleet-dev/aas-sdk-go/api"
)

type AssetAdministrationShell struct {
	AAS       api.AssetAdministrationShell
	Submodels []*Submodel
}

func NewAssetAdministration(idShort string, id Identifier) *AssetAdministrationShell {
	return &AssetAdministrationShell{AAS: api.AssetAdministrationShell{
		IdShort:        &idShort,
		Identification: id.ID,
	}}
}

func NewAssetAdministrationShellFromAPI(aas api.AssetAdministrationShell) *AssetAdministrationShell {
	return &AssetAdministrationShell{AAS: aas}
}

func (aas *AssetAdministrationShell) SetAsset(asset *Asset) {
	aas.AAS.SetAssetInformation(asset.Asset)
}

func (aas *AssetAdministrationShell) Asset() *Asset {
	return &Asset{aas.AAS.AssetInformation}
}

func (aas *AssetAdministrationShell) AddSubmodel(sm *Submodel) {
	aas.Submodels = append(aas.Submodels, sm)

	aas.AAS.Submodels = append(aas.AAS.Submodels, api.Reference{
		Keys: []api.Key{{
			Type:  api.KEYTYPES_SUBMODEL,
			Value: sm.Submodel.Identification.Id,
		}},
		Type: api.REFERENCETYPES_GLOBAL_REFERENCE,
	})
}

// FindSubmodel returns the *first* Submodel with the given IDShort, or nil if it can't be found.
func (aas *AssetAdministrationShell) FindSubmodel(idShort string) *Submodel {
	for i := range aas.AAS.Submodels {
		if idShort == aas.Submodels[i].IDShort() {
			return aas.Submodels[i]
		}
	}

	return nil
}

// FindSubmodelBySemanticID returns the *first* Submodel with the given SemanticID, or nil if it can't be found.
func (aas *AssetAdministrationShell) FindSubmodelBySemanticID(semanticID string) *Submodel {
	for i := range aas.AAS.Submodels {
		semID := aas.Submodels[i].SemanticID()
		if semID == nil {
			continue
		}

		for _, k := range semID.Keys {
			if semanticID == k.Value {
				return aas.Submodels[i]
			}
		}
	}

	return nil
}

func (aas *AssetAdministrationShell) DerivedFrom() *Reference {
	if aas.AAS.DerivedFrom == nil {
		return nil
	}

	return &Reference{
		Ref: api.Reference{
			Keys: aas.AAS.DerivedFrom.Keys,
		},
	}
}

func (aas *AssetAdministrationShell) SetDerivedFrom(ref *Reference) {
	aas.AAS.DerivedFrom = &ref.Ref
}

func (aas *AssetAdministrationShell) SetIdentification(id *Identifier) {
	aas.AAS.SetIdentification(id.ID)
}

func (aas *AssetAdministrationShell) Identification() *Identifier {
	return &Identifier{aas.AAS.Identification}
}
