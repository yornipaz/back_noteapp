package seeder

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/yornifpaz/back_noteapp/app/models"
	"gorm.io/gorm"
)

type Seeder interface {
	Run()
}

type SeederManager struct {
	db *gorm.DB
}

// run implements Seeder.
func (s *SeederManager) Run() {
	// Definir un flag de línea de comandos para el nombre del seeder
	var seederName string
	flag.StringVar(&seederName, "seeder", "", "Nombre del seeder a ejecutar")
	flag.Parse()

	if seederName == "-seeder" {
		fmt.Println("Por favor, proporciona el nombre del seeder usando el flag -seeder")
		os.Exit(1)
	}
	// Ejecutar el seeder correspondiente según el nombre proporcionado
	switch seederName {
	case "roles":
		if err := s.seedRoles(s.db); err != nil {
			panic("Error al sembrar roles: " + err.Error())
		}
		fmt.Println("Seeder de roles ejecutado exitosamente")
	case "permissions":
		if err := s.seedPermissions(s.db); err != nil {
			panic("Error al sembrar permisos: " + err.Error())
		}
		fmt.Println("Seeder de permisos ejecutado exitosamente")
	case "admin_users":
		if err := s.seedAdminUsers(s.db); err != nil {
			panic("Error al sembrar usuarios administradores: " + err.Error())
		}
		fmt.Println("Seeder de usuarios administradores ejecutado exitosamente")
	default:

	}

}
func (s *SeederManager) seedRoles(db *gorm.DB) error {
	roles := []models.Role{
		{ID: "1", Name: "Agricultor"},
		{ID: "2", Name: "Trabajador Agrícola"},
		{ID: "3", Name: "Especialista en Agricultura"},
		{ID: "4", Name: "Administrador"},
	}
	// Agrega más roles según sea necesario

	for _, role := range roles {

		if err := db.Where(models.Role{ID: role.ID}).FirstOrCreate(&role).Error; err != nil {
			return err
		}
	}

	return nil
}
func (s *SeederManager) seedPermissions(db *gorm.DB) error {
	permissions := []models.Permission{
		{ID: "1", Name: "VerProductos"},
		{ID: "2", Name: "CrearProductos"},
		{ID: "3", Name: "EditarProductos"},
		{ID: "4", Name: "EliminarProductos"},
		// Agrega más permisos según sea necesario
	}

	for _, permission := range permissions {
		if err := db.Where(models.Permission{ID: permission.ID}).FirstOrCreate(&permission).Error; err != nil {
			return err
		}
	}

	return nil
}
func (s *SeederManager) seedAdminUsers(db *gorm.DB) error {
	adminUsers := []models.User{
		{
			ID:        "1",
			FirstName: "Admin",
			LastName:  "User",
			Email:     "admin@prueba.com",
			Password:  "hashed_password", // Recuerda almacenar contraseñas de manera segura (con hash y sal)
			IsActive:  true,
			Status:    "created",
			Verified:  true,
			LogoutAt:  time.Now(),
			Roles:     []models.Role{{ID: "1", Name: "Administrator"}}, // Asignar roles según tus necesidades
		},
		// Puedes agregar más usuarios administradores según sea necesario
	}

	for _, adminUser := range adminUsers {
		if err := db.Where(models.User{Email: adminUser.Email}).FirstOrCreate(&adminUser).Error; err != nil {
			return err
		}
	}

	return nil
}

func NewSeederManager(db *gorm.DB) Seeder {
	return &SeederManager{db: db}
}
