package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	urlshorter "github.com/stil4004/url-shorter"
)

type Response struct{
	Message string `json:"message"`
}

func RunTests() {
	fmt.Println("Starting tests")
	err := TestForRequest()
	if err != nil{
		fmt.Printf("Test1 not passed❌: %v\n", err)
	} else {
		fmt.Printf("Test 1 passed ✅")

	}

	err = TestForWrongAdress()
	if err != nil{
		fmt.Printf("Test2 not passed❌: %v\n", err)
	} else {
		fmt.Println("Test 2 passed ✅")
	}
	err = TestForEmptyString()
	if err != nil{
		fmt.Printf("Test3 not passed❌: %v\n", err)
	} else {
		fmt.Println("Test 3 passed ✅")
	}

	if err == nil{
		fmt.Println("All Tests passed!")
		return
	}
	fmt.Println("Some tests failed(")
}

const(
	url = "http://localhost:8083/api/shorter"  
)

func TestForRequest() error{
	
	data := &urlshorter.ShortURL{
		Long_url: `http:\\yandewfawdwadawdwdra.com`,
	}

	jsonValue, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Ошибка отправки POST-запроса:", err)
		return err
	}
	responseBody := bytes.NewBuffer(jsonValue)
	

	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		fmt.Println("Ошибка отправки POST-запроса:", err)
		return err
	}


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return err
	}


	var temp urlshorter.ShortURL
	err = json.Unmarshal(body, &temp)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON-ответа:", err)
		return err
	}

	defer resp.Body.Close()

	res, err := http.Get(url + "/" + temp.Short_url)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return err
	}
	defer res.Body.Close()
	

	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return err
	}
	
	var ans urlshorter.ShortURL
	err = json.Unmarshal(body, &ans)
	
	if ans.Long_url != data.Long_url{
		return errors.New("long url is wrong")
	}
	return nil
}

func TestForWrongAdress() error{
	data := &urlshorter.ShortURL{
		Long_url: "ydwdwad",
	}

	jsonValue, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Ошибка отправки POST-запроса:", err)
		return err
	}
	responseBody := bytes.NewBuffer(jsonValue)
	

	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		fmt.Println("Ошибка отправки POST-запроса:", err)
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return err
	}


	var respon Response
	err = json.Unmarshal(body, &respon)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON-ответа:", err)
		return err
	}
	if respon.Message == "the given string is not url"{
		return nil
	}
	return errors.New("wrong answer in not-valid url")
}

func TestForEmptyString() error{
	data := &urlshorter.ShortURL{
		Long_url: "",
	}

	jsonValue, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Ошибка отправки POST-запроса:", err)
		return err
	}
	responseBody := bytes.NewBuffer(jsonValue)
	

	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		fmt.Println("Ошибка отправки POST-запроса:", err)
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return err
	}


	var respon Response
	err = json.Unmarshal(body, &respon)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON-ответа:", err)
		return err
	}
	if respon.Message == "requested empty url"{
		return nil
	}
	return errors.New("wrong answer in empty string")
}