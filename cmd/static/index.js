fetch("http://localhost:8080/api/v1/quizs")
    .then(response => {console.log(response)})
    .catch(error => console.error(error));
    