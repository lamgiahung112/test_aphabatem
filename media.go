package nft_proxy

import "time"

type Media struct{
	ID              uint      `json:"-" gorm:"primaryKey"`
	Mint            string    `json:"mint" gorm:"uniqueIndex"`
	MintDecimals    uint8     `json:"decimals"`
	ImageUri        string    `json:"imageUri"`
	ImageType       string    `json:"imageType"`
	MediaUri        string    `json:"mediaUri,omitempty"`
	MediaType       string    `json:"mediaType,omitempty"`
	LocalPath       string    `json:"-"`
	Name            string    `json:"name,omitempty"`
	Symbol          string    `json:"symbol,omitempty"`
	UpdateAuthority string    `json:"updateAuthority,omitempty"`
	CreatedAt       time.Time `json:"-"`
}

type IMedia struct {
	GetMedia func() *Media
}

type SolanaMedia struct{
	Media
}

func (m *SolanaMedia) GetMedia() *Media {
	return &Media{
		ID:              m.ID,
		Mint:            m.Mint,
		MintDecimals:    m.MintDecimals,
		ImageUri:        m.ImageUri,
		ImageType:       m.ImageType,
		MediaUri:        m.MediaUri,
		MediaType:       m.MediaType,
		LocalPath:       m.LocalPath,
		Name:            m.Name,
		Symbol:          m.Symbol,
		UpdateAuthority: m.UpdateAuthority,
		CreatedAt:       m.CreatedAt,
	}
}
