const params = new URLSearchParams(window.location.search);
const quizId = params.get('quizId');
console.log("Loaded quiz:", quizId);


const idQuiz = document.getElementById("quiz").getAttribute("data-quiz-id")
console.log("Clicked quiz:", idQuiz);

var quiz;
var selectedQuestion
var score = 0


fetch(`/api/v1/getFullQuiz/${idQuiz}`)
  .then(response => {
    if (!response.ok) {
      throw new Error("Network response was not ok");
    }
    return response.json();  // <-- parse JSON body here
  })
  .then(data => {
    console.log("Quiz data:", data); 
    quiz = data;
    document.getElementById("quizName").textContent = data.quizName;
    document.getElementById("questionText").textContent = quiz.Questions[0].questionText;
    document.getElementById("questionPicture").src = quiz.Questions[0].questionImage;
    selectedQuestion = 1

    
    const container = document.getElementById("option-container");
    container.innerHTML = "";
  
    quiz.Questions[0].options.forEach(post => {
      
      if (post.optionImage = "0x") {
       
        document.getElementById("option-container").insertAdjacentHTML('beforeend', 
          ` <div class="option"  data-option-id="${post.idOption}">
                <div class="option-content">
                    <p>${post.optionText}</p>
                </div>
            </div>`);
      } else {
       
        document.getElementById("option-container").insertAdjacentHTML('beforeend', 
          ` <div class="option" data-option-id="${post.idOption}">
                <img src="forest.jpg">
                <div class="option-content">
                    <p>${post.optionText}</p>
                </div>
            </div>`);

      }

  });
    

    var i = 1
    const container2 = document.getElementById("questionNumbers");
    container2.innerHTML = "";
for(; i<= quiz.Questions.length; i++)
{
        document.getElementById("questionNumbers").insertAdjacentHTML('beforeend', 
          `<div class="question-circle" data-question-order="${i}" data-question-id="${quiz.Questions[i-1].idQuestion}">
                  <p>${i}</p>
            </div>`);
}

const firstQuestion = document.querySelectorAll(".question-circle");
console.log("ima "+ firstQuestion.length)
firstQuestion[0].classList.add("selected-question")
  })
  .catch(error => {
    console.error("Fetch error:", error);
  });



const container = document.getElementById("questionNumbers");



container.addEventListener("click", function(e) {

      const target = e.target.closest(".question-circle");
      if (target) {
        const index = parseInt(target.getAttribute("data-question-order"), 10);
        //alert( quiz.Questions[index-1].questionText);
        document.getElementById("questionText").textContent = quiz.Questions[index-1].questionText;
        document.getElementById("questionPicture").src = quiz.Questions[index-1].questionImage;
        selectedQuestion = index


        const idQuestion = target.getAttribute("data-question-id");
        //alert(idQuestion)
        const idOrder = target.getAttribute("data-question-order");
        //alert(idOrder)
        const questions = document.querySelectorAll(".question-circle");

        questions.forEach(o => {
          o.classList.remove("selected-question");
        })

        target.classList.add("selected-question")


          const container = document.getElementById("option-container");
          container.innerHTML = "";
        
          quiz.Questions[idOrder-1].options.forEach(post => {
            if (post.optionImage = "0x") {
            
              document.getElementById("option-container").insertAdjacentHTML('beforeend', 
                ` <div class="option"  data-option-id="${post.idOption}">
                      <div class="option-content">
                          <p>${post.optionText}</p>
                      </div>
                  </div>`);
            } else {
          
              document.getElementById("option-container").insertAdjacentHTML('beforeend', 
                ` <div class="option"  data-option-id="${post.idOption}">
                      <img src="forest.jpg">
                      <div class="option-content">
                          <p>${post.optionText}</p>
                      </div>
                  </div>`);
      
            }
        });

      }
  });




  
var selectedOptionID
var prevSelectedOptionID
var isAnySelected = false

const optionContainer = document.getElementById("option-container");

optionContainer.addEventListener("click", function(e) {

      const target = e.target.closest(".option");
      if (target) {
       
  
        selectedOptionID = target.getAttribute("data-option-id")
        
        
        const options = document.querySelectorAll(".option");

        options.forEach(o => {
          o.classList.remove("selected");
        })
        
        //console.log("COMPARE " + selectedOptionID + " "+ prevSelectedOptionID)
        if(selectedOptionID != prevSelectedOptionID)
        {
          target.classList.add("selected")
          console.log("SELEKTOVANO")
          isAnySelected = true
        }
        else
        {
          isAnySelected = false
          prevSelectedOptionID = null
          console.log("DEEEESELEKTOVANO")
        }
        
        if(isAnySelected)
          prevSelectedOptionID = target.getAttribute("data-option-id")
      }
  });


  function submit() {
    var isCorrect = false
    alert("HEJ! " + selectedQuestion);
    if(!isAnySelected)
      alert("Please select something!");
    else{
      quiz.Questions[selectedQuestion-1].options.forEach(post => {

        //alert("korekcija "+ post.isCorrect +" "+ post.idOption == selectedOptionID);
        if(post.idOption == selectedOptionID && post.isCorrect)
        {
          isCorrect=true
          
          
        }
      })

      if(isCorrect)
      {
        console.log("Correct!");
        score+=1
      }
      else
        console.log("Incorrect!");


    
    if(selectedQuestion < quiz.Questions.length)
    {
      selectedQuestion += 1
      document.getElementById("questionText").textContent = quiz.Questions[selectedQuestion-1].questionText;
    
    const container = document.getElementById("option-container");
    container.innerHTML = "";
    const questions = document.querySelectorAll(".question-circle");

    questions.forEach(o => {
      o.classList.remove("selected-question");
    })

    questions[selectedQuestion-1].classList.add("selected-question")
    quiz.Questions[selectedQuestion-1].options.forEach(post => {
      
      if (post.optionImage = "0x") {
        
        document.getElementById("option-container").insertAdjacentHTML('beforeend', 
          ` <div class="option"  data-option-id="${post.idOption}">
                <div class="option-content">
                    <p>${post.optionText}</p>
                </div>
            </div>`);
      } else {
      
        document.getElementById("option-container").insertAdjacentHTML('beforeend', 
          ` <div class="option" data-option-id="${post.idOption}">
                <img src="forest.jpg">
                <div class="option-content">
                    <p>${post.optionText}</p>
                </div>
            </div>`);

      }

    })
    }
    else
    {
      console.log("You reached the end. Your score is " + score)
    }
  }
}


  