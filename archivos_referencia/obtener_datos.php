<?php
header("Access-Control-Allow-Origin: *");
header("Access-Control-Allow-Methods: GET, POST,DELETE, PUT");
header("Access-Control-Allow-Headers: Content-Type");
header('Content-Type: application/json');
header('Content-Type: application/json');

$conexion = new mysqli("localhost", "root", "", "ibowl_db"); // Cambia "iot_db" por el nombre real

if ($conexion->connect_error) {
  die("Error de conexión: " . $conexion->connect_error);
}

$sql = "SELECT * FROM `alimentacion`";
$resultado = $conexion->query($sql);

$datos = [];

while ($fila = $resultado->fetch_assoc()) {
  $datos[] = $fila;
}

header('Content-Type: application/json');
echo json_encode(array_reverse($datos)); // Mostramos de más antiguo a más nuevo
?>