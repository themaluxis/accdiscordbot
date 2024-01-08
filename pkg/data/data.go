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
	"kyalami":      "https://www.carmag.co.za/wp-content/uploads/2023/02/FM-1-jpg.webp",
	"imola":        "https://static.bolognawelcome.com/immagini/68/69/7d/6a/20220518143654_landscape_16_9_mobile.jpg",
	"nurburgring":  "https://hips.hearstapps.com/hmg-prod/images/dsc01990-jpg-1656431148-64501d0e8ef7b.jpeg",
	"spa":          "https://www.fia.com/sites/default/files/styles/content_details/public/news/main_image/spa.jpeg",
	"misano":       "https://www.roadracingworld.com/wp-content/uploads/2023/05/43b3d773-bb4a-019b-075d-d4f516cf0341_1685124075.jpg",
	"watkins_glen": "https://www.macny.org/wp-content/uploads/2019/05/WatkinsGlen.png",
	"paul_ricard":  "https://www.lamborghini.com/sites/it-en/files/DAM/lamborghini/facelift_2019/motorsport/circuits/EUROPE/valencia/2023/paul_d.jpg",
	"barcelona":    "https://coachdaveacademy.com/wp-content/uploads/2023/07/blog-image-barcat1112-1024x576.jpg",
	"laguna_seca":  "https://www.montereyherald.com/wp-content/uploads/2019/12/CM5_0542.jpg",
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
