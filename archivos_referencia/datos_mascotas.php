<?php
header("Access-Control-Allow-Origin: *");
header("Access-Control-Allow-Methods: GET, POST, DELETE, PUT");
header("Access-Control-Allow-Headers: Content-Type");
header('Content-Type: application/json');

$servername = "localhost";
$username = "admin";
$password = "admin";
$dbname = "ibowl_db";

try {
    // Crear conexión
    $conn = new PDO("mysql:host=$servername;dbname=$dbname", $username, $password);
    $conn->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
    
    // Determinar el método HTTP
    $method = $_SERVER['REQUEST_METHOD'];
    
    switch ($method) {
        case 'GET':
            // Consulta para obtener todas las mascotas con información del usuario
            $stmt = $conn->prepare("
                SELECT m.id, m.nombre, m.especie, m.raza, m.edad, 
                       m.usuario_id, u.nombre as usuario_nombre 
                FROM mascotas m
                LEFT JOIN usuarios u ON m.usuario_id = u.id
                ORDER BY m.id
            ");
            $stmt->execute();
            
            $mascotas = $stmt->fetchAll(PDO::FETCH_ASSOC);
            
            echo json_encode([
                'success' => true,
                'data' => $mascotas,
                'count' => count($mascotas)
            ]);
            break;
            
        case 'POST':
            // Crear nueva mascota
            $data = json_decode(file_get_contents('php://input'), true);
            
            $stmt = $conn->prepare("
                INSERT INTO mascotas (nombre, especie, raza, edad, usuario_id) 
                VALUES (:nombre, :especie, :raza, :edad, :usuario_id)
            ");
            
            $stmt->bindParam(':nombre', $data['nombre']);
            $stmt->bindParam(':especie', $data['especie']);
            $stmt->bindParam(':raza', $data['raza']);
            $stmt->bindParam(':edad', $data['edad']);
            $stmt->bindParam(':usuario_id', $data['usuario_id']);
            
            $stmt->execute();
            
            echo json_encode([
                'success' => true,
                'message' => 'Mascota creada exitosamente',
                'id' => $conn->lastInsertId()
            ]);
            break;
            
        case 'PUT':
            // Actualizar mascota existente
            $data = json_decode(file_get_contents('php://input'), true);
            
            $stmt = $conn->prepare("
                UPDATE mascotas 
                SET nombre = :nombre, 
                    especie = :especie, 
                    raza = :raza, 
                    edad = :edad, 
                    usuario_id = :usuario_id 
                WHERE id = :id
            ");
            
            $stmt->bindParam(':id', $data['id']);
            $stmt->bindParam(':nombre', $data['nombre']);
            $stmt->bindParam(':especie', $data['especie']);
            $stmt->bindParam(':raza', $data['raza']);
            $stmt->bindParam(':edad', $data['edad']);
            $stmt->bindParam(':usuario_id', $data['usuario_id']);
            
            $stmt->execute();
            
            echo json_encode([
                'success' => true,
                'message' => 'Mascota actualizada exitosamente'
            ]);
            break;
            
        case 'DELETE':
            // Eliminar mascota
            $id = $_GET['id'] ?? null;
            
            if (!$id) {
                echo json_encode([
                    'success' => false,
                    'message' => 'ID de mascota no proporcionado'
                ]);
                exit;
            }
            
            $stmt = $conn->prepare("DELETE FROM mascotas WHERE id = :id");
            $stmt->bindParam(':id', $id);
            $stmt->execute();
            
            echo json_encode([
                'success' => true,
                'message' => 'Mascota eliminada exitosamente'
            ]);
            break;
            
        default:
            echo json_encode([
                'success' => false,
                'message' => 'Método no permitido'
            ]);
            break;
    }
    
} catch(PDOException $e) {
    echo json_encode([
        'success' => false,
        'message' => 'Error de base de datos: ' . $e->getMessage()
    ]);
}

$conn = null;
?>