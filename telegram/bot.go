package telegram

import (
	"encoding/json"
	"english_bot/lib/e"
	"english_bot/telegram/types"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: "bot" + token,
		client:   http.Client{},
	}
}

func (c *Client) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)
	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return e.Wrapper("can't send message", err)
	}
	return nil
}

func (c *Client) Updates() ([]types.Update, error) {
	query := url.Values{}
	query.Set("offset", "0")
	query.Set("limit", "100")

	data, err := c.doRequest(getUpdatesMethod, query)
	if err != nil {
		return nil, err
	}

	var result struct {
		Ok     bool           `json:"ok"`
		Result []types.Update `json:"result"`
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	if !result.Ok {
		return nil, fmt.Errorf("telegram API error")
	}

	return result.Result, nil
}

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	defer func() { err = e.Wrapper("can't do request", err) }()

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath + method),
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = query.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) Process() {
	//u := c.client.Updates()
	//u.Timeout = 60
	//
	//updates, err := c.GetUpdatesChan(u)
	//if err != nil {
	//	log.Fatalf("failed to start updates channel: %v", err)
	//}
	//
	//for update := range updates {
	//	if update.Message == nil {
	//		continue
	//	}
	//}
	//// Обработка входящего сообщения
	//go c.handleMessage("message")

	// ??????????????????????????
	//bot, err := с.New("")
	//if err != nil {
	//	log.Panic(err)
	//}
	//err = handler.HandleMessages(bot)

	// fucking shit
	//u := c.client.NewUpdate(0)
	//u.Timeout = 60
	//updates, err := bot.GetUpdatesChan(u)
	//if err != nil {
	//	log.Panic()
	//}
	//for update := range updates {
	//	switch update.Message.Text {
	//	case "давай":
	//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "feedback")
	//		_, err := c.client..(msg)
	//		if err != nil {
	//			log.Panic()
	//		}
	//	}
	//}
}
