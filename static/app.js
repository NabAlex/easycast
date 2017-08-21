
let ws = new WebSocket('ws://localhost:8080/ws');

ws.addEventListener('message', function(e) {
    let msg = JSON.parse(e.data);
    alert(JSON.stringify(msg));
});