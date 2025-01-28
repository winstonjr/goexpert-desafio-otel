# Concorrência com Golang - Leilão

### Para rodar o projeto realizar os seguintes passos
1. Baixar repositório
2. Criar API Key para consultar temperatura no site: `https://www.weatherapi.com/`
3. Criar dois arquivos `.env` nas pastas `cmd/servico-a/` e `cmd/servico-b/`
4. Editar o arquivo .env na pasta `cmd/servico-a/` e colocar o seguinte conteúdo
```shell
INTERNAL_API_URI=http://api2:8081/
OTEL_COLLECTOR=otel-collector:4317
```
5. Editar o arquivo .env na pasta `cmd/servico-b/` e colocar o seguinte conteúdo
```shell
WEATHER_API_KEY=
OTEL_COLLECTOR=otel-collector:4317
```
6. Preencher a chave `WEATHER_API_KEY=` com o valor da chave criada no passo 2
7. Rodar `docker compose up` na pasta com o arquivo `docker-compose.yaml`

PS: Alternativamente também é possível copiar o arquivo `.env_exemplo` para as pastas `cmd/servico-a/` e `cmd/servico-b/`. Não esquecendo de preencher a variável `WEATHER_API_KEY=` com o valor da chave criada no passo 2. 

### Validando o funcionamento do Zipkin
1. Rodar os testes no arquivo `api/check_weather.http`
2. Entrar no Zipkin através do link `http://localhost:9411/zipkin/`
3. Apertar o botão `RUN QUERY`
4. Expandir os resultados
5. Clicar no botão `SHOW`