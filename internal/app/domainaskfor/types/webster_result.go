package types

type WebsterResult struct {
	Meta struct {
		ID      string `json:"id"`
		UUID    string `json:"uuid"`
		Src     string `json:"src"`
		Section string `json:"section"`
		Target  struct {
			Tuuid string `json:"tuuid"`
			Tsrc  string `json:"tsrc"`
		} `json:"target"`
		Stems     []string      `json:"stems"`
		Syns      [][]string    `json:"syns"`
		Ants      []interface{} `json:"ants"`
		Offensive bool          `json:"offensive"`
	} `json:"meta"`
	Hwi struct {
		Hw string `json:"hw"`
	} `json:"hwi"`
	Fl  string `json:"fl"`
	Def []struct {
		Sseq [][][]interface{} `json:"sseq"`
	} `json:"def"`
	Shortdef []string `json:"shortdef"`
}
