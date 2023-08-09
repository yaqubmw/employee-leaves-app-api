package cli

import (
	"employeeleave/model"
	"employeeleave/usecase"
	"employeeleave/utils/exceptions"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type EmplController struct {
	emplUC usecase.EmplUseCase
}

func (e *EmplController) EmplMenuForm() {
	fmt.Println(`
|==== Master Employee ====|
| 1. Tambah Data     |
| 2. Lihat Data      |
| 3. Detail Data     |
| 4. Update Data     |
| 5. Hapus Data      |
| 6. Kembali ke Menu |
			     `)
	fmt.Print("Pilih Menu (1-6): ")
	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			empl := e.emplCreateForm()
			err := e.emplUC.RegisterNewEmpl(empl)
			exceptions.CheckError(err)
			return
		case "2":
			empls, err := e.emplUC.FindAllEmpl()
			exceptions.CheckError(err)
			e.emplFindAll(empls)
			return
		case "3":
			e.emplGetForm()
			return
		case "4":
			empl := e.emplUpdateForm()
			err := e.emplUC.UpdateEmpl(empl)
			exceptions.CheckError(err)
			return
		case "5":
			id := e.emplDeleteForm()
			err := e.emplUC.DeleteEmpl(id)
			exceptions.CheckError(err)
			return
		case "6":
			return
		default:
			fmt.Println("Menu tidak ditemukan!")
		}
	}
}

func (e *EmplController) emplCreateForm() model.Employee {
	var (
		Id, Name, PhoneNumber, Email, Address, saveConfirmation string
	)
	fmt.Print("Name: ")
	fmt.Scanln(&Name)
	fmt.Print("PhoneNumber: ")
	fmt.Scanln(&PhoneNumber)
	fmt.Print("Email: ")
	fmt.Scanln(&Email)
	fmt.Print("Address: ")
	fmt.Scanln(&Address)
	fmt.Printf("Id: %s, Name: %s, PhoneNumber: %s, Email: %s, Address: %s akan disimpan (y/t)", Id, Name, PhoneNumber, Email, Address)
	fmt.Scanln(&saveConfirmation)
	if saveConfirmation == "y" {
		empl := model.Employee{
			ID:          uuid.New().String(),
			Name:        Name,
			PhoneNumber: PhoneNumber,
			Email:       Email,
			Address:     Address,
		}
		return empl
	}
	return model.Employee{}
}

func (e *EmplController) emplFindAll(empls []model.Employee) {
	for _, empl := range empls {
		fmt.Println("Employee List")
		fmt.Printf("ID: %s \n", empl.ID)
		fmt.Printf("Name: %s \n", empl.Name)
		fmt.Printf("PhoneNumber: %s \n", empl.PhoneNumber)
		fmt.Printf("Email: %s \n", empl.Email)
		fmt.Printf("Address: %s \n", empl.Address)
		fmt.Println()
	}
}

func (e *EmplController) emplUpdateForm() model.Employee {
	var (
		Id, Name, PhoneNumber, Email, Address, saveConfirmation string
	)
	fmt.Print("Employee ID: ")
	fmt.Scanln(&Id)
	fmt.Print("Employee, Name: ")
	fmt.Scanln(&Name)
	fmt.Print("Employee PhoneNumber: ")
	fmt.Scanln(&PhoneNumber)
	fmt.Print("Employee Email: ")
	fmt.Scanln(&Email)
	fmt.Print("Employee Address: ")
	fmt.Scanln(&Address)
	fmt.Printf("Employee Id: %s, Name: %s, PhoneNumber: %s, Email: %s, Address: %s, akan disimpan (y/t)", Id, Name, PhoneNumber, Email, Address)
	fmt.Scanln(&saveConfirmation)
	if saveConfirmation == "y" {
		empl := model.Employee{
			ID:          uuid.New().String(),
			Name:        Name,
			PhoneNumber: PhoneNumber,
			Email:       Email,
			Address:     Address,
		}
		return empl
	}
	return model.Employee{}
}

func (e *EmplController) emplDeleteForm() string {
	var id string
	fmt.Print("Employee ID: ")
	fmt.Scanln(&id)
	return id
}

func (e *EmplController) emplGetForm() {
	var id string
	fmt.Print("Employee ID: ")
	fmt.Scanln(&id)
	empl, err := e.emplUC.FindByIdEmpl(id)
	exceptions.CheckError(err)
	fmt.Printf("Employee ID %s \n", id)
	fmt.Println(strings.Repeat("=", 15))
	fmt.Printf("Employe ID: %s \n", empl.ID)
	fmt.Printf("Name: %s \n", empl.Name)
	fmt.Printf("PhoneNumber: %s \n", empl.PhoneNumber)
	fmt.Printf("Email: %s \n", empl.Email)
	fmt.Printf("Address: %s \n", empl.Address)
	fmt.Println()
}

func NewEmpController(usecase usecase.EmplUseCase) *EmplController {
	return &EmplController{emplUC: usecase}
}
