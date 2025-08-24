# Hiven
The Ultimate Roomate App

## To run the servers in development mode

**Start command:** `docker-compose -f docker-compose.dev.yaml -p hiven-dev up -d`

> Note: The API server might take a few seconds to become available after the container starts because Air performs a rebuild and then starts the application.

**Stop command:** `docker-compose -f docker-compose.dev.yaml -p hiven-dev down`

## Try the API

- Import `Hiven-API.postman_collection.json` into Postman to test the API endpoints.
- The collection uses the base URL: **http://localhost:3000** 

## 🤝 Maintainers

- Nakul Dighe ([Nakul-D](https://github.com/Nakul-D))
- Chirag Bhalotia ([chirag3003](https://github.com/chirag3003))

## 📄 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
