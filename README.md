# Go OTP Verification using Twilio 

**Project Summary:**

This project demonstrates how to implement OTP (One-Time Password) verification using the Twilio API in Go. Users receive an OTP to their phone number, which they can then use to authenticate themselves with your application.

**Key Features:**

- Secure and reliable OTP verification via Twilio's trusted service
- Clean and well-structured Go code for maintainability
- User-friendly API endpoints for sending and verifying OTPs
- Detailed examples and instructions for easy setup and usage

**Setup:**

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/jaybaba8085/Go-SMS-verification-twilio.git
2. **Create a Twilio Account:**

    - Sign up for a free Twilio account at https://www.twilio.com/.
    - Obtain your account SID and auth token from the Account Dashboard.
    - Create a Verify service and note down the service SID.


3. **Create a .env file:**
 
    - Create a .env file in the project's root directory.

    - Add the following lines, replacing the placeholders with your actual credentials:

        ```bash
        TWILIO_ACCOUNT_SID=<ACCOUNT SID>
        TWILIO_AUTHTOKEN=<AUTH TOKEN>
        TWILIO_SERVICES_ID=<SERVICE ID>
        ```

4. **Install dependencies:**

    ```bash
    go mod download
    ```

6. **Running the Server:**

    ```bash
    go run cmd/main.go
    ```

## API Documentation
---
### Send OTP:

Send a POST request to the /otp endpoint with the following body to send an OTP to a user's phone number

### Endpoint:

```bash
POST /otp
```

### Request Body:

```json
{
  "phoneNumber":"<phone-number-with-country-code>"
}
```

```bash
curl -H "Content-Type: application/json" -X POST -d '{"phoneNumber": "+91##########"}' http://localhost:8000/otp
```
OR
### Example using REST API extension in Visual Studio Code:  get-OTP.http file

```bash
POST http://localhost:8000/otp
Content-Type: application/json

{
  "phoneNumber": "+91##########"
}

```

### Response:

```json
{
  "status": 202,
  "message": "success",
  "data": "OTP sent successfully"
}
```
---
### Verify OTP:

Verify a user's OTP by sending a POST request to the /verify endpoint with the following body that contains the phone number and the OTP code received by the user

```bash
POST /verifyOTP
```

### Request Body:

```json
{
  "user": {
    "phoneNumber": "<phone-number-with-country-code>"
  },
  "code": "<code here>"
}
```

```bash
curl -H "Content-Type: application/json" -X POST -d '{"user": {"phoneNumber": "+91##########"}, "code":"123456"}' http://localhost:8000/verifyOTP
```
OR
### Example using REST API extension in Visual Studio Code: verifyOTP.http file


```bash
POST http://localhost:8000/verifyOTP
Content-Type: application/json

{
  "user": {
    "phoneNumber": "+91##########"
  },
  "code":"123456"
}
```


### Response:

```json
{
  "status": 202,
  "message": "success",
  "data": "OTP verified successfully"
}
```
---
### Successfull HTTP Request's Response:
![HTTP REQUEST IMAGE](./image.png)