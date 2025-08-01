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


INSERT INTO `spaceships` (`name`, `class`, `crew`, `image`, `value`, `status`, `created_at`, `updated_at`)
VALUES
	('sp1', 'xd', 0, 'url', 3.00, 'operational', '2025-07-28 16:40:34', '2025-07-30 22:12:27'),
	('sp2', 'xd', 0, 'url', 43.00, 'operational', '2025-07-28 16:40:34', '2025-07-30 22:12:31');

INSERT INTO `armaments` (`spaceship_id`, `title`, `qty`)
VALUES
	(1, 'arm title', 2),
	(1, 'arm2 ', 1),
	(2, 'xdf', 2),
	(2, 'arm1', 1),
	(2, 'sss', 0);

