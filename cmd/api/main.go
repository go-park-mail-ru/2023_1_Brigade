package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/http"
	httpchat "project/internal/chat/delivery/http"
	chatrepository "project/internal/chat/repository"
	chatusecase "project/internal/chat/usecase"
	"project/internal/middleware"
	httpuser "project/internal/user/delivery/http"
	userrepository "project/internal/user/repository"
	userusecase "project/internal/user/usecase"
)

var schema = `
-- DROP TABLE Chat;
-- DROP TABLE Message;
-- DROP TABLE Profile;

	
CREATE TABLE Profile (
	id       integer primary key,
	username    varchar(255),
	name varchar(255),
	email    varchar(255),
	status   varchar(255),
	password varchar(255)
);

CREATE TABLE Message (
    id        integer primary key,  
	author_id  integer NOT NULL,
	body      text,
	media     text,  
	created_at text,  
	is_read    bit
);

CREATE TABLE Chat (
    id        integer primary key,
	name      varchar(255),
	created_at text,
	members   Profile[],
	messages  Message[]
);

-- CREATE TABLE Profile (
-- 	id       integer primary key,
-- 	username    varchar(255),
-- 	name varchar(255),
-- 	email    varchar(255),
-- 	status   varchar(255),
-- 	password varchar(255)
-- );
-- 
-- CREATE TABLE Message (
--     id        integer primary key,  
-- 	author_id  integer NOT NULL,
-- 	body      text,
-- 	media     text,  
-- 	created_at text,  
-- 	is_read    bit
-- );
-- 
-- CREATE TABLE Chat (
--  id        integer primary key,
-- 	name      varchar(255),
-- 	created_at text,
-- 	members   integer REFERENCES Profile (Id),
-- 	messages  integer REFERENCES Message (Id)
-- );
`

func main() {

	connStr := "user=golang password=golang dbname=golang sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	//db.MustExec(schema)
	if err != nil {
		return
	}

	repositoryUserImpl := userrepository.NewUserMemoryRepository(db)
	repositoryChatImpl := chatrepository.NewChatMemoryRepository(db)

	userImpl := userusecase.NewUserUsecase(repositoryUserImpl)
	chatImpl := chatusecase.NewChatUsecase(repositoryChatImpl)

	r := mux.NewRouter()

	r.Use(middleware.RequestResponseMiddleware)

	httpuser.NewUserHandler(r, userImpl)
	httpchat.NewChatHandler(r, chatImpl)

	http.ListenAndServe(":8081", r)

	//var DB *sql.DB
	//result, err := DB.Exec(
	//	"INSERT INTO items (`title`, `description`) VALUES (?, ?)",
	//	r.FormValue("title"),
	//	r.FormValue("description"),
	//)
	////__err_panic(err)
	//result.MustBegin()
	//affected, err := result.RowsAffected()
	////__err_panic(err)
	//lastID, err := result.LastInsertId()
	////__err_panic(err)
	//
	//fmt.Println("Insert - RowsAffected", affected, "LastInsertId: ", lastID)
	//
	//http.Redirect(w, r, "/", http.StatusFound)
	//
	//db, err := sqlx.Connect("postgres", "user=golang password=golang dbname=golang sslmode=disable")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//tmp := db.MustExec(schema)
	//db.MustBegin()
	//tmp.RowsAffected()
	//fmt.Println(tmp)
	//repositoryUserImpl := userrepository.NewUserMemoryRepository(db)
	//repositoryChatImpl := chatrepository.NewChatMemoryRepository(db)
	//
	//userImpl := userusecase.NewUserUsecase(repositoryUserImpl)
	//chatImpl := chatusecase.NewChatUsecase(repositoryChatImpl)
	//
	//r := mux.NewRouter()
	//
	//r.Use(middleware.RequestResponseMiddleware)
	//
	//httpuser.NewUserHandler(r, userImpl)
	//httpchat.NewChatHandler(r, chatImpl)
	//
	//http.ListenAndServe(":8081", r)
}
