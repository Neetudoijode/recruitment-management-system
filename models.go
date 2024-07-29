package models

type User struct {
    Name         string
    Email        string
    Address      string
    UserType     string // "Applicant" or "Admin"
    PasswordHash string
    Headline     string
    Profile      Profile
}

type Profile struct {
    UserID        uint
    ResumeFile    string
    Skills        string
    Education     string
    Experience    string
}

type Job struct {
    ID             uint
    Title          string
    Description    string
    PostedOn       string
    TotalApplications int
    CompanyName    string
    PostedBy       uint
}
