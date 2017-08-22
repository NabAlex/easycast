let canvas = document.getElementById("canvas");

let count = 0;
let history = null;

function generateArray(history) {
    let ar = [];

    let len = history.length;
    for (let i = 0; i < len; i++) {
        ar.push([i, history[i]]);
    }

    return ar;
}

function draw(graph) {
    $.plot($("#placeholder"), [graph]);
}

$.get("/get?size=10", function(data) {
    history = data["history"];

    draw( generateArray(history) );
    let ws = new WebSocket('ws://localhost:8080/ws');

    ws.addEventListener('message', function(e) {
        let msg = JSON.parse(e.data);

        history = history.slice(1);
        history.push(msg["value"]);

        draw( generateArray(history) );
    });

});
