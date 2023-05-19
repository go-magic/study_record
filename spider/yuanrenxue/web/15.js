var window = {}
function loadWebAssembly (path, imports = {}) {
    return fetch(path) // 加载文件
        .then(response => response.arrayBuffer()) // 转成 ArrayBuffer
        .then(buffer => WebAssembly.compile(buffer))
        .then(module => {

        })
}
loadWebAssembly('./main.wasm').then(instance => {
        window.q = instance.exports.encode
        window.m = function (){
        var    t1 = parseInt(Date.parse(new Date())/1000/2);
        var    t2 = parseInt(Date.parse(new Date())/1000/2 - Math.floor(Math.random() * (50) + 1));
            return window.q(t1, t2).toString() + '|' + t1 + '|' + t2;
        }
    })
