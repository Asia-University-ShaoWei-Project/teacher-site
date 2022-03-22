<?php

$link = mysqli_connect(
  "localhost",
  "user",
  "pwd",
  "name"
);

if (!$db_connection) {
  echo "(1)Error : Unable to connect to MySQL : " . PHP_EOL;
  echo "(2)Debugging errno : " . mysqli_connect_errno() . PHP_EOL;
  echo "(3)Debugging error : " . mysqli_connect_error() . PHP_EOL;
  exit;
}
$db_connection->set_charset("utf8");
