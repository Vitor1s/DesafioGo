# DesafioGo

Projeto em Go que implementa dois desafios de programação relacionados a divisibilidade de números.

## Descrição

Este projeto contém duas funcionalidades principais que analisam números de 1 a 100 e identificam padrões de divisibilidade.

## Funcionalidades

### Desafio 1 - Números Divisíveis por 3
- Itera de 1 a 100
- Identifica e exibe apenas os números divisíveis por 3
- Saída: 3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54, 57, 60, 63, 66, 69, 72, 75, 78, 81, 84, 87, 90, 93, 96, 99

### Desafio 2 - Sistema Pin/Pan
- Itera de 1 a 100
- Para números divisíveis por 3: exibe "Pin = [número]"
- Para números divisíveis por 5: exibe "Pan = [número]"
- Para números divisíveis por ambos (3 e 5): exibe "PinPan = [número]"

## Como Executar

### Pré-requisitos
- Go instalado no sistema

### Execução Direta
```bash
go run cmd/main.go
```

### Build e Execução
```bash
go build -o cmd/main cmd/main.go
./cmd/main
```

### Mas como o Build ja existe só rodar naa pasta principal fora do cmd
```bash
./main 
```

## Estrutura do Projeto

```
DesafioGo/
├── cmd/
│   ├── main.go      # Código principal
│   └── main         # Executável compilado
└── README.md        # Este arquivo
```

## Exemplo de Saída

```
3
6
9
12
15
18
21
24
27
30
33
36
39
42
45
48
51
54
57
60
63
66
69
72
75
78
81
84
87
90
93
96
99
--------------------------------
Pin = 3
Pan = 5
Pin = 6
Pin = 9
Pan = 10
Pin = 12
Pin = 15
Pan = 15
PinPan = 15
...
```

## Tecnologias Utilizadas

- **Go** - Linguagem de programação
- **fmt** - Pacote para entrada/saída formatada

## Autor

Desenvolvido como parte de desafios de programação em Go.
