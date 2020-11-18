package easyjson

type JsonInfo struct {
	Success      bool         `json:"success"`
	Message      string       `json:"message"`
	MaterialInfo MaterialInfo `json:"materialInfo"`
	IsLogin      bool         `json:"isLogin"`
	Csrf         string       `json:"csrf"`
	PageName     string       `json:"pageName"`
	URIBroker    URIBroker    `json:"uriBroker"`
}
type TitleInfo struct {
	FundLimit         string `json:"fundLimit"`
	NetValue          string `json:"netValue"`
	NetValueDate      string `json:"netValueDate"`
	ProfitSevenDays   string `json:"profitSevenDays"`
	ProfitTenThousand string `json:"profitTenThousand"`
	DayOfGrowth       string `json:"dayOfGrowth"`
	LastWeek          string `json:"lastWeek"`
	RiskEvaluation    string `json:"riskEvaluation"`
	EstablishmentDate string `json:"establishmentDate"`
	AssetSize         string `json:"assetSize"`
	FundManagerName   string `json:"fundManagerName"`
}
type FundManagerInfoList struct {
	Key        string `json:"key"`
	FundName   string `json:"fundName"`
	OfficeDate string `json:"officeDate"`
	Earnings   string `json:"earnings"`
}
type GeneralInfo struct {
	FundName              string                `json:"fundName"`
	EstablishmentDate     string                `json:"establishmentDate"`
	FundCode              string                `json:"fundCode"`
	AssetSize             string                `json:"assetSize"`
	FundCompanyName       string                `json:"fundCompanyName"`
	TrusteeName           string                `json:"trusteeName"`
	FundManagerBackground string                `json:"fundManagerBackground"`
	FundManagerInfoList   []FundManagerInfoList `json:"fundManagerInfoList"`
	InvestPhilosophy      string                `json:"investPhilosophy"`
	InvestStrategy        string                `json:"investStrategy"`
}
type FundBrief struct {
	FundNameAbbr      string      `json:"fundNameAbbr"`
	FundName          string      `json:"fundName"`
	FundCode          string      `json:"fundCode"`
	EstablishmentDate string      `json:"establishmentDate"`
	ShareSize         string      `json:"shareSize"`
	AssetSize         string      `json:"assetSize"`
	FundManagerName   string      `json:"fundManagerName"`
	SaleStatus        string      `json:"saleStatus"`
	FundCompanyName   string      `json:"fundCompanyName"`
	TrusteeName       string      `json:"trusteeName"`
	ManageRate        string      `json:"manageRate"`
	TrusteeRate       string      `json:"trusteeRate"`
	PurchaseMinMount  string      `json:"purchaseMinMount"`
	RedeemMinMount    string      `json:"redeemMinMount"`
	PurchaseRatio     string      `json:"purchaseRatio"`
	RedeemRatio       string      `json:"redeemRatio"`
	GeneralInfo       GeneralInfo `json:"generalInfo"`
}
type MaterialInfo struct {
	ProductID string    `json:"productId"`
	FundCode  string    `json:"fundCode"`
	FundType  string    `json:"fundType"`
	TitleInfo TitleInfo `json:"titleInfo"`
	FundBrief FundBrief `json:"fundBrief"`
}
type URIBroker struct {
	FaviconIcoURL   string `json:"favicon.ico.url"`
	App404URL       string `json:"app.404.url"`
	ZdrmdataRestURL string `json:"zdrmdata.rest.url"`
	AppErrorpageURL string `json:"app.errorpage.url"`
	AuthcenterURL   string `json:"authcenter.url"`
	AppGotoURL      string `json:"app.goto.url"`
	BumngURL        string `json:"bumng.url"`
	OmeoCheckURL    string `json:"omeo.check.url"`
	OmeoGetURL      string `json:"omeo.get.url"`
	AssetsURL       string `json:"assets.url"`
}

type Datas struct {
	Success bool   `json:"success"`
	List    []List `json:"list"`
}
type List struct {
	BizSeq           int    `json:"bizSeq"`
	Time             int64  `json:"time"`
	ForecastNetValue string `json:"forecastNetValue"`
	ForecastGrowth   string `json:"forecastGrowth"`
}
