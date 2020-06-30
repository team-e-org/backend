package view

type AttachTagsLambdaPayload struct {
	Pin  *AttachTagsLambdaPayloadPin `json:"pin"`
	Tags []string                    `json:"tags"`
}
