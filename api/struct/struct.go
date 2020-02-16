package _struct

type _struct struct {
	Item Response
}

type Response struct {
	Status int
	Item   Item
}

type Item struct {
	Id               int
	Type             string
	Subtype          string
	Title            string
	Year             int
	Cast             string
	Director         string
	Genres           []IdTitle
	Countries        []IdTitle
	Voice            string
	Duration         Duration
	Langs            int
	Quality          int
	Plot             string
	Tracklist        []IdTitle
	Imdb             int
	ImdbRating       float32
	ImdbVotes        float32
	Kinopoisk        int
	KinopoiskRating  float32
	KinopoiskVotes   int
	Rating           int
	RatingVotes      int
	RatingPercentage int
	Views            int
	Comments         int
	Posters          Posters
	Trailer          Trailer
	Finished         bool
	Advert           bool
	PoorQuality      bool
	CreatedAt        int32
	UpdatedAt        int32
	InWatchlist      bool
	Subscribed       bool
	Subtitles        string
	Bookmarks        []IdTitle
	Ac3              int
	Seasons          []Seasons
}

type Seasons struct {
	Title    string
	Number   int
	Watching Watching
	Episodes []Episodes
}

type Episodes struct {
	Id        int
	Title     string
	Thumbnail string
	Duration  int
	Tracks    int
	Number    int
	Ac3       int
	Audios    []Audios
	Watched   int
	Watching  WatchingEpisode
	Subtitles []Subtitles
	Files     []Files
}

type Files struct {
	Codec     string
	W         int
	H         int
	Quality   string
	QualityId int
	Url       FileUrl
}

type FileUrl struct {
	Http string
	Hls  string
	Hls4 string
	Hls2 string
}

type Subtitles struct {
	Lang  string
	Shift int
	Embed bool
	Url   string
}

type Audios struct {
	Id       int
	Index    int
	Codec    string
	Channels int
	Lang     string
	Type     TypeAudio
	Author   Author
}

type Author struct {
	Id         int
	title      string
	ShortTitle string
}

type TypeAudio struct {
	Id         int
	Title      string
	ShortTitle string
}

type WatchingEpisode struct {
	Status int
	Time   int
}

type Watching struct {
	Status int
}

type Trailer struct {
	Id  int
	Url string
}

type Posters struct {
	Small  string
	Medium string
	Big    string
	Wide   string
}

type Duration struct {
	Average float32
	Total   int
}

type IdTitle struct {
	Id    int
	Title string
}
