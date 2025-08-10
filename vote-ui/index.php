<?php
// Simple VoteCast UI
if (isset($_GET['counts'])) {
    echo file_get_contents('http://vote-counter:8080/counts');
    exit;
}
?>
<!DOCTYPE html>
<html>
<head><title>VoteCast</title></head>
<body>
  <h1>VoteCast Demo</h1>
  <button onclick="vote('A')">Vote A</button>
  <button onclick="vote('B')">Vote B</button>
  <pre id="counts"></pre>
  <script>
    async function vote(option) {
      await fetch(`http://localhost:8080/vote?option=${option}`, { method: 'POST' });
      const res = await fetch('http://localhost:8080/counts');
      document.getElementById('counts').textContent = await res.text();
    }
  </script>
</body>
</html>