# HoneyShop

HoneyShop, sahte bir e-ticaret sistemi oluşturarak saldırganların davranışlarını analiz eden bir honeypot uygulamasıdır. Bu proje, e-ticaret API endpoint'leri üzerinden gelen saldırıları loglar ve SQL injection, brute force gibi yaygın saldırı türlerini tespit etmeyi amaçlar.

HoneyShop, özellikle **siber güvenlik araştırmacıları**, **etik hackerlar**, ve **güvenlik ekipleri** için tasarlanmıştır. Gerçek sistemleri koruma altına alarak, saldırganların dikkatini sahte bir hedefe çekmek için ideal bir araçtır.

---

## **Projenin Amacı**
- Gerçek sistemlere yöneltilen saldırıları engellemek için sahte hedefler sunmak.
- Saldırgan davranışlarını analiz ederek daha iyi güvenlik önlemleri geliştirmek.
- Yaygın saldırı türlerini (ör. SQL injection, brute force) kolayca tespit ve analiz etmek.
- Sahte API endpoint'leri üzerinden saldırganları analiz edip, sistem güvenliğini geliştirmek.

---

## **Özellikler**
- **Sahte E-ticaret API'leri:**
  - Kullanıcı Kayıt ve Giriş
  - Ürün Listeleme
  - Sepet İşlemleri
  - Ödeme
  - Admin Paneli
- **Saldırı Tespiti:**
  - **Brute Force:** Tekrarlayan giriş denemelerini algılama.
  - **SQL Injection:** Zararlı giriş stringlerini algılama ve loglama.
- **Loglama ve İzleme:**
  - Gelen isteklerin JSON formatında saklanması.
  - Sistem davranışlarının detaylı analiz edilmesi.
- **Tamamen Mock/Dummy Veri Kullanımı:** Gerçek sistemleri riske atmadan saldırıların taklit edilmesi.

---

## **Kurulum ve Çalıştırma**
HoneyShop’u çalıştırmak için aşağıdaki adımları izleyin:

1. Depoyu klonlayın:
   ```bash
   git clone https://github.com/actanonvebra/honeyshop.git
   cd honeyshop
   ```

2. Gerekli bağımlılıkları yükleyin:
   ```bash
   go mod tidy
   ```

3. MongoDB bağlantısını yapılandırmak için `.env` dosyasını düzenleyin:
   ```env
   MONGO_URI=mongodb://localhost:27017
   ```

4. Uygulamayı başlatın:
   ```bash
   go run cmd/main.go
   ```

5. API endpoint’lerini test etmek için CURL veya Postman kullanabilirsiniz.

---

## **Örnek Kullanım**

### **1. Kullanıcı Girişi**
- **Endpoint:** `/login`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "username": "admin",
    "password": "1234"
  }
  ```

- **CURL Komutu:**
  ```bash
  curl -X POST -H "Content-Type: application/json" -d '{"username":"admin","password":"1234"}' http://localhost:8080/login
  ```

### **2. Kullanıcı Kaydı**
- **Endpoint:** `/register`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "username": "newuser",
    "password": "password123",
    "email": "newuser@example.com"
  }
  ```

---

## **Lisans ve Katkı**

### **Lisans**
Bu proje [MIT Lisansı](LICENSE) ile lisanslanmıştır.

### **Katkı**
Eğer projeye katkıda bulunmak istiyorsanız:
1. Depoyu fork edin.
2. Geliştirmeler yapın.
3. Pull Request gönderin.

