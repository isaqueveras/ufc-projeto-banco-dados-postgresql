package configuracao

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var config *Configuracao

// Carregar tenta ler do arquivo config.json um JSON válido com todas as configurações
func Carregar() {
	path, err := ioutil.ReadFile("/etc/projeto-fbd/config.json")
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(path, &config); err != nil {
		log.Fatal(err)
	}
}

// Obter obtem os dados das configurações
func Obter() *Configuracao {
	return config
}
