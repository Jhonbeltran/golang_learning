# Go?

* Lenguaje de programación multiplataforma
* Compilado, imperativo
* Sintaxis C
* Tipado Estático (pero no tan estricto)
* características de lenguajes interpretados

## Por qué Go?

* Velocidad de ejecución
* Concurrencia
* Librería estandar net/http, net/websocket
* Linter definido desde el principio == Código legible
* Fácil despliegue

***

### Getting Started

* `export GOPATH="/home/jf/Documents/estudio/golang"` Para definir nuestro workspace

### Estructura de Directorios Padre

Para trabajar de la mejor forma con librerias de terceros, vamos a seguir el siguiente patrón:

Dentro del directorio que definimos como $GOPATH crearemos la siguiente estructura de carpetas:

`src/github.com/<username>/<project>`

***
## Paquetes

* Si una función dentro de un paquete está escrita en minuscula `funcitionName` es una función privada. PAra poder llamar una funcion de un paquete esta debe estar en Mayescula `FunctionName`


https://www.godesignpatterns.com/2014/04/new-vs-make.html

> The make() function, on the other hand, is a special built-in function that is used to initialize slices, maps, and channels. Note that make() can only be used to initialize slices, maps, and channels, and that, unlike the new() function, make() does not return a pointer.
