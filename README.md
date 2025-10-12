# Hiven
The Ultimate Roommate App

## Project Setup

> All environments use Docker and AWS CLI. No other dependencies are required.

### Development/Testing environment

**Prerequisites**
- AWS account
- AWS CLI (installed and configured)
- Docker and Docker Compose
- WSL (only on Windows)

**Steps**:
- Change directory into `cd /setup/dev`
- Run bash script
  - On Linux and macOS run `./main.sh`
  - On Windows you need WSL. Run `wsl ./main.sh`

**Try the API**
- Import `backend/postman-collection/Hiven-API.postman_collection.json` into Postman to test the API endpoints.
- The collection uses the base URL: **http://localhost:3000** 

## 🤝 Maintainers

- Nakul Dighe ([Nakul-D](https://github.com/Nakul-D))
- Chirag Bhalotia ([chirag3003](https://github.com/chirag3003))

## 📄 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
