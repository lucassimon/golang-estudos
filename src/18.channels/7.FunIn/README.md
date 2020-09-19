pid 1 -> channel
pid 2 -> channel
pid 3 -> channel

Uma função que recebe os canais anteriores e retorna um canal unico; Como um funil ou afunilamento

Para termos um unico ponto de consulta

```
func pid1, pid2, pid3 {
     return x <- pid1, pid2, pid3
}
```
