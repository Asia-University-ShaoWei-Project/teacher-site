<?php
require_once('db_conn.php');
require_once('../pages/route.php');
function bcryptPassword(string $password): string
{
  $cost = 8;
  echo "bcrypt cost: " . $cost;
  Br();
  $hash = password_hash($password, PASSWORD_BCRYPT, ["cost" => $cost]);
  echo "password bcrypted: " . $hash;
  Br();
  return $hash;
}
function register(DB &$db, string $id, $password)
{
  $secure_key = getenv('SECURE_KEY');
  $hash_salt_password = sha1($password) . $secure_key;
  $bcrypt_password = bcryptPassword($hash_salt_password);
  $token = newToken($id);
  $stmt = $db->prepare('INSERT INTO auths(token,username, password) VALUES(:token, :username, :password)');
  $stmt->bindParam(':token', $token, SQLITE3_TEXT);
  $stmt->bindParam(':username', $username, SQLITE3_TEXT);
  $stmt->bindParam(':password', $bcrypt_password, SQLITE3_TEXT);
  $stmt->execute();
}
function verify(DB &$db, string $id, $password)
{
  $secure_key = getenv('SECURE_KEY');
  $hash_salt_password = sha1($password) . $secure_key;
  echo "salt password: " . $hash_salt_password;
  Br();
  $stmt = $db->prepare('SELECT * FROM auths WHERE id=:id LIMIT 1');

  $stmt->bindValue(':id', $id, SQLITE3_TEXT);
  $account = $stmt->execute();
  $rows = $account->fetchArray();
  if ($rows) {
    $db_password = $rows['password'];

    echo "DB password: " . $db_password;
    Br();
    if (password_verify($hash_salt_password, $db_password)) {
      echo "password confirmed!";
      Br();
    }
  }
  echo "row is empty!";
  Br();
}
?>
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Document</title>
</head>

<body>
  <?php
  echo "Token:" . $_SESSION['token'];
  $db = NewDB("./");
  $id = $_POST['id'];
  $password = $_POST['password'];
  // $id = 'rikki';
  // $pwd = 'password';
  verify($db, $id, $password);
  $db->close();
  ?>


  <form action=<?php echo $Route['login_test'] ?> method="POST">
    <input type="input" name="id" value="rikki">
    <input type="input" name="password" value="password">
    <button type="submit">Login</button>
  </form>
</body>

</html>