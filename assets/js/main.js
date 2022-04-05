function SEND(method, url, data, callback) {
    let xhr = new XMLHttpRequest();
    xhr.open(method, url);
    xhr.onload = (ev) => {
        callback && callback(JSON.parse(ev.target.response));
    };
    xhr.setRequestHeader("Content-Type", "application/json; charset=utf-8");
    xhr.send(JSON.stringify(data));
}

let buttonAddTask = document.querySelector("#AddTask");
if (buttonAddTask) {
    buttonAddTask.onclick = function () {
        let inputs = document.querySelectorAll(".form > input");
        let data = {};

        for (let i = 0; i < inputs.length; i++) {
            data[inputs[i].name] = inputs[i].value;
        }

        SEND("PUT", "/task", data, buildTasks);
    }
}

function buildTasks(r) {
    let blocks = {
        active: document.querySelector(`[data-status="active"]`),
        complete: document.querySelector(`[data-status="complete"]`),
        delete: document.querySelector(`[data-status="delete"]`),
    }

    for (let key in blocks) {
        let block = blocks[key];
        while (block.children.length > 0) {
            block.children[0].remove();
        }
    }

    for (let i = 0; i < r.length; i++) {
        let div = document.createElement("div");
        div.dataset.id = r[i].id;
        div.draggable = true;
        div.className = "item";
        div.textContent = r[i].name;
        console.log(r[i].status);
        blocks[r[i].status] && blocks[r[i].status].append(div);
    }

    setDraggable();
}

// DragAndDrop

function setDraggable() {
    let dragElements = document.querySelectorAll(".task-list .item");
    for (let i = 0; i < dragElements.length; i++) {
        dragElements[i].ondragstart = onDrag;
    }
}

let dropElements = document.querySelectorAll(".task-list > div");
for (let i = 0; i < dropElements.length; i++) {
    dropElements[i].ondrop = onDrop;
    dropElements[i].ondragover = onDragOver;
}

function onDrag(e) {
    e.dataTransfer.setData("item", e.target.dataset.id);
}

function onDrop(e) {
    e.preventDefault();
    let itemId = e.dataTransfer.getData("item");
    let item = document.querySelector('[data-id="' + itemId + '"]');
    item && e.target.append(item);

    SEND("POST", "/task", {
        id: +itemId,
        status: e.target.dataset.status
    }, buildTasks);
}

function onDragOver(e) {
    e.preventDefault();
}

SEND("GET", "/task", null, buildTasks)