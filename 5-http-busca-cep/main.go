package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", RequestHandler)
	http.ListenAndServe(":8080", nil)
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cepResult, error := SearchCEP(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	//Primeira alternativa para retornar sua requisição
	jsonResult, error := json.Marshal(cepResult)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(jsonResult)

	//Segunda alternativa para retornar sua requisição - Aqui está transformando a struct em json e jogando para o "w" para retornar a resposta
	json.NewEncoder(w).Encode(cepResult)
}

func SearchCEP(cep string) (*ViaCEP, error) {
	res, error := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer res.Body.Close()

	body, error := io.ReadAll(res.Body)
	if error != nil {
		return nil, error
	}

	var c ViaCEP
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}

	return &c, nil
}
