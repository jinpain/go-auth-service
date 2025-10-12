UPDATE sessions
SET
    revoked = true,
    revoked_at = NOW()
WHERE
    user_id = $1
    AND id = $2