# Command

O Command é um padrão de projeto comportamental que converte solicitações ou operações simples em objetos.

A conversão permite a execução adiada ou remota de comandos, armazenamento do histórico de comandos, etc.

O Command é um padrão de projeto comportamental que transforma um pedido em um objeto independente que contém toda a informação sobre o pedido. Essa transformação permite que você parametrize métodos com diferentes pedidos, atrase ou coloque a execução do pedido em uma fila, e suporte operações que não podem ser feitas.