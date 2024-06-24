# Estrutura de dados com Go

## Arvores binárias de buscas
Arvores binárias de busca ou BinarySearchTrees (BSTs) são uma estrutura de dados
que permite que uma busca seja feita em O(log n) já que a inserção ordena os elementos

### Operacões em uma BST
Uma arvore binária de busca permite as seguintes operações.
 * Inserção de um valor
    -- Os valores inseridos em uma arvore binária de busca utilizam a seguinte estrutura
    ```
        type Node struct {
        value int
        left  *Node
        right *Node
    }
    ```
    Onde o valor é um inteiro. os nodes left e right representam respectivamente o ponteiro para o valor menor e maior

* Travessia
    -- Essa operação consiste em visitar cada um dos nos de uma arvore percorendo-a de tres formas distintas
    * Em ordem
    * Pré ordem 
    * Pós ordem  

    Cada uma dessas operações consiste em visitar os nos da esquerda pra direita porém na travessia em ordem a operação (no caso um simples fmt.Print) é executada apos a visita do ultimo nó mais a esquerda

* Remoção
    Remover um elemento de uma arvore, é a operação mais complicada. já que deveremos ajustar os demais ponteiros para que não deixemos nenhuma estrutura solta