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

dans ce cas le package ```fmt``` offre la possibilité a go d'intéragire avec le terminal.

---

### fonction

Une fonction permet de donner un nom a une partie de code pour pouvoir le réutiliser plus tard.

Il existe 2 types de fonctions :

- Les fonctions créées par l'utilisateur (comme vous faites tout au long de la piscine).
- Les fonctions déja éxistantes dans le language (built-in), comme par exemple la fonction ```len()``` qui permet de recuperer la longeur d'une string.

Exemple de fonction :

```go
func simpleAdder(a int, b int) int { //la fonction "simpleAdder" prend 2 parametre, a (int) et b (int) et retournera un nouveau int
    return a + b
}

a := 1
b := 2

c := simpleAdder(a, b) //dans ce cas la variable c sera égale a 3
```

> [!IMPORTANT]
> Une fonction peut etre appelé en elle meme, dans ce cas elle est dite récursive.

> [!IMPORTANT]
> La fonction main() sera démaré automatiquement quand un fichier est éxecuté.

---

### exprésion

L'éxprésion est le code en lui meme.

Elle est composée de :

- Variables
- Boucles
- Return (optionel)
- Fonctions

---

#### variable

Une variable permet de stocker une valeur.

Il existe plusieurs types de variables :
- int (nbr entier)
- float (nbr decimal)
- bool (booléen)
- string (chaine de character)
- rune (character unique)

Et une variable peu etre déclarée de plusieurs manieres.

```go
maString := "chaine de character" //le := permet de creer la variable et de stocker la valeur directement
maRune := 'a' //selon le format de la variable elle choisirat le bon type (entre "" pour string, entre '' pour rune, si il y a un chiffre int, si il y a une virgule float ...)

var maFloat float //si le mot clef var est utilisé la variable est crée mais vide (int = 0, string = "", bool = false...)
maFloat = 7.2 //puis une valeur peut lui etre atribué plus tard avec le signe =

var a, b, c int //on peut en déclarer plusieurs du meme type
var ( //ou de differents types
    x int
    y float
    z bool
)
```
> [!IMPORTANT]
> Retenez la syntaxe de tout ca pour les éxams

---

