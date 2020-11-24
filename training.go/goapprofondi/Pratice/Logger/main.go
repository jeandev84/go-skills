package main

/*
 FILENAME
 log.txt     : nom du fichier

 FLAGS
 os.O_CREATE : creer le fichier s'il n'existe pas
 os.O_APPEND : ecrire et ajouter des informations dans le fichier
 os.O_WRONLY : mode ecriture seulement

 PERMISSION
 0666 (ouvert pour tout le monde)

 =========================================
 $ go build && ./goapprofondi
*/
func main() {

	// log.SetOutput(file)
	// log.Println("Hello gophers!")
	// log.Println("Output in logs.txt")

	initLoggers()
	InfoLogger.Println("Message d'info")
	WarningLogger.Println("Attention, le fichier n'existe pas")
	ErrorLogger.Println("Connexion a la BDD impossible")

	// Bibliotheque recommandee pour les logs:
	// https://github.com/Siruspen/logrus
}
