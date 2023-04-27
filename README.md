# Guía: Programación Dinámica

Usando las estrategias vistas en clase resolver:

## Ejercicios

### [**Camino de costo mínimo**](./caminominimo/)

Dada una matriz con valores, donde cada celda contiene el costo de pasar por esa celda. Se desea encontrar un camino que conecte la celda de la esquina superior izquierda con la esquina inferior derecha, cuyo coste total es mínimo. Devolver el costo del camino.

> [!IMPORTANT]
> Resolver utilizando una estrategia _Top/Down_.

### [**Serie de Fibonacci**](./fibonacci/)

Modificar el algoritmo de la serie de Fibonacci por tabulación para que utilice un arreglo de tamaño 2 en lugar de un arreglo de tamaño `n+1`. Esto se puede lograr utilizando dos variables para almacenar los dos últimos valores de la serie y actualizarlos en cada iteración.

### [**Subsecuencia común más larga (LCS)**](./lcs/)

Dadas dos cadenas `s1` y `s2`, encontrar el largo de su subsecuencia en común más larga. Una subsecuencia de un string `s` es un subconjunto de sus caracteres, no necesariamente adyacentes, pero que tienen el mismo orden.

> [!IMPORTANT]
> Resolver utilizando estrategia _Bottom/Up_.

### [**Formas de subir una escalera**](./escalera/) (_Climbing Stairs_)

Dada una escalera con `n` escalones, y una lista de formas posibles en que se puede subir, encontrar el número total de formas de llegar al último escalón, comenzando desde el suelo.

> [!IMPORTANT]
> Resolver utilizando una estrategia _Top/Down_.

### [**Problema de la mochila 0/1**](./mochila/)

Dado el valor y peso de cada uno de $n$ objetos, queremos maximizar el valor de los elementos que colocamos en la mochila sin exceder su capacidad $k$. Se pide encontrar el valor total máximo que no exceda un peso total de $k$. Cada elemento puede ser utilizado a lo sumo una vez.

> [!IMPORTANT]
> Resolver utilizando estrategia _Bottom/Up_.

### [**Suma de subconjuntos**](./subconjuntos/)

Dada una lista de enteros estrictamente positivos, y un entero $k$, crear una función que retorne el número de subconjuntos de la lista que suman $k$.

> [!IMPORTANT]
> Resolver utilizando estrategia _Top/Down_.
