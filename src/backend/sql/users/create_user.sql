INSERT INTO
    users (
        email,
        phone,
        password
    )
VALUES ($1, $2, $3) RETURNING id