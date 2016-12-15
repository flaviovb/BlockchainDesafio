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
	"strconv"
	"encoding/json"	
	"github.com/hyperledger/fabric/core/chaincode/shim"	
)

// BoletoPropostaChaincode - implementacao do chaincode
type BoletoPropostaChaincode struct {

}

// Definição da Struct Proposta e parametros para exportação para JSON
type Boleto struct {
    ID  int      `json:"id"`
    Name string `json:"name"`	
}



// consts associadas à tabela de Propostas








// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(BoletoPropostaChaincode))
	if err != nil {
		fmt.Printf("Error starting BoletoPropostaChaincode chaincode: %s", err)
	}
}


// GetTable returns the table for the specified table name or ErrTableNotFound
// if the table does not exist.
func (stub *ChaincodeStub) GetTable(tableName string) (*Table, error) {
	return getTable(stub, tableName)
}

func getTable(stub ChaincodeStubInterface, tableName string) (*Table, error) {

	tableName, err := getTableNameKey(tableName)
	if err != nil {
		return nil, err
	}

	tableBytes, err := stub.GetState(tableName)
	if tableBytes == nil {
		return nil, ErrTableNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("Error fetching table: %s", err)
	}
	table := &Table{}
	err = proto.Unmarshal(tableBytes, table)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling table: %s", err)
	}

	return table, nil
}

// ============================================================================================================================
// Init
// 		Inicia/Reinicia a tabela de propostas
// ============================================================================================================================
func (t *BoletoPropostaChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Init Chaincode...")

	// Verificação da quantidade de argumentos recebidos

	if len(args) != 2 {
		return nil, errors.New("Please, I need 2 arguments for BoletoProposta")
	}


	// Verifica se a tabela 'Proposta' existe
	fmt.Println("Verificando se a tabela 'Proposta' existe...")
	
	err := stub.GetTable("BoletoPropostaAsset") 
	
	//tabela existe
	if err != nil {
	
		//exclui tabela
		
		
	}

	
	// Create ownership table
	err := stub.CreateTable("BoletoPropostaAsset", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "ID", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "Name", Type: shim.ColumnDefinition_BYTES, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating AssetBoleto table.")
	}
	


	// Se a tabela 'Proposta' já existir, excluir a tabela
	





	// Criar tabela de Propostas
	fmt.Println("Criando a tabela 'Proposta'...")
	














	fmt.Println("Tabela 'Proposta' criada com sucesso.")

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
func (t *BoletoPropostaChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Invoke Chaincode...")
	fmt.Println("invoke is running " + function)

	// Estrutura de Seleção para escolher qual função será executada, 
	// de acordo com a funcao chamada
	





	fmt.Println("invoke não encontrou a func: " + function) //error

	return nil, errors.New("Invocação de função desconhecida: " + function)
}

// registrarProposta: função Invoke para registrar uma nova proposta, recebendo os seguintes argumentos:
// args[0]: Id. Hash que identificará a proposta
// args[1]: cpfPagador. CPF do Pagador
// args[2]: pagadorAceitou. Status de aceite do Pagador da proposta
// args[3]: beneficiarioAceitou. Status de aceite do Beneficiario da proposta
// args[4]: boletoPago. Status do Pagamento do Boleto
func (t *BoletoPropostaChaincode) registrarProposta(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("registrarProposta...")

	// Verifica se a quantidade de argumentos recebidas corresponde a esperada





	// Obtem os valores da array de arguments (args) e 
	// os converte no tipo necessário para salvar na tabela 'Proposta'














	// Registra a proposta na tabela 'Proposta'
	
















	// Caso a proposta já exista
	
		// Trecho para atualizar uma proposta existente
		// substitui um registro existente em uma linha com o registro associado ao idProposta recebido nos argumentos

		




		


	fmt.Println("Proposta criada!")

	return nil, nil
}


// ============================================================================================================================
// Query
// ============================================================================================================================

// Query is our entry point for queries

// Query - Ponto de entrada para chamadas do tipo Query.
// Funções suportadas:
// "consultarProposta(Id)": para consultar uma proposta existente
func (t *BoletoPropostaChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query Chaincode...")

	fmt.Println("query is running " + function)

	// Estrutura de Seleção para escolher qual função será executada, 
	// de acordo com a funcao chamada






	fmt.Println("query encontrou a func: " + function) 

	return nil, errors.New("Query de função desconhecida: " + function)
}

// consultarProposta: função Query para consultar uma proposta existente, recebendo os seguintes argumentos
// args[0]: Id. Hash da proposta
func (t *BoletoPropostaChaincode) consultarProposta(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
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
