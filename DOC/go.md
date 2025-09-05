# Documentation go

## Syntaxe
Le go nécésite une syntaxe qui nécésite un ordre précis dans le code pour fonctionner :

1. déclaration de package
2. import de packages (optionel)
3. fonction(s)
4. exprésion(s)

```go
package main //déclaration de package

import "fmt" //import de packages

func main() { //fonction (dans ce cas une seul)

    //exprésion (en gros le code)
    name := "John Doe"
    fmt.Println(name)
}
```

> [!IMPORTANT]
> Retenez la syntaxe de tout ca pour les éxams