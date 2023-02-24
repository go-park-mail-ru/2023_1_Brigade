package http

import (
	"example.com/m/user/repository"
	"example.com/m/user/usecase"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestGetUser(t *testing.T) {
	//ctrl := gomock.NewController(t)
	//defer ctrl.Finish()
	//
	//usecasee := user.NewMockUsecase(ctrl)
	//
	//usecasee.EXPECT().GetUserById(gomock.Eq(1)).Return(model.User{
	//	1, "marcussss", "danila", "123456", "88005553535",
	//}, nil)

	//handler := userHandler{usecase: usecase}

	//req := httptest.NewRequest("GET", url, nil)
	//w := httptest.NewRecorder()

	//handler.GetUserById(1)

	repositoryImpl := repository.NewUserMemoryRepository()
	userImpl := usecase.NewUserUsecase(repositoryImpl)
	handl := NewUserHandler(userImpl)

	http.HandleFunc("/", handl)
	http.HandleFunc("/user/", handl)

	url := "http://127.0.0.1:8081/user/"

	go func() {
		time.Sleep(100 * time.Millisecond)
		client := &http.Client{}
		// Create request
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Fetch Request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		// Read Response Body
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Display Results
		fmt.Println("response Status : ", resp.Status)
		fmt.Println("response Headers : ", resp.Header)
		fmt.Println("response Body : ", string(respBody))
	}()

	http.ListenAndServe(":8081", nil)

	//fmt.Println(1)

	//time.Sleep(100 * time.Millisecond)

	//http.NewRequest("GET", url, nil)

	//time.Sleep(100 * time.Millisecond)
	//
	//time.Sleep(100 * time.Second)

	//GetUser(w, req)

	//e := echo.New()

	//req := httptest.NewRequest(http.MethodGet, "/user/ivan", nil)
	//rec := httptest.NewRecorder()

	//c := e.NewContext(req, rec)
	//c.SetPath("/user/:username")
	//c.SetParamNames("username")
	//c.SetParamValues("ivan")

	//_, err := handler.GetUserById(1)

	//if err != nil {
	//	t.Errorf("err is not nil: %s", err)
	//}
	//
	//body, _ := ioutil.ReadAll(rec.Body)
	//
	//if strings.Trim(string(body), "\n") != `{"Username":"Ivan"}` {
	//	t.Errorf("Expected: %s, got: %s", `{"Username":"Ivan"}`, string(body))
	//}

}
