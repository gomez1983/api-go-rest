package models

type Personalidade struct { /** Define uma struct chamada Personalidade **/
	Id       int    `json:"id"`       /** Campo Id do tipo int, será convertido para JSON com o nome "id" **/
	Nome     string `json:"nome"`     /** Campo Nome do tipo string, será convertido para JSON com o nome "nome" **/
	Historia string `json:"historia"` /** Campo Historia do tipo string, será convertido para JSON com o nome "historia" **/
}

var Personalidades []Personalidade /** Declara uma variável slice para armazenar várias personalidades **/
