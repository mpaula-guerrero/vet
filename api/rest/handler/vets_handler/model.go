package vets_handler

type VetRequest struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
}

type VetResponse struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
}
