package models

type Question struct {
    ID            int      `json:"id"`
    Question      string   `json:"question"`
    Options       []string `json:"options"`
    CorrectAnswer int      `json:"correct_answer"`
}
