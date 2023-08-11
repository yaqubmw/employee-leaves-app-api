package repository

import (
	"employeeleave/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var productDummy = []model.Employee{
	{
		ID:          "1",
		Name:        "imron",
		Email:       "iyaron@gmail.com",
		PhoneNumber: "0812281515",
		Address:     "jakarta",
	},
	{
		ID:          "2",
		Name:        "imam",
		Email:       "iyamam@gmail.com",
		PhoneNumber: "0812281525",
		Address:     "tanggerang",
	},
	{
		ID:          "3",
		Name:        "ayu",
		Email:       "iyayu@gmail.com",
		PhoneNumber: "0812281555",
		Address:     "depok",
	},
	{
		ID:          "4",
		Name:        "joko",
		Email:       "joko@gmail.com",
		PhoneNumber: "0812345678",
		Address:     "bandung",
	},
}

type EmployeeRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo EmployeeRepository
}

func (suite *EmployeeRepositoryTestSuite) SetupSuite() {
	// GORM setup dengan SQLite (basis data in-memory)
	gormDB, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	suite.Require().NoError(err, "seharusnya tidak ada kesalahan saat membuka koneksi database")

	// Migrasi tabel berdasarkan model Employee
	err = gormDB.AutoMigrate(&model.Employee{})
	suite.NoError(err, "tidak seharusnya mengembalikan kesalahan saat migrasi tabel")

	// Buat repository menggunakan koneksi database GORM
	suite.repo = NewEmployeeRepository(gormDB)

	// Isi database dengan datadumy
	for _, emp := range productDummy {
		err := suite.repo.Create(emp)
		suite.NoError(err)
	}
}

func (suite *EmployeeRepositoryTestSuite) TearDownSuite() {
	// Tutup koneksi database
	sqlDB, err := suite.db.DB()
	if err != nil {
		fmt.Printf("an error '%s' occurred when accessing underlying DB", err)
		return
	}
	sqlDB.Close()
}

// TestCreateSuccess menguji Create pada repository dengan kasus berhasil.
func (suite *EmployeeRepositoryTestSuite) TestCreateSuccess() {
	// Membuat data karyawan baru untuk disimpan
	newEmployee := model.Employee{
		ID:          "16",
		Name:        "joko",
		Email:       "joko@gmail.com",
		PhoneNumber: "0812345678",
		Address:     "bandung",
	}

	// Melakukan operasi Create pada repository
	err := suite.repo.Create(newEmployee)

	// Penegasan (assertions)
	suite.NoError(err, "seharusnya tidak mengembalikan kesalahan")
	// Anda juga bisa menambahkan pengecekan lain, misalnya memeriksa apakah ID karyawan baru benar-benar telah tersimpan di database
}

// TestCreateFail menguji Create pada repository dengan kasus gagal (contoh: data yang sudah ada).
func (suite *EmployeeRepositoryTestSuite) TestCreateFail() {
	// Membuat data karyawan yang sudah ada untuk disimpan kembali
	duplicateEmployee := model.Employee{
		ID:          "1", // ID ini sudah ada dalam contoh data productDummy
		Name:        "imron",
		Email:       "iyaron@gmail.com",
		PhoneNumber: "0812281515",
		Address:     "jakarta",
	}

	// Melakukan operasi Create pada repository dengan data yang sudah ada
	err := suite.repo.Create(duplicateEmployee)

	// Penegasan (assertions)
	suite.Error(err, "seharusnya mengembalikan kesalahan karena data sudah ada")
}

func (suite *EmployeeRepositoryTestSuite) TestGetByIDSuccess() {
	// Melakukan operasi GetByID pada repository untuk ID yang ada
	employee, err := suite.repo.Get("1")

	suite.NoError(err, "seharusnya tidak mengembalikan kesalahan")
	suite.Equal("imron", employee.Name, "nama seharusnya sesuai")
	suite.Equal("iyaron@gmail.com", employee.Email, "alamat email seharusnya sesuai")
}

// TestGetByIDFail menguji GetByID pada repository dengan kasus gagal (ID tidak ditemukan).
func (suite *EmployeeRepositoryTestSuite) TestGetByIDFail() {
	// Melakukan operasi GetByID pada repository untuk ID yang tidak ada
	_, err := suite.repo.Get("999")

	suite.Error(err, "seharusnya mengembalikan kesalahan karena ID tidak ditemukan")
}
func TestEmployeeRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(EmployeeRepositoryTestSuite))
}
