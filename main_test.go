package main

import (
	"bareksa-interview-project/interfaces/http/middlewares"
	"bareksa-interview-project/interfaces/http/routes"
	log "bareksa-interview-project/util/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/klauspost/compress/gzhttp"
	"github.com/uptrace/bunrouter"
)

func UnitTestSetup() *http.Server {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[!] Please setup a .env file first")
		os.Exit(1)
	}

	// Get LOG_FILE from .env
	filename := fmt.Sprintf("%s.txt", os.Getenv("LOG_FILE"))
	if filename == ".txt" {
		fmt.Println("[!] Log file should not be empty (make sure LOG_FILE property in .env is set correctly)")
		os.Exit(1)
	}

	// Setup log file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("[!] Error creating log file: %s\n", err)
		os.Exit(1)
	}

	// Initialize logger
	customLogger := log.CreateLogger(file)

	defer func() {
		if r := recover(); r != nil {
			customLogger.WriteLog(log.ERROR, r)
		}
	}()

	dbPass := os.Getenv("POSTGRES_PASSWORD")
	if dbPass == "" {
		fmt.Println("[!] POSTGRES_PASSWORD property in .env is not set correctly")
		os.Exit(1)
	}

	migrationPass := os.Getenv("MIGRATION_PASSWORD")
	if migrationPass == "" {
		fmt.Println("[!] MIGRATION_PASSWORD property in .env is not set correctly")
		os.Exit(1)
	}

	redisPass := os.Getenv("REDIS_PASSWORD")
	if redisPass == "" {
		fmt.Println("[!] REDIS_PASSWORD property in .env is not set correctly")
		os.Exit(1)
	}

	router := bunrouter.New()

	handler := http.Handler(router)
	handler = gzhttp.GzipHandler(handler)

	router.Use(middlewares.ErrorMiddleware(&customLogger)).
		WithGroup("", func(group *bunrouter.Group) {
			apiGroups, apiSlice := routes.ApiRoutes(dbPass, migrationPass, redisPass)
			group.WithGroup("/api/v1", apiGroups)
			group.WithGroup("/", routes.BaseRoutes(apiSlice))
		})

	httpServer := &http.Server{
		Addr:    ":8888",
		Handler: handler,
	}

	return httpServer
}

var r = UnitTestSetup()

func JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}

	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}

	return reflect.DeepEqual(j2, j), nil
}

func AssertBodyEqual(expected, result []byte, t *testing.T) {
	res, err := JSONBytesEqual(expected, result)

	str_a := strings.TrimSpace(string(expected))
	str_b := strings.TrimSpace(string(result))

	if res {
		t.Logf("Expected to get body: %v, got: %v [PASS]", str_a, str_b)
	} else {
		if err != nil {
			if str_a == str_b {
				t.Logf("Expected to get body: %v, got: %v [PASS]", str_a, str_b)
			} else {
				t.Errorf("Expected to get body: %v, got: %v [FAIL]", str_a, str_b)
			}
		} else {
			t.Errorf("Expected to get body: %v, got: %v [FAIL]", str_a, str_b)
		}
	}
}

func AssertResponseCodeEqual(expected, result int, t *testing.T) {
	if expected == result {
		t.Logf("Expected to get status code: %v, got: %v [PASS]\n", expected, result)
	} else {
		t.Errorf("Expected to get status code: %v, got: %v [FAIL]\n", expected, result)
	}
}

func CreateAndSendRequest(method string, url string, body io.Reader, t *testing.T) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Content-Type", "application/json")
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.Handler.ServeHTTP(w, req)

	return w
}

func TestHealthCheck(t *testing.T) {
	w := CreateAndSendRequest(http.MethodGet, "/api/v1/health_check", nil, t)
	AssertResponseCodeEqual(w.Code, http.StatusOK, t)

	expected := []byte(`{"message":"🤖: Ayy sir, service is currently healthy, you may want to continue enjoy your life now"}`)
	result := w.Body.Bytes()
	AssertBodyEqual(expected, result, t)
}

func TestRefreshMigrationWrongSecret(t *testing.T) {
	migrationPass := os.Getenv("MIGRATION_PASSWORD") + "_dummy"
	jsonStr := fmt.Sprintf(`{"secret":"%s"}`, migrationPass)
	jsonBytes := []byte(jsonStr)

	w := CreateAndSendRequest(http.MethodPost, "/api/v1/refresh_migration", bytes.NewBuffer(jsonBytes), t)
	AssertResponseCodeEqual(w.Code, http.StatusInternalServerError, t)

	expected := []byte(`The secret is wrong, dont try any harder if you are not the admin`)
	result := w.Body.Bytes()
	AssertBodyEqual(expected, result, t)
}
