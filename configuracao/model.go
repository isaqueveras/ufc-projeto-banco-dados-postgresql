package configuracao

// Configuracao estrutura principal da aplicação
type Configuracao struct {
	NomeServidor  string   `json:"nome_servidor,omitempty"`
	PortaServidor string   `json:"porta_servidor,omitempty"`
	Database      database `json:"banco_dados,omitempty"`
}

type database struct {
	Driver string `json:"driver"`
	Url    string `json:"url"`
}
