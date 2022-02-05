package categorias

import "time"

// DadosCategoria modela os dados de uma categoria
type DadosCategoria struct {
	ID          *int64
	Nome        *string
	Descricao   *string
	DataCriacao *time.Time
}

// ListaCategorias modela uma lista de categorias
type ListaCategorias struct {
	TotalCategorias *int64
	Dados           []DadosCategoria
}
