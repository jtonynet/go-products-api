# [<img src="./docs/assets/images/icons/go-original.svg" width="30px" height="30px" alt="Go Logo" title="Go">](https://go.dev/) Go Products API Challenge
<a id="go-products-api"></a>
<!-- 
    Logo image generated by Bing IA: https://www.bing.com/images/create/
-->
<img 
  src="./docs/assets/images/layout/header.png" 
  alt="Gopher azul, simbolo da linguagem Golang empurrando um carrinho de compras com caixas de produtos"
  title="Microsservices GO Products API"
/>

<!-- 
    icons by:
    https://devicon.dev/
    https://simpleicons.org/
-->
[<img src="./docs/assets/images/icons/postman.svg" width="25px" height="25px" alt="Postman Logo" title="Postman">](https://www.postman.com/) [<img src="./docs/assets/images/icons/go.svg" width="25px" height="25px" alt="go" title="Go">](https://go.dev/) [<img src="./docs/assets/images/icons/echo.png" width="25px" height="25px" alt="Echo Framework" title="Echo Framework">](https://echo.labstack.com/) [<img src="./docs/assets/images/icons/mysql.svg" width="50px" height="50px" alt="MySQL Logo" title="MySQL">](https://www.mysql.com/) [<img src="./docs/assets/images/icons/docker.svg" width="25px" height="25px" alt="Docker Logo" title="Docker">](https://www.docker.com/) [<img src="./docs/assets/images/icons/ubuntu.svg" width="25px" height="25px Logo" title="Ubunto" alt="Ubunto" />](https://ubuntu.com/) [<img src="./docs/assets/images/icons/dotenv.svg" width="25px" height="25px" alt="Viper DotEnv Logo" title="Viper DotEnv">](https://github.com/spf13/viper) [<img src="./docs/assets/images/icons/github.svg" width="25px" height="25px" alt="GitHub Logo" title="GitHub">](https://github.com/jtonynet) [<img src="./docs/assets/images/icons/mermaidjs.svg" width="25px" height="25px" alt="MermaidJS Logo" title="MermaidJS">](https://mermaid.js.org/) [<img src="./docs/assets/images/icons/visualstudiocode.svg" width="25px" height="25px" alt="VsCode Logo" title="VsCode">](https://code.visualstudio.com/) [<img src="./docs/assets/images/icons/swagger.svg" width="25px" height="25px" alt="Swagger Logo" title="Swagger">](https://swagger.io/) [<img src="./docs/assets/images/icons/githubactions.svg" width="25px" height="25px" alt="GithubActions Logo" title="GithubActions">](https://docs.github.com/en/actions) [<img src="./docs/assets/images/icons/prometheus.svg" width="25px" height="25px" alt="Prometheus Logo" title="Prometheus">](https://prometheus.io/) [<img src="./docs/assets/images/icons/grafana.svg" width="25px" height="25px" alt="Grafana Logo" title="Grafana">](https://grafana.com/) <!-- [<img src="./docs/assets/images/icons/gatling.svg" width="35px" height="35px" alt="Gatling Logo" title="Gatling">](https://gatling.com/) [<img src="./docs/assets/images/icons/redis.svg" width="25px" height="25px" alt="Redis Logo" title="Redis">](https://redis.com/) [<img src="./docs/assets/images/icons/rabbitmq.svg" width="25px" height="25px" alt="RabbitMQ Logo" title="RabbitMQ">](https://rabbitmq.com/) -->


![Badge Status](https://img.shields.io/badge/STATUS-ENCERRADO_🔐-blue) [![Github Project](https://img.shields.io/badge/PROJECT%20VIEW%20KANBAM-GITHUB-green?logo=github&logoColor=white)](https://github.com/users/jtonynet/projects/5/views/1) ![Badge GitHubActions](https://github.com/jtonynet/go-products-api/actions/workflows/main.yml/badge.svg?branch=main) 

_*Status `ENCERRADO` - Primeiro commit dia: 07/02/2024 Ultimo commit dia 12/02/2024 - 5 dias corridos_

---

## 🕸️ Encontre-me na web
[![linkedin](https://img.shields.io/badge/Linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/jos%C3%A9-r-99896a39/) [![dev.to](https://img.shields.io/badge/dev.to-0A0A0A?style=for-the-badge&logo=devdotto&logoColor=white)](https://dev.to/learningenuity) [![gmail](https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:learningenuity@gmail.com) [![Twitter](https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://twitter.com/learningenuity) [![instagram](https://img.shields.io/badge/Instagram-E4405F?style=for-the-badge&logo=instagram&logoColor=white)](https://www.instagram.com/learningenuity) 

---

## 📁 O Projeto

<a id="index"></a>
### ⤴️ index

[Go Products API](#go-products-api)<br/>
  1.   ⤴️ [Índice](#arrow_heading_up-index)
  2.   📗 [Sobre](#about)
  3.   💻 [Rodando o Projeto](#run)
  4.   📰 [Documentação da API](#api-docs)
  5.   📊 [Diagramas](#diagrams)
  6.   ✅ [Testes](#tests)
  7.   🪲 [Debug](#debug)
  8.   🕵️ [Observabilidade](#observability)
  9.   👏 [Boas Práticas](#best-practices)
  10.  🧠 [ADR - Architecture Decision Records](#adr)
  11.  🔢 [Versões](#versions)
  12.  🧰 [Ferramentas](#tools)
  13.  🤖 [Uso de AI](#ia)
  14.  🏁 [Conclusão](#conclusion)

<br/>

[⤴️ de volta ao índice](#index)

---

<a id="about"></a>
### 📗 Sobre

Este repositório foi criado com a intenção de propor uma possível solução para o seguinte desafio:

> 👨‍💻 __Case Dev Backend:__
> 
> __Vamos construir uma API?__
> A API a ser desenvolvida deve conter rotas C.R.U.D. de produtos, seguindo um bom design de API.
> - Requisitos:
>   - Rota para criação de produtos.
>   - Rota para consulta de produtos.
>   - Rota para atualização de produtos.
>   - Rota para exclusão de produtos.
> 
> Sinta-se livre para aprimorar a API e demonstrar seus conhecimentos.
>
> - Stack Back-end
>   - Postman
>   - GoLang
>   - Echo Framework
>   - MySQL
>
> - Padrões de qualidade.
>   - API Design
>   - Código Limpo (Clean Code)
>   - Padrões e convenções GoLang

Dada sua simplicidade, uma vez que se trata de um __CRUD simples__. Faz sentido utilizar __Arquitetura de Duas Camadas__ o que aumenta o ritmo do desenvolvimento porém mantendo a qualidade do resultado final. Construido junto a uma visualização estilo __Kanbam__ no [Github Project](https://github.com/users/jtonynet/projects/5/views/1)

Foco em garantir estabilidade com __TDD__ e uma implementação minima de __CI__ no [GitHub Actions](https://github.com/jtonynet/go-products-api/actions), boas práticas de design de código e __Observabilidade__ com Prometheus e Grafana garantindo a robustez necessária e devidos aprimoramentos

Exatamente por parecer ser "simples", é necessário ter um bom nível de qualidade.

O desenvolvimento e os testes foram executados no sistema __Linux Ubuntu 20.04.6 LTS__ e testada em __Linux Ubuntu 22.04.3 LTS__.


<br/>

[⤴️ de volta ao índice](#index)

---

<a id="run"></a>
### 💻 Rodando o Projeto

Crie uma cópia do arquivo `sample.env` com o nome `.env` e rode o comando `docker compose` (de acordo com sua versão do `docker compose`) no diretorio raiz do projeto:
```bash
$ docker compose build
$ docker compose up
```

<br/>

####  <img src="./docs/assets/images/icons/postman.svg" width="20px" height="20px" alt="Swagger" title="Swagger">  Postman collection

<details>
  <summary>Dentro da pasta <a href="./scripts/postman-collections/go-products-api.postman_collection.json">./scripts/postman-collection/go-products-api.postman_collection.json</a> encontra-se o arquivo JSON básico que pode ser importado no seu <i>Postman</i> para auxiliar em testes manuais e desenvolvimento.</summary>
  <img src="./docs/assets/images/screen_captures/postman.png">
</details>

<br/>

[⤴️ de volta ao índice](#index)

---
<a id="api-docs"></a>
### 📰  Documentação da API

Com a aplicação em execução, a rota de documentação Swagger fica disponível em http://localhost:8080/swagger/index.html#/ .

Acesse-a para realizar validações, caso prefira ao usar o Postman. Utilizar o Swagger-API também é uma boa maneira de tornar a aplicação aderente às boas práticas e ao design de API.

O cliente deve informar o UUID do recurso, seguindo as modernas práticas de desenvolvimento. Para validações, você pode utilizar um [site gerador de UUIDs](https://www.uuidgenerator.net/). Outras restrições e características dos recursos que podem ser criados, obtidos, listados e deletados podem ser visualizadas na seção 'Models' do Swagger.

####  <img src="./docs/assets/images/icons/swagger.svg" width="20px" height="20px" alt="Swagger" title="Swagger"/> Swagger docs

<img src="./docs/assets/images/screen_captures/swagger_routes.png"/>

<img src="./docs/assets/images/screen_captures/swagger_models.png"/>


<br/>

__Gerando a Documentação:__

Para gerar a documentação, você precisa ter o Golang, o [Swaggo](https://github.com/swaggo/swag) e o projeto instalados localmente e executar o seguinte comando no diretório raiz do projeto:

```bash
swag init --generalInfo cmd/api/main.go -o ./api
```

<br/>

> :writing_hand: **Nota**:
>
> Existe um bug na geração da documentação do Echo Swagger. No arquivo [./api/docs.go](./api/docs.go) está gerando duas propriedades na struct `swag.Spec` que não são reconhecidas. A solução encontrada no momento é remover manualmente essas propriedades a cada geração da documentação.
>
> ```go
>	LeftDelim:        "{{",
>	RightDelim:       "}}",
>```
> <i>*A geração da documentação necessita de mais testes no S.O. Windows</i>
>

<br/>

[⤴️ de volta ao índice](#index)

---

<a id="diagrams"></a>
### 📊 Diagramas do Sistema

<!-- 
    diagrams by:
    https://mermaid.js.org/
-->
```mermaid
graph LR
    subgraph User Flow
      ADMIN(["👤 User"])

      ADMIN --> CREATE_PRODUCT("💻 Create Product")
      ADMIN --> RETRIEVE_PRODUCT_LIST("💻 Retrieve Product List")
      ADMIN --> RETRIEVE_PRODUCT("💻 Retrieve Product")
      ADMIN --> UPDATE_PRODUCT("💻 Update Product")
      ADMIN --> REMOVE_PRODUCT("💻 Remove Product")
    end

    subgraph go-products-api - Two Tier Architecture -
      subgraph Handlers
        API_CREATE_PRODUCT("🖥️ Create Product")
        API_GET_PRODUCTS("🖥️ Get Products")
        API_GET_PRODUCT("🖥️ Get Product by UUID")
        API_UPDATE_PRODUCT("🖥️ Update Product by UUID")
        API_DELETE_PRODUCT_BY_ID("🖥️ Delete Product by UUID")
      end


      subgraph Entities
        ENTITY_PRODUCT("📄 Product")
      end

      subgraph DATABASE
        CATALOGO_DB[("🗄️ MySQL <br/> catalogo-db")]
      end 
    end
  

  CREATE_PRODUCT -->|http POST| API_CREATE_PRODUCT
  RETRIEVE_PRODUCT_LIST -->|http GET| API_GET_PRODUCTS
  RETRIEVE_PRODUCT -->|http GET| API_GET_PRODUCT
  UPDATE_PRODUCT -->|http PATCH| API_UPDATE_PRODUCT
  REMOVE_PRODUCT -->|http DELETE| API_DELETE_PRODUCT_BY_ID


  API_CREATE_PRODUCT-->ENTITY_PRODUCT
  API_GET_PRODUCTS-->ENTITY_PRODUCT
  API_GET_PRODUCT-->ENTITY_PRODUCT
  API_UPDATE_PRODUCT-->ENTITY_PRODUCT 
  API_DELETE_PRODUCT_BY_ID-->ENTITY_PRODUCT




  ENTITY_PRODUCT-->CATALOGO_DB
```
_*Diagrama geral com baixo nível de fidelidade_

<br/>

[⤴️ de volta ao índice](#index)

---

<a id="tests"></a>
### :white_check_mark: Testes



Para testar localmente, é necessário ter o Go v1.21.1 instalado. Execute os para garantir o funcionamento correto da API e do banco de dados. Inicie o banco de dados na raiz do projeto usando docker-compose.

```bash
docker compose up mysql-go-products-api
```

Em outro terminal mas ainda na raiz do projeto, execute o comando:
```bash
go test -v ./ ./internal/handlers/
```

<!-- go test -v -->

obtendo uma saida similar a seguinte:<br/>
<img src="./docs/assets/images/screen_captures/testing.png">

<br/>

<details>
  <summary>Os testes também são executados como parte da rotina minima de <b>CI</b> do <a href="https://github.com/jtonynet/go-products-api/actions">GitHub Actions</a>, garantindo que versões estáveis sejam mescladas na branch principal. O badge <i>TESTS_CI</i> no cabeçalho do arquivo readme é uma ajuda visual para verificar rapidamente a integridade do desenvolvimento.</summary>
  <img src="./docs/assets/images/screen_captures/testing_ci.png">
</details>

<br/>

[⤴️ de volta ao índice](#index)

---

<a id="debug"></a>
### :beetle: Debug
Utilizando o VSCode como editor de código ([maiores informações aqui](https://code.visualstudio.com/docs/languages/go#_debugging)) com a seguinte configuração no arquivo `.vscode/launch.json`:

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch go-products-api",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/cmd/api/main.go",
            "cwd": "${workspaceFolder}",
            "trace": "verbose",
        },
        {
            "name": "Test Smoke Happy Path",
            "type": "go",
            "request": "launch",
            "mode": "test",
            "program":"${workspaceFolder}/main_smoke_test.go",
            "trace": "verbose",
        },
        {
            "name": "Test Integration Corner Cases",
            "type": "go",
            "request": "launch",
            "mode": "test",
            "program":"${workspaceFolder}/internal/handlers/productHandler_integration_test.go",
            "trace": "verbose",
        }
    ]
}
```

Altere a connection com a database para localhost em seu arquivo `.env`
```
DATABASE_HOST=localhost  # mysql-go-products-api | localhost - To run localy (without all docker compose dependency) or debug use localhost
```

E suba o Banco de Dados antes de iniciar seu debug:
```bash
docker compose up mysql-go-products-api
```

<details>
  <summary>Uma seção de Depuração de Testes da aplicação pode ser vista aqui:</summary>
  <img src="./docs/assets/images/screen_captures/testing_debug.png">
</details>

<br/>

[⤴️ de volta ao índice](#index)

---

<a id="observability"></a>
### 🕵️ Observabilidade

__Prometheus:__

Apos rodar com sucesso o `docker compose up` como visto anteriormente, acesse:
- Prometheus - http://localhost:9090/

<img src="./docs/assets/images/screen_captures/prometheus.png">

<br/>

__Configurando o Grafana:__

A primeira vez que executarmos o Grafana, entramos com `usuário/senha` padrão de `admin/admin`. Ele solicita a alteração da senha, para facilitar o desenvolvimento local, alteramos para `admin/12345`.
- Grafana - http://localhost:3000/ (usuário/senha: admin/admin | admin/12345)
  
<details>
  <summary>Uma vez dentro do Grafana em sua primeira execução, também precisamos criar uma conexão Datasource com o Prometheus (que acessamos acima). Procure por <i>`Connections > Add New Connection`</i> digite <i>Prometheus</i> no campo de Search, selecione-o, clique em <i>`Add New Datasource`</i> e configure-o com a URL: <i>http://prometheus-go-products-api:9090</i> e clique no botão <i>Save & test</i> no final da página</summary>
  <img src="./docs/assets/images/screen_captures/grafana_create_prometheus_conn.png">
</details>

<br/>

<details>
  <summary>Agora você pode usar o menu <i>`Dashboards > New > Import`</i> para importar o arquivo <b>dash-go-products-api.json</b> que está localizado no diretório: <a href="./scripts/grafana-dashboards/">./scripts/grafana-dashboards</a>. Acesse o diretório em seu computador, clique e arraste o arquivo para o campo correto especificado pela tela <b>Upload Dashboard JSON File</b></summary>
  <img src="./docs/assets/images/screen_captures/grafana_import_dashboard.png">
</details>

<br/>

<details>
  <summary>Vincule o Dashboard a conexão previamente criada e acesse-o</summary>
  <img src="./docs/assets/images/screen_captures/grafana_import_dashboard_prometheus.png">
</details>

<br/>

Quando adequadamente importado, o Dashboard estará disponível e responderá às solicitações que você pode simular pelo `Postman` ou `Swagger`.

<img src="./docs/assets/images/screen_captures/grafana_red.png">

A partir dessas métricas dos dashboards, temos uma ideia da saúde da API e quais são as reais necessidades de escala que ela deve ter em produção, o que nos dá uma ideia de quais arquiteturas e abordagens poderão ser utilizadas para atender às suas demandas, incluindo testes de carga, possíveis caches defensivos, filas, etc.

Nossas decisões de `Arquitetura` e `Design de Sistemas` devem sempre ser baseadas em dados!

<br/>

[⤴️ de volta ao índice](#index)

---

<a id="best-practices"></a>
### 👏 Boas Práticas

- [Swagger](https://swagger.io/)
- [Github Project - Kanbam](https://github.com/users/jtonynet/projects/5/views/1)
- [Layout padrão de projetos em Go](https://github.com/golang-standards/project-layout/blob/master/README_ptBR.md)
- [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
- [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)
- [ADR - Architecture Decision Records](https://cognitect.com/blog/2011/11/15/documenting-architecture-decisions)
- [Mermaid Diagrams](https://mermaid.js.org)
- [Observabilidade](https://en.wikipedia.org/wiki/Observability_(software)) com:
  - [Prometheus](https://prometheus.io/docs/guides/go-application/)
  - [Grafana](https://grafana.com/)

<br/>

[⤴️ de volta ao índice](#index)

---

<a id="adr"></a> 
### 🧠 ADR - Architecture Decision Records

- [0001: Registro de Decisões de Arquitetura (ADR)](./docs/architecture/decisions/0001-registro-de-decisoes-de-arquitetura.md)
- [0002: Echo, Gorm e MySQL em API TDD com Arquitetura de Duas camadas](./docs/architecture/decisions/0002-echo-gorm-e-mysql-com-arquitetura-de-api-tdd-em-duas-camadas.md)
- [0003: Observabilidade com Prometheus e Grafana](./docs/architecture/decisions/0003-observabilidade-com-prometheus-e-api-tdd-em-duas-camadas.md)


<br/>

[⤴️ de volta ao índice](#index)

---

<a id="versions"></a>
### 🔢 Versões

As tags de versões estão sendo criadas manualmente a medida que o projeto avança com melhorias notáveis. Cada funcionalidade é desenvolvida em uma branch a parte (Branch Based, [feature branch](https://www.atlassian.com/git/tutorials/comparing-workflows/feature-branch-workflow)) quando finalizadas é gerada tag e mergeadas em master.

Para obter mais informações, consulte o [Histórico de Versões](./CHANGELOG.md).

<br/>

[⤴️ de volta ao índice](#index)

---

<a id="tools"></a>
### 🧰 Ferramentas

- Linguagem:
  - [Go v1.21.1](https://go.dev/)
  - [GVM v1.0.22](https://github.com/moovweb/gvm)

- Framework & Libs:
  - [Echo](https://echo.labstack.com/)
  - [Testify](github.com/stretchr/testify)
  - [GORM](https://gorm.io/index.html)
  - [Gjson](https://github.com/tidwall/gjson)
  - [Viper](https://github.com/spf13/viper)
  - [uuid](https://github.com/google/uuid)
  - [Delve](https://github.com/go-delve/delve)
  - [Echo-Swagger](https://github.com/swaggo/echo-swagger)
  - [Exponential Backoff](https://github.com/cenkalti/backoff)
  - [Gopsutil]("github.com/shirou/gopsutil)
  - [Client Prometheus](https://github.com/prometheus/client_golang)
  - [GORM Prometheus](https://github.com/go-gorm/prometheus)
  - [Go Playground Validator](https://github.com/go-playground/validator)
  <!-- [Zap log](https://github.com/uber-go/zap)  -->

- Infra & Technologias
  - [Docker v24.0.6](https://www.docker.com/)
  - [Docker compose v2.21.0](https://www.docker.com/)
  - [MySQL](https://www.postgresql.org/)
  - [Prometheus](https://prometheus.io/docs/guides/go-application)
  - [Grafana](https://grafana.com/)


- GUIs:
  - [VsCode](https://code.visualstudio.com/)
  - [Postman](https://blog.postman.com/introducing-the-postman-vs-code-extension/)
  - [MySQL Workbench](https://www.mysql.com/products/workbench/)

<br/>

[⤴️ de volta ao índice](#index)

---

<a id="ia"></a>
### 🤖 Uso de AI

A [imagem do cabeçalho](#go-products-api) desta página foi criada com o auxílio de inteligência artificial e um mínimo de retoque e construção no Gimp [<img src="./docs/assets/images/icons/gimp.svg" width="30" height="30 " title="Gimp" alt="Gimp Logo" />](https://www.gimp.org/)

__Foi utilizado os seguinte prompt para sua criação no [Bing IA:](https://www.bing.com/images/create/)__

<details>
  <summary><b>Gopher com carrinho de compras</b></summary>
<i>"gopher azul, simbolo da linguagem golang, empurrando um carinho de compras com caixas escrito REST, API, ECHO e Swagger dentro desse carrinho, estilo cartoon, historia em quadrinhos, fundo branco chapado para facilitar remoção"<b>(sic)</b></i>
</details>

<br/>

IA também é utilizada em minhas pesquisas e estudos como ferramenta de apoio; no entanto, __artes e desenvolvimento são, sobretudo, atividades criativas humanas.__

Contrate artistas para projetos comerciais ou mais elaborados e Aprenda Engenhosidade!

<br/>

[⤴️ de volta ao índice](#index)

---

<a id="conclusion"></a>
### 🏁 Conclusão

Conduzi o desafio no espírito do TDD e das melhores práticas que conheço. Preocupando-me com a coleta de métricas visando expansão futura, caso os dados apontem essa necessidade, e utilização de ferramentas que fazem parte do parque de tecnologias do proponente do desafio. Mantive-me focado nos prazos das tarefas e na qualidade do código desenvolvido.

Aqui deixo alguns pontos que imagino serem melhorias possíveis (alguns merecem seus próprios ADRs):

- Bug do Swagger, cuja solução manual que adotei impediu a geração da documentação automaticamente a cada subida da aplicação.
- Melhoria no script do Postman para melhor compartilhamento entre equipes.
- Montagem do `docker-compose` que poderia incluir automaticamente o Dashboard Grafana.
- Melhorias no Log e utilização de algum centralizador de Logs como Grafana Loki ou Logstash
- Utilização de algum APM
- Processo de CD
- Testes de Performance e Carga que poderiam rodar em conjunto com nossa monitoria local. Gostaria de ter utilizado o Gatling com essa finalidade.
- A partir de tais testes, poderíamos decidir escalar:
  - Implementação de um cache defensivo com Redis para consultas.
  - Criação de produtos em fila RabbitMQ ou outro, com workers, permitindo o aumento de requisições na rota POST, se necessário.
  
Algumas das melhorias citadas (Gatling, Cache Defensivo e Filas) desenvolvi como POC em [outro projeto](https://github.com/jtonynet/api-gin-rest) que não está na solução "ótima", mas servem para nortear possibilidades.

Terei prazer em discutir sobre o assunto e sinta-se à vontade para visitar outros dos meus repositórios.

😊🚀

<br/>

[⤴️ de volta ao índice](#index)
