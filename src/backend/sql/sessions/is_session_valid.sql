SELECT EXISTS (
        SELECT 1
        FROM sessions
        WHERE
            user_id = $1
            AND id = $2
            AND revoked = false
    );