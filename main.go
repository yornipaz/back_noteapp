package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yornifpaz/back_noteapp/app/http/middleware"
	"github.com/yornifpaz/back_noteapp/config"
	"github.com/yornifpaz/back_noteapp/database/seeder"
	"github.com/yornifpaz/back_noteapp/routes"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {

	newConfigurationApplication := config.NewConfigurationApplication()
	dbApplication, errors := newConfigurationApplication.Init()
	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Println(err.Error())
		}

	}
	seederManager := seeder.NewSeederManager(dbApplication)
	seederManager.Run()
	db = dbApplication

}

func main() {
	app := gin.Default()
	app.Use(middleware.CORSMiddleware())
	routes.NewApplicationRouter(app, db).Setup()
	app.Run()
	// go func() {
	// 	if err := app.Run(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()
	// Inicia el demonio de recarga
	// startReloader()
}

// func startReloader() {
// 	// Espera por cambios en el código fuente
// 	watcher, err := fsnotify.NewWatcher()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer watcher.Close()

// 	// Encuentra la ruta del directorio actual
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Agrega el directorio actual al watcher
// 	err = filepath.Walk(dir, func(path string, info os.FileInfo, errfile error) error {
// 		if info.IsDir() {
// 			return watcher.Add(path)
// 		}
// 		return errfile
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Observa los cambios en el código fuente
// 	go func() {
// 		for {
// 			select {
// 			case event := <-watcher.Events:
// 				if event.Op&fsnotify.Write == fsnotify.Write {
// 					fmt.Println("Se detectaron cambios. Reiniciando el servidor...")
// 					restartServer()
// 				}
// 			case err := <-watcher.Errors:
// 				fmt.Println("Error en el watcher:", err)
// 			}
// 		}
// 	}()

// 	// Espera a que se interrumpa la ejecución (por ejemplo, Ctrl+C)
// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, os.Interrupt)
// 	<-c
// }

// func restartServer() {
// 	// Encuentra el comando para ejecutar este programa
// 	executable, err := os.Executable()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Construye el comando para reiniciar el programa
// 	cmd := exec.Command(executable)

// 	// Pasa las variables de entorno actuales al comando
// 	cmd.Env = os.Environ()

// 	// Pasa los argumentos de la línea de comandos actuales al comando
// 	cmd.Args = os.Args

// 	// Configura la salida estándar y la salida de error del comando para que vayan a la consola actual
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr

// 	// Espera a que el comando termine antes de continuar
// 	err = cmd.Run()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
