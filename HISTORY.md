<div id="top"></div>

<div align="center" style="margin-bottom: 50px;">
  <img src="docs/assets/bbb-icon.png" alt="Flow Icon" width="150" height="auto"/>

  <h1>BBB Voting System </h1>

   <h4>
    <a href="https://github.com/Jav4Script/pulls">Request Feature</a>
    <span> . </span>
    <a href="https://github.com/Jav4Script/issues">Report Issue</a>
  </h4>
</div>

# Índice

- [Índice](#índice)
  - [Decisões de Arquitetura](#decisões-de-arquitetura)
    - [Arquitetura e Implementações](#arquitetura-e-implementações)
    - [Tecnologias Utilizadas](#tecnologias-utilizadas)
  - [Introdução ao Golang](#introdução-ao-golang)
    - [Características](#características)
    - [Concorrência](#concorrência)
    - [Gerenciamento de Memória](#gerenciamento-de-memória)
    - [Simplicidade](#simplicidade)
    - [Recursos Adicionais](#recursos-adicionais)
  - [Possíveis Evoluções do Projeto](#possíveis-evoluções-do-projeto)
    - [Monitoramento com Prometheus e Grafana](#monitoramento-com-prometheus-e-grafana)
    - [Implementação de Circuit Breaker](#implementação-de-circuit-breaker)
    - [Configuração de Timeouts](#configuração-de-timeouts)
    - [Redundância para Banco de Dados e Serviços](#redundância-para-banco-de-dados-e-serviços)
    - [Uso de Proxy Reverso](#uso-de-proxy-reverso)
    - [Auto Scale](#auto-scale)
    - [Comunicação com gRPC](#comunicação-com-grpc)
      - [Vantagens do gRPC](#vantagens-do-grpc)
      - [Recursos Adicionais](#recursos-adicionais-1)
    - [Implementação de MultiCaptcha](#implementação-de-multicaptcha)
      - [Benefícios do MultiCaptcha](#benefícios-do-multicaptcha)
    - [Recursos Adicionais](#recursos-adicionais-2)
    - [Segurança e Análise de Fraudes](#segurança-e-análise-de-fraudes)
      - [Análise de Fraudes](#análise-de-fraudes)
      - [Prevenção de Ataques](#prevenção-de-ataques)
      - [Criptografia](#criptografia)
      - [Auditoria e Logging](#auditoria-e-logging)
    - [Recursos Adicionais](#recursos-adicionais-3)

## Decisões de Arquitetura

O projeto de votação de participantes do BBB foi desenvolvido com foco em alta resiliência e performance, utilizando uma arquitetura robusta e tecnologias modernas para garantir a escalabilidade e a eficiência do sistema. Abaixo estão as principais decisões de arquitetura, implementações realizadas e tecnologias utilizadas:

<div align="right"><a style="font-weight: 500;" href="#top">Back to Top</a></div>

![-](/docs/assets/rainbow-divider.png)

### Arquitetura e Implementações

- **API REST**: Desenvolvida em Golang, a API gerencia as requisições de votação, aproveitando a alta concorrência e eficiência da linguagem.
- **Redis**: Utilizado para armazenamento temporário de resultados parciais, permitindo acesso rápido e eficiente aos dados.
- **RabbitMQ**: Implementado para buffering e processamento assíncrono das requisições de votação, evitando sobrecarga e possibilitando lidar com grandes picos de requisições.
- **PostgreSQL**: Banco de dados relacional utilizado para persistência de dados históricos, garantindo integridade e consistência.

### Tecnologias Utilizadas

- **Golang**: Escolhido pela sua simplicidade, eficiência e suporte nativo à concorrência, essencial para lidar com um grande volume de requisições simultâneas.
- **Redis**: Armazenamento em memória para resultados parciais, proporcionando acesso rápido e eficiente aos dados.
- **RabbitMQ**: Sistema de filas para gerenciamento de picos de tráfego e processamento assíncrono de mensagens.
- **PostgreSQL**: Banco de dados relacional para persistência de dados históricos.
- **Swag**: Ferramenta para geração de documentação Swagger, facilitando a integração e o uso da API.
- **Wire**: Utilizado para injeção de dependências, promovendo um código mais modular e testável.

<div align="right"><a style="font-weight: 500;" href="#top">Back to Top</a></div>

![-](/docs/assets/rainbow-divider.png)

## Introdução ao Golang

Golang, também conhecido como Go, é uma linguagem de programação desenvolvida pela Google. Ela é conhecida por sua simplicidade e eficiência, especialmente em sistemas concorrentes.

### Características
- **Concorrência**: Go possui um modelo de concorrência robusto baseado em goroutines e channels.
- **Gerenciamento de Memória**: Go tem um garbage collector eficiente que facilita o gerenciamento de memória.
- **Simplicidade**: A linguagem é simples e fácil de aprender, com uma sintaxe clara e concisa.
- **Compilada**: Go é uma linguagem compilada, o que resulta em uma execução mais rápida e eficiente.

### Concorrência
Go é especialmente conhecida por seu modelo de concorrência, que permite a execução de múltiplas tarefas de forma eficiente e segura. As goroutines são leves e podem ser gerenciadas facilmente pelo runtime da linguagem. Além disso, o uso de channels facilita a comunicação segura entre goroutines, evitando problemas comuns de concorrência como race conditions.

### Gerenciamento de Memória
O garbage collector do Go é projetado para ser eficiente e minimizar pausas, o que é crucial para aplicações de alta performance. Ele realiza a coleta de lixo de forma incremental, permitindo que o programa continue executando com interrupções mínimas.

### Simplicidade
A simplicidade do Go é uma de suas maiores vantagens. A linguagem foi projetada para ser fácil de aprender e usar, com uma sintaxe que é direta e sem ambiguidades. Isso permite que os desenvolvedores escrevam código limpo e mantenível, reduzindo a complexidade e o tempo de desenvolvimento.

---

### Recursos Adicionais
- [Cheatsheet do Go](https://devhints.io/go)
- [Tour do Go](https://tour.golang.org/)

<div align="right"><a style="font-weight: 500;" href="#top">Back to Top</a></div>

![-](/docs/assets/rainbow-divider.png)

## Possíveis Evoluções do Projeto

### Monitoramento com Prometheus e Grafana
Para melhorar o monitoramento do sistema, podemos integrar o Prometheus para coleta de métricas e o Grafana para visualização dessas métricas. Isso permitirá uma melhor análise de desempenho e identificação de gargalos.

Para mais informações, consulte a [documentação oficial do Prometheus](https://prometheus.io/docs/introduction/overview/) e a [documentação oficial do Grafana](https://grafana.com/docs/grafana/latest/).

### Implementação de Circuit Breaker
A implementação de um Circuit Breaker ajudará a aumentar a resiliência do sistema, prevenindo falhas em cascata e melhorando a tolerância a falhas. Ferramentas como Hystrix podem ser utilizadas para essa finalidade.

Para mais informações, consulte a [documentação oficial do Hystrix](https://github.com/Netflix/Hystrix/wiki).

### Configuração de Timeouts
Definir timeouts apropriados para chamadas de rede e operações de I/O é crucial para evitar que o sistema fique bloqueado indefinidamente. Isso também ajuda a melhorar a responsividade do sistema.

Para mais informações, consulte a [documentação oficial do Go sobre contextos](https://golang.org/pkg/context/).

### Redundância para Banco de Dados e Serviços
Adicionar redundância para bancos de dados e serviços críticos pode aumentar a disponibilidade e a confiabilidade do sistema. Isso pode ser feito através de replicação de dados e balanceamento de carga.

Para mais informações, consulte a [documentação oficial do PostgreSQL sobre replicação](https://www.postgresql.org/docs/current/high-availability.html).

### Uso de Proxy Reverso
A utilização de um proxy reverso, como Nginx ou HAProxy, pode ajudar a distribuir a carga de forma eficiente, além de fornecer funcionalidades adicionais como cache e SSL termination.

Para mais informações, consulte a [documentação oficial do Nginx](https://nginx.org/en/docs/) e a [documentação oficial do HAProxy](https://www.haproxy.org/documentation/).

### Auto Scale
Implementar auto scale permitirá que o sistema ajuste automaticamente a quantidade de recursos computacionais com base na demanda. Isso pode ser configurado em plataformas de cloud como AWS, GCP ou Azure.

Para mais informações, consulte a [documentação oficial da AWS sobre Auto Scaling](https://docs.aws.amazon.com/autoscaling/), a [documentação oficial do GCP sobre Auto Scaling](https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-autoscaler), e a [documentação oficial do Azure sobre Auto Scaling](https://docs.microsoft.com/en-us/azure/azure-monitor/autoscale/autoscale-get-started).

### Comunicação com gRPC
Para melhorar a performance e eficiência na comunicação entre micro serviços, podemos adotar o gRPC. Ele utiliza HTTP/2 para transporte, Protobuf para serialização de mensagens e oferece suporte a comunicação bidirecional.

#### Vantagens do gRPC
- **Performance**: gRPC é mais rápido e eficiente em comparação com REST, devido ao uso de HTTP/2 e Protobuf.
- **Contratos Fortes**: Utiliza arquivos `.proto` para definir contratos de serviço, garantindo que todos os serviços sigam a mesma especificação.
- **Streaming**: Suporta streaming de dados bidirecional, permitindo comunicação mais eficiente em tempo real.

Para mais informações, consulte a [documentação oficial do gRPC](https://grpc.io/docs/).

#### Recursos Adicionais
- [Cheatsheet do gRPC](https://devhints.io/grpc)
- [Guia de Introdução ao gRPC](https://grpc.io/docs/what-is-grpc/introduction/)

<div align="right"><a style="font-weight: 500;" href="#top">Back to Top</a></div>

![-](/docs/assets/rainbow-divider.png)

### Implementação de MultiCaptcha
A implementação de MultiCaptcha pode aumentar a segurança do sistema, prevenindo ataques automatizados e garantindo que apenas usuários legítimos possam acessar determinados recursos. MultiCaptcha combina diferentes tipos de desafios, como reCAPTCHA, hCaptcha, e outros, para fornecer uma camada adicional de proteção.

#### Benefícios do MultiCaptcha
- **Segurança Aumentada**: Ao combinar múltiplos métodos de verificação, torna-se mais difícil para bots automatizados burlarem o sistema.
- **Flexibilidade**: Permite a escolha do tipo de captcha mais adequado para cada situação, melhorando a experiência do usuário.
- **Resiliência**: Se um tipo de captcha for comprometido, os outros ainda fornecem uma camada de segurança adicional.

Para mais informações, consulte a [documentação oficial do reCAPTCHA](https://developers.google.com/recaptcha) e a [documentação oficial do hCaptcha](https://docs.hcaptcha.com/).

---

### Recursos Adicionais
- [Guia de Implementação do reCAPTCHA](https://developers.google.com/recaptcha/docs/display)
- [Guia de Implementação do hCaptcha](https://docs.hcaptcha.com/gettingstarted)

<div align="right"><a style="font-weight: 500;" href="#top">Back to Top</a></div>

![-](/docs/assets/rainbow-divider.png)

### Segurança e Análise de Fraudes

Para garantir a segurança do sistema e proteger contra fraudes e ataques, podemos implementar diversas estratégias e ferramentas. Abaixo estão algumas sugestões de melhorias que podem ser adotadas:

#### Análise de Fraudes
Implementar um sistema de análise de fraudes pode ajudar a identificar e prevenir atividades suspeitas. Isso pode ser feito através de:
- **Machine Learning**: Utilizar algoritmos de machine learning para detectar padrões anômalos e comportamentos suspeitos.
- **Regras de Negócio**: Definir regras específicas para identificar transações ou atividades fora do comum.
- **Monitoramento Contínuo**: Implementar monitoramento contínuo para detectar fraudes em tempo real.

Para mais informações, consulte a [documentação oficial do TensorFlow](https://www.tensorflow.org/).


#### Prevenção de Ataques
Para proteger o sistema contra ataques, podemos adotar as seguintes medidas:
- **WAF (Web Application Firewall)**: Utilizar um WAF para filtrar e monitorar o tráfego HTTP, protegendo contra ataques comuns como SQL Injection e Cross-Site Scripting (XSS).
- **Rate Limiting**: Implementar rate limiting para limitar o número de requisições que um usuário pode fazer em um determinado período de tempo, prevenindo ataques de força bruta.
- **DDoS Protection**: Utilizar serviços de proteção contra DDoS para mitigar ataques de negação de serviço distribuída.

Para mais informações, consulte a [documentação oficial do Cloudflare](https://developers.cloudflare.com/waf/) e a [documentação oficial do AWS Shield](https://docs.aws.amazon.com/waf/latest/developerguide/ddos-overview.html).

#### Criptografia
Garantir que todos os dados sensíveis sejam criptografados tanto em trânsito quanto em repouso. Isso pode ser feito através de:
- **TLS (Transport Layer Security)**: Utilizar TLS para criptografar dados em trânsito.
- **Criptografia de Dados**: Utilizar algoritmos de criptografia como AES para proteger dados armazenados.

Para mais informações, consulte a [documentação oficial do Let's Encrypt](https://letsencrypt.org/docs/) e a [documentação oficial do AWS KMS](https://docs.aws.amazon.com/kms/latest/developerguide/overview.html).

#### Auditoria e Logging
Implementar auditoria e logging detalhados para rastrear todas as atividades no sistema. Isso pode ajudar a identificar e investigar incidentes de segurança. As práticas recomendadas incluem:
- **Logs Centralizados**: Utilizar uma solução de logging centralizada para coletar e analisar logs de diferentes componentes do sistema.
- **Alertas em Tempo Real**: Configurar alertas para atividades suspeitas ou anômalas.

Para mais informações, consulte a [documentação oficial do ELK Stack](https://www.elastic.co/what-is/elk-stack) e a [documentação oficial do Splunk](https://docs.splunk.com/Documentation/Splunk).

---

### Recursos Adicionais
- [Guia de Segurança do OWASP](https://owasp.org/www-project-top-ten/)
- [Documentação do TensorFlow para Análise de Fraudes](https://www.tensorflow.org/tutorials/structured_data/imbalanced_data)
- [Guia de Implementação do Let's Encrypt](https://letsencrypt.org/getting-started/)

<div align="right"><a style="font-weight: 500;" href="#top">Back to Top</a></div>

![-](/docs/assets/rainbow-divider.png)
