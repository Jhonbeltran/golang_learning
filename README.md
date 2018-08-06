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

* `export GOPATH=~/Documents/estudio/golang` Para definir nuestro workspace

### Estructura de Directorios Padre

Para trabajar de la mejor forma con librerias de terceros, vamos a seguir el siguiente patrón:

Dentro del directorio que definimos como $GOPATH crearemos la siguiente estructura de carpetas:

`src/github.com/<username>/<project>`
