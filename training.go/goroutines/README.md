 # GOROUTINE 
 - Permet de faire de la programmation parallele
 Elle consiste a separer toutes les taches en petits morceaux
 en les executants part a part


``function LongOperation(i int) {``
``    // Take many seconds``
``    fmt.Println("Done", i)``
``}``

``go LongOperation(1)``
``go LongOperation(2)``
``go LongOperation(3)``
``output: Done 2, Done 3, Done 1``