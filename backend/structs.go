package main

type Origin struct {
	ShopName     string `json:"shopName" sql:"ShopName"`
	ShopLocation string `json:"shopLocation" sql:"ShopLocation"`
}

type Tea struct {
	Id               string `json:"id" sql:"Id"`
	Origin           Origin `json:"origin" sql:"Origin"`
	Temperature      int    `json:"temperature" sql:"Temperature"`
	PortionWeight    int    `json:"portionWeight" sql:"PortionWeight"`
	ContainerWeight  int    `json:"containerWeight" sql:"ContainerWeight"`
	InitialWeight    int    `json:"initialWeight" sql:"InitialWeight"`
	BrewingDuration  int    `json:"brewingDuration" sql:"BrewingDuration"`
	TeaType          string `json:"teaType" sql:"TeaType"`
	TeaName          string `json:"teaName" sql:"TeaName"`
	Color            string `json:"color" sql:"Color"`
	Size             string `json:"size" sql:"Size"`
	InUse            int    `json:"inUse" sql:"InUse"`
	BlendDescription string `json:"blendDescription" sql:"BlendDescription"`
}
