

fetch("/api/v1/quizs")
  .then(response => {
    if (!response.ok) {
      throw new Error("Network response was not ok");
    }
    return response.json();  // <-- parse JSON body here
  })
  .then(data => {
    console.log("Quiz data:", data); // now you see the actual JSON array/object

    
    data.forEach(post => {
          document.getElementById("quizGrid").insertAdjacentHTML('beforeend', 
            `<div class="card">
              <img src="/static/ARTPOP.jpg">

              <div class="card-content">
                <h3>${post.quizName}</h3>
                <p>${post.description}</p>
                <a href="" type="button" class="btn" data-quiz-id="${post.idQuiz}">Play</a>
              </div>
            </div>`);
    });
  })
  .catch(error => {
    console.error("Fetch error:", error);
  });



  const container = document.getElementById("quizGrid");

  container.addEventListener("click", function(e) {
      // Check if the clicked element has the class we want
      if (e.target && e.target.classList.contains("btn")) {
          e.preventDefault(); // stop page reload
          const quizId = e.target.dataset.quizId;
          console.log("Clicked quiz:", quizId);
          
          // Fetch quiz data or open editor
          fetch(`/api/v1/quiz/${quizId}`)
              .then(res => res.json())
              .then(data => console.log("Quiz data:", data))
              .catch(error => {
                console.error("Fetch error:", error);
              });
            window.location.href = `/static/quizWindow.html?quizId=${quizId}`
      }
  });




  function decodeJwtPayload(token) {
    const payload = token.split('.')[1]; // get the middle part
    if (!payload) throw new Error("Invalid token");
  
    // Replace URL-safe characters and add padding if needed
    const base64 = payload.replace(/-/g, '+').replace(/_/g, '/');
    const padded = base64.padEnd(base64.length + (4 - base64.length % 4) % 4, '=');
  
    // Decode base64 to string
    const jsonPayload = atob(padded);
  
    // Parse JSON
    return JSON.parse(jsonPayload);
  }

  document.getElementById("loginForm").addEventListener("submit", function(event) {
    event.preventDefault(); // prevent page reload on form submit
  
    const formData = new FormData(event.target);
    const data = Object.fromEntries(formData.entries());
  
    // If you want to check data before sending:
    console.log("Form data object:", data);
    console.log("json",JSON.stringify(data))

  
    fetch("/api/v1/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    })
    .then(async response => {
      if (!response.ok)
      {
        // Try to parse JSON error message from response body
        let errorMsg = "Unknown error";
        try 
        {
          const errorData = await response.json();
          errorMsg = errorData.message || JSON.stringify(errorData);
        } 
        catch 
        {
          // If response isn't JSON, try text
          try 
          {
            errorMsg = await response.text();
          } catch {}
        }
        throw new Error(`HTTP ${response.status}: ${errorMsg}`);
      }

      return response.json();
    })
    .then(result => {
      console.log("Server response:", result);

      const token = result.token; 
      const decoded = decodeJwtPayload(token);
      console.log("Username is", decoded.username);
      document.getElementById("welcomeParag").textContent = "Welcome "+ decoded.username
      // Here you can redirect user or show a success message
    })
    .catch(error => {
      console.error("Fetch error:", error);
      // Here you can show an error message to the user
    });
  });