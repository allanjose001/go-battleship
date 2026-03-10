# Estrutura do Projeto (Provisório)

Este documento descreve a organização atual do projeto e as
responsabilidades de cada parte.\
O objetivo é manter separação clara entre **lógica do jogo** e
**interface/renderização**.

------------------------------------------------------------------------

## Visão Geral

O projeto segue o padrão recomendado em Go para aplicações maiores,
adaptado para jogos com **Ebiten**.

    go-battleship/
    ├── assets/
    ├── game/
    ├── internal/
    ├── cmd/ 
    ├   └── battleship
    ├       └── main.go
    ├── go.mod
    ├── go.sum
    └── README.md

------------------------------------------------------------------------


## cmd/ (Entrypoints)

Responsável por conter **pontos de entrada executáveis** da aplicação.

Uso: - Cada subpasta gera um binário diferente - Ideal para separar
versões, modos ou ferramentas


------------------------------------------------------------------------

## game/ (Camada de Jogo / UI)

Responsável por: - Integração com Ebiten - Loop do jogo (`Update`,
`Draw`, `Layout`) - Cenas e estados - Input do jogador - Chamada dos
serviços do domínio

Arquivos típicos:

    game/
    ├── game.go
    ├── scene.go
    └── state.go

Regra: - `game` pode importar `internal` - `internal` nunca importa
`game`

------------------------------------------------------------------------

## internal/ (Domínio do Jogo)

Contém toda a lógica do Battleship, sem dependência de Ebiten.

### internal/entity/

Entidades do domínio.

Exemplos:

    entity/
    ├── board.go
    ├── ship.go
    ├── fleet.go
    └── position.go

------------------------------------------------------------------------

### internal/service/

Regras de negócio e coordenação do jogo.

Exemplos:

    service/
    ├── match_service.go
    ├── board_service.go
    └── attack_service.go

------------------------------------------------------------------------

### internal/system/

Sistemas que orquestram fluxo e eventos.

Exemplos:

    system/
    ├── input_system.go
    └── turn_system.go

------------------------------------------------------------------------

### internal/ai/

Lógica da IA.

Exemplos:

    ai/
    ├── random_ai.go
    └── smart_ai.go

------------------------------------------------------------------------

## Dependências

O projeto usa Go Modules.

Após clonar:

``` bash
go mod tidy
```

Para rodar o jogo:

``` bash
go run cmd/battleship/main.go
```

------------------------------------------------------------------------

Integrantes:
* **[Hugo Matheus Costa Araújo](https://github.com/hugomtths)**
* **[Luís Henrique Domingos da Silva](https://github.com/LuisH07)**
* **[Maria Luiza Bezerra dos Santos](https://github.com/marialuizab11)**
* **[Allan José Marinho](https://github.com/allanjose001)**
* **[Alcielma Luzinete da Silva](https://github.com/Alcielma)**
* **[Cauã de Souza Santos](https://github.com/cauaofsouza)**
* **[Gabriel Saraiva](https://github.com/gabrielsaraiv4)**
