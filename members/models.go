package members

// Member is struct to represent member
type Member struct {
	Birthday         string   `json:"birthday"`
	BloodType        string   `json:"blood_type"`
	EnglishFirstname string   `json:"english_first_name"`
	EnglishLastname  string   `json:"english_last_name"`
	Facebook         string   `json:"facebook"`
	Height           int      `json:"height"`
	Hobby            string   `json:"hobby"`
	Instagram        string   `json:"instagram"`
	Likes            []string `json:"like"`
	Nickname         string   `json:"nickname"`
	Province         string   `json:"province"`
	ThaiFirstname    string   `json:"thai_first_name"`
	ThaiLastname     string   `json:"thai_last_name"`
}
