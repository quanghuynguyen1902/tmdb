package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tmdb/pkg/config"

	"github.com/gin-gonic/gin"
)

type githubUserResponse struct {
	Login     string `json:"login"`
	Id        int    `json:"id"`
	AvatarUrl string `json:"avatar_url"`
}

func (s *Server) rootHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, `<a href="/login/github/">LOGIN</a>`)
	}
}

func (s *Server) loginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		githubData, ok := c.Get("user")
		fmt.Printf("data %v", githubData)
		if !ok {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}
		var prettyJSON bytes.Buffer
		// json.indent is a library utility function to prettify JSON indentation
		err := json.Indent(&prettyJSON, []byte(githubData.(string)), "", "\t")
		if err != nil {
			log.Fatalf("JSON parse error", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		// Return the prettified JSON as a string
		c.JSON(http.StatusOK, prettyJSON.Bytes())
	}
}

func (s *Server) githubLoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg, err := config.NewConfig()
		if err != nil {
			log.Fatalf("create new config error %s", err)
			return
		}

		redirectURL := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s", cfg.ClientId, "http://localhost:8080/login/github/callback")

		c.Redirect(http.StatusMovedPermanently, redirectURL)
	}
}

func (s *Server) githubCallbackHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		code := c.Query("code")

		githubAccessToken := getGithubAccessToken(code)

		githubData := getGithubData(githubAccessToken)
		//var prettyJSON bytes.Buffer
		//// json.indent is a library utility function to prettify JSON indentation
		//err := json.Indent(&prettyJSON, []byte(githubData), "", "")
		//if err != nil {
		//	log.Fatalf("JSON parse error", err)
		//	c.JSON(http.StatusBadRequest, nil)
		//	return
		//}

		// Return the prettified JSON as a string
		c.JSON(http.StatusOK, githubData)
	}
}

func getGithubData(accessToken string) githubUserResponse {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		log.Panic("API Request creation failed")
	}

	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Panic("Request failed")
	}

	respBody, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("data %v", string(respBody))

	var ghResponse githubUserResponse
	json.Unmarshal(respBody, &ghResponse)

	return ghResponse
}

func getGithubAccessToken(code string) string {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("create new config error %s", err)
	}

	requestBodyMap := map[string]string{"client_id": cfg.ClientId, "client_secret": cfg.ClientSecret, "code": code}
	requestJSON, _ := json.Marshal(requestBodyMap)

	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(requestJSON))
	if err != nil {
		log.Fatalf("Request creation failed", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Request failed", err)
	}

	responseBody, _ := ioutil.ReadAll(resp.Body)

	// Represents the response received from Github
	type githubAccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	var ghResponse githubAccessTokenResponse
	json.Unmarshal(responseBody, &ghResponse)

	return ghResponse.AccessToken
}
