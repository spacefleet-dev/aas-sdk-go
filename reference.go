package aassdkgo

import (
	"github.com/spacefleet-dev/aas-sdk-go/api"
)

type Reference struct {
	Ref api.Reference
}

func NewReference(ref *api.Reference) *Reference {
	if ref == nil {
		return nil
	}

	return &Reference{*ref}
}

func (r Reference) FirstKey() *api.Key {
	if len(r.Ref.Keys) != 0 {
		return &r.Ref.Keys[0]
	}

	return nil
}

func (r Reference) NumKeys() int {
	return len(r.Ref.Keys)
}

func (r *Reference) ptr() *api.Reference {
	if r == nil {
		return nil
	}

	return &r.Ref
}

func refSliceToAPIRefSlice(refs []Reference) []api.Reference {
	res := make([]api.Reference, len(refs))

	for i := range refs {
		res[i] = refs[i].Ref
	}

	return res
}

func apiRefSliceToRefSlice(refs []api.Reference) []Reference {
	res := make([]Reference, len(refs))

	for i := range refs {
		res[i].Ref = refs[i]
	}

	return res
}
