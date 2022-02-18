namespace Network {

    function Get(url: string, data: any, callback: (r: any) => void ) {

        let xhr = new XMLHttpRequest();
        xhr.open("POST", url);

        xhr.onload = function (event) {
            callback(JSON.parse(this.response));
        }

        xhr.setRequestHeader("Content-Type", "application/json; charset=utf-8")
        xhr.send(JSON.stringify(data));
    }

    Get("/user", {Date: "sdasfd"}, function (r) {
        
    })
}