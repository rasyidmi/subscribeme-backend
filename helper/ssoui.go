package helper

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"projects-subscribeme-backend/models"

	"github.com/jinzhu/copier"
	ssojwt "github.com/ristekoss/golang-sso-ui-jwt"
)

func ValidateSSOTicket(ticket string, serviceUrl string) (*models.ServiceResponse, error) {
	url := fmt.Sprintf("https://sso.ui.ac.id/cas2/serviceValidate?ticket=%s&service=%s", ticket, serviceUrl)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	var response models.ServiceResponse

	err = xml.Unmarshal(bodyBytes, &response)
	if err != nil {
		err = fmt.Errorf("error in unmarshaling: %w", err)
		return nil, err
	}

	data := ssojwt.ReadOrgcode()

	jurusan := data[response.AuthenticationSuccess.Attributes.Kd_org]

	if jurusan.ShortFaculty != "Fasilkom" {
		return nil, errors.New("401")
	}

	copier.Copy(&response.AuthenticationSuccess.Attributes.Jurusan, jurusan)

	return &response, nil
}
