package data

var trackName = map[string]string{
	"barcelona":      "Circuit de Barcelona-Catalunya",
	"brands_hatch":   "Circuit de Brands Hatch",
	"cota":           "Circuit des Amériques",
	"donington":      "Circuit de Donington Park",
	"hungaroring":    "Circuit du Hungaroring",
	"imola":          "Autodromo Enzo e Dino Ferrari",
	"indianapolis":   "Indianapolis Motor Speedway",
	"kyalami":        "Circuit de Kyalami",
	"laguna_seca":    "WeatherTech Raceway Laguna Seca",
	"misano":         "Circuit de Misano",
	"monza":          "Autodromo Nazionale Monza",
	"mount_panorama": "Circuit de Mount Panorama",
	"nurburgring":    "Nürburgring",
	"oulton_park":    "Circuit d'Oulton Park",
	"paul_ricard":    "Circuit Paul Ricard",
	"silverstone":    "Circuit de Silverstone",
	"snetterton":     "Circuit de Snetterton",
	"spa":            "Circuit de Spa-Francorchamps",
	"suzuka":         "Circuit de Suzuka",
	"valencia":       "Circuit de Valencia",
	"watkins_glen":   "Watkins Glen International",
	"zandvoort":      "Circuit Park Zandvoort",
	"zolder":         "Circuit de Zolder",
}

func GetTrackName() map[string]string {
	return trackName
}

var trackImage = map[string]string{
	"kyalami":        "https://www.carmag.co.za/wp-content/uploads/2023/02/FM-1-jpg.webp",
	"mount_panorama": "https://www.bathurst.nsw.gov.au/files/oc-templates/00000000-0000-0000-0000-000000000000/95d1acf0-3fb4-45d7-9991-a0a7dd537984/v/89/background-image-min.jpg",
	"oulton_park":    "https://www.endurance-info.com/sites/default/files/2021-09/British%20GT%20Race%201%20Sunday%2012%2009%2021%200013.jpg",
	"spa":            "https://www.fia.com/sites/default/files/styles/content_details/public/news/main_image/spa.jpeg",
	"suzuka":         "https://asset.japan.travel/image/upload/v1678260718/mie/M_00241_002.jpg",
	"watkins_glen":   "https://www.macny.org/wp-content/uploads/2019/05/WatkinsGlen.png",
	"zandvoort":      "https://www.autohebdo.fr/app/uploads/2021/06/wri2_00005324-002-753x494.jpg",
	"zolder":         "https://www.moto80.be/app/uploads/2021/12/18081_big.jpg",
}

func GetTrackImage() map[string]string {
	return trackImage
}

var carList = map[int]string{
	0:  "Porsche 991 GT3 R",
	1:  "Mercedes-AMG GT3",
	2:  "Ferrari 488 GT3",
	3:  "Audi R8 LMS",
	4:  "Lamborghini Huracán GT3",
	5:  "McLaren 650S GT3",
	6:  "Nissan GT-R Nismo GT3",
	7:  "BMW M6 GT3",
	8:  "Bentley Continental GT3",
	9:  "Porsche 991 II GT3 Cup",
	10: "Nissan GT-R Nismo GT3",
	11: "Bentley Continental GT3",
	12: "AMR V12 Vantage GT3",
	13: "Reiter Engineering R-EX GT3",
	14: "Emil Frey Jaguar G3",
	15: "Lexus RC F GT3",
	16: "Lamborghini Huracan GT3 Evo",
	17: "Honda NSX GT3",
	18: "Lamborghini Huracan SuperTrofeo",
	19: "Audi R8 LMS Evo",
	20: "AMR V8 Vantage",
	21: "Honda NSX GT3 Evo",
	22: "McLaren 720S GT3",
	23: "Porsche 991 II GT3 R",
	24: "Ferrari 488 GT3 Evo",
	25: "Mercedes-AMG GT3",
	26: "BMW M4 GT3",
	27: "BMW M2 Club Sport Racing",
	28: "Porsche 992 GT3 Cup",
	29: "Lamborghini Huracán SuperTrofeo EVO2",
	30: "Ferrari 488 Challenge Evo",
	31: "Audi R8 LMS GT3 Evo 2",
	32: "Ferrari 296 GT3",
	33: "Lamborghini Huracan GT3 Evo 2",
	34: "Porsche 992 GT3 R",
	35: "McLaren 720S GT3 Evo",
	50: "Alpine A110 GT4",
	51: "Aston Martin Vantage GT4",
	52: "Audi R8 LMS GT4",
	53: "BMW M4 GT4",
	55: "Chevrolet Camaro GT4",
	56: "Ginetta G55 GT4",
	57: "KTM X-Bow GT4",
	58: "Maserati MC GT4",
	59: "McLaren 570S GT4",
	60: "Mercedes AMG GT4",
	61: "Porsche 718 Cayman GT4 Clubsport",
}

func GetCarList() map[int]string {
	return carList
}

func PlayerIDWhitelist(lookup string) bool {
	switch lookup {
	case
		"S76561198078913852",
		"S76561198268126242",
		"S76561197986835266",
		"S76561198017272968",
		"S76561198028877013",
		"S76561197996323054":
		return true
	}
	return false
}
