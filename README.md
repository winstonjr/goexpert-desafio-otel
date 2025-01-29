# Concorrência com Golang - Leilão

### Para rodar o projeto realizar os seguintes passos
1. Baixar repositório
2. Criar API Key para consultar temperatura no site: `https://www.weatherapi.com/`
3. Acessar a pasta `cmd/servico-a/`
4. Renomear o arquivo `.env_example` para `.env`
5. Acessar a pasta `cmd/servico-b/`
6. Renomear o arquivo `.env_example` para `.env`
7. Preencher a chave `WEATHER_API_KEY=` com o valor da chave criada no passo 2
8. Rodar `docker compose up` na pasta com o arquivo `docker-compose.yaml`

PS: Alternativamente também é possível copiar o arquivo `.env_exemplo` para as pastas `cmd/servico-a/` e `cmd/servico-b/`. Não esquecendo de preencher a variável `WEATHER_API_KEY=` com o valor da chave criada no passo 2. 

### Validando o funcionamento do Zipkin
1. Rodar os testes no arquivo `api/check_weather.http`
2. Entrar no Zipkin através do link `http://localhost:9411/zipkin/`
3. Apertar o botão `RUN QUERY`
4. Expandir os resultados
5. Clicar no botão `SHOW`