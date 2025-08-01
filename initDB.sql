CREATE DATABASE IF NOT EXISTS imperial_fleet;
USE imperial_fleet;

CREATE TABLE IF NOT EXISTS spaceships (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    class VARCHAR(255) NOT NULL,
    crew INT NOT NULL,
    image VARCHAR(500),
    value DECIMAL(10,2) NOT NULL,
    status ENUM('operational', 'damaged', 'destroyed', 'maintenance') NOT NULL DEFAULT 'operational',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS armaments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    spaceship_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    qty INT NOT NULL,
    FOREIGN KEY (spaceship_id) REFERENCES spaceships(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS imperial_officers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    rank VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert sample data
INSERT IGNORE INTO spaceships (id, name, class, crew, image, value, status) VALUES
(1, 'Devastator', 'Star Destroyer', 35000, 'https://url.to.image', 1999.99, 'operational'),
(2, 'Red Five', 'X-Wing Fighter', 1, 'https://url.to.xwing', 149.99, 'damaged');

INSERT IGNORE INTO armaments (spaceship_id, title, qty) VALUES
(1, 'Turbo Laser', 60),
(1, 'Ion Cannons', 60),
(1, 'Tractor Beam', 10),
(2, 'Laser Cannons', 4),
(2, 'Proton Torpedoes', 6);

-- Insert sample imperial officer (password: "empire123")
INSERT IGNORE INTO imperial_officers (username, password_hash, rank) VALUES
('r3d3', '$2a$10$rZ3Q8ZKvyQn9h5YqT6mJbeXgJH8fJ2m3rK9Q2vC8xH7wJ4pL5nM6a', 'General');
