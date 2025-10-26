// --- Navigation Buttons (from scoreboard.html) ---
document.getElementById('index-btn').onclick = () => window.location.href = '/';
document.getElementById('completed-btn').onclick = () => window.location.href = '/completed';
document.getElementById('scoreboard-btn').onclick = () => window.location.href = '/scoreboard';
document.getElementById('settings-btn').onclick = () => window.location.href = '/settings';

// --- Scoreboard Elements ---
const scoreboardList = document.getElementById('scoreboard-list');
const loadingText = document.getElementById('loading-text');

// --- Auth & API Config ---
const AUTH_TOKEN = localStorage.getItem('AUTH_TOKEN');
const CURRENT_USERNAME = localStorage.getItem('USERNAME');

// --- Main Scoreboard Loader ---
async function loadScoreboard() {
  try {
    // 1. Fetch all user IDs
    const res = await fetch('http://127.0.0.1:8081/users', {
      headers: { 'auth': AUTH_TOKEN }
    }); // [cite: api.go, get_user.go]
    
    if (!res.ok) throw new Error('Failed to fetch user list');
    
    const { usersIDs } = await res.json(); // Gets {"usersIDs": [1, 2, ...]} [cite: get_user.go]

    // 2. Create an array of promises to fetch details for each user
    const userPromises = usersIDs.map(id => {
      return fetch(`http://127.0.0.1:8081/users/${id}`, {
        headers: { 'auth': AUTH_TOKEN }
      }) // [cite: api.go, get_user.go]
      .then(res => res.json());
    });

    // 3. Wait for all user details to be fetched
    const users = await Promise.all(userPromises);
    
    // 4. Display the scores
    displayScores(users, CURRENT_USERNAME);

  } catch (err) {
    console.error(err);
    loadingText.textContent = `Error loading scoreboard: ${err.message}`;
    loadingText.style.color = 'red';
  }
}

// --- Sort and display users ---
function displayScores(users, currentUser) {
  loadingText.style.display = "none";
  scoreboardList.innerHTML = "";

  // Sort by lowercase 'points'
  users.sort((a, b) => b.points - a.points);

  users.forEach((user, index) => {
    const li = document.createElement("li");
    li.classList.add("score-item");

    // Highlight by lowercase 'name'
    if (user.name === currentUser) {
      li.classList.add("current-user");
    }

    // This is the calculated position (#1, #2, #3...)
    const position = document.createElement("span");
    position.classList.add("rank");
    position.textContent = `#${index + 1}`;

    // Username (lowercase 'name')
    const username = document.createElement("span");
    username.classList.add("username");
    username.textContent = user.name; 

    // --- REMOVED: The (Rank: ${user.rank}) element ---

    // Points (lowercase 'points')
    const points = document.createElement("span");
    points.classList.add("points");
    points.textContent = `${user.points} pts`; 

    li.appendChild(position);
    li.appendChild(username);
    // --- REMOVED: li.appendChild(userRank) ---
    li.appendChild(points);

    scoreboardList.appendChild(li);
  });
}

// --- On page load ---
document.addEventListener("DOMContentLoaded", () => {
  // Check if user is logged in
  if (!AUTH_TOKEN || !CURRENT_USERNAME) {
    loadingText.textContent = "Please log in on the Settings page to view the scoreboard.";
    return;
  }

  // Load the scoreboard data
  loadScoreboard();
});

