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

Dada sua simplicidade, uma vez que se trata de um __CRUD simples__. Faz sentido utilizar __Arquitetura de Duas Camadas__ o que aumenta o ritmo do desenvolvimento com menos partes para se preocupar porém mantendo a qualidade do resultado final.

__GORM__ foi escolhido como ORM por facilitar a integração aos principais BDs e oferecer facilidades como observabilidade, que poderemos adotar no futuro.

Utilizaremos __Smoke Test__ (mais próximo do Integration do que do Unit) e iniciamos a construção de nosso sistema através dele, no espírito do __TDD__. Para que haja forte integração ao GithubActions, que adotaremos como processo de __CI__. [Consulte este artigo para mais informações](https://novateus.com/blog/8-functional-testing-types-explained-with-examples/)

Também faz sentido adotar ferramentas de documentação e design de APIs, como o __Swagger__, utilizando a implementação __OpenAPI__, devido ao seu amplo histórico de utilização.

Decidi também utilizar o [Github Projects](https://github.com/users/jtonynet/projects/5/views/1) como ferramenta para auxiliar na abordagem __Kanban__ para conclusão das tarefas, visando alcançar um cenário próximo ao de um time de desenvolvimento real.

## Consequências

Uma API bem documentada (__Swagger, Readme, ADRs e Diagramas__) que atende às propostas da área de negócios e às __Guidelines and Guardrails__ (definições prévias da área de tecnologia), conforme descrito no [arquivo README.md](../../../README.md), com capacidade de desenvolvimento e expansão de forma __Iterativa e Incremental__.