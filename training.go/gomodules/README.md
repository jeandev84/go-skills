# GO MODULES (Est UTILSER POUR GENERER LES DEPENDANCES DU PROJECT)
- install module

``go get``

# CREER UN PROJET AVEC GO MODULE

``go mod init nom_du_package_comme_si_on_etait_a_l'aborescence_de_go``

``go mod init training.go/helloword``
``cat go.mod``
``module training.go/helloword``

``go 1.13``

- Importer module
``go get``
``go build``
``./helloword``

* go.mod : like composer.json
* go.sum : like composer.lock

# MISE AJOUR D'UN PACKAGE

``go get github.com/sirupsen/logrus``

# LISTER TOUS LES MODULES
- List tous les modules
``go list -u -m all``

# SUPPRIMER UN MODULE
- Nettoie le module
``go mod tidy``

