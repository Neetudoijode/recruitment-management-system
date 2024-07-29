package utils

import (
	"bytes"
	"encoding/json"
	//"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type ResumeData struct {
    Education  []Education  `json:"education"`
    Experience []Experience `json:"experience"`
    Skills     []string     `json:"skills"`
    Name       string       `json:"name"`
    Email      string       `json:"email"`
    Phone      string       `json:"phone"`
}

type Education struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

type Experience struct {
    Dates []string `json:"dates"`
    Name  string   `json:"name"`
    URL   string   `json:"url"`
}

func ParseResume(filePath string) (*ResumeData, error) {
    apiKey := "gNiXyflsFu3WNYCz1ZCxdWDb7oQg1Nl1" // Use the actual API key
    url := "https://api.apilayer.com/resume_parser/upload"

    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var b bytes.Buffer
    writer := multipart.NewWriter(&b)
    part, err := writer.CreateFormFile("file", filePath)
    if err != nil {
        return nil, err
    }

    _, err = io.Copy(part, file)
    if err != nil {
        return nil, err
    }
    writer.Close()

    req, err := http.NewRequest("POST", url, &b)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", writer.FormDataContentType())
    req.Header.Set("apikey", apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var resumeData ResumeData
    err = json.Unmarshal(body, &resumeData)
    if err != nil {
        return nil, err
    }

    return &resumeData, nil
}
