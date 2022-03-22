<?php
session_start();
require_once('../pages/route.php');
header('refresh: 2; url = http://' . $Route['home']);
?>

</html>
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="stylesheet" href="static/css/loading.css">
  <title>Login</title>
</head>

<body id="loading_body">
  <div class="loader"></div>
</body>

</html>
<?php
require_once('db_conn.php');
$id = $_POST['id'];
$password = $_POST['password'];
// $id = 'rikki';
// $password = 'password';
Login($id, $password);

function Login(string $id, $password)
{
  if (empty($id) || empty($password)) return;
  $db = NewDB("./");
  $valid = verifyAuth($db, $id, $password);
  if ($valid) {
    $token = getToken($db, $id);
    $_SESSION['token'] = $token;
  }
  $db->close();
}
function verifyAuth(DB &$db, string $id, $password): bool
{
  $secure_key = getenv('SECURE_KEY');
  $hash_salt_password = sha1($password) . $secure_key;

  $stmt = $db->prepare('SELECT * FROM auths WHERE id=:id LIMIT 1');
  $stmt->bindValue(':id', $id, SQLITE3_TEXT);
  $account = $stmt->execute();
  $rows = $account->fetchArray();
  if ($rows) {
    $db_password = $rows['password'];
    if (password_verify($hash_salt_password, $db_password)) {
      return true;
    }
  }
  return false;
}
function getToken(DB &$db, string $id): string
{
  $token = '';
  $stmt = $db->prepare('SELECT token FROM tokens WHERE auth_id=:auth_id LIMIT 1');
  $stmt->bindValue(':auth_id', $id, SQLITE3_TEXT);
  $account = $stmt->execute();
  $row = $account->fetchArray();
  if ($row) {
    $token = $row['token'];
  }
  return $token;
}
function bcryptPassword(string $password): string
{
  $cost = 8;
  $hash = password_hash($password, PASSWORD_BCRYPT, ["cost" => $cost]);
  return $hash;
}
?>