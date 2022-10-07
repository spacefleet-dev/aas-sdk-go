package aassdkgo

import (
	"github.com/spacefleet-dev/aas-sdk-go/api"
)

type AnySubmodelElement interface {
	GetModelType() api.ModelType
	GetIDShort() string
}

type SubmodelElement struct {
	Extensions                 []api.Extension
	Category                   string
	Checksum                   string
	Description                []api.LangString
	DisplayName                []api.LangString
	IDShort                    string
	Kind                       *api.ModelingKind
	SemanticID                 *Reference
	ModelType                  api.ModelType
	SupplementalSemanticIDs    []Reference
	Qualifiers                 []api.Qualifier
	EmbeddedDataSpecifications []api.EmbeddedDataSpecification
}

func NewSubmodelElement(sme api.ISubmodelElement) (*SubmodelElement, bool) {
	if sme.SubmodelElement == nil {
		return nil, false
	}

	return &SubmodelElement{
		Extensions:                 sme.SubmodelElement.Extensions,
		Category:                   strFromPtr(sme.SubmodelElement.Category),
		Checksum:                   strFromPtr(sme.SubmodelElement.Checksum),
		Description:                sme.SubmodelElement.Description,
		DisplayName:                sme.SubmodelElement.DisplayName,
		IDShort:                    strFromPtr(sme.SubmodelElement.IdShort),
		Kind:                       sme.SubmodelElement.Kind,
		SemanticID:                 NewReference(sme.SubmodelElement.SemanticId),
		ModelType:                  sme.SubmodelElement.ModelType,
		SupplementalSemanticIDs:    apiRefSliceToRefSlice(sme.SubmodelElement.SupplementalSemanticIds),
		Qualifiers:                 sme.SubmodelElement.Qualifiers,
		EmbeddedDataSpecifications: sme.SubmodelElement.EmbeddedDataSpecifications,
	}, true
}

func (sme SubmodelElement) ISubmodelElement() api.ISubmodelElement {
	return api.ISubmodelElement{
		SubmodelElement: &api.SubmodelElement{
			Extensions:                 sme.Extensions,
			Category:                   strPtr(sme.Category),
			Checksum:                   strPtr(sme.Checksum),
			Description:                sme.Description,
			DisplayName:                sme.DisplayName,
			IdShort:                    strPtr(sme.IDShort),
			Kind:                       sme.Kind,
			SemanticId:                 sme.SemanticID.ptr(),
			ModelType:                  sme.ModelType,
			SupplementalSemanticIds:    refSliceToAPIRefSlice(sme.SupplementalSemanticIDs),
			Qualifiers:                 sme.Qualifiers,
			EmbeddedDataSpecifications: sme.EmbeddedDataSpecifications,
		},
	}
}

func (sme SubmodelElement) GetModelType() api.ModelType {
	return sme.ModelType
}

func (sme SubmodelElement) GetIDShort() string {
	return sme.IDShort
}

func NewAnySubmodelElement(sme api.ISubmodelElement) AnySubmodelElement {
	if sme.AnnotatedRelationshipElement != nil {
		a, _ := NewAnnotatedRelationshipElement(sme)
		return a
	}

	if sme.BasicEventElement != nil {
		a, _ := NewBasicEventElement(sme)
		return a
	}

	if sme.Blob != nil {
		a, _ := NewBlob(sme)
		return a
	}

	if sme.Entity != nil {
		a, _ := NewEntity(sme)
		return a
	}

	if sme.File != nil {
		a, _ := NewFile(sme)
		return a
	}

	if sme.MultiLanguageProperty != nil {
		a, _ := NewMultiLanguageProperty(sme)
		return a
	}

	if sme.Operation != nil {
		a, _ := NewOperation(sme)
		return a
	}

	if sme.Property != nil {
		a, _ := NewProperty(sme)
		return a
	}

	if sme.Range != nil {
		a, _ := NewRange(sme)
		return a
	}

	if sme.ReferenceElement != nil {
		a, _ := NewReferenceElement(sme)
		return a
	}

	if sme.RelationshipElement != nil {
		a, _ := NewRelationshipElement(sme)
		return a
	}

	if sme.SubmodelElement != nil {
		a, _ := NewSubmodelElement(sme)
		return a
	}

	if sme.SubmodelElementCollection != nil {
		a, _ := NewSubmodelElementCollection(sme)
		return a
	}

	if sme.SubmodelElementList != nil {
		a, _ := NewSubmodelElementList(sme)
		return a
	}

	return nil
}

type AnnotatedRelationshipElement struct {
	SubmodelElement
	First       Reference
	Second      Reference
	Annotations []api.ISubmodelElement
}

func NewAnnotatedRelationshipElement(sme api.ISubmodelElement) (*AnnotatedRelationshipElement, bool) {
	if sme.AnnotatedRelationshipElement == nil {
		return nil, false
	}

	return &AnnotatedRelationshipElement{
		SubmodelElement: SubmodelElement{
			Extensions:                 sme.AnnotatedRelationshipElement.Extensions,
			Category:                   strFromPtr(sme.AnnotatedRelationshipElement.Category),
			Checksum:                   strFromPtr(sme.AnnotatedRelationshipElement.Checksum),
			Description:                sme.AnnotatedRelationshipElement.Description,
			DisplayName:                sme.AnnotatedRelationshipElement.DisplayName,
			IDShort:                    strFromPtr(sme.AnnotatedRelationshipElement.IdShort),
			Kind:                       sme.AnnotatedRelationshipElement.Kind,
			SemanticID:                 NewReference(sme.AnnotatedRelationshipElement.SemanticId),
			ModelType:                  sme.AnnotatedRelationshipElement.ModelType,
			SupplementalSemanticIDs:    apiRefSliceToRefSlice(sme.AnnotatedRelationshipElement.SupplementalSemanticIds),
			Qualifiers:                 sme.AnnotatedRelationshipElement.Qualifiers,
			EmbeddedDataSpecifications: sme.AnnotatedRelationshipElement.EmbeddedDataSpecifications,
		},
		First:       *NewReference(&sme.AnnotatedRelationshipElement.First),
		Second:      *NewReference(&sme.AnnotatedRelationshipElement.Second),
		Annotations: sme.AnnotatedRelationshipElement.Annotations,
	}, true
}

func (are AnnotatedRelationshipElement) ISubmodelElement() api.ISubmodelElement {
	return api.ISubmodelElement{
		AnnotatedRelationshipElement: &api.AnnotatedRelationshipElement{
			Extensions:                 are.Extensions,
			Category:                   strPtr(are.Category),
			Checksum:                   strPtr(are.Checksum),
			Description:                are.Description,
			DisplayName:                are.DisplayName,
			IdShort:                    strPtr(are.IDShort),
			Kind:                       are.Kind,
			SemanticId:                 are.SemanticID.ptr(),
			SupplementalSemanticIds:    refSliceToAPIRefSlice(are.SupplementalSemanticIDs),
			Qualifiers:                 are.Qualifiers,
			EmbeddedDataSpecifications: are.EmbeddedDataSpecifications,
			// AnnotatedRelationshipElement
			ModelType:   api.MODELTYPE_ANNOTATED_RELATIONSHIP_ELEMENT,
			First:       are.First.Ref,
			Second:      are.Second.Ref,
			Annotations: are.Annotations,
		},
	}
}

type BasicEventElement struct {
	SubmodelElement
	Direction     api.Direction
	LastUpdate    string
	MaxInterval   string
	MessageBroker *Reference
	MessageTopic  string
	MinInterval   string
	Observed      Reference
	State         api.StateOfEvent
}

func NewBasicEventElement(sme api.ISubmodelElement) (*BasicEventElement, bool) {
	if sme.BasicEventElement == nil {
		return nil, false
	}

	return &BasicEventElement{
		SubmodelElement: SubmodelElement{
			Extensions:                 sme.BasicEventElement.Extensions,
			Category:                   strFromPtr(sme.BasicEventElement.Category),
			Checksum:                   strFromPtr(sme.BasicEventElement.Checksum),
			Description:                sme.BasicEventElement.Description,
			DisplayName:                sme.BasicEventElement.DisplayName,
			IDShort:                    strFromPtr(sme.BasicEventElement.IdShort),
			Kind:                       sme.BasicEventElement.Kind,
			SemanticID:                 NewReference(sme.BasicEventElement.SemanticId),
			ModelType:                  sme.BasicEventElement.ModelType,
			SupplementalSemanticIDs:    apiRefSliceToRefSlice(sme.BasicEventElement.SupplementalSemanticIds),
			Qualifiers:                 sme.BasicEventElement.Qualifiers,
			EmbeddedDataSpecifications: sme.BasicEventElement.EmbeddedDataSpecifications,
		},
		Direction:     sme.BasicEventElement.Direction,
		LastUpdate:    strFromPtr(sme.BasicEventElement.LastUpdate),
		MaxInterval:   strFromPtr(sme.BasicEventElement.MaxInterval),
		MessageBroker: NewReference(sme.BasicEventElement.MessageBroker),
		MessageTopic:  strFromPtr(sme.BasicEventElement.MessageTopic),
		MinInterval:   strFromPtr(sme.BasicEventElement.MinInterval),
		Observed:      *NewReference(&sme.BasicEventElement.Observed),
		State:         sme.BasicEventElement.State,
	}, true
}

func (bee BasicEventElement) ISubmodelElement() api.ISubmodelElement {
	return api.ISubmodelElement{
		BasicEventElement: &api.BasicEventElement{
			Extensions:                 bee.Extensions,
			Category:                   strPtr(bee.Category),
			Checksum:                   strPtr(bee.Checksum),
			Description:                bee.Description,
			DisplayName:                bee.DisplayName,
			IdShort:                    strPtr(bee.IDShort),
			Kind:                       bee.Kind,
			SemanticId:                 bee.SemanticID.ptr(),
			SupplementalSemanticIds:    refSliceToAPIRefSlice(bee.SupplementalSemanticIDs),
			Qualifiers:                 bee.Qualifiers,
			EmbeddedDataSpecifications: bee.EmbeddedDataSpecifications,
			// BasicEventElement
			ModelType:     api.MODELTYPE_BASIC_EVENT_ELEMENT,
			Direction:     bee.Direction,
			LastUpdate:    strPtr(bee.LastUpdate),
			MaxInterval:   strPtr(bee.MaxInterval),
			MessageBroker: bee.MessageBroker.ptr(),
			MessageTopic:  strPtr(bee.MessageTopic),
			MinInterval:   strPtr(bee.MinInterval),
			Observed:      bee.Observed.Ref,
			State:         bee.State,
		},
	}
}

type Blob struct {
	SubmodelElement
	ContentType string
	Value       string
}

func NewBlob(sme api.ISubmodelElement) (*Blob, bool) {
	if sme.Blob == nil {
		return nil, false
	}

	return &Blob{
		SubmodelElement: SubmodelElement{
			Extensions:                 sme.Blob.Extensions,
			Category:                   strFromPtr(sme.Blob.Category),
			Checksum:                   strFromPtr(sme.Blob.Checksum),
			Description:                sme.Blob.Description,
			DisplayName:                sme.Blob.DisplayName,
			IDShort:                    strFromPtr(sme.Blob.IdShort),
			Kind:                       sme.Blob.Kind,
			SemanticID:                 NewReference(sme.Blob.SemanticId),
			ModelType:                  sme.Blob.ModelType,
			SupplementalSemanticIDs:    apiRefSliceToRefSlice(sme.Blob.SupplementalSemanticIds),
			Qualifiers:                 sme.Blob.Qualifiers,
			EmbeddedDataSpecifications: sme.Blob.EmbeddedDataSpecifications,
		},
		ContentType: sme.Blob.ContentType,
		Value:       strFromPtr(sme.Blob.Value),
	}, true
}

func (b Blob) ISubmodelElement() api.ISubmodelElement {
	return api.ISubmodelElement{
		Blob: &api.Blob{
			Extensions:                 b.Extensions,
			Category:                   strPtr(b.Category),
			Checksum:                   strPtr(b.Checksum),
			Description:                b.Description,
			DisplayName:                b.DisplayName,
			IdShort:                    strPtr(b.IDShort),
			Kind:                       b.Kind,
			SemanticId:                 b.SemanticID.ptr(),
			SupplementalSemanticIds:    refSliceToAPIRefSlice(b.SupplementalSemanticIDs),
			Qualifiers:                 b.Qualifiers,
			EmbeddedDataSpecifications: b.EmbeddedDataSpecifications,
			// Blob
			ModelType:   api.MODELTYPE_BLOB,
			ContentType: b.ContentType,
			Value:       strPtr(b.Value),
		},
	}
}

type Entity struct {
	SubmodelElement
	EntityType      api.EntityType
	GlobalAssetId   *Reference
	SpecificAssetId *api.SpecificAssetId
	Statements      []api.SubmodelElement
}

func NewEntity(sme api.ISubmodelElement) (*Entity, bool) {
	if sme.Entity == nil {
		return nil, false
	}

	return &Entity{
		SubmodelElement: SubmodelElement{
			Extensions:                 sme.Entity.Extensions,
			Category:                   strFromPtr(sme.Entity.Category),
			Checksum:                   strFromPtr(sme.Entity.Checksum),
			Description:                sme.Entity.Description,
			DisplayName:                sme.Entity.DisplayName,
			IDShort:                    strFromPtr(sme.Entity.IdShort),
			Kind:                       sme.Entity.Kind,
			SemanticID:                 NewReference(sme.Entity.SemanticId),
			ModelType:                  sme.Entity.ModelType,
			SupplementalSemanticIDs:    apiRefSliceToRefSlice(sme.Entity.SupplementalSemanticIds),
			Qualifiers:                 sme.Entity.Qualifiers,
			EmbeddedDataSpecifications: sme.Entity.EmbeddedDataSpecifications,
		},
		EntityType:      sme.Entity.EntityType,
		GlobalAssetId:   NewReference(sme.Entity.GlobalAssetId),
		SpecificAssetId: sme.Entity.SpecificAssetId,
		Statements:      sme.Entity.Statements,
	}, true
}

func (e Entity) ISubmodelElement() api.ISubmodelElement {
	return api.ISubmodelElement{
		Entity: &api.Entity{
			Extensions:                 e.Extensions,
			Category:                   strPtr(e.Category),
			Checksum:                   strPtr(e.Checksum),
			Description:                e.Description,
			DisplayName:                e.DisplayName,
			IdShort:                    strPtr(e.IDShort),
			Kind:                       e.Kind,
			SemanticId:                 e.SemanticID.ptr(),
			SupplementalSemanticIds:    refSliceToAPIRefSlice(e.SupplementalSemanticIDs),
			Qualifiers:                 e.Qualifiers,
			EmbeddedDataSpecifications: e.EmbeddedDataSpecifications,
			// Entity
			ModelType:       api.MODELTYPE_ENTITY,
			EntityType:      e.EntityType,
			GlobalAssetId:   e.GlobalAssetId.ptr(),
			SpecificAssetId: e.SpecificAssetId,
			Statements:      e.Statements,
		},
	}
}

type File struct {
	SubmodelElement
	ContentType string
	Value       string
}

func NewFile(sme api.ISubmodelElement) (*File, bool) {
	if sme.File == nil {
		return nil, false
	}

	return &File{
		SubmodelElement: SubmodelElement{
			Extensions:                 sme.File.Extensions,
			Category:                   strFromPtr(sme.File.Category),
			Checksum:                   strFromPtr(sme.File.Checksum),
			Description:                sme.File.Description,
			DisplayName:                sme.File.DisplayName,
			IDShort:                    strFromPtr(sme.File.IdShort),
			Kind:                       sme.File.Kind,
			SemanticID:                 NewReference(sme.File.SemanticId),
			ModelType:                  sme.File.ModelType,
			SupplementalSemanticIDs:    apiRefSliceToRefSlice(sme.File.SupplementalSemanticIds),
			Qualifiers:                 sme.File.Qualifiers,
			EmbeddedDataSpecifications: sme.File.EmbeddedDataSpecifications,
		},
		ContentType: sme.File.ContentType,
		Value:       strFromPtr(sme.File.Value),
	}, true
}

func (f File) ISubmodelElement() api.ISubmodelElement {
	return api.ISubmodelElement{
		File: &api.File{
			Extensions:                 f.Extensions,
			Category:                   strPtr(f.Category),
			Checksum:                   strPtr(f.Checksum),
			Description:                f.Description,
			DisplayName:                f.DisplayName,
			IdShort:                    strPtr(f.IDShort),
			Kind:                       f.Kind,
			SemanticId:                 f.SemanticID.ptr(),
			SupplementalSemanticIds:    refSliceToAPIRefSlice(f.SupplementalSemanticIDs),
			Qualifiers:                 f.Qualifiers,
			EmbeddedDataSpecifications: f.EmbeddedDataSpecifications,
			// File
			ModelType:   api.MODELTYPE_FILE,
			ContentType: f.ContentType,
			Value:       strPtr(f.Value),
		},
	}
}

type MultiLanguageProperty struct {
	SubmodelElement
	Value   []api.LangString
	ValueID *Reference
}

func NewMultiLanguageProperty(sme api.ISubmodelElement) (*MultiLanguageProperty, bool) {
	if sme.MultiLanguageProperty == nil {
		return nil, false
	}

	return &MultiLanguageProperty{
		SubmodelElement: SubmodelElement{
			Extensions:                 sme.MultiLanguageProperty.Extensions,
			Category:                   strFromPtr(sme.MultiLanguageProperty.Category),
			Checksum:                   strFromPtr(sme.MultiLanguageProperty.Checksum),
			Description:                sme.MultiLanguageProperty.Description,
			DisplayName:                sme.MultiLanguageProperty.DisplayName,
			IDShort:                    strFromPtr(sme.MultiLanguageProperty.IdShort),
			Kind:                       sme.MultiLanguageProperty.Kind,
			SemanticID:                 NewReference(sme.MultiLanguageProperty.SemanticId),
			ModelType:                  sme.MultiLanguageProperty.ModelType,
			SupplementalSemanticIDs:    apiRefSliceToRefSlice(sme.MultiLanguageProperty.SupplementalSemanticIds),
			Qualifiers:                 sme.MultiLanguageProperty.Qualifiers,
			EmbeddedDataSpecifications: sme.MultiLanguageProperty.EmbeddedDataSpecifications,
		},
		Value:   sme.MultiLanguageProperty.Value,
		ValueID: NewReference(sme.MultiLanguageProperty.ValueId),
	}, true
}

func (mlp MultiLanguageProperty) ISubmodelElement() api.ISubmodelElement {
	return api.ISubmodelElement{
		MultiLanguageProperty: &api.MultiLanguageProperty{
			Extensions:                 mlp.Extensions,
			Category:                   strPtr(mlp.Category),
			Checksum:                   strPtr(mlp.Checksum),
			Description:                mlp.Description,
			DisplayName:                mlp.DisplayName,
			IdShort:                    strPtr(mlp.IDShort),
			Kind:                       mlp.Kind,
			SemanticId:                 mlp.SemanticID.ptr(),
			SupplementalSemanticIds:    refSliceToAPIRefSlice(mlp.SupplementalSemanticIDs),
			Qualifiers:                 mlp.Qualifiers,
			EmbeddedDataSpecifications: mlp.EmbeddedDataSpecifications,
			// MultiLanguageProperty
			ModelType: api.MODELTYPE_MULTI_LANGUAGE_PROPERTY,
			Value:     mlp.Value,
			ValueId:   mlp.ValueID.ptr(),
		},
	}
}

type Operation struct {
	SubmodelElement
	InoutputVariables []api.ISubmodelElement
	InputVariables    []api.ISubmodelElement
	OutputVariables   []api.ISubmodelElement
}

func NewOperation(sme api.ISubmodelElement) (*Operation, bool) {
	if sme.Operation == nil {
		return nil, false
	}

	inout := make([]api.ISubmodelElement, len(sme.Operation.InoutputVariables))
	for i := range sme.Operation.InoutputVariables {
		inout[i] = sme.Operation.InoutputVariables[i].Value
	}

	input := make([]api.ISubmodelElement, len(sme.Operation.InputVariables))
	for i := range sme.Operation.InputVariables {
		input[i] = sme.Operation.InputVariables[i].Value
	}

	output := make([]api.ISubmodelElement, len(sme.Operation.OutputVariables))
	for i := range sme.Operation.OutputVariables {
		output[i] = sme.Operation.OutputVariables[i].Value
	}

	return &Operation{
		SubmodelElement: SubmodelElement{
			Extensions:                 sme.File.Extensions,
			Category:                   strFromPtr(sme.File.Category),
			Checksum:                   strFromPtr(sme.File.Checksum),
			Description:                sme.File.Description,
			DisplayName:                sme.File.DisplayName,
			IDShort:                    strFromPtr(sme.File.IdShort),
			Kind:                       sme.File.Kind,
			SemanticID:                 NewReference(sme.File.SemanticId),
			ModelType:                  sme.File.ModelType,
			SupplementalSemanticIDs:    apiRefSliceToRefSlice(sme.File.SupplementalSemanticIds),
			Qualifiers:                 sme.File.Qualifiers,
			EmbeddedDataSpecifications: sme.File.EmbeddedDataSpecifications,
		},
		InoutputVariables: inout,
		InputVariables:    input,
		OutputVariables:   output,
	}, true
}

func (op Operation) ISubmodelElement() api.ISubmodelElement {
	inout := make([]api.OperationVariable, len(op.InoutputVariables))
	for i := range op.InoutputVariables {
		inout[i].Value = op.InoutputVariables[i]
	}

	input := make([]api.OperationVariable, len(op.InputVariables))
	for i := range op.InputVariables {
		input[i].Value = op.InputVariables[i]
	}

	output := make([]api.OperationVariable, len(op.OutputVariables))
	for i := range op.OutputVariables {
		output[i].Value = op.OutputVariables[i]
	}

	return api.ISubmodelElement{
		Operation: &api.Operation{
			Extensions:                 op.Extensions,
			Category:                   strPtr(op.Category),
			Checksum:                   strPtr(op.Checksum),
			Description:                op.Description,
			DisplayName:                op.DisplayName,
			IdShort:                    strPtr(op.IDShort),
			Kind:                       op.Kind,
			SemanticId:                 op.SemanticID.ptr(),
			SupplementalSemanticIds:    refSliceToAPIRefSlice(op.SupplementalSemanticIDs),
			Qualifiers:                 op.Qualifiers,
			EmbeddedDataSpecifications: op.EmbeddedDataSpecifications,
			// Operation
			ModelType:         api.MODELTYPE_OPERATION,
			InoutputVariables: inout,
			InputVariables:    input,
			OutputVariables:   output,
		},
	}
}

type Property struct {
	SubmodelElement
	Value     any
	ValueID   *Reference
	ValueType api.DataTypeDefXsd
}

func NewProperty(sme api.ISubmodelElement) (*Property, bool) {
	if sme.Property == nil {
		return nil, false
	}

	return &Property{
		SubmodelElement: SubmodelElement{
			Extensions:                 sme.Property.Extensions,
			Category:                   strFromPtr(sme.Property.Category),
			Checksum:                   strFromPtr(sme.Property.Checksum),
			Description:                sme.Property.Description,
			DisplayName:                sme.Property.DisplayName,
			IDShort:                    strFromPtr(sme.Property.IdShort),
			Kind:                       sme.Property.Kind,
			SemanticID:                 NewReference(sme.Property.SemanticId),
			ModelType:                  sme.Property.ModelType,
			SupplementalSemanticIDs:    apiRefSliceToRefSlice(sme.Property.SupplementalSemanticIds),
			Qualifiers:                 sme.Property.Qualifiers,
			EmbeddedDataSpecifications: sme.Property.EmbeddedDataSpecifications,
		},
		Value:     sme.Property.Value,
		ValueID:   NewReference(sme.Property.ValueId),
		ValueType: sme.Property.ValueType,
	}, true
}

func (p Property) ISubmodelElement() api.ISubmodelElement {
	val, typ := encodePropertyValue(p.Value)

	if p.ValueType != "" {
		typ = p.ValueType
	}

	return api.ISubmodelElement{
		Property: &api.Property{
			Extensions:                 p.Extensions,
			Category:                   strPtr(p.Category),
			Checksum:                   strPtr(p.Checksum),
			Description:                p.Description,
			DisplayName:                p.DisplayName,
			IdShort:                    strPtr(p.IDShort),
			Kind:                       p.Kind,
			SemanticId:                 p.SemanticID.ptr(),
			SupplementalSemanticIds:    refSliceToAPIRefSlice(p.SupplementalSemanticIDs),
			Qualifiers:                 p.Qualifiers,
			EmbeddedDataSpecifications: p.EmbeddedDataSpecifications,
			// Property
			ModelType: api.MODELTYPE_PROPERTY,
			Value:     strPtr(val),
			ValueId:   p.ValueID.ptr(),
			ValueType: typ,
		},
	}
}

type Range struct {
	SubmodelElement
	Max       any
	Min       any
	ValueType api.DataTypeDefXsd
}

func NewRange(sme api.ISubmodelElement) (*Range, bool) {
	if sme.Range == nil {
		return nil, false
	}

	return &Range{
		SubmodelElement: SubmodelElement{
			Extensions:                 sme.Range.Extensions,
			Category:                   strFromPtr(sme.Range.Category),
			Checksum:                   strFromPtr(sme.Range.Checksum),
			Description:                sme.Range.Description,
			DisplayName:                sme.Range.DisplayName,
			IDShort:                    strFromPtr(sme.Range.IdShort),
			Kind:                       sme.Range.Kind,
			SemanticID:                 NewReference(sme.Range.SemanticId),
			ModelType:                  sme.Range.ModelType,
			SupplementalSemanticIDs:    apiRefSliceToRefSlice(sme.Range.SupplementalSemanticIds),
			Qualifiers:                 sme.Range.Qualifiers,
			EmbeddedDataSpecifications: sme.Range.EmbeddedDataSpecifications,
		},
		Min:       sme.Range.Min,
		Max:       sme.Range.Min,
		ValueType: sme.Range.ValueType,
	}, true
}

func (r Range) ISubmodelElement() api.ISubmodelElement {
	max, typ := encodePropertyValue(r.Max)
	min, _ := encodePropertyValue(r.Min)

	if r.ValueType != "" {
		typ = r.ValueType
	}

	return api.ISubmodelElement{
		Range: &api.Range{
			Extensions:                 r.Extensions,
			Category:                   strPtr(r.Category),
			Checksum:                   strPtr(r.Checksum),
			Description:                r.Description,
			DisplayName:                r.DisplayName,
			IdShort:                    strPtr(r.IDShort),
			Kind:                       r.Kind,
			SemanticId:                 r.SemanticID.ptr(),
			SupplementalSemanticIds:    refSliceToAPIRefSlice(r.SupplementalSemanticIDs),
			Qualifiers:                 r.Qualifiers,
			EmbeddedDataSpecifications: r.EmbeddedDataSpecifications,
			// Propery
			ModelType: api.MODELTYPE_RANGE,
			Max:       strPtr(max),
			Min:       strPtr(min),
			ValueType: typ,
		},
	}
}

type ReferenceElement struct {
	SubmodelElement
	Value *Reference
}

func NewReferenceElement(sme api.ISubmodelElement) (*ReferenceElement, bool) {
	if sme.ReferenceElement == nil {
		return nil, false
	}

	return &ReferenceElement{
		SubmodelElement: SubmodelElement{
			Extensions:                 sme.ReferenceElement.Extensions,
			Category:                   strFromPtr(sme.ReferenceElement.Category),
			Checksum:                   strFromPtr(sme.ReferenceElement.Checksum),
			Description:                sme.ReferenceElement.Description,
			DisplayName:                sme.ReferenceElement.DisplayName,
			IDShort:                    strFromPtr(sme.ReferenceElement.IdShort),
			Kind:                       sme.ReferenceElement.Kind,
			SemanticID:                 NewReference(sme.ReferenceElement.SemanticId),
			ModelType:                  sme.ReferenceElement.ModelType,
			SupplementalSemanticIDs:    apiRefSliceToRefSlice(sme.ReferenceElement.SupplementalSemanticIds),
			Qualifiers:                 sme.ReferenceElement.Qualifiers,
			EmbeddedDataSpecifications: sme.ReferenceElement.EmbeddedDataSpecifications,
		},
		Value: NewReference(sme.ReferenceElement.Value),
	}, true
}

func (f ReferenceElement) ISubmodelElement() api.ISubmodelElement {
	return api.ISubmodelElement{
		ReferenceElement: &api.ReferenceElement{
			Extensions:                 f.Extensions,
			Category:                   strPtr(f.Category),
			Checksum:                   strPtr(f.Checksum),
			Description:                f.Description,
			DisplayName:                f.DisplayName,
			IdShort:                    strPtr(f.IDShort),
			Kind:                       f.Kind,
			SemanticId:                 f.SemanticID.ptr(),
			SupplementalSemanticIds:    refSliceToAPIRefSlice(f.SupplementalSemanticIDs),
			Qualifiers:                 f.Qualifiers,
			EmbeddedDataSpecifications: f.EmbeddedDataSpecifications,
			// ReferenceElement
			ModelType: api.MODELTYPE_REFERENCE_ELEMENT,
			Value:     f.Value.ptr(),
		},
	}
}

type RelationshipElement struct {
	SubmodelElement
	First  Reference
	Second Reference
}

func NewRelationshipElement(sme api.ISubmodelElement) (*RelationshipElement, bool) {
	if sme.RelationshipElement == nil {
		return nil, false
	}

	return &RelationshipElement{
		SubmodelElement: SubmodelElement{
			Extensions:                 sme.RelationshipElement.Extensions,
			Category:                   strFromPtr(sme.RelationshipElement.Category),
			Checksum:                   strFromPtr(sme.RelationshipElement.Checksum),
			Description:                sme.RelationshipElement.Description,
			DisplayName:                sme.RelationshipElement.DisplayName,
			IDShort:                    strFromPtr(sme.RelationshipElement.IdShort),
			Kind:                       sme.RelationshipElement.Kind,
			SemanticID:                 NewReference(sme.RelationshipElement.SemanticId),
			ModelType:                  sme.RelationshipElement.ModelType,
			SupplementalSemanticIDs:    apiRefSliceToRefSlice(sme.RelationshipElement.SupplementalSemanticIds),
			Qualifiers:                 sme.RelationshipElement.Qualifiers,
			EmbeddedDataSpecifications: sme.RelationshipElement.EmbeddedDataSpecifications,
		},
		First:  *NewReference(&sme.RelationshipElement.First),
		Second: *NewReference(&sme.RelationshipElement.Second),
	}, true
}

func (r RelationshipElement) ISubmodelElement() api.ISubmodelElement {
	return api.ISubmodelElement{
		RelationshipElement: &api.RelationshipElement{
			Extensions:                 r.Extensions,
			Category:                   strPtr(r.Category),
			Checksum:                   strPtr(r.Checksum),
			Description:                r.Description,
			DisplayName:                r.DisplayName,
			IdShort:                    strPtr(r.IDShort),
			Kind:                       r.Kind,
			SemanticId:                 r.SemanticID.ptr(),
			SupplementalSemanticIds:    refSliceToAPIRefSlice(r.SupplementalSemanticIDs),
			Qualifiers:                 r.Qualifiers,
			EmbeddedDataSpecifications: r.EmbeddedDataSpecifications,
			// RelationshipElement
			ModelType: api.MODELTYPE_RELATIONSHIP_ELEMENT,
			First:     r.First.Ref,
			Second:    r.Second.Ref,
		},
	}
}
