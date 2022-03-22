<?php session_start();
require_once('db_conn.php');

header('refresh: 1; url = http://localhost/index.php');
$db = NewDB("./");
$token = $_SESSION['token'];
if (empty($token)) return;
$new_token = newToken();
updateTokenID($db, $token, $new_token);
unset($_SESSION['token']);
$db->close();
function newToken(): string
{
  $token = uniqid();
  return $token;
}
function updateTokenID(DB &$db, string $token, $new_token)
{
  $stmt = $db->prepare('UPDATE tokens SET token=:new_token WHERE token=:token');
  $stmt->bindParam(':new_token', $new_token, SQLITE3_TEXT);
  $stmt->bindParam(':token', $token, SQLITE3_TEXT);
  $stmt->execute();
}