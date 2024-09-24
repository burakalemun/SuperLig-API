package main

// EN: Importing necessary libraries for our operations.
// TR: İşlemlerimiz için gerekli kütüphaneleri içe aktarıyoruz.
import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// EN: TeamData represents the season statistics of a football team (matches played, results, goals, and points).
// TR: Bir futbol takımının sezon istatistiklerini (oynanan maçlar, sonuçlar, goller ve puanlar) temsil eder.
type TeamData struct {
	Name         string `json:"name"`
	Played       int    `json:"played"`
	Wins         int    `json:"wins"`
	Draws        int    `json:"draws"`
	Losses       int    `json:"losses"`
	GoalsFor     int    `json:"goals_for"`
	GoalsAgainst int    `json:"goals_against"`
	Average      int    `json:"average"`
	Points       int    `json:"points"`
}

// EN: leagueData fetches league data from the given URL and returns it as a slice of TeamData.
// TR: Verilen URL'den lig verilerini alır ve bir dizi TeamData olarak döner.
func leagueData(dataType string, url string) ([]TeamData, error) {
	// EN: Make an HTTP GET request to retrieve the page content; return an error if the site is not found.
	// TR: HTTP GET isteği yaparak sayfa içeriğini alır; site bulunamazsa hata döner.
	response, err := http.Get(url)
	if err != nil {
		return nil, errors.New("Site not found!")
	}
	defer response.Body.Close()

	// EN: Decode the page content using Windows-1254 character set.
	// TR: Sayfa içeriğini Windows-1254 karakter seti ile çözümler.
	decoder := transform.NewReader(response.Body, charmap.Windows1254.NewDecoder())

	// EN: Read the decoded content into a byte slice; return an error if decoding fails.
	// TR: Çözümlenen içeriği byte dizisine okur; çözümleme başarısız olursa hata döner.
	decodedBytes, err := ioutil.ReadAll(decoder)
	if err != nil {
		return nil, errors.New("Decoding error")
	}

	// EN: Process the decoded bytes into an HTML document; return an error if document loading fails.
	// TR: Çözümlenen byte'lar ile HTML dokümanını işler; doküman yüklenemezse hata döner.
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(decodedBytes)))
	if err != nil {
		return nil, errors.New("Failed to load document")
	}

	var teams []TeamData

	// EN: Find the league table in the HTML by class name; return an error if the table is not found.
	// TR: HTML içinde lig tablosunu sınıf adına göre bulur; tablo bulunamazsa hata döner.
	table := doc.Find("table.s-table")
	if table.Length() == 0 {
		return nil, errors.New("Table not found")
	}

	// EN: If there is a <tbody> element, narrow down the selection to it.
	// TR: Eğer <tbody> elemanı varsa, seçimi ona göre daraltır.
	if tbody := table.Find("tbody"); tbody.Length() > 0 {
		table = tbody
	}

	// EN: Find all rows in the table, excluding the header row.
	// TR: Tablo içindeki tüm satırları, başlık satırı hariç olacak şekilde bulur.
	rows := table.Find("tr")
	rows = rows.Slice(1, rows.Length())

	// EN: Iterate over each row and extract team data.
	// TR: Her bir satırı dolaşarak takım verilerini çıkartır.
	rows.Each(func(i int, row *goquery.Selection) {
		cells := row.Find("td")
		teamData := TeamData{
			// EN: Get and clean the team name from the first cell.
			// TR: İlk hücreden takım adını alır ve temizler.
			Name:         strings.TrimSpace(strings.Split(cells.Eq(0).Text(), ".")[1]),
			Played:       parseInt(cells.Eq(1).Text()),
			Wins:         parseInt(cells.Eq(2).Text()),
			Draws:        parseInt(cells.Eq(3).Text()),
			Losses:       parseInt(cells.Eq(4).Text()),
			GoalsFor:     parseInt(cells.Eq(5).Text()),
			GoalsAgainst: parseInt(cells.Eq(6).Text()),
			Average:      parseInt(cells.Eq(7).Text()),
			Points:       parseInt(cells.Eq(8).Text()),
		}
		teams = append(teams, teamData)
	})
	return teams, nil
}

// EN: parseInt converts a string to an integer; logs an error if the conversion fails.
// TR: Bir string'i tam sayıya dönüştürür; dönüşüm başarısız olursa hatayı loglar.
func parseInt(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		log.Printf("Error: %v. Failed to convert string to integer. Cell value: %s", err, s)
		return -1
	}
	return i
}

// EN: liveHandler retrieves live league data and returns it in JSON format.
// TR: Canlı lig verilerini alıp JSON formatında döner.
func liveHandler(w http.ResponseWriter, r *http.Request) {
	// EN: Fetch league data from the specified URL; return an internal server error if data retrieval fails.
	// TR: Belirtilen URL'den lig verilerini alır; veri çekimi başarısız olursa iç sunucu hatası döner.
	teams, err := leagueData("live", "https://www.tff.org/default.aspx?pageID=198")
	if err != nil {
		http.Error(w, "Internal Server Error, unable to fetch data", http.StatusInternalServerError) // TR: İç sunucu hatası, veri çekemiyorum
		return
	}

	// EN: Marshal the team data into a pretty-printed JSON format; return an internal server error if JSON marshaling fails.
	// TR: Takım verilerini güzel bir JSON formatında kodlar; JSON kodlaması başarısız olursa iç sunucu hatası döner.
	jsonStr, err := json.MarshalIndent(teams, "", "    ")
	if err != nil {
		http.Error(w, "Internal Server Error, unable to write data", http.StatusInternalServerError) // TR: İç sunucu hatası, yazdıramıyorum
		return
	}

	// EN: Set the content type to application/json with UTF-8 encoding.
	// TR: İçerik türünü application/json ve UTF-8 olarak ayarlar.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// EN: Write the JSON response back to the client.
	// TR: JSON yanıtını istemciye yazar.
	w.Write(jsonStr)
}

// EN: main function starts the HTTP server and sets up routing.
// TR: HTTP sunucusunu başlatıp yönlendirmeleri ayarlayan ana fonksiyon.
func main() {
	// EN: Create a new router using gorilla/mux.
	// TR: gorilla/mux kullanarak yeni bir yönlendirici oluşturur.
	r := mux.NewRouter()

	// EN: Define the route for live data retrieval.
	// TR: Canlı veri alımı için rotayı tanımlar.
	r.HandleFunc("/live", liveHandler)
	http.Handle("/", r)

	// EN: Log that the server is running on port 8000.
	// TR: Sunucunun 8000 portunda çalıştığını loglar.
	fmt.Println("Server is running on port 8000")

	// EN: Start the HTTP server. (http://localhost:8000/live)
	// TR: HTTP sunucusunu başlatır. (http://localhost:8000/live)
	http.ListenAndServe(":8000", nil)
}
