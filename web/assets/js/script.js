const hamburger = document.getElementById('hamburger');
const taskbar = document.getElementById('taskbar');
document.getElementById('index-btn').addEventListener('click', () => {
  window.location.href = '/';
});
// Redirect Scoreboard button
document.getElementById('scoreboard-btn').addEventListener('click', () => {
  window.location.href = '/scoreboard';
});
document.getElementById('settings-btn').addEventListener('click', () => {
  window.location.href = '/settings';
});

hamburger.addEventListener('click', () => {
  hamburger.classList.toggle('active');
  taskbar.classList.toggle('hidden');
});