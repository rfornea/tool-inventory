package models

import (
	"github.com/gin-gonic/gin"
	"github.com/rfornea/library/pkg/backend/utils"
	"testing"
)

const testDatabaseErrorString = "should only be calling this method on test database"

func setupTests(t *testing.T) {
	utils.SetTesting("../.env")
	Connect(utils.Env.TestDatabaseURL)
	DeleteTripsForTest(t)
	gin.SetMode(gin.TestMode)
}

func ReturnValidBookDataForTest(t *testing.T) []BookData {
	utils.AbortIfNotTesting(t)
	return []BookData{
		{
			Title:           "Anna Karenina",
			AuthorFirstName: "Leo",
			AuthorLastName:  "Tolstoy",
			Isbn:            "0-8659-3759-1",
			Description:     "Anna Karenina tells of the doomed love affair between the sensuous and rebellious Anna and the dashing officer, Count Vronsky. Tragedy unfolds as Anna rejects her passionless marriage and thereby exposes herself to the hypocrisies of society. Set against a vast and richly textured canvas of nineteenth-century Russia, the novel's seven major characters create a dynamic imbalance, playing out the contrasts of city and country life and all the variations on love and family happiness.",
		},
		{
			Title:           "Don Quixote",
			AuthorFirstName: "Miguel",
			AuthorLastName:  "de Cervantes",
			Isbn:            "0-4821-4331-2",
			Description:     "Don Quixote has become so entranced by reading chivalric romances, that he determines to become a knight-errant himself. In the company of his faithful squire, Sancho Panza, his exploits blossom in all sorts of wonderful ways.",
		},
	}
}

func MakeBooksForTest(t *testing.T) []Book {
	utils.AbortIfNotTesting(t)

	var books []Book

	bookData := ReturnValidBookDataForTest(t)

	for _, val := range bookData {
		book := Book{
			BookData: val,
		}
		DB.Create(&book)
		books = append(books, book)
	}
	return books
}

func DeleteTripsForTest(t *testing.T) {
	deleteForTest(t, "books")
}

func deleteForTest(t *testing.T, table string) {
	if utils.Env.DatabaseURL != utils.Env.TestDatabaseURL {
		t.Fatalf(testDatabaseErrorString)
	} else {
		cmd := "DELETE from " + table + ";"
		DB.Exec(cmd)
	}
}
