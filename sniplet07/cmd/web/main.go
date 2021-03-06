package main
import (
	"context"
	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os"
)
type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}
func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	pool, err:= pgxpool.Connect(context.Background(), "user=postgres password=789321456POEBEN host=localhost port=5432 dbname=Snippet07 sslmode=disable pool_max_conns=10")
if err!= nil{
	log.Fatalf("Unable to connect to database %v\n", err)
}

	defer pool.Close();

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(), // Call the new app.routes() method
	}
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}







