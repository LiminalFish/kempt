// // Get D








// Get DOM elements
const loginBtn = document.getElementById('login-btn');
const logoutBtn = document.getElementById('logout-btn');
const authStatus = document.getElementById('auth-status');
const loginForm = document.getElementById('login-form');
const usernameInput = document.getElementById('username');
const passwordInput = document.getElementById('password');

// Function to show logged-in state
function showLoggedIn(username) {
  authStatus.textContent = `Logged in as ${username}`;
  loginForm.style.display = "none";
  logoutBtn.style.display = "inline-block";
}

// Function to show logged-out state
function showLoggedOut() {
  authStatus.textContent = "You are currently logged out.";
  logoutBtn.style.display = "none";
  loginForm.style.display = "block";
  usernameInput.value = "";
  passwordInput.value = "";
}

// Check localStorage on page load
document.addEventListener("DOMContentLoaded", () => {
  const storedUsername = localStorage.getItem("username");
  if (storedUsername) {
    showLoggedIn(storedUsername);
  } else {
    showLoggedOut();
  }
});

// Login button click
loginBtn.addEventListener('click', () => {
  const username = usernameInput.value.trim();
  const password = passwordInput.value.trim();

  if (username && password) {
    // TODO: call backend login API
    localStorage.setItem("username", username); // persist login
    showLoggedIn(username);
  } else {
    alert("Please enter both username and password");
  }
});

// Logout button click
logoutBtn.addEventListener('click', () => {
  localStorage.removeItem("username"); // remove persisted login
  showLoggedOut();
});
