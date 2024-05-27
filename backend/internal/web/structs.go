package web

type Origin struct {
	ShopName     string `json:"shopName" sql:"shopName"`
	ShopLocation string `json:"shopLocation" sql:"shopLocation"`
}

type Tea struct {
	Id               string `json:"id" sql:"id"`
	Origin           Origin `json:"origin" sql:"origin"`
	Temperature      int    `json:"temperature" sql:"temperature"`
	PortionWeight    int    `json:"portionWeight" sql:"portionWeight"`
	ContainerWeight  int    `json:"containerWeight" sql:"containerWeight"`
	InitialWeight    int    `json:"initialWeight" sql:"initialWeight"`
	BrewingDuration  int    `json:"brewingDuration" sql:"brewingDuration"`
	TeaType          string `json:"teaType" sql:"teaType"`
	TeaName          string `json:"teaName" sql:"teaName"`
	Color            string `json:"color" sql:"color"`
	Size             string `json:"size" sql:"size"`
	InUse            int    `json:"inUse" sql:"inUse"`
	BlendDescription string `json:"blendDescription" sql:"blendDescription"`
}

type TeaStatus struct {
	Active   []interface{} `json:"active"`
	Inactive []interface{} `json:"inactive"`
}
