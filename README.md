# Music-weather

API criada para sugerir uma playlist no spotify com base na temperatura atual da cidade escolhida.

## Requisitos

* Golang: [Install Guide](https://golang.org/doc/install)
* Gin: [Documentation](https://pkg.go.dev/github.com/gin-gonic/gin)
* OpenWeatherMap: [Documentation](https://openweathermap.org/guide)
* Spotify: [Documentation](https://developer.spotify.com/documentation/web-api)

## Justificativa das Escolhas Tecnológicas

### Padrão de API
Optei pelo padrão RESTful para o serviço por sua simplicidade e ampla adoção na indústria. REST é um padrão leve, fácil de implementar e compatível com diversas tecnologias, permitindo que o serviço seja consumido por diferentes clientes, como web, mobile, ou outros serviços. Além disso, RESTful APIs são escaláveis e podem ser facilmente versionadas, garantindo a evolução do serviço sem impactar os consumidores existentes.

### Linguagem de Programação
Escolhi a linguagem Go para o desenvolvimento do serviço devido à sua eficiência e simplicidade. Go é uma linguagem compilada, o que resulta em um desempenho superior, especialmente em aplicações que precisam lidar com um alto volume de requisições, como APIs. Além disso, Go tem um suporte nativo a concorrência através de goroutines, permitindo a construção de serviços altamente performáticos e escaláveis. A simplicidade da linguagem também facilita a manutenção e leitura do código.

### Frameworks e Ferramentas
Utilizei o framework Gin para o desenvolvimento da API devido à sua simplicidade e performance. Gin é um microframework que oferece uma sintaxe simples e intuitiva, facilitando o desenvolvimento rápido de APIs. Ele também é conhecido por ser extremamente performático, o que é crucial para aplicações que exigem baixa latência e alta taxa de requisições. Além disso, Gin oferece suporte nativo para middlewares e manipulação de rotas, o que torna o desenvolvimento de APIs RESTful mais eficiente.

### Serviços de Terceiros
Para integrar dados externos ao serviço, decidi utilizar as APIs do Spotify e OpenWeatherMap. A API do OpenWeatherMap foi escolhida para obter dados precisos e em tempo real sobre a temperatura das cidades, o que é essencial para o funcionamento do serviço. Já a API do Spotify foi selecionada por sua ampla base de dados de músicas e playlists, permitindo sugerir playlists baseadas em diferentes gêneros musicais conforme a temperatura. Ambas as APIs são amplamente usadas e possuem boas documentações, o que facilita a integração.

## Deployment
GitHub Actions para Integração Contínua e Deploy Automático
Para garantir a integração contínua e o deploy automático do serviço, configurei um workflow do GitHub Actions que é acionado sempre que há um push ou um pull request na branch main. O objetivo desse workflow é automatizar o processo de build, execução de testes, e deploy da aplicação na plataforma Render.

Aqui está o workflow configurado:

```bash
name: Go

on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]

jobs:

  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21.3'

    - name: Build
      run: go build -v ./...

    - name: Test
      run: go test -v ./...
```

## Justificativa das Escolhas de Infraestrutura

1. GitHub Actions: Escolhi o GitHub Actions como ferramenta de CI/CD (Integração Contínua/Entrega Contínua) devido à sua integração nativa com o GitHub, facilidade de uso e custo zero para repositórios públicos. O workflow definido automatiza o processo de build e testes, garantindo que o código na branch main esteja sempre funcionando corretamente antes de ser implantado.

2. Render: A escolha do Render como plataforma de hospedagem foi motivada por vários fatores:

Custo: Render oferece uma camada gratuita generosa, o que é ideal para projetos em estágio inicial ou de pequeno porte, como este.

Facilidade de Uso: O Render proporciona um processo de deploy simplificado, onde a aplicação é automaticamente implantada a partir do repositório GitHub sempre que há um push na branch main. Isso elimina a necessidade de configurar um pipeline de deploy complexo.

Adequação ao Escopo do Projeto: Para uma aplicação como a music-weather, que não exige uma infraestrutura altamente complexa, o Render se encaixa perfeitamente. Ele oferece suporte nativo para Go e integrações fáceis com GitHub, facilitando o processo de deploy e gerenciamento da aplicação.

Essa configuração garante que o processo de desenvolvimento seja ágil e automatizado, desde a codificação até o deploy em produção, permitindo que novas funcionalidades e correções de bugs sejam disponibilizadas de maneira contínua e confiável.

## Instalação

Para configurar e rodar o projeto localmente, siga os passos abaixo:

1. Certifique-se de ter o Go instalado em sua máquina. Você pode seguir o Guia de Instalação do Go para configurar o ambiente.

2. Clonando o Repositório:

```bash
git clone https://github.com/Moreira-Henrique-Pedro/music-weather.git
cd music-weather
```

3. Instalando dependências:

```bash
go mod tidy
```

4. Rodando o projeto:

```bash
go run main.go
```

## Variáveis de Ambiente

As seguintes variáveis de ambiente são necessárias para a aplicação funcionar corretamente:

OPENWEATHER_API_KEY - Chave de API para acessar o OpenWeatherMap.
SPOTIFY_CLIENT_ID - ID do cliente para autenticação com a API do Spotify.
SPOTIFY_CLIENT_SECRET - Secret do cliente para autenticação com a API do Spotify.

Você pode criar um arquivo .env na raiz do projeto com as variáveis mencionadas e seus respectivos valores:

OPENWEATHER_API_KEY=<sua-_openweather-chave-aqui>
SPOTIFY_CLIENT_ID=<seu-spotify-id-aqui>
SPOTIFY_CLIENT_SECRET=<sua-spotify-secret-aqui>

## Uso

Para utilizar a aplicação, você pode fazer uma requisição POST para o endpoint /music_weather com o seguinte payload:

```bash
    {
    "City": "Santo André"
    }
```

URL de Produção

```bash
https://music-weather.onrender.com/music_weather
```

URL Local

```bash
http://localhost:8000/music_weather
```

A aplicação irá retornar uma playlist sugerida baseada na temperatura atual da cidade informada.

## Documentação Swagger

A documentação da API é gerada automaticamente pelo Swagger e está disponível em uma interface interativa que permite explorar e testar os endpoints da API.
Para acessar a documentação Swagger:

URL de Produção
Documentação Swagger: https://music-weather.onrender.com/swagger/index.html

URL Local
Documentação Swagger: http://localhost:8000/swagger/index.html

### Acessando a Documentação

Abrir a URL: Acesse a URL fornecida acima no seu navegador para visualizar a documentação.

Explorar Endpoints: Use a interface do Swagger para explorar os endpoints disponíveis, ver exemplos de requisições e respostas, e testar as APIs diretamente no navegador.
