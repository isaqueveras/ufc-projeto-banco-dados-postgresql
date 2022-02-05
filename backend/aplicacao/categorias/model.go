package categorias

import "time"

// DadosCategoria modela os dados de uma categoria
type DadosCategoria struct {
	ID          *int64     `json:"id,omitempty"`
	Nome        *string    `json:"nome,omitempty"`
	Descricao   *string    `json:"descricao,omitempty"`
	DataCriacao *time.Time `json:"data_criacao,omitempty"`
}

// ListaCategoriasRes modela uma lista de categorias
type ListaCategoriasRes struct {
	TotalCategorias *int64           `json:"total_categorias,omitempty"`
	Dados           []DadosCategoria `json:"dados,omitempty"`
}
