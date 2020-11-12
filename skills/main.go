package main 



import (
	"fmt"
	"goskills/router"
)



func main () {

	/* fmt.Println("Hello, playground") */
    /* var router *gin.Engine равно router := value */

    
    /* Create new Router */
    r := router.New()

    
    /* Registering routes */
    r.GET("/", indexHandler)
    r.GET("/search/:category", collectHandler)
    r.GET("/result/:category", resultHandler)
    

    // Run Router (запускать router)
    r.Run()
}


/*
// package block
// main function
// import block

Lunch code

/var/www/goskills/hillel$ go build
/var/www/goskills/hillel$ ./hillel

*/
/* 
# https://newsapi.org
# https://github.com/gin-gonic/gin
Language core

- functions
- structs
- pointers
- methods
- interfaces
- errors
- don't panic!
- goroutines
- channels
- simple web server
*/

/*
Compile binary "hillel бинарный"
$ go build ( will generate file hillel in our case )
можно под любой платформе компилировать программу
это делается через $ go build и указать какие-то параметры
*/