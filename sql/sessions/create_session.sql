INSERT INTO
    sessions (user_id, device, ip_address)
VALUES ($1, $2, $3) RETURNING id;