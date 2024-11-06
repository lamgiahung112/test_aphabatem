package nft_proxy

import "strings"

type NFTFile struct {
	URL  string `json:"URL"`
	Type string `json:"type"`
}

type NFTCreatorSimple struct {
	Address string `json:"address"`
}

type NFTPropertiesSimple struct {
	Category string             `json:"category"`
	Creators []NFTCreatorSimple `json:"creators"`
}

type NFTMetadataSimple struct {
	Name     		string				`json:"name"`
	Symbol   		string				`json:"symbol"`
	Decimals 		uint8 				`json:"-"`
	Image        	string				`json:"image"`
	AnimationURL 	string				`json:"animation_url"`
	ExternalURL  	string				`json:"external_url"`
	Properties 		NFTPropertiesSimple `json:"properties"`
	Files			[]NFTFile			`json:"files"`
	UpdateAuthority string				`json:"updateAuthority"`
}

func (m *NFTMetadataSimple) AnimationFile() *NFTFile {
	for _, f := range m.Files {
		isImage := f.URL == m.Image || strings.Contains(f.Type, "image")
		isNotGif := !strings.Contains(f.Type, "gif")
		
		if isImage && isNotGif {
			continue
		}
		return &f
	}
	return nil
}

func (m *NFTMetadataSimple) ImageFile() *NFTFile {
	for _, f := range m.Files {
		if f.URL == m.Image {
			return &f
		}
	}
	return nil
}
