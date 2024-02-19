package utilQuery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/JubaerHossain/golang-ddd/pkg/utils"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Pagination(query *gorm.DB, queryValues map[string][]string) *gorm.DB {
	q := url.Values(queryValues)
	page, _ := strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(q.Get("pageSize"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	query = query.Offset(offset).Limit(pageSize) // Pagination

	return query

}

func HashPassword(password string) (string, error) {
	bp := []byte(password)
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hp), nil
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func BodyParse(s interface{}, w http.ResponseWriter, r *http.Request, isValidation bool) error {
	err := json.NewDecoder(r.Body).Decode(s)
	if err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid JSON")
		return err
	}

	if isValidation {
		validate := validator.New()
		validateErr := validate.Struct(s)
		fmt.Println(validateErr)
		if validateErr != nil {
			utils.WriteJSONEValidation(w, http.StatusBadRequest, validateErr.(validator.ValidationErrors))
			return validateErr
		}
	}
	return nil
}

