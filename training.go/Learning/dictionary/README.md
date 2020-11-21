# PROJECT DICTIONARY (COMMAND LINE INTERFACE - CLI IN GO)
``STUDY THE CLI in GO (COMMAND LINE INTERFACE)``

- Recherche sur google (golang database key value)

# INSTALL PACKAGE "BadgerDB"
``On indique (3) petits points (...) pour telecharger les 3 sous petites dependences du package``

``go get github.com/dgraph-io/badger/...``
``go get github.com/dgraph-io/badger/v2``

# DOC POUR SE DOCUMENTER SUR LES PACKAGES
Inserer dans la bare de recherche le nom du package et on obtient toutes les informations sur le package
``https://godoc.org``


# GO BUILD (Lance la commande afin d'executer la CLI : go build)
- Quand on lance go build, go essaie de lancer le fichier binaire qui a le meme nom que notre fichier
``jeandev@jeandev-Lenovo-G580:~/go/src/training.go/dictionary$ go build``
``go build: build output "dictionary" already exists and is a directory``
``jeandev@jeandev-Lenovo-G580:~/go/src/training.go/dictionary$``

- Help : ``go build -h``
``jeandev@jeandev-Lenovo-G580:~/go/src/training.go/dictionary$ go build -h``
``usage: go build [-o output] [-i] [build flags] [packages]``
``Run 'go help build' for details.``
``jeandev@jeandev-Lenovo-G580:~/go/src/training.go/dictionary$``


- Pour resoudre ``go build: build output "dictionary" already exists and is a directory``
- On lance la commande ``go build -o dict`` et il sera creer un fichier 'dict'
- Pour lancher la cli alors on fait :
``./dict -action list``

jeandev@jeandev-Lenovo-G580:~/go/src/training.go/dictionary$ ./dict -action list
badger 2020/11/10 12:27:20 INFO: All 2 tables opened in 1ms
badger 2020/11/10 12:27:20 INFO: Discard stats nextEmptySlot: 0
badger 2020/11/10 12:27:20 INFO: Set nextTxnTs to 3
badger 2020/11/10 12:27:20 INFO: Deleting empty file: ./badger/000002.vlog
Dictionary content
                                                                  Jan  1 00:00:00
                                                                  Jan  1 00:00:00
Golang          A wonderful language                              Nov  9 02:31:13
Python          An interpreted language                           Nov 10 11:52:38
badger 2020/11/10 12:27:20 INFO: Lifetime L0 stalled for: 0s
badger 2020/11/10 12:27:20 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 02. Size: 600 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level Done


- Lancer une commande inconnue par example : ``./dict -action toto``
jeandev@jeandev-Lenovo-G580:~/go/src/training.go/dictionary$ ./dict -action toto
badger 2020/11/10 12:29:00 INFO: All 2 tables opened in 2ms
badger 2020/11/10 12:29:00 INFO: Discard stats nextEmptySlot: 0
badger 2020/11/10 12:29:00 INFO: Set nextTxnTs to 3
badger 2020/11/10 12:29:00 INFO: Deleting empty file: ./badger/000003.vlog
Unknown action: toto
badger 2020/11/10 12:29:00 INFO: Lifetime L0 stalled for: 0s
badger 2020/11/10 12:29:00 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 02. Size: 600 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level Done

jeandev@jeandev-Lenovo-G580:~/go/src/training.go/dictionary$ go build -o dict
jeandev@jeandev-Lenovo-G580:~/go/src/training.go/dictionary$ ./dict -action add Kotlin "A great contender to Java"
badger 2020/11/10 16:31:54 INFO: All 3 tables opened in 0s
badger 2020/11/10 16:31:54 INFO: Discard stats nextEmptySlot: 0
badger 2020/11/10 16:31:54 INFO: Set nextTxnTs to 4
badger 2020/11/10 16:31:54 INFO: Deleting empty file: ./badger/000009.vlog
'Kotlin' added to the dictionary
badger 2020/11/10 16:31:54 INFO: Lifetime L0 stalled for: 0s
badger 2020/11/10 16:31:54 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 03. Size: 930 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level Done

jeandev@jeandev-Lenovo-G580:~/go/src/training.go/dictionary$ ./dict -action list
badger 2020/11/10 16:30:47 INFO: All 3 tables opened in 2ms
badger 2020/11/10 16:30:48 INFO: Discard stats nextEmptySlot: 0
badger 2020/11/10 16:30:48 INFO: Set nextTxnTs to 4
badger 2020/11/10 16:30:48 INFO: Deleting empty file: ./badger/000008.vlog
Dictionary content
                                                                  Jan  1 00:00:00
                                                                  Jan  1 00:00:00
                                                                  Jan  1 00:00:00
Golang          A wonderful language                              Nov  9 02:31:13
Kotlin          A great contender to Java                         Nov 10 16:29:12
Python          An interpreted language                           Nov 10 11:52:38
badger 2020/11/10 16:30:48 INFO: Lifetime L0 stalled for: 0s
badger 2020/11/10 16:30:48 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 03. Size: 902 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level Done

=======================================================================
jeandev@jeandev-Lenovo-G580:~/go/src/training.go/dictionary$ go build -o dict
jeandev@jeandev-Lenovo-G580:~/go/src/training.go/dictionary$ ./dict -action list
badger 2020/11/10 16:59:39 INFO: All 3 tables opened in 1ms
badger 2020/11/10 16:59:39 INFO: Discard stats nextEmptySlot: 0
badger 2020/11/10 16:59:39 INFO: Set nextTxnTs to 5
badger 2020/11/10 16:59:39 INFO: Deleting empty file: ./badger/000010.vlog
Dictionary content
                                                                  Jan  1 00:00:00
                                                                  Jan  1 00:00:00
                                                                  Jan  1 00:00:00
Golang          A wonderful language                              Nov  9 02:31:13
Kotlin          A great contender to Java                         Nov 10 16:31:54
Python          An interpreted language                           Nov 10 11:52:38
badger 2020/11/10 16:59:39 INFO: Lifetime L0 stalled for: 0s
badger 2020/11/10 16:59:39 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 03. Size: 930 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level Done
jeandev@jeandev-Lenovo-G580:~/go/src/training.go/dictionary$ ./dict -action define python
badger 2020/11/10 17:00:22 INFO: All 3 tables opened in 1ms
badger 2020/11/10 17:00:22 INFO: Discard stats nextEmptySlot: 0
badger 2020/11/10 17:00:22 INFO: Set nextTxnTs to 5
badger 2020/11/10 17:00:22 INFO: Deleting empty file: ./badger/000011.vlog
Python          An interpreted language                           Nov 10 11:52:38
badger 2020/11/10 17:00:22 INFO: Lifetime L0 stalled for: 0s
badger 2020/11/10 17:00:22 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 03. Size: 930 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level Done
jeandev@jeandev-Lenovo-G580:~/go/src/training.go/dictionary$ ./dict -action remove python
badger 2020/11/10 17:00:51 INFO: All 3 tables opened in 0s
badger 2020/11/10 17:00:51 INFO: Discard stats nextEmptySlot: 0
badger 2020/11/10 17:00:51 INFO: Set nextTxnTs to 5
badger 2020/11/10 17:00:51 INFO: Deleting empty file: ./badger/000012.vlog
'v%!'(string=python) was removed from the dictionary
badger 2020/11/10 17:00:51 INFO: Lifetime L0 stalled for: 0s
badger 2020/11/10 17:00:51 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 03. Size: 950 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level Done
jeandev@jeandev-Lenovo-G580:~/go/src/training.go/dictionary$ ./dict -action remove python
badger 2020/11/10 17:01:57 INFO: All 3 tables opened in 1ms
badger 2020/11/10 17:01:57 INFO: Discard stats nextEmptySlot: 0
badger 2020/11/10 17:01:57 INFO: Set nextTxnTs to 6
badger 2020/11/10 17:01:57 INFO: Deleting empty file: ./badger/000013.vlog
'python' was removed from the dictionary
badger 2020/11/10 17:01:57 INFO: Lifetime L0 stalled for: 0s
badger 2020/11/10 17:01:57 INFO: 
Level 0 [ ]: NumTables: 00. Size: 0 B of 0 B. Score: 0.00->0.00 Target FileSize: 64 MiB
Level 1 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 2 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 3 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 4 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 5 [ ]: NumTables: 00. Size: 0 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level 6 [B]: NumTables: 03. Size: 801 B of 10 MiB. Score: 0.00->0.00 Target FileSize: 2.0 MiB
Level Done