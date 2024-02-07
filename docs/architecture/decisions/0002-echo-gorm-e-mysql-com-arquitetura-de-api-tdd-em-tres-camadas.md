# 2. Echo, Gorm e MySQL em API TDD com arquitetura de Três Camadas

Data: 07 de fevereiro de 2024

## Status

Aceito

## Contexto

Precisamos definir o melhor fluxo de trabalho e testes para a `go-api-products`.

Não existem muitas opções para a arquitetura base, pois elementos cruciais para o projeto já foram definidos, conforme descrito na seção [Sobre do arquivo README.md](../../../README.md). Segue o principal trecho para nosso contexto:

>
> - Stack Back-end
>   - Postman
>   - GoLang
>   - Echo Framework
>   - MySQL
>

Temos diversas abordagens para a construção dessa API, incluindo `Arquitetura Hexagonal`, `DDD` e `Três Camadas`, além da abordagem de testes na já conhecida Pirâmide de Testes, que pode auxiliar no fluxo de `CI`, componente crucial para a estabilidade em um [Microsserviço Pronto para Produção](https://www.amazon.com.br/Microsservi%C3%A7os-Prontos-Para-Produ%C3%A7%C3%A3o-Padronizados/dp/8575226215).

<img src="../../assets/images/layout/graphics/test_pyramid.jpg">

_[Imagem retirada do artigo: The Testing Pyramid: Simplified for One and All](https://www.headspin.io/blog/the-testing-pyramid-simplified-for-one-and-all)_

## Decisão

Este documento determina o fluxo de trabalho __Branch Based com Feature Branch__, design estrutural e a abordagem de testes para garantir um padrão e a estabilidade da aplicação.

Considerando ser um __CRUD simples__, entre as várias abordagens estruturais, a __Arquitetura em Três Camadas__ faz mais sentido, por se adequar bem à simplicidade do desafio proposto.

__GORM__ foi escolhido como ORM por facilitar integração aos principais BDs e facilidades como observabilidade que poderemos adotar no futuro.

Como não há regras de negócio vinculadas às operações dos endpoints propostos, não parece fazer muito sentido utilizar testes unitários (a base da pirâmide), pois não desejamos testar as operações do nosso banco de dados/ORM, mas sim a integridade dos dados nessas operações. Portanto, iremos utilizar __Testes de Integração__ e iniciar a construção de nosso sistema através deles, no espírito do __TDD__. Para que exista forte integração ao GithubActions que adotaremos como processo de __CI__.

Também faz sentido adotar ferramentas de documentação e design de APIs, como o __Swagger__, utilizando a implementação __OpenAPI__, devido ao seu amplo histórico de utilização.

Decidi também utilizar o __Github Projects__ como ferramenta para auxiliar na abordagem __Kanban__ para conclusão das tarefas, visando alcançar um cenário próximo ao de um time de desenvolvimento real.

## Consequências

Uma API bem documentada (__Swagger, Readme, ADRs e Diagramas__) que atende às propostas da área de negócios e às __Guidelines and Guardrails__ (definições prévias da área de tecnologia), conforme descrito no [arquivo README.md](../../../README.md), com capacidade de desenvolvimento e expansão de forma __Iterativa e Incremental__.