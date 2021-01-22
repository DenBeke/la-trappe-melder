package mailer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type listmonkLists struct {
	Data struct {
		Results []struct {
			ID              int       `json:"id"`
			CreatedAt       time.Time `json:"created_at"`
			UpdatedAt       time.Time `json:"updated_at"`
			UUID            string    `json:"uuid"`
			Name            string    `json:"name"`
			Type            string    `json:"type"`
			Optin           string    `json:"optin"`
			Tags            []string  `json:"tags"`
			SubscriberCount int       `json:"subscriber_count"`
		} `json:"results"`
		Total   int `json:"total"`
		PerPage int `json:"per_page"`
		Page    int `json:"page"`
	} `json:"data"`
}

func GetListID(listmonkURL string, ListmonkListID string, ListmonkUsername string, ListmonkPassword string) (int, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, listmonkURL+"/api/lists", nil)
	req.SetBasicAuth(ListmonkUsername, ListmonkPassword)

	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("couldn't get lists: %w", err)
	}

	defer resp.Body.Close()

	lists := listmonkLists{}

	err = json.NewDecoder(resp.Body).Decode(&lists)
	if err != nil {
		return 0, fmt.Errorf("couldn't get lists: %w", err)
	}

	for _, list := range lists.Data.Results {
		if list.UUID == ListmonkListID {
			return list.ID, nil
		}
	}

	return 0, fmt.Errorf("list not found with give ListmonkListID")
}

type listmonkCampaignReq struct {
	Name        string        `json:"name"`
	Subject     string        `json:"subject"`
	Lists       []int         `json:"lists"`
	FromEmail   string        `json:"from_email"`
	ContentType string        `json:"content_type"`
	Messenger   string        `json:"messenger"`
	Type        string        `json:"type"`
	Tags        []interface{} `json:"tags"`
	TemplateID  int           `json:"template_id"`
}

type listmonkCampaignResp struct {
	Data struct {
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Views     int       `json:"views"`
		Clicks    int       `json:"clicks"`
		Lists     []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"lists"`
		StartedAt   interface{}   `json:"started_at"`
		ToSend      int           `json:"to_send"`
		Sent        int           `json:"sent"`
		UUID        string        `json:"uuid"`
		Type        string        `json:"type"`
		Name        string        `json:"name"`
		Subject     string        `json:"subject"`
		FromEmail   string        `json:"from_email"`
		Body        string        `json:"body"`
		SendAt      interface{}   `json:"send_at"`
		Status      string        `json:"status"`
		ContentType string        `json:"content_type"`
		Tags        []interface{} `json:"tags"`
		TemplateID  int           `json:"template_id"`
		Messenger   string        `json:"messenger"`
	} `json:"data"`
}

func CreateCampaign(listmonkURL string, ListmonkListID int, ListmonkUsername string, ListmonkPassword string) (int, error) {

	data := listmonkCampaignReq{
		Name:        "La Trappe Test",
		Subject:     "Batch X",
		Lists:       []int{ListmonkListID},
		FromEmail:   "no-reply@denbeke.be",
		ContentType: "richtext",
		Messenger:   "email",
		Type:        "regular",
		Tags:        []interface{}{},
		TemplateID:  1,
	}

	jsonValue, err := json.Marshal(&data)
	if err != nil {
		return 0, fmt.Errorf("couldn't marshall json: %w", err)
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, listmonkURL+"/api/campaigns", bytes.NewBuffer(jsonValue))
	req.SetBasicAuth(ListmonkUsername, ListmonkPassword)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("couldn't create campaign: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		newStr := buf.String()

		fmt.Printf(newStr)

		return 0, fmt.Errorf("unexecpted status code: %d", resp.StatusCode)
	}

	campaign := listmonkCampaignResp{}

	err = json.NewDecoder(resp.Body).Decode(&campaign)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse campaign: %w", err)
	}

	fmt.Printf("%+v", campaign)

	// Set body

	return campaign.Data.ID, nil

}

func SendMail(listmonkURL string, ListmonkListID string, ListmonkUsername string, ListmonkPassword string) error {

	listID, err := GetListID(listmonkURL, ListmonkListID, ListmonkUsername, ListmonkPassword)
	if err != nil {
		return err
	}

	fmt.Println("listID: ", listID)

	campaignID, err := CreateCampaign(listmonkURL, listID, ListmonkUsername, ListmonkPassword)
	if err != nil {
		return err
	}

	fmt.Println("campaignID: ", campaignID)

	return nil
}
