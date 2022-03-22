DROP TABLE IF EXISTS auths;
DROP TABLE IF EXISTS tokens;
DROP TABLE IF EXISTS teachers;

-- INSERT INTO token_auth(token, auth_id)VALUES('user623550b7f05c4', 'rikki');

-- UPDATE auths
-- SET id = 'rikki'
-- WHERE id IN(
--   SELECT auth_id
--   FROM tokens
--   WHERE token='user623550b7f05c4'
-- );

-- UPDATE tokens
-- SET auth_id = 'rikki'
-- WHERE token='user623550b7f05c4';

--  when update auth with id
  -- SET auth_id=:new_token
  -- WHERE token=:token'
  -- $stmt_process = array(
  --   "auth" => 'UPDATE auths
  --   SET id=:new_token
  --   WHERE id IN (
  --     SELECT auth_id
  --     FROM tokens
  --     WHERE token=:token)',
  --   "token" => 'UPDATE tokens
  --   SET auth_id=:new_token
  --   WHERE token=:token'
  -- );