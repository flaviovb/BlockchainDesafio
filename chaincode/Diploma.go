/*
Copyright IBM Corp 2016 All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
		 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
Implementação iniciada por Caue Garcia Polimanti e Vitor Diego dos Santos de Sousa
*/

// nome do package
package main

// lista de imports
// "encoding/json"	-> enconding para json
// "strconv" -> conversor de strings
import (
	"errors"
	"fmt"


	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// BoletoPropostaChaincode - implementacao do chaincode
type DiplomaGeneralChaincode struct {
}



// Definição da Struct Proposta e parametros para exportação para JSON
type Diploma struct {
	ID					string	`json:"id_diploma"`
	IDStudent				string 	`json:"id_student"`
	NameStudent				string 	`json:"name_student"`
	Approved		 		bool 	`json:"approved"`
	IDInstitution 				string 	`json:"id_institution"`
}

const (
	idDiploma		=	"id"
	nameTableDiploma	=	"Diploma"
	colIdStudent		=	"idStudent"
	colApproved		=	"approved"
	colNameStudent		=	"nameStudent"
	colIDInstitution	=	"idInstitution"
)

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(DiplomaGeneralChaincode))
	if err != nil {
		fmt.Printf("Error starting DiplomaGeneralChaincode chaincode: %s", err)
	}
}

// ============================================================================================================================
// Init
// 		Inicia/Reinicia a tabela de propostas
// ============================================================================================================================
func (t *DiplomaGeneralChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	//myLogger.Debug("Init Chaincode...")
	fmt.Println("Init Chaincode...")

	// Verificação da quantidade de argumentos recebidos
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	// Verifica se a tabela 'Proposta' existe
	fmt.Println("Verificando se a tabela " + nameTableDiploma + " existe...")
	tbProposta, err := stub.GetTable(nameTableDiploma)
	if err != nil {
		fmt.Println("Falha ao executar stub.GetTable para a tabela " + nameTableDiploma + ". [%v]", err)
	}
	// Se a tabela 'Proposta' já existir, excluir a tabela
	if tbProposta != nil {
		err = stub.DeleteTable(nameTableDiploma)
		fmt.Println("Tabela " + nameTableDiploma + " excluída.")
	}


	// Criar tabela de Propostas
	fmt.Println("Criando a tabela " + nameTableDiploma + "...")


	err = stub.CreateTable(nameTableDiploma, []*shim.ColumnDefinition{
		// Id Diploma (hash)
		&shim.ColumnDefinition{Name: idDiploma, Type: shim.ColumnDefinition_STRING, Key: true},
		// Student ID
		&shim.ColumnDefinition{Name: colIdStudent, Type: shim.ColumnDefinition_STRING, Key: false},
		// Flag identifying if the student was approved
		&shim.ColumnDefinition{Name: colApproved, Type: shim.ColumnDefinition_BOOL, Key: false},
		// Student name
		&shim.ColumnDefinition{Name: colNameStudent, Type: shim.ColumnDefinition_STRING, Key: false},
		// Institution ID
		&shim.ColumnDefinition{Name: colIDInstitution, Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, fmt.Errorf("Falha ao criar a tabela " + nameTableDiploma + ". [%v]", err)
	}
	fmt.Println("Table " + nameTableDiploma + " was created.")

	fmt.Println("Init Chaincode... Finalizado!")

	return nil, nil
}



// ============================================================================================================================
// Invoke Functions
// ============================================================================================================================

// Invoke - Ponto de entrada para chamadas do tipo Invoke.
// Funções suportadas:
// "init": inicializa o estado do chaincode, também utilizado como reset
// "registrarProposta(Id, cpfPagador, pagadorAceitou, beneficiarioAceitou,
// boletoPago)": para registrar uma nova proposta ou atualizar uma já existente.
func (t *DiplomaGeneralChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Invoke Chaincode...")
	fmt.Println("invoke is running " + function)

	// Estrutura de Seleção para escolher qual função será executada,
	// de acordo com a funcao chamada
	if function == "init" {
		fmt.Println("Firing init")
		return t.Init(stub,"init",args)

	}else if function == "issueDiploma" {
		fmt.Println("Firing issueDiploma")
		//Create an asset with some value
		return t.issueDiploma(stub, args)
	}

	fmt.Println("invoke não encontrou a func: " + function) //error

	return nil, errors.New("Invocação de função desconhecida: " + function)
}

// registrarProposta: função Invoke para registrar uma nova proposta, recebendo os seguintes argumentos:
// args[0]: Id. Hash que identificará a proposta
// args[1]: cpfPagador. CPF do Pagador
// args[2]: pagadorAceitou. Status de aceite do Pagador da proposta
// args[3]: beneficiarioAceitou. Status de aceite do Beneficiario da proposta
// args[4]: boletoPago. Status do Pagamento do Boleto
func (t *DiplomaGeneralChaincode) issueDiploma(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("issueDiploma...")

	// Verifica se a quantidade de argumentos recebidas corresponde a esperada





	// Obtem os valores da array de arguments (args) e
	// os converte no tipo necessário para salvar na tabela 'Proposta'














	// Registra a proposta na tabela 'Proposta'

















	// Caso a proposta já exista

	// Trecho para atualizar uma proposta existente
	// substitui um registro existente em uma linha com o registro associado ao idProposta recebido nos argumentos









	fmt.Println("Diploma created")

	return nil, nil
}


// ============================================================================================================================
// Query
// ============================================================================================================================

// Query is our entry point for queries

// Query - Ponto de entrada para chamadas do tipo Query.
// Funções suportadas:
// "consultarProposta(Id)": para consultar uma proposta existente
func (t *DiplomaGeneralChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query Chaincode...")

	fmt.Println("query is running " + function)

	// Estrutura de Seleção para escolher qual função será executada,
	// de acordo com a funcao chamada






	fmt.Println("query encontrou a func: " + function)

	return nil, errors.New("Query de função desconhecida: " + function)
}

// consultarProposta: função Query para consultar uma proposta existente, recebendo os seguintes argumentos
// args[0]: Id. Hash da proposta
func (t *DiplomaGeneralChaincode) consultarProposta(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("consultarProposta...")

	//var resProposta Proposta		// Proposta
	var propostaAsBytes []byte		// retorno do json em bytes

	// Verifica se a quantidade de argumentos recebidas corresponde a esperada





	// Obtem os valores dos argumentos e os prepara para salvar na tabela 'Proposta'




	// Define o valor de coluna do registro a ser buscado




	// Consultar a proposta na tabela 'Proposta'





	// Tratamento para o caso de não encontrar nenhuma proposta correspondente






	// Criação do objeto Proposta









	// Converter o objeto da Proposta para Bytes, para retorná-lo em formato JSON





	// retorna o objeto em bytes
	return propostaAsBytes, nil
}
