# Technical Challenge Go-Bootcamp

Golang api to say hello to someone or the whole world.
It can also get the weather, min and max temperatures and coordinates of a City.

### Requirements

* Go 1.15

### Framework

This project utilizes Gorilla web toolkit.

### 1. Install dependencies

The project needs specific version of openweathermap module for it to be able to get the metereological information.
As this project utilizes go modules, the dependencies can be easily downloaded executing the following line:
```
go mod download
```

### 2. Run Tests

To run unit tests execute:
```
go test ./...
```

### 3. Environment variable

The environment variable OWM_API_KEY must be aded for being able to use openweathermap api. 
To add it simply execute the next line:
```
export OWM_API_KEY=0809093061a9fa551dd6e36ab91a180c
```

### 4. Usage

1. Enter the app folder
   ```
   cd app
   ```

2. Start the API with the next command, it will run on http://localhost:8000
   ```
   go run .
   ```

3. Hello Endpoint
   The hello endpoint is a basic hello world, it will display a "Hello, World!" message when you are on the base URL.
   You can add a "name" argument to display Hello and the name of the person you want to greet. You can add the argument and its value
   at the end of the URL.
   ```
   Eg. http://localhost:8000/?name=Sam
   ```
   
4. Weather Endpoint (/get-weather)
   The weather endpoint get real time information about the weather, coordinate locations and temperature on a specific city of the world.
   You must add a "city" argument as well as the name of the city that you will look for. If the city is incorrect or not valid, a message will notify you.
   ```
   Eg. http://localhost:8000/get-weather?city=Kioto
   ```
   Also you can add a comma and the country abbreviation after the city, this to make sure you are looking for the desired place.
   ```
   Eg. http://localhost:8000/get-weather?city=Guadalajara,mx
   ```
