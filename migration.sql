-- 1. Tabel Users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('admin', 'nasabah', 'karyawan')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 2. Tabel Master Account (dibutuhkan oleh jurnal_umum)
CREATE TABLE master_account (
    id SERIAL PRIMARY KEY,
    nama_akun VARCHAR(100) NOT NULL,
    kode_akun VARCHAR(100) UNIQUE NOT NULL,
    tipe VARCHAR(10) NOT NULL CHECK (tipe IN ('Aktiva', 'Pasiva', 'Ekuitas', 'Revenue', 'Beban', 'Dividen'))
);

-- 3. Tabel Nasabah (dibutuhkan oleh tabungan)
CREATE TABLE nasabah (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    nama VARCHAR(100) NOT NULL,
    nik VARCHAR(20) UNIQUE NOT NULL,
    no_hp VARCHAR(15) UNIQUE NOT NULL,
    no_pin VARCHAR(15) NOT NULL,
    no_rekening VARCHAR(20) UNIQUE NOT NULL,
    saldo DECIMAL(15,2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 4. Tabel Karyawan (dibutuhkan oleh jurnal_umum)
CREATE TABLE karyawan (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    nama VARCHAR(100) NOT NULL,
    nip VARCHAR(20) UNIQUE NOT NULL,
    posisi VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 5. Tabel Produk Tabungan
CREATE TABLE produk_tabungan (
    id SERIAL PRIMARY KEY,
    nama_produk VARCHAR(100) NOT NULL,
    deskripsi TEXT,
    bunga DECIMAL(5,2) DEFAULT 0,
    biaya_administrasi DECIMAL(10,2) DEFAULT 0,
    min_saldo DECIMAL(15,2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 6. Tabel Tabungan (harus dibuat sebelum jurnal_umum)
CREATE TABLE tabungan (
    id SERIAL PRIMARY KEY,
    no_rekening VARCHAR(20) UNIQUE NOT NULL,  -- Tambahkan UNIQUE
    id_produk_tabungan INT REFERENCES produk_tabungan(id) ON DELETE CASCADE,
    nominal DECIMAL(15,2) NOT NULL CHECK (nominal >= 0),
    tipe CHAR(1) NOT NULL CHECK (tipe IN ('K', 'D')), -- K (Kredit), D (Debet)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- 7. Tabel Jurnal Umum (dibutuhkan oleh log_aktivitas)
-- Tabel Jurnal Umum (Pencatatan Akuntansi)
CREATE TABLE jurnal_umum (
    id SERIAL PRIMARY KEY,
    tanggal TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    no_rekening VARCHAR(20) REFERENCES tabungan(no_rekening) ON DELETE CASCADE,
    id_master_account INT NOT NULL REFERENCES master_account(id) ON DELETE CASCADE,
    tipe CHAR(1) NOT NULL CHECK (tipe IN ('D', 'K')),
    nominal DECIMAL(15,2) NOT NULL CHECK (nominal >= 0),
    keterangan TEXT,
    id_karyawan INT REFERENCES karyawan(id) ON DELETE SET NULL,
    input_by VARCHAR(50) NOT NULL
);


-- 8. Tabel Log Aktivitas (harus dibuat setelah jurnal_umum)
CREATE TABLE log_aktivitas (
    id SERIAL PRIMARY KEY,
    id_jurnal INT REFERENCES jurnal_umum(id) ON DELETE CASCADE,
    id_karyawan INT REFERENCES karyawan(id) ON DELETE SET NULL,
    input_by VARCHAR(50) NOT NULL,
    aksi VARCHAR(10) NOT NULL CHECK (aksi IN ('INSERT', 'UPDATE', 'DELETE')),
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    keterangan TEXT
);
