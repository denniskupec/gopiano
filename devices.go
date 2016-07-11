package gopiano

// Android OS client
var AndroidClient = ClientDescription{
	DeviceModel: "android-generic",
	Username:    "android",
	Password:    "AC7IBG09A3DTSYM4R41UJWL07VLN8JI7",
	BaseURL:     "tuner.pandora.com/services/json/",
	EncryptKey:  "6#26FRL$ZWD",
	DecryptKey:  "R=U!LH$O2B#",
	Version:     "5",
}

// iOS client
var IOSClient = ClientDescription{
	DeviceModel: "IP01",
	Username:    "iphone",
	Password:    "P2E4FC0EAD3*878N92B2CDp34I0B1@388137C",
	BaseURL:     "tuner.pandora.com/services/json/",
	EncryptKey:  "721^26xE22776",
	DecryptKey:  "20zE1E47BE57$51",
	Version:     "5",
}

// Palm WebOS client
var PalmClient = ClientDescription{
	DeviceModel: "pre",
	Username:    "palm",
	Password:    "IUC7IBG09A3JTSYM4N11UJWL07VLH8JP0",
	BaseURL:     "tuner.pandora.com/services/json/",
	EncryptKey:  "%526CBL$ZU3",
	DecryptKey:  "E#U$MY$O2B=",
	Version:     "5",
}

// Windows Mobile
var WinMoClient = ClientDescription{
	DeviceModel: "VERIZON_MOTOQ9C",
	Username:    "winmo",
	Password:    "ED227E10a628EB0E8Pm825Dw7114AC39",
	BaseURL:     "tuner.pandora.com/services/json/",
	EncryptKey:  "v93C8C2s12E0EBD",
	DecryptKey:  "7D671jt0C5E5d251",
	Version:     "5",
}

// Windows Vista Widget client (Pandora One)
var VistaClient = ClientDescription{
	DeviceModel: "WG01",
	Username:    "windowsgadget",
	Password:    "EVCCIBGS9AOJTSYMNNFUML07VLH8JYP0",
	BaseURL:     "internal-tuner.pandora.com/services/json/",
	EncryptKey:  "%22CML*ZU$8YXP[1",
	DecryptKey:  "E#IO$MYZOAB%FVR2",
	Version:     "5",
}

// Adobe Air Desktop client (Pandora One)
var AirClient = ClientDescription{
	DeviceModel: "D01",
	Username:    "pandora one",
	Password:    "TVCKIBGS9AO9TSYLNNFUML0743LH82D",
	BaseURL:     "internal-tuner.pandora.com/services/json/",
	EncryptKey:  "2%3WCL*JU$MP]4",
	DecryptKey:  "U#IO$RZPAB%VX2",
	Version:     "5",
}
