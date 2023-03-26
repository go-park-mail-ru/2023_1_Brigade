package repository

//func TestAddValueToRedis(t *testing.T) {
//	rdb, mock := redismock.NewClientMock()
//
//	// Ожидаем, что будет вызваны методы Set и Expire
//	mock.ExpectSet("my-key", "my-value").ExpectExpire("my-key", 3600*time.Second)
//
//	// Вызываем функцию добавления значения в Redis
//	err := addValueToRedis(rdb, "my-key", "my-value")
//
//	// Проверяем, что добавление значения прошло успешно
//	assert.Nil(t, err)
//	assert.NoError(t, mock.ExpectationsWereMet())
//
//	// Закрываем mock-клиент
//	err = mock.Close()
//	assert.Nil(t, err)
//}
//
//func addValueToRedis(r *redis.Client, key string, value string) error {
//	// Добавляем значение в Redis-хранилище
//	err := r.Set(key, value, 3600*time.Second).Err()
//	if err != nil {
//		return err
//	}
//
//	// Устанавливаем время жизни ключа
//	err = r.Expire(key, 3600*time.Second).Err()
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

//func TestPostgres_GetSessionByCookie_True(t *testing.T) {
//	var ctx echo.Context
//	cookie := uuid.New().String()
//
//	//db, mock, err := sqlmock.New()
//	rdb, mock := redismock.NewNiceMock()
//	mock.E("my-key", "my-value").ExpectExpire("my-key", 3600*time.Second)
//	//require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
//	//defer rdb
//
//	rowMain := sqlmock.NewRows([]string{"cookie"}).
//		AddRow(cookie)
//
//	mock.
//		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM session WHERE cookie=$1`)).
//		WithArgs(cookie).
//		WillReturnRows(rowMain)
//
//	dbx := sqlx.NewDb(db, "sqlmock")
//	repo := NewAuthSessionMemoryRepository(dbx)
//
//	_, err = repo.GetSessionByCookie(ctx, cookie)
//	require.NoError(t, err)
//
//	err = mock.ExpectationsWereMet()
//	require.NoError(t, err)
//}

//
//func TestPostgres_GetSessionByCookie_False(t *testing.T) {
//	var ctx echo.Context
//	cookie := uuid.New().String()
//
//	db, mock, err := sqlmock.New()
//	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
//	defer db.Close()
//
//	rowMain := sqlmock.NewRows([]string{"cookie"})
//
//	mock.
//		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM session WHERE cookie=$1`)).
//		WithArgs(cookie).
//		WillReturnRows(rowMain)
//
//	dbx := sqlx.NewDb(db, "sqlmock")
//	repo := NewAuthMemoryRepository(dbx)
//
//	_, err = repo.GetSessionByCookie(ctx, cookie)
//	require.Error(t, err, myErrors.ErrSessionNotFound)
//
//	err = mock.ExpectationsWereMet()
//	require.NoError(t, err)
//}
//
//func TestPostgres_CreateUser_True(t *testing.T) {
//	var ctx echo.Context
//	username := "marcussss"
//	email := "marcussss@gmail.com"
//	status := "my status"
//	password := "baumanka"
//	user := model.User{
//		Username: username,
//		Email:    email,
//		Status:   status,
//		Password: password,
//	}
//
//	db, mock, err := sqlmock.New()
//	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
//	defer db.Close()
//
//	rowMain := sqlmock.NewRows([]string{"username", "email", "status", "password"}).
//		AddRow(username, email, status, password)
//
//	mock.
//		ExpectQuery(regexp.QuoteMeta(`INSERT INTO profile (username, email, status, password)`)).
//		WithArgs(username, email, status, password).
//		WillReturnRows(rowMain)
//
//	dbx := sqlx.NewDb(db, "sqlmock")
//	repo := NewAuthMemoryRepository(dbx)
//
//	_, err = repo.CreateUser(ctx, user)
//	require.NoError(t, err)
//
//	err = mock.ExpectationsWereMet()
//	require.NoError(t, err)
//}
//
//func TestPostgres_DeleteSession_True(t *testing.T) {
//	var ctx echo.Context
//	cookie := uuid.New().String()
//
//	db, mock, err := sqlmock.New()
//	require.Nil(t, err, fmt.Errorf("cant create mock: %s", err))
//	defer db.Close()
//
//	rowMain := sqlmock.NewRows([]string{"cookie"}).
//		AddRow(cookie)
//
//	mock.
//		ExpectQuery(regexp.QuoteMeta(`DELETE FROM session WHERE cookie=$1`)).
//		WithArgs(cookie).
//		WillReturnRows(rowMain)
//
//	dbx := sqlx.NewDb(db, "sqlmock")
//	repo := NewAuthMemoryRepository(dbx)
//
//	err = repo.DeleteSession(ctx, cookie)
//	require.NoError(t, err)
//
//	err = mock.ExpectationsWereMet()
//	require.NoError(t, err)
//}
