-- Tabel master barang
CREATE TABLE todo.items (
    id SERIAL PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(100),
    stock INT NOT NULL DEFAULT 0,
    created_by VARCHAR(100),
    updated_by VARCHAR(100),
    deleted_by VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Tabel user (untuk login, audit log, dsb)
CREATE TABLE todo.users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL, -- contoh: superadmin, admin, staff
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel event barang masuk/keluar
CREATE TABLE todo.stock_events (
    id SERIAL PRIMARY KEY,
    item_code VARCHAR(50) NOT NULL REFERENCES todo.items(code),
    event_type VARCHAR(20) NOT NULL, -- 'IN' atau 'OUT'
    quantity INT NOT NULL,
    event_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    responsible VARCHAR(100) NOT NULL, -- penanggung jawab
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel audit log
CREATE TABLE todo.audit_log (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES todo.users(id),
    action VARCHAR(255) NOT NULL,
    details TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabel event kalender gudang
CREATE TABLE todo.event_calendar (
    id SERIAL PRIMARY KEY,
    event_title VARCHAR(255) NOT NULL,
    event_description TEXT,
    event_start TIMESTAMP NOT NULL,
    event_end TIMESTAMP NOT NULL,
    external_id VARCHAR(255), -- untuk sinkronisasi dengan API eksternal (Google/TeamUp)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
