package services

import (
	"fmt"
	"net/http"

	"github.com/joaosoft/errors"

	"github.com/joaosoft/web"
)

func NewExecutorRabbitMq(service *CmdService) *ExecutorRabbitMq {
	return &ExecutorRabbitMq{service: service}
}

type ExecutorRabbitMq struct {
	client  *web.Client
	service *CmdService
}

func (e *ExecutorRabbitMq) Open() (err error) {
	e.client, err = web.NewClient()
	return err
}

func (e *ExecutorRabbitMq) Close() error {
	return nil
}

func (e *ExecutorRabbitMq) Begin() error {
	return nil
}

func (e *ExecutorRabbitMq) Commit() error {
	return nil
}

func (e *ExecutorRabbitMq) Rollback() error {
	return nil
}

func (e *ExecutorRabbitMq) Execute(arg interface{}, args ...interface{}) error {
	request, err := e.client.NewRequest(web.MethodPost, fmt.Sprintf("%s/api/definitions", e.service.config.RabbitMq.Host))
	if err != nil {
		return err
	}

	request.Headers["Authorization"] = []string{"Basic Zm91cnNvdXJjZTpmNHMwdTQ1ZQ=="}

	response, err := request.WithBody([]byte(arg.(string)), web.ContentTypeApplicationJSON).Send()
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/api/definitions", e.service.config.RabbitMq.Host), nil)
	req.Header.Add("Authorization", `Basic Zm91cnNvdXJjZTpmNHMwdTQ1ZQ==`)
	resp, err := client.Do(req)
	fmt.Println(resp)

	if response.Status >= web.StatusBadRequest {
		return errors.New(errors.ErrorLevel, 0, "error importing configurations to rabbitmq [status: %d, error: %s]", response.Status, string(response.Body))
	}

	return nil
}
