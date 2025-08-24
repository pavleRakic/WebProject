const idQuiz = document.getElementById("quiz").getAttribute("data-quiz-id")
console.log("Clicked quiz:", idQuiz);



// Fetch quiz data or open editor

fetch(`/api/v1/quiz/${idQuiz}`)
  .then(response => {
    if (!response.ok) {
      throw new Error("Network response was not ok");
    }
    return response.json();  // <-- parse JSON body here
  })
  .then(data => {
    console.log("Quiz data:", data); 
    document.getElementById("quizName").textContent = data.quizName;
  })
  .catch(error => {
    console.error("Fetch error:", error);
  });

  var questions;


  fetch(`/api/v1/quiz/${idQuiz}/questions`)
  .then(response => {
    if (!response.ok) {
      throw new Error("Network response was not ok");
    }
    return response.json();  // <-- parse JSON body here
  })
  .then(data => {
    console.log("Quiz data:", data); 
    const numberOfQuestions = data.length
    questions = data;

    var i = 1
  data.forEach(post => {
    alert(data[i-1].idQuestion);
    document.getElementById("questionNumbers").insertAdjacentHTML('beforeend', 
      `<div class="question-circle" data-question-order="${i}" data-question-id="${data[i-1].idQuestion}">
              <p>${i++}</p>
        </div>`);
        
});
    document.getElementById("questionText").textContent = data[0].questionText;

  })
  .catch(error => {
    console.error("Fetch error:", error);
  });


  const container = document.getElementById("questionNumbers");

  container.addEventListener("click", function(e) {
      // Check if the clicked element has the class we want

      
      const target = e.target.closest(".question-circle");
      if (target) {
        const index = parseInt(target.getAttribute("data-question-order"), 10);
        alert( questions[index-1].questionText);
        document.getElementById("questionText").textContent = questions[index-1].questionText;


        const idQuestion = target.getAttribute("data-question-id");
        alert(idQuestion)

        fetch(`/api/v1/question/${idQuestion}/options`)
        .then(response => {
          if (!response.ok) {
            throw new Error("Network response was not ok");
          }
          return response.json();  // <-- parse JSON body here
        })
        .then(data => {
          console.log("Quiz data:", data); 
          const container = document.getElementById("option-container");
          container.innerHTML = "";

          data.forEach(post => {
            document.getElementById("option-container").insertAdjacentHTML('beforeend', 
              ` <div class="option">
                    <img src="forest.jpg">
                    <div class="option-content">
                        <p>${post.optionText}</p>
                    </div>
                </div>`);
        });
        })
        .catch(error => {
        console.error("Fetch error:", error);
        });
      }
  });






