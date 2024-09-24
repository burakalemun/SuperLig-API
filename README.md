# SuperLig-API


# Football League Data Fetcher

This project is a Go application that scrapes live football league data from a specific website (Turkish Football Federation) and serves it in JSON format via a web API. Libraries such as `gorilla/mux` for routing and `goquery` for web scraping have been used.

---

## Features

- Fetches football league data (team name, matches played, wins, draws, losses, goals, and points).
- Serves the data in JSON format via an HTTP API.
- Decodes website content encoded in Windows-1254 character set.

---

## Libraries Used

- **Gorilla Mux**: HTTP router for defining routes and handling HTTP requests.
- **GoQuery**: For parsing and querying the HTML content from the target website.
- **Encoding**: For handling character set conversions (e.g., Windows-1254).
- **JSON**: For formatting and serving the data in a structured JSON format.

---

## Installation

1. Clone this repository:
    ```bash
    git clone https://github.com/burakalemun/SuperLig-API.git
    cd SuperLig-API
    ```

2. Download and install the required dependencies:
    ```bash
    go mod download
    ```

---

## Running the Application

1. Start the HTTP server:
    ```bash
    go run main.go
    ```

2. The application will be running at `http://localhost:8000/live`. You can access the live football data by making a GET request to this endpoint.

Example:
```bash
curl http://localhost:8000/live
```

---

# Futbol Ligi Veri Çekici

Bu proje, belirli bir web sitesinden (Türkiye Futbol Federasyonu) canlı futbol ligi verilerini kazıyan ve bunları bir web API'si aracılığıyla JSON formatında sunan bir Go uygulamasıdır. Yönlendirme için `gorilla/mux` ve web kazıma için `goquery` gibi kütüphaneler Kullanılmıştır.

---

## Özellikler

- Futbol ligi verilerini çeker (takım adı, oynanan maçlar, galibiyetler, beraberlikler, mağlubiyetler, goller ve puanlar).
- Verileri bir HTTP API aracılığıyla JSON formatında sunar.
- Windows-1254 karakter setiyle kodlanmış web sitesi içeriğini çözümler.

---

## Kullanılan Kütüphaneler

- **Gorilla Mux**: HTTP isteklerini yönlendirmek ve rotaları tanımlamak için kullanılır.
- **GoQuery**: Hedef web sitesinden HTML içeriğini ayrıştırmak ve sorgulamak için kullanılır.
- **Encoding**: Karakter seti dönüştürmelerini (örneğin, Windows-1254) işlemek için kullanılır.
- **JSON**: Verileri yapılandırılmış bir JSON formatında biçimlendirmek ve sunmak için kullanılır.

---

## Kurulum

1. Bu depoyu klonlayın:
    ```bash
    git clone https://github.com/burakalemun/SuperLig-API.git
    cd SuperLig-API
    ```

2. Gerekli bağımlılıkları indirin ve yükleyin:
    ```bash
    go mod download
    ```

---

## Uygulamayı Çalıştırma

1. HTTP sunucusunu başlatın:
    ```bash
    go run main.go
    ```

2. Uygulama `http://localhost:8000/live` adresinde çalışıyor olacak. Bu uç noktaya bir GET isteği yaparak canlı futbol verilerine erişebilirsiniz.

Örnek:
```bash
curl http://localhost:8000/live
```

---

## JSON Response Format / JSON Yanıt Formatı

[
    {
        "name": "Takım A",
        "played": 34,
        "wins": 20,
        "draws": 10,
        "losses": 4,
        "goals_for": 60,
        "goals_against": 30,
        "average": 30,
        "points": 70
    },
    ...
]

---

## License / Lisans

This project is licensed under the MIT License - see the [LICENSE](https://github.com/burakalemun/SuperLig-API/blob/main/LICENSE) file for details. </br>
Bu proje MIT Lisansı altında lisanslanmıştır - ayrıntılar için [LICENSE](https://github.com/burakalemun/SuperLig-API/blob/main/LICENSE) dosyasına bakın.



