package model

type GraphPath []interface {}

type GraphErr struct {
	Message string `json:"message"`
	Path GraphPath `json:"path"`
}
