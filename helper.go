package aassdkgo

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/spacefleet-dev/aas-sdk-go/api"
)

func findSubmodelElement(smes []api.ISubmodelElement, idShort string) AnySubmodelElement {
	for i := range smes {
		if idShort == getIDShortFromISubmodelElement(smes[i]) {
			return NewAnySubmodelElement(smes[i])
		}
	}

	return nil
}

func findAllSubmodelElements(smes []api.ISubmodelElement, idShort string) []AnySubmodelElement {
	var res []AnySubmodelElement

	for _, sme := range smes {
		if idShort == getIDShortFromISubmodelElement(sme) {
			res = append(res, NewAnySubmodelElement(sme))
		}
	}

	return res
}

func findAllSubmodelElementsByModelType(smes []api.ISubmodelElement, modelType api.ModelType) []AnySubmodelElement {
	var res []AnySubmodelElement

	for _, sme := range smes {
		if modelType == getModelTypeFromISubmodelElement(sme) {
			res = append(res, NewAnySubmodelElement(sme))
		}
	}

	return res
}

func encodePropertyValue(v any) (string, api.DataTypeDefXsd) {
	switch v := v.(type) {
	case string:
		return fmt.Sprint(v), api.DATATYPEDEFXSD_STRING
	case int, int8, int16, int32:
		return fmt.Sprint(v), api.DATATYPEDEFXSD_INT
	case int64:
		return fmt.Sprint(v), api.DATATYPEDEFXSD_LONG
	case uint, uint16, uint32, uint64:
		return fmt.Sprint(v), api.DATATYPEDEFXSD_NON_NEGATIVE_INTEGER
	case bool:
		return fmt.Sprint(v), api.DATATYPEDEFXSD_BASE64_BINARY
	case float32:
		return fmt.Sprint(v), api.DATATYPEDEFXSD_FLOAT
	case float64:
		return fmt.Sprint(v), api.DATATYPEDEFXSD_DOUBLE
	case byte:
		return string(v), api.DATATYPEDEFXSD_BYTE
	case []byte:
		return base64.StdEncoding.EncodeToString(v), api.DATATYPEDEFXSD_BASE64_BINARY
	case time.Time:
		return fmt.Sprint(v.Unix()), api.DATATYPEDEFXSD_DATE_TIME_STAMP
	case time.Duration:
		return v.String(), api.DATATYPEDEFXSD_DAY_TIME_DURATION
	}

	return fmt.Sprint(v), api.DATATYPEDEFXSD_STRING
}

func strPtr(str string) *string {
	if str != "" {
		return &str
	}

	return nil
}

func strFromPtr(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}

func getIDShortFromISubmodelElement(sme api.ISubmodelElement) string {
	if sme.AnnotatedRelationshipElement != nil {
		return strFromPtr(sme.AnnotatedRelationshipElement.IdShort)
	}

	if sme.BasicEventElement != nil {
		return strFromPtr(sme.BasicEventElement.IdShort)
	}

	if sme.Blob != nil {
		return strFromPtr(sme.Blob.IdShort)
	}

	if sme.Entity != nil {
		return strFromPtr(sme.Entity.IdShort)
	}

	if sme.File != nil {
		return strFromPtr(sme.File.IdShort)
	}

	if sme.MultiLanguageProperty != nil {
		return strFromPtr(sme.MultiLanguageProperty.IdShort)
	}

	if sme.Operation != nil {
		return strFromPtr(sme.Operation.IdShort)
	}

	if sme.Property != nil {
		return strFromPtr(sme.Property.IdShort)
	}

	if sme.Range != nil {
		return strFromPtr(sme.Range.IdShort)
	}

	if sme.ReferenceElement != nil {
		return strFromPtr(sme.ReferenceElement.IdShort)
	}

	if sme.RelationshipElement != nil {
		return strFromPtr(sme.RelationshipElement.IdShort)
	}

	if sme.SubmodelElement != nil {
		return strFromPtr(sme.SubmodelElement.IdShort)
	}

	if sme.SubmodelElementCollection != nil {
		return strFromPtr(sme.SubmodelElementCollection.IdShort)
	}

	if sme.SubmodelElementList != nil {
		return strFromPtr(sme.SubmodelElementList.IdShort)
	}

	return ""
}

func getModelTypeFromISubmodelElement(sme api.ISubmodelElement) api.ModelType {
	if sme.AnnotatedRelationshipElement != nil {
		return sme.AnnotatedRelationshipElement.ModelType
	}

	if sme.BasicEventElement != nil {
		return sme.BasicEventElement.ModelType
	}

	if sme.Blob != nil {
		return sme.Blob.ModelType
	}

	if sme.Entity != nil {
		return sme.Entity.ModelType
	}

	if sme.File != nil {
		return sme.File.ModelType
	}

	if sme.MultiLanguageProperty != nil {
		return sme.MultiLanguageProperty.ModelType
	}

	if sme.Operation != nil {
		return sme.Operation.ModelType
	}

	if sme.Property != nil {
		return sme.Property.ModelType
	}

	if sme.Range != nil {
		return sme.Range.ModelType
	}

	if sme.ReferenceElement != nil {
		return sme.ReferenceElement.ModelType
	}

	if sme.RelationshipElement != nil {
		return sme.RelationshipElement.ModelType
	}

	if sme.SubmodelElement != nil {
		return sme.SubmodelElement.ModelType
	}

	if sme.SubmodelElementCollection != nil {
		return sme.SubmodelElementCollection.ModelType
	}

	if sme.SubmodelElementList != nil {
		return sme.SubmodelElementList.ModelType
	}

	return ""
}
