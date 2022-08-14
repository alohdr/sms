package usecase

import (
	"context"
	"fmt"
	"hanoman-id/xendit-payment/internal/apis/operations/sms"
	"hanoman-id/xendit-payment/internal/models"
	"hanoman-id/xendit-payment/pkg/utils"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

func (uc *useCase) GetSmsHistory(ctx context.Context) (*models.SmsHistory, error) {
	var resp models.SmsHistory
	data, err := uc.repo.GetSmsHistory(ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range data {
		resp = append(resp, &models.SmsHistoryItems0{
			ID:          v.ID,
			PhoneNumber: v.PhoneNumber,
			Request: &models.SmsHistoryItems0Request{
				Sender:      v.Sender,
				TypeMessege: v.Type,
			},
			Responses: &models.SmsHistoryItems0Responses{
				Status: v.Status,
			},
		})
	}
	return &resp, nil

}

func (uc *useCase) CreateSms(ctx context.Context, params sms.PostSmsParams) (*string, error) {
	id := uuid.New().String()

	var notifNumber string

	var url string
	var body map[string]interface{}
	var header map[string]string
	data, err := uc.repo.IsSelectProvider(ctx)
	if err != nil {
		return nil, err
	}

	if params.Body.Type == "notif" {
		notifNumber = "1"
	} else if params.Body.Type == "otp" {
		notifNumber = "2"

	}

	template, err := uc.repo.GetTemplateBid(ctx, notifNumber, params.Body.AppsName)
	if err != nil {
		return nil, err
	}

	if strings.EqualFold(data.ID, utils.IdProviderA) {
		body = map[string]interface{}{
			"source":          params.Body.AppsName,
			"destination":     params.Body.PhoneNumber,
			"text":            template.Text + params.Body.Token,
			"clientMessageId": "Teknologi Informasi",
			"encoding":        "",
		}

		if strings.EqualFold(params.Body.Type, utils.TypeNotif) {
			url = os.Getenv("URL_A_OTP")
		} else {
			url = os.Getenv("URL_A_NOTIF")
		}

	} else {
		var channel string

		body = map[string]interface{}{
			"Userid":    "hackraya",
			"Password":  "hackraya123",
			"Sender":    params.Body.AppsName,
			"Msisdn":    params.Body.PhoneNumber,
			"message":   "temp1",
			"division":  "Teknologi Informasi",
			"batchname": "",
			"uploadby":  "API",
			"channel":   channel,
		}
		url = os.Getenv("URL_B")
	}

	hasil := utils.NewHttpRequest(http.MethodPost, url, body, header, nil)
	fmt.Println(hasil)

	err = uc.repo.CreateSms(ctx, params, id)
	if err != nil {
		return nil, err
	}

	resp := utils.SuccessCreateTemplate

	return &resp, nil
}
