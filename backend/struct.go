package main

type Setting struct {
	BackEnd BackEndSetting  `json:"Backend"`
	DB      DatabaseSetting `json:"DataBase"`
}
type BackEndSetting struct {
	BackendPort string `json:"BACKENDPORT"`
	UserName    string `json:"USERNAME"`
	PassWord    string `json:"PASSWORD"`
	NetWorkType string `json:"NETWORKTYPE"`
}
type DatabaseSetting struct {
	DBServerAddr string `json:"DBSERVERADDR"`
	DBPort       int    `json:"DBPORT"`
	DataBase     string `json:"DATABASE"`
}
type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Token_btof struct {
	AccessToken string
	Signature   string
}
type Token_ftob struct {
	Token string `json:"token"`
}
type Response struct {
	Imagebase64 string
}

// type PriceData struct {
// 	VanishingJourney struct {
// 		First int `json:"First"`
// 		Last  int `json:"Last"`
// 	} `json:"Vanishing_Journey"`
// 	ChuChuIsland struct {
// 		First int `json:"First"`
// 		Last  int `json:"Last"`
// 	} `json:"Chu_Chu_Island"`
// 	Lachelein struct {
// 		First int `json:"First"`
// 		Last  int `json:"Last"`
// 	} `json:"Lachelein"`
// 	Arcana struct {
// 		First int `json:"First"`
// 		Last  int `json:"Last"`
// 	} `json:"Arcana"`
// 	Morass struct {
// 		First int `json:"First"`
// 		Last  int `json:"Last"`
// 	} `json:"Morass"`
// 	Esfera struct {
// 		First int `json:"First"`
// 		Last  int `json:"Last"`
// 	} `json:"Esfera"`
// 	Cernium struct {
// 		First int `json:"First"`
// 		Last  int `json:"Last"`
// 	} `json:"Cernium"`
// 	Arcus struct {
// 		First int `json:"First"`
// 		Last  int `json:"Last"`
// 	} `json:"Arcus"`
// 	Odium struct {
// 		First int `json:"First"`
// 		Last  int `json:"Last"`
// 	} `json:"Odium"`
// 	ShangriLa struct {
// 		First int `json:"First"`
// 		Last  int `json:"Last"`
// 	} `json:"Shangri_La"`
// 	Arteria struct {
// 		First int `json:"First"`
// 		Last  int `json:"Last"`
// 	} `json:"Arteria"`
// 	Carcion struct {
// 		First int `json:"First"`
// 		Last  int `json:"Last"`
// 	} `json:"Carcion"`
// }

// 符文金額PriceData與星力StarForce的JSON結構
type PriceData struct {
	VanishingJourney Range `json:"Vanishing_Journey"`
	ChuChuIsland     Range `json:"Chu_Chu_Island"`
	Lachelein        Range `json:"Lachelein"`
	Arcana           Range `json:"Arcana"`
	Morass           Range `json:"Morass"`
	Esfera           Range `json:"Esfera"`
	Cernium          Range `json:"Cernium"`
	Arcus            Range `json:"Arcus"`
	Odium            Range `json:"Odium"`
	ShangriLa        Range `json:"Shangri_La"`
	Arteria          Range `json:"Arteria"`
	Carcion          Range `json:"Carcion"`
}

type StarForceData struct {
	Scope    string `json:"Scope"`
	Range    Range  `json:"Range"`
	IsWeapon bool   `json:"IsWeapon"`
}

type Range struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// 星火數據
type BonusStatData struct {
	Level     int    `json:"level"`
	WeaponATT int    `json:"weaponATT"`
	Tier      string `json:"Tier"`
}

// 各種小算盤的JSON結構
type CalculatorData struct {
	Calculator1 Calculator1 `json:"Calculator1"`
	Calculator2 Calculator2 `json:"Calculator2"`
	Calculator3 Calculator3 `json:"Calculator3"`
	Calculator4 Calculator4 `json:"Calculator4"`
}

type Calculator1 struct {
	CombatEffectiveness int `json:"Combat Effectiveness"`
}

type Calculator2 struct {
	CurrentTotalOfBossDamage int `json:"Current Total of Boss Damage"`
	BossDamageEfficiency     int `json:"Boss Damage Efficiency"`
	AttackDamageEfficiency   int `json:"Attack Damage Efficiency"`
}

type Calculator3 struct {
	Original int `json:"Original"`
	Increase int `json:"Increase"`
}

type Calculator4 struct {
	Original int `json:"Original"`
	Decrease int `json:"Decrease"`
}
