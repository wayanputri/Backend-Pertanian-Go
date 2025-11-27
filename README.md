# Backend-Pertanian-Go

kalo mau buka minio bisa
url= http://localhost:9001
trus di cmd di ketik ( minio.exe server D:/minio --console-address ":9001" )

kalo mau buka swagger
```http://localhost:8088/docs/```

swagger custom ```
	"/api/profile": {
  "post": {
    "summary": "Update profile",
    "description": "This API is for updating the profile including profile picture.",
    "operationId": "ExampleService_UpdateProfile",
    "responses": {
      "200": {
        "description": "A successful response.",
        "schema": {
          "$ref": "#/definitions/exampleUploadProfileResponse"
        }
      },
      "default": {
        "description": "An unexpected error response.",
        "schema": {
          "$ref": "#/definitions/runtimeError"
        }
      }
    },
    "parameters": [
      {
        "name": "profile_picture",
        "in": "formData",
        "description": "Profile Picture",
        "required": true,
        "type": "file"
      }
    ],
    "tags": ["User"]
  }
},
```