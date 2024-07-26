package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/venture-technology/vtx-account-manager/config"
	"github.com/venture-technology/vtx-account-manager/internal/repository"
	"github.com/venture-technology/vtx-account-manager/models"
)

type ResponsibleService struct {
	responsiblerepository repository.IResponsibleRepository
}

func NewResponsibleService(repo repository.IResponsibleRepository) *ResponsibleService {
	return &ResponsibleService{
		responsiblerepository: repo,
	}
}

func (rs *ResponsibleService) CreateSponsor(ctx context.Context, sponsor *models.Sponsor) error {
	return rs.responsiblerepository.CreateSponsor(ctx, sponsor)
}

func (rs *ResponsibleService) GetPartners(ctx context.Context, cpf *string) ([]models.Sponsor, error) {
	return rs.responsiblerepository.GetPartners(ctx, cpf)
}

func (rs *ResponsibleService) FindAllDriverAtSchool(ctx context.Context, cnpj *string) ([]models.Driver, error) {
	return rs.responsiblerepository.FindAllDriverAtSchool(ctx, cnpj)
}

func (rs *ResponsibleService) BreachSponsor(ctx context.Context, record *int) error {
	return rs.responsiblerepository.BreachSponsor(ctx, record)
}

func (rs *ResponsibleService) CreateEmploymentContract(ctx context.Context, employmentContract *models.EmploymentContract) {

}

func getDistance(ctx context.Context, origin, destination string) (*float64, error) {

	conf := config.Get()

	endpoint := conf.GoogleCloudSecret.EndpointMatrixDistance

	params := url.Values{
		"units":        {"metric"},
		"origins":      {origin},
		"destinations": {destination},
		"key":          {conf.GoogleCloudSecret.ApiKey},
	}

	url := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	log.Print(url)

	resp, err := http.Get(url)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	var data models.DistanceMatrixResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Print(err.Error())
		return nil, err
	}

	if data.Status != "OK" {
		log.Print("Erro na API:", data.Status)
	}

	distance := data.Rows[0].Elements[0].Distance.Text

	distance = strings.TrimSpace(strings.Replace(distance, "km", "", 1))

	kmFloat, err := strconv.ParseFloat(distance, 64)
	if err != nil {
		return nil, err
	}

	return &kmFloat, err

}
