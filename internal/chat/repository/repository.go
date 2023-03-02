package repository

//func NewChatMemoryRepository(db *sqlx.DB) chat.Repository {
//	return &repositoryImpl{db: db}
//}
//
//type repositoryImpl struct {
//	db *sqlx.DB
//}
//
//func (r *repositoryImpl) InsertChatInDB(ctx context.Context, chat model.Chat) http_utils.Response {
//	fmt.Println("POST REPOSITORY")
//	return http_utils.Response{}
//	//r.db.Begin()
//	//
//	//query := "SELECT Name,Skill FROM USER WHERE Skill IN (?,?)"
//	//// searchSkills := []string{"go", "python"}
//	//searchSkills := []interface{}{"go", "python"}
//	//
//	//var records []User
//	//err := r.db.Select(&records, query, searchSkills...)
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//
//	//result, err := r.db.Exec(
//	//	"INSERT INTO Chat (`id`, `name`, `created_at`, `members`, `messages`) VALUES ($1, $2, $3, PARSE_JSON($4), $5)",
//	//	chat.Id,
//	//	chat.Name,
//	//	chat.CreatedAt,
//	//	chat.Members,
//	//	chat.Messages,
//	//)
//	//
//	//if err != nil {
//	//	fmt.Println(err)
//	//	fmt.Println("DB ERROR 1")
//	//	return err
//	//}
//	//
//	//affected, err := result.RowsAffected()
//	////__err_panic(err)
//	//if err != nil {
//	//	fmt.Println("DB ERROR 2")
//	//	return err
//	//}
//	//
//	//lastID, err := result.LastInsertId()
//	//if err != nil {
//	//	fmt.Println("DB ERROR 3")
//	//	return err
//	//}
//	//
//	//fmt.Println("Insert - RowsAffected", affected, "LastInsertId: ", lastID)
//	//
//	//return nil
//}
//
//func (r *repositoryImpl) GetChatInDB(ctx context.Context, chatID int) http_utils.Response {
//	//fmt.Println("GET ID CHAT")
//	//return model.Chat{}, nil
//	return http_utils.Response{}
//}
//
//func (r *repositoryImpl) GetAllChatsInDB(ctx context.Context) http_utils.Response {
//	fmt.Println("GET ALL CHATS")
//	//return []model.Chat{}, nil
//	return http_utils.Response{}
//}
//
//func (r *repositoryImpl) DeleteChatInDB(ctx context.Context, chatID int) http_utils.Response {
//	fmt.Println("DELETE CHAT")
//	//return nil
//	return http_utils.Response{}
//}
