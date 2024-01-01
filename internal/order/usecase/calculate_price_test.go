package usecase

import (
	"database/sql"
	"fmt"
	"math/rand"
	"module/internal/order/config"
	"module/internal/order/entity"
	"module/internal/order/infra/database"
	"strconv"
	"testing"

	//mysql
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

/*O suite faz parte do pacote de testes */
type CalculatePriceUseCaseTestSuite struct {
	suite.Suite
	OrderRepository database.OrderRepository
	Db *sql.DB
}

/*O SetupSuite vai rodar antes de executar um teste. O intuito é gerenciar o processo de conexão, gerenciamento
e encerramento do banco de dados. nesse caso, como o banco está em memória, ao rodar o teste a tabela será criada 
e depois fechad*/
func (suite *CalculatePriceUseCaseTestSuite) SetupSuite() {
	config.Load()
	db, err := sql.Open("mysql", config.ConnectionString)
	if err != nil {
		fmt.Print("Erro na conexão!!")
		fmt.Println(db)
	}

	suite.NoError(err)

	sqlStmt := "CREATE TABLE IF NOT EXISTS orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))"
	
	_, err = db.Exec(sqlStmt)
	if err != nil {
		fmt.Print("Erro na criação da tabela!!")
		fmt.Printf("%q: %s\n", err, sqlStmt)
		return
	}
	suite.Db = db
	suite.OrderRepository = *database.NewOrderRepository(db)
}

/*O TearDownTest vai rodar depois de executar um teste*/
func (suite *CalculatePriceUseCaseTestSuite) TearDownTest() {
	//Fechando a conexão
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new (CalculatePriceUseCaseTestSuite))
}

func (suite *CalculatePriceUseCaseTestSuite) TestCalculateFinalPrice() {
	id := rand.Intn(1000 - 10 + 1) + 10
	price := rand.Intn(600 - 5 + 1) + 10
	order, err := entity.NewOrder(strconv.Itoa(id), float64(price), 2.0)
	suite.NoError(err)
	order.CalculateFinalPrice()

	//Utilizando o DTO para validar os dados
	calculateFinalPriceInput := OrderInputDTO{
		ID: order.ID,
		Price: order.Price,
		Tax: order.Tax,
	}

	calculateFinalPriceUseCase := NewCalculateFinalPriceUseCase(suite.OrderRepository)
	output, err := calculateFinalPriceUseCase.Execute(calculateFinalPriceInput)
	suite.NoError(err)

	suite.Equal(order.ID, output.ID)
	suite.Equal(order.Price, output.Price)
	suite.Equal(order.FinalPrice, output.FinalPrice)
}