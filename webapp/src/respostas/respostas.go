package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErroAPI struct {
	Erro string `json:"erro"`
}

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent{
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
	

}

func TratarStatusCodeErro(w http.ResponseWriter, r *http.Request, resp *http.Response) {
	var erro ErroAPI
	json.NewDecoder(resp.Body).Decode(&erro)
	if erro.Erro == "Token is expired" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	JSON(w, resp.StatusCode, erro)
}
