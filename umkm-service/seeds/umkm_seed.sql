START TRANSACTION;
SET FOREIGN_KEY_CHECKS = 0;
TRUNCATE TABLE umkm_ratings;
TRUNCATE TABLE umkm_work_hours;
TRUNCATE TABLE umkm_gallerys;
TRUNCATE TABLE umkm_locations;
TRUNCATE TABLE umkms;
TRUNCATE TABLE umkm_owners;
SET FOREIGN_KEY_CHECKS = 1;

INSERT INTO umkm_owners (id, name, image, phone, email, website, facebook, twitter, instagram, whatsapp) VALUES
(1, 'Budi Santoso', 'https://cdn.example.com/owners/budi.jpg', '+62 812-1111-2222', 'budi@example.com', 'https://budishop.id', NULL, NULL, 'https://instagram.com/budi_shop', 'https://wa.me/6281211112222'),
(2, 'Sari Dewi', 'https://cdn.example.com/owners/sari.jpg', '+62 813-2222-3333', 'sari@example.com', NULL, 'https://facebook.com/sari.dewi', NULL, 'https://instagram.com/sari_craft', NULL),
(3, 'Andi Wijaya', 'https://cdn.example.com/owners/andi.jpg', '+62 814-3333-4444', 'andi@example.com', 'https://andikopi.com', NULL, 'https://twitter.com/andi_kopi', NULL, 'https://wa.me/6281433334444'),
(4, 'Rina Putri', 'https://cdn.example.com/owners/rina.jpg', '+62 815-4444-5555', 'rina@example.com', NULL, NULL, NULL, 'https://instagram.com/rina_batik', NULL);

INSERT INTO umkms (id, name, about, description, icon, slug, type, owner_id) VALUES
(1, 'Warung Kopi Andalas', 'Warung kopi dengan biji pilihan.', 'Tempat nongkrong nyaman dengan berbagai varian kopi lokal dan snack.', 'https://cdn.example.com/icons/kopi.png', 'warung-kopi-andalas', 'kuliner', 3),
(2, 'Batik Rina', 'UMKM batik tulis dan cap.', 'Koleksi batik modern dan klasik, menerima pesanan motif khusus.', 'https://cdn.example.com/icons/batik.png', 'batik-rina', 'fashion', 4),
(3, 'Kerajinan Sari', 'Kerajinan tangan ramah lingkungan.', 'Produk anyaman dan dekorasi rumah dari bahan daur ulang.', 'https://cdn.example.com/icons/craft.png', 'kerajinan-sari', 'kerajinan', 2),
(4, 'Toko Kue Budi', 'Kue rumahan segar setiap hari.', 'Aneka kue basah dan kering, menerima pesanan ulang tahun.', 'https://cdn.example.com/icons/kue.png', 'toko-kue-budi', 'kuliner', 1);

INSERT INTO umkm_gallerys (umkm_id, url) VALUES
(1, 'https://cdn.example.com/umkm/1/1.jpg'),
(1, 'https://cdn.example.com/umkm/1/2.jpg'),
(1, 'https://cdn.example.com/umkm/1/3.jpg'),
(2, 'https://cdn.example.com/umkm/2/1.jpg'),
(2, 'https://cdn.example.com/umkm/2/2.jpg'),
(3, 'https://cdn.example.com/umkm/3/1.jpg'),
(3, 'https://cdn.example.com/umkm/3/2.jpg'),
(4, 'https://cdn.example.com/umkm/4/1.jpg');

INSERT INTO umkm_locations (umkm_id, url, text, short_text, longitude, latitude) VALUES
(1, 'https://maps.google.com/?q=-6.200000,106.816666', 'Jl. Melati No. 10, Jakarta', 'Jakarta', 106.816666, -6.200000),
(2, 'https://maps.google.com/?q=-7.250445,112.768845', 'Jl. Kenanga No. 5, Surabaya', 'Surabaya', 112.768845, -7.250445),
(3, 'https://maps.google.com/?q=-6.914744,107.609810', 'Jl. Anggrek No. 7, Bandung', 'Bandung', 107.609810, -6.914744),
(4, 'https://maps.google.com/?q=-7.801389,110.364722', 'Jl. Mawar No. 12, Yogyakarta', 'Yogyakarta', 110.364722, -7.801389);

INSERT INTO umkm_work_hours (umkm_id, day, hours) VALUES
(1, 'monday', '08.00 - 22.00'),
(1, 'tuesday', '08.00 - 22.00'),
(1, 'wednesday', '08.00 - 22.00'),
(1, 'thursday', '08.00 - 22.00'),
(1, 'friday', '08.00 - 23.00'),
(1, 'saturday', '09.00 - 23.00'),
(1, 'sunday', '09.00 - 21.00'),
(2, 'monday', '09.00 - 17.00'),
(2, 'tuesday', '09.00 - 17.00'),
(2, 'wednesday', '09.00 - 17.00'),
(2, 'thursday', '09.00 - 17.00'),
(2, 'friday', '09.00 - 17.00'),
(2, 'saturday', '10.00 - 16.00'),
(2, 'sunday', 'Tutup'),
(3, 'monday', '08.00 - 18.00'),
(3, 'tuesday', '08.00 - 18.00'),
(3, 'wednesday', '08.00 - 18.00'),
(3, 'thursday', '08.00 - 18.00'),
(3, 'friday', '08.00 - 18.00'),
(3, 'saturday', '09.00 - 15.00'),
(3, 'sunday', 'Tutup'),
(4, 'monday', '07.00 - 19.00'),
(4, 'tuesday', '07.00 - 19.00'),
(4, 'wednesday', '07.00 - 19.00'),
(4, 'thursday', '07.00 - 19.00'),
(4, 'friday', '07.00 - 19.00'),
(4, 'saturday', '08.00 - 14.00'),
(4, 'sunday', 'Tutup');

INSERT INTO umkm_ratings (umkm_id, value) VALUES
(1, 5), (1, 4), (1, 5), (1, 3), (1, 4),
(2, 4), (2, 4), (2, 4),
(3, 5), (3, 5), (3, 5), (3, 5), (3, 5),
(4, 2), (4, 3);

COMMIT;