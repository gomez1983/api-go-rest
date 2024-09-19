package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() { /** Define a função para conectar com o banco de dados **/
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5433 sslmode=disable" /** Define a string de conexão com os dados do banco **/
	DB, err = gorm.Open(postgres.Open(stringDeConexao))                                               /** Tenta abrir a conexão com o banco de dados usando GORM **/
	if err != nil {                                                                                   /** Verifica se ocorreu algum erro na conexão **/
		log.Panic("Erro ao conectar com banco de dados") /** Exibe uma mensagem de erro e interrompe o programa se a conexão falhar **/
	}
}
