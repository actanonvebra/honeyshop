# HoneyShop

HoneyShop, sahte bir e-ticaret sistemi oluşturarak saldırganların davranışlarını analiz eden bir **honeypot** uygulamasıdır. Bu proje, e-ticaret temalı API endpoint'leri üzerinden gelen saldırıları loglayarak **SQL injection**, **brute force** gibi yaygın saldırı türlerini tespit etmeyi ve analiz etmeyi amaçlar.

## Projenin Amacı
- Saldırgan davranışlarını gözlemlemek.
- Olası saldırı türlerini (ör. SQL injection, brute force) analiz etmek.
- Sahte API endpoint'leri ile saldırganları gerçek sistemlerden uzaklaştırmak.
- Saldırılardan elde edilen verilerle güvenlik iyileştirmeleri yapmak.

## Özellikler
- Sahte e-ticaret API'leri:
  - Kullanıcı Kayıt ve Giriş
  - Ürün Listeleme
  - Sepet İşlemleri
  - Ödeme
  - Admin Paneli
- Gelen istekleri JSON/SQL loglama.
- Saldırı tespiti:
  - **Brute Force**: Tekrarlayan giriş denemeleri.
  - **SQL Injection**: Zararlı giriş stringlerini tespit etme.
- Tamamen mock/dummy veri kullanımı.