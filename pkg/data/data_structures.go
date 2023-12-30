package data

type BestLapInfo struct {
	TrackName string
	CarModel  int
	CarName   string
	BestLap   int
	DriverID  string
}

type TrackData struct {
	SessionType   string        `json:"sessionType"`
	TrackName     string        `json:"trackName"`
	SessionResult SessionResult `json:"sessionResult"`
}

type SessionResult struct {
	LeaderBoardLines []LeaderBoardLine `json:"leaderBoardLines"`
}

type LeaderBoardLine struct {
	Car    Car    `json:"car"`
	Timing Timing `json:"timing"`
}

type Car struct {
	CarModel int      `json:"carModel"`
	Drivers  []Driver `json:"drivers"`
}

type Driver struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	PlayerId  string `json:"playerId"`
}

type Timing struct {
	BestLap  int `json:"bestLap"`
	LapCount int `json:"lapCount"`
}

type MessageEmbedThumbnail struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}
