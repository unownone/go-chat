
# <p color=#29BEB0>Go-Chat</p>

## Fast Fiber-ous Chat Application

# Deployed At [heroku](https://go-chatify.herokuapp.com)


# Currently Supports:
- Creation of chatrooms
- Adding users to chatroom
- Messaging to chatroom using Pure Websockets
# Upcoming Features:
- Invite to Chatrooms
- Delete Messages
- Tag Users
- Search for Users
- Email Verification
- Client Website and Application

## Last Update:
- Added New Chat API (v2) with a different Auth Method

# API Routes:

### Users ```/api/v1/auth```

- `/register` : Register a new user
    - Method : `POST`
    - Body : 
        ```
            {
                "username": "username",
                "email": "abc@xyz.com",
                "password": "password"
            }
        ```
    - Response :
        ```
            {
                "error": false,
                "message": "User Registered Successfully",
                "access_token":"",
                "refresh_token":""
            }
        ```
- `/login` : Login a user
    - Method : `POST`
    - Body : 
        ```
            {
                "email": "abc@xyz.com",
                "password": "password"
            }
        ```
    - Response :
        ```
            {
                "error": false,
                "message": "Logged in Successfully",
                "access_token":"",
                "refresh_token":""
            }
        ```
- `/refresh` : Refresh a JWT Token
    - Method : `POST`
    - Body : 
        ```
            {
                "refresh_token": "refresh_token"
            }
        ```
    - Response :
        ```
            {
                "access_token":"",
                "refresh_token":""
            }
        ```
- `/current-user` : Get the current user
    - Method : `GET`
    - Response :
        ```
            {
                "user":"User Email ID"
            }
        ```
- `/user` : Get a user by their email
    - Method : `GET`
    - Query : `?email=<user email>`
    - Response :
        ```
            {
                "id": "User ID",
                "email":"User Email ID",
                "name":"User Name",
            }
        ```





### Chats ```/api/v1/chat```

- `/chats` : Get all chats
    - Method : `GET`
    - Response :
        ```
            {
                "chats": [
                    {
                        "chat_id": "chat_id",
                        "chat_name": "chat_name",
                        "chat_users": [
                            "user_id",
                            "user_id"
                        ]
                    },
                    {
                        "chat_id": "chat_id",
                        "chat_name": "chat_name",
                        "chat_users": [
                            "user_id",
                            "user_id"
                        ]
                    }
                ]
            }
        ```
- `/chats` : Create a Chat
    - Method : `POST`
    - Body : 
        ```
            {
                "name": "Chat Name",
                "users": ["User1", "User2"] // User IDs 
            }
        ```
    - Response :
        ```
            {
                "error": false,
                "message": "Chat Created Successfully",
                "chat": {
                    "id": "Chat ID",
                    "name": "Chat Name",
                    "users": ["User1", "User2"] // User IDs 
                }
            }
        ```
- `/chats` : Update a Chat
    - Method : `PUT`
    - Body : 
        ```
            {
                "name": "Chat Name", //[Optional]
                "users": ["User1", "User2"] // User IDs [Optional]
            }
        ```
    - Response :
        ```
            {
                "error": false,
                "message": "Chat Updated Successfully",
                "chat": {
                    "id": "Chat ID",
                    "name": "Chat Name", // Updated Name
                    "users": ["User1", "User2"] // User IDs , Updated Valid USERIDS
                }
            }
        ```
### Chats ```/api/v1/chat``` [WebSockets]
- /<sess_id>/<chat_id>
    `<sess_id>` : User's JWT access Token
    `<chat_id>` : Chat ID to join.

Creates a Websocket Connection.

### Chats ```/api/v2/chat``` [WebSockets] [NEW]
- /<sess_id>
    `<sess_id>` : User's JWT access Token
    User is authenticated
    To connect to chatroom send TEXTMessage: "!startChat <chat_id>"
        
# How to Run:

- Clone the Repo Using 
    ```
    git clone https://github.com/unownone/fiberous.git
    ```
- Get Packages using :
    ```
    go get
    cd go-chat
    ```
- Run:
    ```
    go build .
    ./go-chat
    ```
