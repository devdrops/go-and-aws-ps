# Go e AWS Parameter Store

Código feito para entender como usar o AWS Parameter Store com Go para leitura
de valores do SSM.

## Execução

### Pré Requisitos

Antes de executar, certifique-se de que tem os requisitos:

- Docker
- Docker Compose
- AWS AMI com permissões para Systems Manager
- Region, Access Key ID e Secret Access Key dessa AMI
  - Esses valores deverão estar escritos no arquivo `.env`, criado como uma
  cópia do arquivo `.env.dist` deste repositório, na raiz do repositório.

Uma vez verificado, clone este repositório e execute o seguinte comando no seu
terminal, a partir da pasta raiz deste repositório:

```shell
docker-compose run --rm shell
```

Uma vez executado o comando acima, basta executar o arquivo `main.go` com o
comando abaixo:

```shell
go run main.go
```

As dependências serão instaladas e a requisição será realizada. Se tudo for
feito com sucesso, você verá algo parecido na saída:

```shell
go: downloading github.com/aws/aws-sdk-go v1.39.1
go: downloading github.com/jmespath/go-jmespath v0.4.0
{
  Parameter: {
    ARN: "arn:aws:ssm:REDACTED:parameterREDACTED",
    DataType: "text",
    LastModifiedDate: REDACTED,
    Name: "REDACTED",
    Type: "REDACTED",
    Value: "REDACTED",
    Version: 1
  }
}
```

## Referências

- https://docs.aws.amazon.com/sdk-for-go/api/
- https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html
