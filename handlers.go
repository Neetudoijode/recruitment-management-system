package handlers

import (
	"io"
	"net/http"
	"os"
	//"recruitment-system/internal/utils"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/signup", SignupHandler).Methods("POST")
    r.HandleFunc("/login", LoginHandler).Methods("POST")
    r.HandleFunc("/uploadResume", UploadResumeHandler).Methods("POST")
    r.HandleFunc("/admin/job", CreateJobHandler).Methods("POST")
    r.HandleFunc("/admin/job/{job_id}", GetJobHandler).Methods("GET")
    r.HandleFunc("/admin/applicants", GetApplicantsHandler).Methods("GET")
    r.HandleFunc("/admin/applicant/{applicant_id}", GetApplicantHandler).Methods("GET")
    r.HandleFunc("/jobs", GetJobsHandler).Methods("GET")
    r.HandleFunc("/jobs/apply", ApplyToJobHandler).Methods("POST")
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
    // Handle user signup
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    // Handle user login
}

func UploadResumeHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(10 << 20) // 10 MB file size limit

    file, _, err := r.FormFile("resume")
    if err != nil {
        http.Error(w, "Invalid file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    tempFile, err := os.CreateTemp("", "resume-*.docx")
    if err != nil {
        http.Error(w, "Failed to save file", http.StatusInternalServerError)
        return
    }
    defer os.Remove(tempFile.Name())

    _, err = io.Copy(tempFile, file)
    if err != nil {
        http.Error(w, "Failed to save file", http.StatusInternalServerError)
        return
    }

    //resumeData, err := utils.ParseResume(tempFile.Name())
    if err != nil {
        http.Error(w, "Failed to parse resume", http.StatusInternalServerError)
        return
    }

    // Store the extracted data in the database (example shown)
    //db.SaveUserProfile(resumeData)

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Resume uploaded and processed successfully"))
}

func CreateJobHandler(w http.ResponseWriter, r *http.Request) {
    // Handle job creation
}

func GetJobHandler(w http.ResponseWriter, r *http.Request) {
    // Handle fetching job details
}

func GetApplicantsHandler(w http.ResponseWriter, r *http.Request) {
    // Handle fetching applicants
}

func GetApplicantHandler(w http.ResponseWriter, r *http.Request) {
    // Handle fetching a specific applicant's details
}

func GetJobsHandler(w http.ResponseWriter, r *http.Request) {
    // Handle fetching job listings
}

func ApplyToJobHandler(w http.ResponseWriter, r *http.Request) {
    // Handle job application
}

