-- Buat tabel Users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('admin', 'nasabah', 'karyawan')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Buat tabel Nasabah
CREATE TABLE nasabah (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    nama VARCHAR(100) NOT NULL,
    nik VARCHAR(20) UNIQUE NOT NULL,
    no_hp VARCHAR(15) UNIQUE NOT NULL,
    no_rekening VARCHAR(20) UNIQUE NOT NULL,
    saldo DECIMAL(15,2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Buat tabel Karyawan
CREATE TABLE karyawan (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    nama VARCHAR(100) NOT NULL,
    nip VARCHAR(20) UNIQUE NOT NULL,
    posisi VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Buat tabel Produk Tabungan
CREATE TABLE produk_tabungan (
    id SERIAL PRIMARY KEY,
    nama_produk VARCHAR(100) NOT NULL,
    deskripsi TEXT,
    bunga DECIMAL(5,2) DEFAULT 0,
    biaya_administrasi DECIMAL(10,2) DEFAULT 0,
    min_saldo DECIMAL(15,2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Buat tabel Tabungan
CREATE TABLE tabungan (
    id SERIAL PRIMARY KEY,
    no_rekening VARCHAR(20) REFERENCES nasabah(no_rekening) ON DELETE CASCADE,
    id_produk_tabungan INT REFERENCES produk_tabungan(id) ON DELETE CASCADE,
    nominal DECIMAL(15,2) NOT NULL CHECK (nominal >= 0),
    tipe CHAR(1) NOT NULL CHECK (tipe IN ('K', 'D')), -- K (Kredit), D (Debet)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Jurnal Umum (Pencatatan Akuntansi)
CREATE TABLE jurnal_umum (
    id SERIAL PRIMARY KEY,
    tanggal TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    no_rekening VARCHAR(20) REFERENCES tabungan(no_rekening) ON DELETE CASCADE,
    kode_akun VARCHAR(20) NOT NULL REFERENCES master_account(kode_akun) ON DELETE CASCADE,
    tipe CHAR(1) NOT NULL CHECK (tipe IN ('D', 'K')), -- D (Debit), K (Kredit)
    nominal DECIMAL(15,2) NOT NULL CHECK (nominal >= 0),
    keterangan TEXT,
    id_karyawan INT REFERENCES karyawan(id) ON DELETE SET NULL,
    input_by VARCHAR(50) NOT NULL
);


-- Buat tabel Log Aktivitas
CREATE TABLE log_aktivitas (
    id SERIAL PRIMARY KEY,
    id_jurnal INT REFERENCES jurnal_umum(id) ON DELETE CASCADE,
    id_karyawan INT REFERENCES karyawan(id) ON DELETE SET NULL,
    input_by VARCHAR(50) NOT NULL,
    aksi VARCHAR(10) NOT NULL CHECK (aksi IN ('INSERT', 'UPDATE', 'DELETE')),
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    keterangan TEXT
);

-- Tabel Master Account (Chart of Accounts)
CREATE TABLE master_account (
    id SERIAL PRIMARY KEY,
    nama_akun VARCHAR(100) NOT NULL,
    kode_akun VARCHAR(100) NOT NULL,
    tipe VARCHAR(10) NOT NULL CHECK (tipe IN ('Aktiva', 'Pasiva', 'Ekuitas', 'Revenue', 'Beban', 'Dividen'))
);
