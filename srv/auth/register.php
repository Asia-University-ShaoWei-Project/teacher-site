<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Document</title>
</head>

<body>
</body>

</html>
<?php
require_once('db_conn.php');
$db = NewDB("./");
$id = 'rikki';
$pwd = 'password';
// $hash_pwd = register($db, $id, $pwd);
$secure_key = getenv('SECURE_KEY');
$hash_salt_password = sha1($pwd) . $secure_key;
// $bcrypt = bcryptPassword($hash_salt_password);

function register(DB &$db, string $username, $password)
{
  $secure_key = $_ENV['SECURE_KEY'];
  $hash_salt_password = sha1($password) . $secure_key;
  $bcrypt_password = bcryptPassword($hash_salt_password);
  $token = newToken($username);
  $stmt = $db->prepare('INSERT INTO auths(token,username, password) VALUES(:token, :username, :password)');
  $stmt->bindParam(':token', $token, SQLITE3_TEXT);
  $stmt->bindParam(':username', $username, SQLITE3_TEXT);
  $stmt->bindParam(':password', $bcrypt_password, SQLITE3_TEXT);
  $stmt->execute();
}
function bcryptPassword(string $password): string
{
  $cost = 8;
  $hash = password_hash($password, PASSWORD_BCRYPT, ["cost" => $cost]);
  return $hash;
}
?>