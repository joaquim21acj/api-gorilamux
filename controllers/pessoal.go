package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pessoalAPI-gorilamux/db"

	"github.com/gorilla/mux"
)

type Servidor struct {
	ID                 int    `db:"id, primarykey, autoincrement" json:"id"`
	Siape              string `db:"siape" json:"siape"`
	Id_pessoa          int    `db:"id_pessoa" json:"id_pessoa"`
	Nome               string `db:"nome" json:"nome"`
	Matricula_interna  string `db:"matricula_interna" json:"matricula_interna"`
	Nome_identificacao string `db:"nome_identificacao" json:"nome_identificacao"`
	Data_nascimento    string `db:"data_nascimento" json:"data_nascimento"`
	Sexo               string `db:"sexo" json:"sexo"`
}

type ServidorController struct{}

func (ctrl ServidorController) GetServidor(w http.ResponseWriter, r *http.Request) {
	q := `select s.id_servidor, s.siape, s.id_pessoa, s.matricula_interna, s.nome_identificacao,
		p.nome, p.data_nascimento, p.sexo from rh.servidor s
	inner join comum.pessoa p on (s.id_pessoa = p.id_pessoa)`

	rows, err := db.GetDB().Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pessoas []Servidor

	var id, id_pessoa int
	var siape, nome, matricula_interna, nome_identificacao, data_nascimento, sexo string
	for rows.Next() {
		err := rows.Scan(&id, &siape, &id_pessoa, &matricula_interna, &nome_identificacao, &nome, &data_nascimento, &sexo)
		if err != nil {
			log.Fatal(err)
		}

		pessoas = append(pessoas, Servidor{
			ID:                 id,
			Siape:              siape,
			Id_pessoa:          id_pessoa,
			Nome:               nome,
			Matricula_interna:  matricula_interna,
			Nome_identificacao: nome_identificacao,
			Data_nascimento:    data_nascimento,
			Sexo:               sexo,
		})
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(&pessoas)

	if err != nil {
		return
	}

	return
}

func (ctrl ServidorController) GetServidorMat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mat := vars["matricula"] // URL parameter
	// Data security checking to be insterted here

	q := fmt.Sprintf(`select s.id_servidor, s.siape, s.id_pessoa, s.matricula_interna, s.nome_identificacao,
		p.nome, p.data_nascimento, p.sexo from rh.servidor s
		inner join comum.pessoa p on (s.id_pessoa = p.id_pessoa) where s.matricula_interna = %s`, mat) //String formating

	rows, err := db.GetDB().Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pessoas []Servidor

	var id, id_pessoa int
	var siape, nome, matricula_interna, nome_identificacao, data_nascimento, sexo string
	for rows.Next() {
		err := rows.Scan(&id, &siape, &id_pessoa, &matricula_interna, &nome_identificacao, &nome, &data_nascimento, &sexo)
		if err != nil {
			log.Fatal(err)
		}

		pessoas = append(pessoas, Servidor{
			ID:                 id,
			Siape:              siape,
			Id_pessoa:          id_pessoa,
			Nome:               nome,
			Matricula_interna:  matricula_interna,
			Nome_identificacao: nome_identificacao,
			Data_nascimento:    data_nascimento,
			Sexo:               sexo,
		})
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(&pessoas)

	if err != nil {
		return
	}

	return
}
