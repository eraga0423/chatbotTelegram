package models

type GeminiReqest struct {
	Contents []Content `json:"contents"`
}
type Content struct {
	Parts []Part `json:"parts"`
}
type Part struct {
	Text string `json:"text"`
}

type GeminiResponse struct {
	Outputs []Output `json:"outputs"`
}
type Output struct {
	Parts []Part `json:"parts"`
}
