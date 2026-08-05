# Objetivo

Criar uma mesa de Poker Texas Hold'em.

Ao iniciar o servidor, a mesa é criada e todo estado da partida é persistido em memória até que o servidor seja encerrado. Ao iniciar novamente, nenhum estado anterior é salvo, mas uma partida totalmente nova é criada.

Sempre que um jogador entra na mesa, ele começa com um número fixo de dinheiro. Sempre que ele sai, todo dinheiro ganho/perdido é limpado da memória. Não existe criação de conta para salvar progresso.

Todos jogadores devem estar no mesmo wi-fi local do servidor.

# TO-DO

## Frontend

- [ ] trocar Material por Spartan

## Backend

- [ ] criação de conta e persistência de dados
- [ ] permitir jogadores de fora da LAN
