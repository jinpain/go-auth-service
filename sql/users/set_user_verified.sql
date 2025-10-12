UPDATE users
SET verified = true
WHERE id = $1;
