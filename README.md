# Desafio 02

Multithreading

Serão realizadas duas requisições simultaneamente para as seguintes APIs:

https://cdn.apicep.com/file/apicep/" + cep + ".json

http://viacep.com.br/ws/" + cep + "/json/

Onde: 

- Será acatada a API que entregar a resposta mais rápida e descartada a resposta mais lenta.

- O resultado da request será exibido no command line, bem como qual API a enviou.

- Existe um limite do tempo de resposta em 1 segundo. Caso esse limite seja atingido, o erro de timeout será exibido.

### Como usar

Clonar o repositório e entrar na pasta do projeto

```bash
#clonar o repositório
git clone github.com/virb30/go-challenge-02 .
# entrar no diretório
cd go-challenge-02
```

Executar o código da aplicação informando o CEP desejado

```bash
# executar a aplicação
go run main.go [cep]
```