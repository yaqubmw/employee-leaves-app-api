package delivery

import (
	"employeeleave/usecase"
	"fmt"
	"os"
)

type Console struct {
	// semua usecase taruh disini
	emplUC usecase.EmplUseCase
}

func (c *Console) mainMenuForm() {
	fmt.Println(`
|++++ Enigma Laundry Menu ++++|
| 1. Master Pegawai           |
| 2. Master Jabatan           |
| 3. Master Jenis Cuti        |
| 4. Master Approval Cuti     |              |
| 5. Keluar                   |
		     `)
	fmt.Print("Pilih Menu (1-6): ")
}

func (c *Console) Run() {
	for {
		c.mainMenuForm()
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":

		case "2":
			fmt.Println("Master Jabatan")
		case "3":
			fmt.Println("Master Jenis Cuti")
		case "4":
			fmt.Println("Master Approval Cuti")
		case "5":
			os.Exit(0)
		default:
			fmt.Println("Menu tidak ditemukan")
		}
	}
}

// func NewConsole() *Console {
// 	cfg, err := config.NewConfig()
// 	exceptions.CheckError(err)
// 	dbConn, _ := manager.NewInfraManager(cfg)
// 	gormDB, err := gorm.Open(dbConn.Dialect(), dbConn.Conn())
// 	exceptions.CheckError(err)

// 	emplRepo := repository.NewEmployeeRepository(gormDB) // Use the GORM DB instance
// 	emplUseCase := usecase.NewEmplUseCase(emplRepo)
// 	return &Console{
// 		emplUC: emplUseCase,
// 	}
// }
