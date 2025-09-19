# Hiven
The Ultimate Roomate App

## Project Setup

> All environments use Docker and Make. No other dependencies are required.

### Development environment
**Steps**:
- `cd /setup/dev`
- `make start` to start development environment
- `make stop` to stop development environment

**Try the API**
- Import `backend/postman-collection/Hiven-API.postman_collection.json` into Postman to test the API endpoints.
- The collection uses the base URL: **http://localhost:3000** 

### Testing environment
**Steps**:
- `cd /setup/test`
- `make tests` will run tests in an isolated environment and output the results

## 🤝 Maintainers

- Nakul Dighe ([Nakul-D](https://github.com/Nakul-D))
- Chirag Bhalotia ([chirag3003](https://github.com/chirag3003))

## 📄 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
