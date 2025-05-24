<?php
header("Access-Control-Allow-Origin: *");
header("Access-Control-Allow-Methods: GET, POST,DELETE,PUT");
header("Access-Control-Allow-Headers: Content-Type");
header('Content-Type: application/json');
header('Content-Type: application/json');
$servername = "localhost";
$username = "admin";
$password = "admin";
$dbname = "sensores";

$conn = new mysqli($servername, $username, $password, $dbname);

// Obtener el último registro
$result = $conn->query("SELECT distancia, movimiento FROM registros ORDER BY fecha DESC LIMIT 1");
$data = $result->fetch_assoc();

echo json_encode([
    'distancia' => $data['distancia'],
    'movimiento' => (bool)$data['movimiento']
]);

$conn->close();
?>