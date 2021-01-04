# Tutorial #6 : Git_GitHub

Este tutorial es inspirado del tutorial de [bluuweb](https://bluuweb.github.io/tutorial-github/guia/). Introduce los conceptos de repositorios GIT y GITHUB para facilitar el manejo y respaldo de códigos al momento de trabajar de forma local o remota, tanto de forma individual como en equipo. 

 1. Introducción 
 2. Despliegar un primer contenedor Docker básico en práctica
 3. Crear sus propias imágenes con el archivo *Dockerfile*
 4. Combinar varias imágenes para construir software complejos con *Docker-compose*
 5. Preguntas
 6. Ejercicio

## 1. Introducción
### 1.1 ¿Qué es Git?

Es un software de control de versiones, cuyo propósito es llevar registro de los cambios realizados en archivos de computadora y coordinar el trabajo que una o varias personas realizan sobre archivos compartidos. Tambien existe la posibilidad de trabajar de forma remota y una opción es **GitHub**.

### 1.2 ¿Qué es GitHub?

Es una plataforma de desarrollo colaborativo para alojar proyectos en la nube utilizando el sistema de control de versiones Git, además cuenta con una herramienta muy útil que es GitHub Pages donde podemos publicar nuestros proyectos estáticos (HTML, CSS y JS) gratis.

### 1.3 ¿Qué es un repositorio?

Un repositorio es un espacio centralizado donde se almacena, organiza, mantiene y difunde información digital, habitualmente archivos informáticos, que pueden contener trabajos científicos, conjuntos de datos o software. 

### 1.4 Plataformas soportadas

Git es un software multiplataforma, es soportado por Linux/Unix, Windows y Mac OS X

### 1.5 Instalación

Puede consultar las instrucciones de instalación de Git [aquí](https://git-scm.com/downloads).

Una vez instalado Git, lo primero que debemos hacer es ver si Git se instaló de manera correcta, y una forma de hacerlo es usando el siguiente comando:
```
//indica la versión de Git que está instalada en el equipo.
git version
```
Si la consola nos arroja el siguiente mensaje, podemos aumir que se instalño correctamente, en caso contrario reinstalar el software.
```
git version 2.24.3
```
## 2. Fundamentos de Git
### 2.1 Comandos básicos

Acontinuación se enlistarán los comandos básicos de **Git**

```
// Ayuda sobre los comandos
git help
```
```
// Iniciar un nuevo repositorio
// Crear la carpeta oculta .git
// Se usa solo una vez
git init
```
```
// Ver que archivos no han sido registrados
git status
// Solo muestra los archivos modificados
git status -s
```
```
// Agrega todos los archivos del repositorio que esten con cambios, para poder hacer "commit"
git add .
```
```
// Crear commit
// -m "detalle" se usa para agregar una pequeña información sobre el commit
git commit -m "primer commit"
```
**Nota:** Un "commit" es un guardado del proyecto en un momento específico, en otras palabras, cuando uno hace commit, guarda los cambios y así permite que nos podamos mover entre las versiones guardadas.
```
// Muestra la lista de commit del mas reciente al más antigüo
git log
// Muestra en una línea los commit realizados
git log --oneline
// Muestra en una línea los commit realizados pero más elegante
git log --oneline --decorate --all --graph
```

### 2.2 Viajes a través de los commit (control de versiones)

Se tienen los siguientes commit
```
93a6e68 (HEAD -> main) Last commit
fb07222 Third commit
ada910e Second commit
34ba199 First commit
696bb0d Initial commit
```

Para mayor información visite [aquí](https://git-scm.com/book/es/v2/Herramientas-de-Git-Reiniciar-Desmitificado)

### 2.3 Renombrar archivos

Puede que queramos renombrar un archivo, es recomendable hacerlo directamente en la línea de comandos para registrar los cambios con git.
```
// Cambiar nombre
git mv nombreOriginal.vue nombreNuevo.vue
```

### 2.4 Eliminar archivos
```
// Cambiar nombre
git rm nombreArchivo.vue
```
### 2.5 Ignorando Archivos
```
arhivo.js // Ignora el archivo en cuestion
*.js // Ignora todos los arhivos con extensión .js
node_modules/ //Ignora toda la carpeta
```
### 2.6 Ramas o branch

Las ramas o branch son caminos paralelos que toma el código, un repositorio puede tener muchos brazos y al hacer los commit correspondientes, estos brazos se pueden fusionar con el frazo prinncipal o brazo master.

```
// Crea una nueva rama
git branch nombreRama
```
```
// Nos muestra en que rama estamos
git branch
```
```
// Nos movemos a la nueva rama
git checkout nombreRama
```
Para ejecutar los siguientes comandos, debe estar en la rama master.

```
// Fusionamos la nueva rama con la rama master
git merge nombreRama
```
```
// Eliminar una rama
git branch -d nombreRama
```
```
// Crear la nuevaRama y viajar a ella
git checkout -b nuevaRama
```
### 2.7 Tags

Los Tags se utilizan para crear versiones y facilitar la descarga de archivos.

```
// Crear un tags con nombre "versión alpha"
git tag versionAlpha -m "versión alpha"
```
```
// Listar tags
git tag
```
```
// Borrar tags
git tag -d nombreTags
```
```
// Hacer una versión en un commit anterior ej: f52f3da
git tag -a nombreTag f52f3da -m "version alpha"
```
```
// Mostrar información del tag
git show nombreTag
```

## 3. GitHub

GitHub es una plataforma online y gratuita que permite el trabajo remoto en proyectos Git. Recordar que Git y GitHub son cosas distintas

### 3.1 Crear una cuenta en Github

Para crear una cuenta haga click [aquí](https://github.com/join?ref_cta=Sign+up&ref_loc=header+logged+out&ref_page=%2F&source=header-home). Recuerde que crearse una cuenta en GitHub es totalmente gratis.

### 3.1 Crear una nuevo repositorio
### 3.1 Subir los tags
### 3.1 Push
### 3.1 Pull
### 3.1 Fetch
### 3.1 Clonar repositorio
## 4. Ejemplo práctico
