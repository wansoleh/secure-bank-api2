-- Insert ke tabel Jurnal Umum (Simulasi transaksi setor)
WITH akun AS (
    SELECT id, kode_akun FROM master_account
)
INSERT INTO jurnal_umum (no_rekening, id_master_account, tipe, nominal, keterangan, input_by)
VALUES 
('100001', (SELECT id FROM akun WHERE kode_akun = '101'), 'D', 500000, 'Setoran Awal', 'system'),
('100001', (SELECT id FROM akun WHERE kode_akun = '201'), 'K', 500000, 'Setoran Awal', 'system'),
('100002', (SELECT id FROM akun WHERE kode_akun = '101'), 'D', 750000, 'Setoran Awal', 'system'),
('100002', (SELECT id FROM akun WHERE kode_akun = '201'), 'K', 750000, 'Setoran Awal', 'system');



-- Insert ke tabel Users
INSERT INTO users (username, password, role, created_at) VALUES
('admin1', '$2y$12$EXAMPLEHASHEDPASSWORD1', 'admin', NOW()), -- Password harus di-hash
('nasabah1', '$2y$12$EXAMPLEHASHEDPASSWORD2', 'nasabah', NOW()),
('nasabah2', '$2y$12$EXAMPLEHASHEDPASSWORD3', 'nasabah', NOW()),
('karyawan1', '$2y$12$EXAMPLEHASHEDPASSWORD4', 'karyawan', NOW());

-- Insert ke tabel Nasabah
INSERT INTO nasabah (user_id, nama, nik, no_hp, no_rekening, no_pin, saldo, created_at) VALUES
(2, 'Ahmad Susanto', '1234567890123456', '081234567890', '100001', '$2y$12$EXAMPLEHASHEDPIN1', 500000, NOW()),
(3, 'Rina Kartika', '6543210987654321', '081298765432', '100002', '$2y$12$EXAMPLEHASHEDPIN2', 750000, NOW());

-- Insert ke tabel Karyawan
INSERT INTO karyawan (user_id, nama, nip, posisi, created_at) VALUES
(4, 'Budi Santoso', 'KRY001', 'Teller', NOW()),
(4, 'Siti Aminah', 'KRY002', 'Customer Service', NOW());

-- Insert ke tabel Produk Tabungan
INSERT INTO produk_tabungan (nama_produk, deskripsi, bunga, biaya_administrasi, min_saldo, created_at) VALUES
('Tabungan Reguler', 'Tabungan dengan bunga standar', 1.5, 5000, 50000, NOW()),
('Tabungan Premium', 'Tabungan dengan bunga tinggi', 2.5, 10000, 100000, NOW());

-- Insert ke tabel Tabungan (Menggunakan produk tabungan)
INSERT INTO tabungan (no_rekening, id_produk_tabungan, nominal, tipe, created_at) VALUES
('100001', 1, 500000, 'D', NOW()),
('100002', 2, 750000, 'D', NOW());

-- Insert ke tabel Master Account (Chart of Accounts)
INSERT INTO master_account (kode_akun, nama_akun, tipe) VALUES 
('101', 'Kas', 'Aktiva'),
('102', 'Bank', 'Aktiva'),
('201', 'Piutang Usaha', 'Aktiva'),
('301', 'Hutang Usaha', 'Pasiva'),
('401', 'Modal Pemilik', 'Ekuitas'),
('501', 'Pendapatan Jasa', 'Revenue'),
('601', 'Beban Operasional', 'Beban'),
('701', 'Dividen Dibayarkan', 'Dividen');

-- Insert ke tabel Jurnal Umum (Simulasi transaksi setor)
WITH akun AS (
    SELECT id, kode_akun FROM master_account
)
INSERT INTO jurnal_umum (no_rekening, id_master_account, tipe, nominal, keterangan, input_by, tanggal)
VALUES 
('100001', (SELECT id FROM akun WHERE kode_akun = '101'), 'D', 500000, 'Setoran Awal', 'system', NOW()),
('100001', (SELECT id FROM akun WHERE kode_akun = '201'), 'K', 500000, 'Setoran Awal', 'system', NOW()),
('100002', (SELECT id FROM akun WHERE kode_akun = '101'), 'D', 750000, 'Setoran Awal', 'system', NOW()),
('100002', (SELECT id FROM akun WHERE kode_akun = '201'), 'K', 750000, 'Setoran Awal', 'system', NOW());

-- Insert ke tabel Log Aktivitas (Audit Trail)
INSERT INTO log_aktivitas (id_jurnal, id_karyawan, input_by, aksi, keterangan, timestamp) VALUES
(1, 1, 'system', 'INSERT', 'Setoran awal nasabah 100001', NOW()),
(2, 1, 'system', 'INSERT', 'Setoran awal nasabah 100002', NOW());
