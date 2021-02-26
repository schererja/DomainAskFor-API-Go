package types

type WhoIsResult struct {
	DomainName  string `json:"domainName"`
	IsAvailable bool   `json:"isAvailable"`
}
