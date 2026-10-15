# 001 - **Usar envconfig para gerenciamento de configuração**

## Contexto
O gateway vai precisar ler variaveis de ambiente(porta, ambiente,chaves e etc), de forma confiável. no Go tem várias opções: `os.Getenv` muito manual...`godotenv` só lê, não valida tipos...Enfim...Preciso de algo que valide tipos, tenha defaults e funcione igual em dev e prod.

## Decisão
Usar `github.com/kelseyhightower/envconfig`

### Alternativas consideradas

**`github.com/joho/godotenv`**
- RECHAZADO: apenas carrega arquivo `.env` para `os.Environ`, não
  valida tipos e em produção (Docker/K8s) o `.env` nem é lido.

**`os.Getenv` FULL manual **
- RECHAZADO: verboso, sem validação de tipos, fácil de esquecer
  defaults, difícil de auditar quais variáveis a aplicação usa.

**`github.com/spf13/viper`**
- RECHAZADO (por agora): muito poderoso para o caso atual, traz
  complexidade (ler de arquivo, env, flags, remote config) que
  não precisamos ainda. Reconsiderar se surgir necessidade de
  múltiplas fontes de configuração e etc...

## Consequências
- Configuração centralizada em struct
- Validação automática de tipos
- Defaults via tags
- Funciona igual em dev e stagin e prod
- Mais uma dependência... (mas é madura e leve então esta piola)