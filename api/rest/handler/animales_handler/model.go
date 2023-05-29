package animales_handler

type AnimalRequest struct {
	ID       int    `json:"id"`
	Id_vet   int    `json:"id_vet"`
	Usuario  string `json:"usuario"`
	Password string `json:"password"`
	Nombres  string `json:"nombres"`
	Raza     string `json:"raza" `
	Edad     string `json:"edad" `
}

type AnimalResponse struct {
	ID       int    `json:"id"`
	Id_vet   int    `json:"id_vet"`
	Usuario  string `json:"usuario"`
	Password string `json:"password"`
	Nombres  string `json:"nombres"`
	Raza     string `json:"raza" `
	Edad     string `json:"edad" `
}
