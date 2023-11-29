package usecase

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type TranslatorUseCase struct {
}

func NewTranslatorUsecase() *TranslatorUseCase {
	return &TranslatorUseCase{}
}

func (u *TranslatorUseCase) Translate(data string) (string, error) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"q":       data,
		"source":  "auto",
		"target":  "pl",
		"format":  "text",
		"api_key": "",
	})
	if err != nil {
		return "", err
	}

	// Send the HTTP request
	resp, err := http.Post(
		"https://libretranslate.com/translate",
		"application/json",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read and print the server response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Println(responseBody)
	//return responseBody["translatedText"], nil
	return "", nil
}

//
//func (u *TranslatorUseCase) Translate(data string) (string, error) {
//	body := &bytes.Buffer{}
//	writer := multipart.NewWriter(body)
//
//	writer.WriteField("q", data)
//	writer.WriteField("source", "auto")
//	writer.WriteField("target", "en")
//	writer.WriteField("format", "text")
//	writer.WriteField("api_key", "")
//	writer.WriteField("secret", "3SYSZ1W")
//	writer.Close()
//
//	// Создаем HTTP-запрос
//	req, err := http.NewRequest("POST", "URL_ЗАПРОСА", body)
//	if err != nil {
//		fmt.Println("Ошибка при создании запроса:", err)
//		return "", err
//	}
//
//	// Устанавливаем заголовок Content-Type для формата form data
//	req.Header.Set("Content-Type", writer.FormDataContentType())
//
//	// Отправляем запрос
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		fmt.Println("Ошибка при отправке запроса:", err)
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	// Читаем и выводим ответ сервера
//	responseBody, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println("Ошибка чтения ответа:", err)
//		return "", err
//	}
//	return string(responseBody), nil
//}
