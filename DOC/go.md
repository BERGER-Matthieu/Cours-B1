# DOCUMENTATION GO

## Syntaxe
Le Go nécessite une syntaxe qui nécessite un ordre précis dans le code pour fonctionner :
1. **déclaration de package**
2. **import de packages** (optionnel)
3. **fonction(s)**
4. **expression(s)**

```go
package main //déclaration de package
import "fmt" //import de packages
func main() { //fonction (dans ce cas une seule)
    //expression (en gros le code)
    name := "John Doe"
    fmt.Println(name)
}
```

> [!IMPORTANT]
> Retenez la syntaxe de tout ça pour les examens

---

### Déclaration de package
Un package est comme une boîte à outils vous permettant de facilement importer et exporter vos fonctions entre différents projets. 
Vous pouvez préciser dans quelle boîte grâce à la ligne...

```go
package main //dans ce cas tout le code ira dans la boîte à outils "main"
```

Cependant il y a une différence entre le package `main` et tous les autres.<br>
Le package main est réservé pour du code qui va être exécuté (`go run nom_du_fichier.go`).<br>
Là où les autres sont prévus pour être importés.

> [!IMPORTANT]
> Deux fichiers go avec des packages différents ne peuvent pas être dans le même dossier.

---

### Import de package
Une fois qu'un package a été créé, il peut être importé dans d'autres fichiers pour être utilisé.<br>
Dans l'exemple du début, un package est importé à la ligne...

```go
import "fmt"
```

Dans ce cas, le package `fmt` offre la possibilité à Go d'interagir avec le terminal.

---

### Fonction
Une fonction permet de donner un nom à une partie de code pour pouvoir la réutiliser plus tard.
Il existe 2 types de fonctions :
- Les fonctions créées par l'utilisateur (comme vous le faites tout au long de la piscine).
- Les fonctions déjà existantes dans le langage (built-in), comme par exemple la fonction `len()` qui permet de récupérer la longueur d'une string.

Exemple de fonction :
```go
func simpleAdder(a int, b int) int { //la fonction "simpleAdder" prend 2 paramètres, a (int) et b (int) et retournera un nouveau int
    return a + b
}
a := 1
b := 2
c := simpleAdder(a, b) //dans ce cas la variable c sera égale à 3
```

#### Fonctions récursives
Une fonction peut s'appeler elle-même, dans ce cas elle est dite récursive.

```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1) //la fonction s'appelle elle-même
}

result := factorial(5) //result sera égal à 120 (5 * 4 * 3 * 2 * 1)
```

> [!IMPORTANT]
> La fonction main() sera démarrée automatiquement quand un fichier est exécuté.

---

### Expression
L'expression est le code en lui-même.
Elle est composée de :
- Variables
- Boucles
- Return (optionnel)
- Fonctions

---

#### Variable
Une variable permet de stocker une valeur.
Il existe plusieurs types de variables :
- int (nombre entier)
- float (nombre décimal)
- bool (booléen)
- string (chaîne de caractères)
- rune (caractère unique)

Et une variable peut être déclarée de plusieurs manières.

```go
maString := "chaîne de caractères" //le := permet de créer la variable et de stocker la valeur directement
maRune := 'a' //selon le format de la variable elle choisira le bon type (entre "" pour string, entre '' pour rune, si il y a un chiffre int, si il y a une virgule float...)
var maFloat float64 //si le mot-clef var est utilisé la variable est créée mais vide (int = 0, string = "", bool = false...)
maFloat = 7.2 //puis une valeur peut lui être attribuée plus tard avec le signe =
var a, b, c int //on peut en déclarer plusieurs du même type
var ( //ou de différents types
    x int
    y float64
    z bool
)
```

##### Index dans les variables
Certaines variables peuvent être indexées pour accéder à leurs éléments individuels :

**String (chaîne de caractères) :**
```go
maString := "Hello"
premierCaractere := maString[0] //récupère le premier caractère (H)
deuxiemeCaractere := maString[1] //récupère le deuxième caractère (e)
fmt.Println(premierCaractere) //affiche 72 (valeur ASCII de 'H')
fmt.Printf("%c", premierCaractere) //affiche H
```

**Array (tableau) :**
```go
monArray := [5]int{10, 20, 30, 40, 50}
premierElement := monArray[0] //récupère 10
dernierElement := monArray[4] //récupère 50
monArray[2] = 35 //modifie le troisième élément (30 devient 35)
```

**Slice (tranche) :**
```go
monSlice := []string{"Go", "Python", "Java"}
premierLangage := monSlice[0] //récupère "Go"
monSlice[1] = "JavaScript" //modifie "Python" en "JavaScript"
```

**Plages d'index :**
```go
maString := "Bonjour"
sousChaine := maString[1:4] //récupère "onj" (de l'index 1 à 3)
debut := maString[:3] //récupère "Bon" (du début à l'index 2)
fin := maString[3:] //récupère "jour" (de l'index 3 à la fin)

monSlice := []int{1, 2, 3, 4, 5}
portion := monSlice[1:4] //récupère [2, 3, 4]
```

> [!IMPORTANT]
> Les index commencent toujours à 0 en Go. Le premier élément est à l'index 0, le deuxième à l'index 1, etc.

---

#### Return
Le mot-clef `return` permet de retourner une valeur depuis une fonction et d'arrêter son exécution.

```go
func divide(a, b int) int {
    if b == 0 {
        return 0 //arrête la fonction et retourne 0
    }
    return a / b //retourne le résultat de la division
}
```

Une fonction peut retourner plusieurs valeurs :
```go
func divideWithRemainder(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder //retourne deux valeurs
}

q, r := divideWithRemainder(10, 3) //q = 3, r = 1
```

---

#### Boucles

Go propose deux types de boucles principales :

##### Boucle for (équivalent du while et for classique)
```go
// Boucle for classique
for i := 0; i < 5; i++ {
    fmt.Println(i) //affiche 0, 1, 2, 3, 4
}

// Boucle while (utilise for sans initialisation ni incrément)
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}

// Boucle infinie
for {
    fmt.Println("Boucle infinie")
    break //utiliser break pour sortir de la boucle
}

// Boucle sur une string ou un slice
name := "Go"
for index, character := range name {
    fmt.Printf("Index: %d, Character: %c\n", index, character)
}
```

##### Mot-clefs de contrôle de boucle
```go
for i := 0; i < 10; i++ {
    if i == 3 {
        continue //passe à l'itération suivante
    }
    if i == 7 {
        break //sort de la boucle
    }
    fmt.Println(i) //affiche 0, 1, 2, 4, 5, 6
}
```

> [!IMPORTANT]
> Go n'a pas de boucle `while` traditionnelle, mais utilise `for` pour tous les types de boucles.

> [!IMPORTANT]
> Retenez la syntaxe de tout ça pour les examens
