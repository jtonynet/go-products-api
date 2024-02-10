# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [Unreleased]
### Added

### Fixed

---

## [0.0.6] - 2023-02-10
### Fixed

- Corrigido bug de a aplicação iniciar mas o banco ainda não estar respondendo a solicitações, bug conhecido do docker-compose.Paratanto foi adicionado [Exponential Backoff](https://github.com/cenkalti/backoff), que tenta se conectar ao banco por `API_RETRY_MAX_ELAPSED_TIME_IN_MS` __[Issue-14](https://github.com/jtonynet/go-products-api/issues/14)__ 


---

## [0.0.5] - 2023-02-10
### Added

- Adicionado swagger __[Issue-12](https://github.com/jtonynet/go-products-api/issues/12)__ 

---

## [0.0.4] - 2023-02-10
### Added

- Criação do `cmd/api/main.go` entrypoint da `API` com suas rotas integrado as variáveis de ambiente com `Viper` e sendo Dockerizada na __[Issue-10](https://github.com/jtonynet/go-products-api/issues/10)__ do projeto.
- Script Postman básico

---

## [0.0.3] - 2023-02-09
### Added

- Criação do handler de remoção de Produto conforme descrito na __[Issue-8](https://github.com/jtonynet/go-products-api/issues/8)__ do projeto.

---

## [0.0.2] - 2023-02-09
### Added

- Utilizando GithubAction para performar o Smoke Test a cada commit e push na main conforma descrito na __[Issue-4](https://github.com/jtonynet/go-products-api/issues/4)__ do projeto.
- Criação dos Endpoints de Alteração e Listagem de um Produtos conforme descrito na __[Issue-6](https://github.com/jtonynet/go-products-api/issues/6)__ do projeto.

---

## [0.0.1] - 2023-02-09
### Added

- Criação dos Endpoints de Criação e Obtenção de um Produto conforme descrito na __[Issue-2](https://github.com/jtonynet/go-products-api/issues/2)__ do projeto.
- Utilizando TDD e Smoke Tests para validar o correto funcionamento dos Handlers e recursos a eles vinculados (entity e response, request DTOs).
- Criado `docker-compose.yml` para acesso à base de dados local, de `DEV`, também utilizada no `Smoke Test`.
---

## [0.0.0] - 2023-02-07
### Added

- Iniciado o [Projeto](https://github.com/users/jtonynet/projects/3) com o commit inicial. Documentação base: Readme Rico, ADRs: [0001: Registro de Decisões de Arquitetura (ADR)](./docs/architecture/decisions/registro-de-decisoes-de-arquitetura.md) e [0002: Echo, Gorm e MySQL em API TDD com arquitetura de Duas camadas](./docs/architecture/decisions/0002-echo-gorm-e-mysql-com-arquitetura-de-api-tdd-em-duas-camadas.md), e [Diagramas Mermaid](https://github.com/jtonynet/go-products-api/tree/main#diagrams)
- Sabemos o que fazer, graças às definições do arquivo __README.md__. Sabemos como fazer graças aos __ADRs__ e documentações vinculadas. Devemos nos organizar em estrutura __Kanban__, guiados pelo modelo Agile, em nosso __Github Project__, e dar o devido prosseguimento às tarefas.


[0.0.6]: https://github.com/jtonynet/go-products-api/compare/v0.0.5...v0.0.6
[0.0.5]: https://github.com/jtonynet/go-products-api/compare/v0.0.4...v0.0.5
[0.0.4]: https://github.com/jtonynet/go-products-api/compare/v0.0.3...v0.0.4
[0.0.3]: https://github.com/jtonynet/go-products-api/compare/v0.0.2...v0.0.3
[0.0.2]: https://github.com/jtonynet/go-products-api/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/jtonynet/go-products-api/compare/v0.0.0...v0.0.1
[0.0.0]: https://github.com/jtonynet/go-products-api/releases/tag/v0.0.0
