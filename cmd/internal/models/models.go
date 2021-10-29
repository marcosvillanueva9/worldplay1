package models

type Category struct {
	Id uint64 `json:"id"`
	Label string `json:"label"`
}

type Product struct {
	Id uint64 `json:"id"`
	Label string `json:"label"`
	Type string `json:"type"`
	DownloadURL string `json:"download_url,omitempty"`
	Weight uint64 `json:"weight,omitempty"`
}

type Cart struct {
	Products []Product `json:"products"`
}

type Order struct {
	Id int `json:"id"`
	Products []Product `json:"products"`
}