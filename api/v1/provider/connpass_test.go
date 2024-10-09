package provider_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/yashikota/chronotes/pkg/provider"
	"github.com/yashikota/chronotes/pkg/utils"
)

func TestConpassHandler(t *testing.T) {
	w := httptest.NewRecorder()
	err := godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))
	if err != nil && !os.IsNotExist(err) {
		t.Error(err)
	}

	userID := os.Getenv("CONNPASS_USER_ID")
	pass := os.Getenv("CONNPASS_PASS")

	if userID == "" {
		userID = "CONNPASS_USER_ID"
	}

	if pass == "" {
		pass = "CONNPASS_PASS"
	}

	if userID == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("CONNPASS_USER_ID is not set"))
		return
	}

	if pass == "" {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("CONNPASS_PASS is not set"))
		return
	}

	summaries, err := provider.ConnpassProvider()
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	if summaries == nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, errors.New("could not fetch commits"))
		return
	}

	fmt.Println(summaries)

	utils.SuccessJSONResponse(w, summaries)
}
