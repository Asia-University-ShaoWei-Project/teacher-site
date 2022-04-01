<?php
class DB extends SQLite3
{
  function __construct(string $db_file_path)
  {
    $this->open($db_file_path);
  }
}
function NewDB(string $path): DB
{
  $db_file_path = $path . getenv('DB_FILE_NAME') . ".db";
  return new DB($db_file_path);
}
function Br()
{
  echo "<br />";
}