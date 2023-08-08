package delivery

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/jutionck/enigma-laundry-apps/config"
	"github.com/jutionck/enigma-laundry-apps/model"
	"github.com/jutionck/enigma-laundry-apps/repository"
	"github.com/jutionck/enigma-laundry-apps/usecase"
	"github.com/jutionck/enigma-laundry-apps/utils/execptions"
)

type Console struct {
	// semua usecase taruh disini
	uomUC usecase.UomUseCase
}

func (c *Console) mainMenuForm() {
	fmt.Println(`
|++++ Employee Leave Menu ++++|
| 1. Master Employee          |
| 2. Transaksi                |
| 3. Keluar                   |
		     `)
	fmt.Print("Pilih Menu (1-3): ")
}

func (c *Console) uomMenuForm() {
	fmt.Println(`
|==== Master Employee ====|
| 1. Tambah Data     |
| 2. Lihat Data      |
| 3. Update Data     |
| 4. Hapus Data      |
| 5. Kembali ke Menu |
			     `)
	fmt.Print("Pilih Menu (1-4): ")
	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			uom := c.uomCreateForm()
			err := c.uomUC.RegisterNewUom(uom)
			execptions.CheckErr(err)
			return
		case "2":
			uoms, err := c.uomUC.FindAllUom()
			execptions.CheckErr(err)
			for _, uom := range uoms {
				fmt.Println("UOM List")
				fmt.Printf("ID: %s \n", uom.Id)
				fmt.Printf("Name: %s \n", uom.Name)
				fmt.Println()
			}
			return
		case "3":
			fmt.Println("Update Data")
		case "4":
			fmt.Println("Hapus Data")
		case "5":
			return
		default:
			fmt.Println("Menu tidak ditemukan!")
		}
	}
}

func (c *Console) uomCreateForm() model.Uom {
	var (
		uomId, uomName, saveConfirmation string
	)
	fmt.Print("UOM Name: ")
	fmt.Scanln(&uomName)
	fmt.Printf("UOM Id: %s, Name: %s akan disimpan (y/t)", uomId, uomName)
	fmt.Scanln(&saveConfirmation)
	if saveConfirmation == "y" {
		uom := model.Uom{
			Id:   uuid.New().String(),
			Name: uomName,
		}
		return uom
	}
	return model.Uom{}
}

func (c *Console) Run() {
	for {
		c.mainMenuForm()
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			c.uomMenuForm()
		case "6":
			os.Exit(0)
		default:
			fmt.Println("Menu tidak ditemukan")
		}
	}
}

func NewConsole() *Console {
	cfg, err := config.NewConfig()
	execptions.CheckErr(err)
	dbConn, _ := config.NewDbConnection(cfg)
	db := dbConn.Conn()
	uomRepo := repository.NewUomRepository(db)
	uomUseCase := usecase.NewUomUseCase(uomRepo)
	return &Console{
		uomUC: uomUseCase,
	}
}
