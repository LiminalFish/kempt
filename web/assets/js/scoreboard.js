const scoreboardList = document.getElementById('scoreboard-list');
const loadingText = document.getElementById('loading-text');

// Mock data (replace with backend API call)
const mockGroupScores = [
  { username: "alice", score: 120 },
  { username: "bob", score: 95 },
  { username: "charlie", score: 150 },
  { username: "you", score: 110 }
];

// Sort and display users
function displayScores(users, currentUser) {
  loadingText.style.display = "none";
  scoreboardList.innerHTML = "";

  // Sort by descending score
  users.sort((a, b) => b.score - a.score);

  users.forEach((user, index) => {
    const li = document.createElement("li");
    li.classList.add("score-item");

    const rank = document.createElement("span");
    rank.classList.add("rank");
    rank.textContent = `#${index + 1}`;

    const username = document.createElement("span");
    username.classList.add("username");
    username.textContent = user.username === currentUser ? `${user.username} (You)` : user.username;

    const points = document.createElement("span");
    points.classList.add("points");
    points.textContent = `${user.score} pts`;

    li.appendChild(rank);
    li.appendChild(username);
    li.appendChild(points);

    scoreboardList.appendChild(li);
  });
}

// On page load
document.addEventListener("DOMContentLoaded", () => {
  const currentUser = localStorage.getItem("username");

  if (!currentUser) {
    loadingText.textContent = "Please log in to view your group's scoreboard.";
    return;
  }

  // TODO: Replace mockGroupScores with real backend API call like:
  // fetch(`/api/scoreboard?user=${currentUser}`)
  //   .then(res => res.json())
  //   .then(data => displayScores(data, currentUser));

  displayScores(mockGroupScores, currentUser);
});
