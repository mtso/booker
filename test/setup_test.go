package test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mtso/booker/server/config"
)

const (
	DropTables = `DO $$ DECLARE
	    r RECORD;
	BEGIN
	    -- if the schema you operate on is not "current", you will want to
	    -- replace current_schema() in query with 'schematodeletetablesfrom'
	    -- *and* update the generate 'DROP...' accordingly.
	    FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = current_schema()) LOOP
	        EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
	    END LOOP;
	END $$;`
)

func createTestDb() {
	app := config.InitializeApp()
	ts := httptest.NewServer(app.Handler)
	defer app.Db.Close()
	defer ts.Close()

	for _, u := range testusers {
		log.Println("Creating user", u)
		buf := BufferUser(u.Username, u.Password)
		_, err := http.Post(ts.URL+"/auth/signup", "application/json", buf)
		if err != nil {
			panic(err)
		}
	}

	client := MakeCookieMonster()
	err := AuthenticateSession(ts, client, User1, Pass1)
	if err != nil {
		panic(err)
	}

	for _, b := range testbooks {
		log.Println("Creating book", b.Title)
		js, err := json.Marshal(b)
		if err != nil {
			panic(err)
		}
		buf := bytes.NewBuffer(js)
		req, err := http.NewRequest("POST", ts.URL+"/api/book", buf)
		if err != nil {
			panic(err)
		}
		_, err = client.Do(req)
		if err != nil {
			panic(err)
		}
	}
}

func dropTables() {
	app := config.InitializeApp()
	defer app.Db.Close()

	log.Println("Dropping tables...")
	_, err := app.Db.Exec(DropTables)
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	// drop tables and create test users 1 and 2
	createTestDb()

	code := m.Run()
	// drop tables
	dropTables()

	os.Exit(code)
}
