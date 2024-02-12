# 3. Observabilidade com Prometheus e Grafana

Data: 10 de fevereiro de 2024

## Status

Aceito

## Contexto

Dado que o projeto está sendo desenvolvido em duas camadas, aderente ao desafio, e não temos dados sobre a escala que ele deve alcançar, ter algum nível de __Observabilidade__ desde o início é muito útil. As métricas fornecidas por essa abordagem nos orientam sobre como nossos projetos crescem, tornando-os mais escaláveis e sustentáveis.

Essas métricas nos ajudam a escolher as melhores abordagens e arquiteturas. Decisões como o uso de filas, caches defensivos e outras alternativas devem ser tomadas com base em __métricas__.

## Decisão

Como o [Prometheus](https://prometheus.io/) e o [Grafana](https://grafana.com/) são amplamente utilizados no mercado e já possuo experiência e conhecimento prévio com eles, além do fato de o Grafana já fazer parte do conjunto de ferramentas do proponente do desafio, faz sentido utilizá-los neste contexto.

Farei a implementação de um middleware usando o [Prometheus](https://github.com/prometheus/client_golang) e o [Gorm-Prometheus](https://github.com/go-gorm/prometheus) customizado no framework Echo, juntamente com suas devidas configurações em variáveis de ambiente. Também preparamos nosso `docker-compose.yml` para atender à estrutura dessas ferramentas, permitindo um desenvolvimento na máquina local que esteja alinhado com essa diretriz.

Durante minhas pesquisas, também me deparei com o [Echo-Prometheus](https://github.com/globocom/echo-prometheus) da Globo.com, mas optei pela abordagem customizada devido à baixa popularidade (estrelas) que o projeto possui e minha familiaridade com essa abordagem no `Gin Framework`.

## Consequências

Obteremos um conjunto básico de métricas `RED (Rate, Errors e Duration)` e `API Basics`, amplamente utilizadas no mercado, que são indicadores do crescimento e escalabilidade de nossas APIs. Balizadores do desenvolvimento.