package utils

import "time"

// PonteiroInt retorna um ponteiro para um inteiro indicado
func PonteiroInt(i int) *int {
	return &i
}

// PonteiroInt64 retorna um ponteiro para um inteiro de 64 bits indicado
func PonteiroInt64(i int64) *int64 {
	return &i
}

// PonteiroString retorna um ponteiro para uma string dada
func PonteiroString(i string) *string {
	return &i
}

// PonteiroMapInt retorna um ponteiro para um mapa de string interface indicado
func PonteiroMapInt(m map[string]interface{}) *map[string]interface{} {
	return &m
}

// PonteiroFloat64 retorna um ponteiro para o um float de 64 bits indicado
func PonteiroFloat64(f float64) *float64 {
	return &f
}

// PonteiroBool retorna um ponteiro para o booleano especificado
func PonteiroBool(i bool) *bool {
	return &i
}

// PonteiroTime retorna um ponteiro para um time.Time especificado
func PonteiroTime(t time.Time) *time.Time {
	return &t
}
