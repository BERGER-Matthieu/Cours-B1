# DOCUMENTATION GO

## Syntaxe
Le go nécésite une syntaxe qui nécésite un ordre précis dans le code pour fonctionner :

1. **déclaration de package**
2. **import de packages** (optionel)
3. **fonction(s)**
4. **exprésion(s)**

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

---
### déclaration de package
Un package est comme une boite a outils vous permetant de facilement importer et exporter vos fonctions entre diférents projets. 

Vous pouvez préciser dans quelle boite grace a la ligne ...
```go
package main //dans ce cas tout le code ira dans la boite a outil "main"
```

Cependant il y a une différence entre le package ```main``` et tout les autre.<br>
Le package main est réservé pour du code qui va etre éxécuté (go run nom_du_fichié.go).<br>
La ou les autres sont prévue pour etre importé.

> [!IMPORTANT]
> Deux fichier go avec des package différents ne peuvent pas etre dans le meme dossier.

---

### import de package

Une fois qu'un package a été créé il peut etre importé dans d'autre fichiers pour être utiliser.<br>
Dans l'éxemple du début un package est importé a la ligne ...

```go
import "fmt"
```

dans ce cas le package offre la possibilité a go d'intéragire avec le terminal.