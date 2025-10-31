package tests

import (
	"encoding/json"
	"keceox_modules/controllers"
	"keceox_modules/initializers"
	"keceox_modules/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetOriginalURL(t *testing.T) {
	gin.SetMode(gin.TestMode)

	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer sqlDB.Close()

	dialector := postgres.New(postgres.Config{
		Conn:       sqlDB,
		DriverName: "postgres",
	})

	initializers.DB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open gorm db: %v", err)
	}

	rows := sqlmock.NewRows([]string{"id", "short_code", "full_url", "enabled", "expiry_date", "usage_count"}).
		AddRow(1, "abc123", "https://example.com/original-url", true, time.Now().Add(24*time.Hour), 0)

	mock.ExpectQuery("SELECT (.+) FROM \"url_mappings\"").
		WithArgs("abc123", 1).
		WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"url_mappings\"").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{gin.Param{Key: "id", Value: "abc123"}}

	controllers.GetOriginalURL(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "https://example.com/original-url", response["URL"])

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllMyURLS(t *testing.T) {
	gin.SetMode(gin.TestMode)

	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer sqlDB.Close()

	dialector := postgres.New(postgres.Config{
		Conn:       sqlDB,
		DriverName: "postgres",
	})

	initializers.DB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open gorm db: %v", err)
	}

	testUser := models.User{
		Email: "test@example.com",
		Name:  "Test User",
	}
	testUser.ID = 123

	rows := sqlmock.NewRows([]string{"id", "short_code", "full_url", "enabled", "expiry_date", "usage_count", "user_id"}).
		AddRow(1, "abc123", "https://example.com/url1", true, time.Now().Add(24*time.Hour), 5, 123).
		AddRow(2, "xyz789", "https://example.com/url2", true, time.Now().Add(48*time.Hour), 10, 123)

	mock.ExpectQuery("SELECT (.+) FROM \"url_mappings\"").
		WithArgs(true, 123, sqlmock.AnyArg()).
		WillReturnRows(rows)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Set("user", testUser)

	controllers.GetAllMyURLS(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string][]models.Url_mappings
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, 2, len(response["Code"]))
	assert.Equal(t, "abc123", response["Code"][0].ShortCode)
	assert.Equal(t, "xyz789", response["Code"][1].ShortCode)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDisableURL(t *testing.T) {
	gin.SetMode(gin.TestMode)

	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer sqlDB.Close()

	dialector := postgres.New(postgres.Config{
		Conn:       sqlDB,
		DriverName: "postgres",
	})

	initializers.DB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open gorm db: %v", err)
	}

	testUser := models.User{
		Email: "test@example.com",
		Name:  "Test User",
	}
	testUser.ID = 123

	rows := sqlmock.NewRows([]string{"id", "short_code", "full_url", "enabled", "expiry_date", "usage_count", "user_id"}).
		AddRow(1, "abc123", "https://example.com/url1", true, time.Now().Add(24*time.Hour), 5, 123)

	mock.ExpectQuery("SELECT (.+) FROM \"url_mappings\" WHERE Short_Code = (.+) ORDER BY").
		WithArgs("abc123", 1).
		WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE \"url_mappings\"").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = gin.Params{gin.Param{Key: "id", Value: "abc123"}}

	c.Set("user", testUser)

	controllers.DisableURL(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Succesfully disabled the URL permanently!", response["Message"])

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
