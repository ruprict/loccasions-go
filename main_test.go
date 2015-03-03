package main

import (
	"encoding/json"
	"fmt"
	"github.com/erikstmartin/go-testdb"
	"github.com/jinzhu/gorm"
	"github.com/ruprict/loccasions-go/api"
	"github.com/ruprict/loccasions-go/app"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUsersIndex(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/users", nil)
	testDB, err := gorm.Open("testdb", "")
	sql := "SELECT  * FROM `users`  WHERE (`users`.deleted_at IS NULL OR `users`.deleted_at <= '0001-01-02')"
	columns := []string{"id", "name", "email"}
	result := `
	1,tim,tim@tim.com
	2,joe,joe@joe.com
	3,bob,bob@bob.com
	`
	testdb.StubQuery(sql, testdb.RowsFromCSVString(columns, result))
	sql = "SELECT  * FROM `loccasions`  WHERE (`user_id` IN (?,?,?))"
	columns2 := []string{"id", "name", "description", "user_id"}
	result2 := `
	1,Loccasion 1,The first loccasion,1
	2,Loccasion 2,The second loccasion,2
	`
	testdb.StubQuery(sql, testdb.RowsFromCSVString(columns2, result2))
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	mockContext := app.NewContext(&testDB, req)
	server := httptest.NewServer(mockContext.Handler(api.UsersIndexHandler))
	defer server.Close()
	resp, err := http.Get(server.URL + "/api/users")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	var js []api.User
	if err = json.Unmarshal(b, &js); err != nil {
		t.Fatalf("reading reponse body: %v", err)
	}
	fmt.Printf("%v", js[0])
	if js[0].ID != 1 {
		t.Error("Bad JSON")
	}

}
