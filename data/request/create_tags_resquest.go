package request

type CreateTagsRequest struct {
	Name string `validate:"required, min=2,max=10" json:"name"`
}
