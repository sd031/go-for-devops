export MYSQL_DB_URI="username:password@tcp(db_url:3306)/db_name"

mysql -h DB_URL -u root -p

CREATE DATABASE demodb;

USE demodb;


CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL
);


Install redis:
Linux:
sudo snap install redis

or:
sudo apt install lsb-release curl gpg
curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg

echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list

sudo apt-get update
sudo apt-get install redis

Mac:
brew install redis
brew services start redis
if want to stop:
brew services stop redis
check status:
brew services info redis


Connect to Redis:
redis-cli
Basic opertions:
SET key value: Sets the string value of a key.
GET key: Gets the value of a key.
DEL key: Deletes a key.
EXISTS key: Checks if a key exists.
KEYS pattern: Finds all keys matching the given pattern.
