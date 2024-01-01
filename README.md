# Golang MS

Microsserviço em Golang com o objetivo de realizar o envio de pedidos na fila e consumir essas mensagens retornando o valor do pedido com a taxa de 10% aplicada e registrar esses valores no banco.

##Primeiros passos para rodar o projeto:

- Ligar o rabbitmq e grafana:
`docker-compose up -d`

Inserir no ***From exchange*** o valor ***amq.direct* **em bindings na página local do seu rabbitmq para que seja possível ver as mensagens na fila:

local page rabbitmq > bindings > from exgange = amq.direct

Acessar o local do rabbitmq: http://localhost:15672/
Acessar o local do grafana: http://localhost:3000/

Iniciando o producer e o consumer